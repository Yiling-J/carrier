# carrier - A Test Fixture Replacement for Go
![example workflow](https://github.com/Yiling-J/carrier/actions/workflows/go.yml/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/Yiling-J/carrier?style=flat-square)

- **Statically Typed** - 100% statically typed using code generation
- **Developer Friendly API** - explicit API with method chaining support
- **Feature Rich** - Default/Sequence/SubFactory/PostHook/Trait
- **Ent Support** - [ent: An Entity Framework For Go](https://github.com/ent/ent)

A snippet show how carrier works:

- *You have a model*
```go
type User struct {
	Name  string
	Email string
	Group *Group
}

```
- *Add carrier schema*
```go
Schemas := []carrier.Schema{
	&carrier.StructSchema{
		To: model.User{},
	},
}
```
- *Generate Structs* 🎉
```go
userMetaFactory := carrier.UserMetaFactory()
userFactory := userMetaFactory.
	SetNameDefault("carrier").
	SetEmailLazy(func(ctx context.Context, i *model.User) (string, error) {
		return fmt.Sprintf("%s@carrier.go", i.Name), nil
	}).
	SetGroupFactory(groupFactory.Create).
	Build()
user, err := userFactory.Create(ctx)
users, err := userFactory.CreateBatch(5, ctx)
```

## Installation
```console
go get github.com/Yiling-J/carrier/cmd
```
After installing `carrier` codegen, go to the root directory(or the directory you think carrier should stay) of your project, and run:
```console
go run github.com/Yiling-J/carrier/cmd init
```
The command above will generate `carrier` directory under current directory:
```console {12-20}
└── carrier
    └── schema
        └── schema.go
```
It's up to you where the carrier directory should be, just remember to use the right directory in **MetaFactory Generation** step.

## Add Schema
Edit `schema.go` and add some schemas:

**> struct**
```go
package schema

import (
	"github.com/Yiling-J/carrier"
)

var (
	Schemas = []carrier.Schema{
		{
			To: model.User{},
		},
	}
)
```

**> ent**

To support ent, you need to provide the `SchemaCreate` struct to schema, so carrier can get enough information.
```go
package schema

import (
	"github.com/Yiling-J/carrier"
	"your/ent"
)

var (
	Schemas = []carrier.Schema{
		{
			To: &ent.UserCreate{},
		},
	}
)
```

The `To` field only accept struct/struct pointer, carrier will valid that on generation step. [Schema definition reference](#schema-definition)

## MetaFactory Generation
Run code generation from the root directory of the project as follows:
```console
# this will use default schema path ./carrier/schema
go run github.com/Yiling-J/carrier/cmd generate
```
Or can use custom schema path:
```console
go run github.com/Yiling-J/carrier/cmd generate ./your/carrier/schema
```

This produces the following files:
```console {12-20}
└── carrier
    ├── factory
    │   ├── base.go
    │   ├── ent_user.go
    │   └── user.go
    ├── schema
    │   └── schema.go
    └── factory.go
```
Here `factory.go` include all meta factories you need.
Also all ent files and meta factories will have `ent` prefix to avoid name conflict.

If you update schemas, just run `generate` again.

## Build Factory and Generate Structs
To construct a real factory for testing:

**Create MetaFactory struct**
```go
userMetaFactory := carrier.UserMetaFactory()
```
**Build factory from meta factory**
```go
userFactory := userMetaFactory.SetNameDefault("carrier").Build()
```
MetaFactory provide several methods to help you initial field values automatically.

[MetaFactory API Reference](#metafactory-api)

**Create structs**

**> struct**
```go
user, err := userFactory.Create(context.TODO())
users, err := userFactory.CreateBatch(context.TODO(), 3)
```
**> ent**
```go
// need ent client
user, err := userFactory.Client(entClient).Create(context.TODO())
user, err := userFactory.Client(entClient).CreateBatch(context.TODO(), 3)
```
[Factory API Reference](#factory-api)

**Use factory wrapper**

Carrier also include a wrapper where you can put all your factories in:

**> struct**
```go
factory := &carrier.Factory{}
factory.SetUserFactory(userFactory)
factory.UserFactory().Create(context.TODO())
```
**> ent**
```go
factory := &carrier.EntFactory{}
factory.SetClient(entClient).SetUserFactory(userFactory)

// no .Client(entClient) for userFactory
// because we already set that in wrapper factory and userFactory will inherit it
factory.UserFactory().Create(context.TODO())
```
## Schema Definition
There are 2 kinds of schemas `StructSchema` and `EntSchema`,
both of them implement `carrier.Schema` interface so you can put them in the schema slice.

Each schema has 4 fields:
- **Alias**: Optional. If you have 2 struct type from different package, but have same name, add alias for them.
Carrier will use alias directly as factory name.

- **To**: Required. For `StructSchema`, this is the struct factory should generate. Carrier will get struct type from it and used in code generation, Only public fields are concerned.
For `EntSchema`, this field should be the `{SchemaName}Create` struct which `ent` generated. Carrier will look up all `Set{Field}` methods
and generate factory based on them. Both struct and pointer of struct are OK.

- **Traits**: Optional. String slice of trait names. Traits allow you to group attributes together and override them at once.

- **Posts**: Optional. Slice of `carrier.PostField`. Each `carrier.PostField` require `Name`(string) and `Input`(any interface{}), and map to a post function after code generation.
Post function will run after struct created, with input value as param.

## MetaFactory API
MetaFactory API can be categorized into 7 types of method:
- Each field in `To` struct has 4 types:
	- **Default**: `Set{Field}Default`
	- **Sequence**: `Set{Field}Sequence`
	- **Lazy**: `Set{Field}Lazy`
	- **Factory**: `Set{Field}Factory`

- Each field in `[]Posts` has 1 type:
	- **Post**: `Set{PostField}PostFunc`

- Each name in `[]Traits` has 1 type:
	- **Trait**: `Set{TraitName}Trait`

- Each `MetaFactory` has 1 type:
	- **AfterCreate**: `SetAfterCreateFunc`

The evaluation order of these methods are:
```
Trait -> Default/Sequence/Factory -> Lazy -> Create -> AfterCreate -> Post
```
`Create` only exists in ent factory, will call ent builder `Save` method.

Put `Trait` first because `Trait` can override other types.

All methods except `Default` and `Trait` use function as input and it's fine to set it to `nil`. This is very useful in `Trait` override.

#### Default
Set a fixed default value for field.
```go
userMetaFactory.SetNameDefault("carrier")
```

#### Sequence
If a field should be unique, and thus different for all built structs, use a sequence.
Sequence counter is shared by all fields in a factory, not a single field.
```go
// i is the current sequence counter
userMetaFactory.SetNameSequence(
	func(ctx context.Context, i int) (string, error) {
		return fmt.Sprintf("user_%d", i), nil
	},
),
```
The sequence counter is concurrent safe and increase by 1 each time factory's `Create` method called.

#### Lazy
For fields whose value is computed from other fields, use lazy attribute. Only Default/Sequence/Factory values are accessible in the struct.
```go
userMetaFactory.SetEmailLazy(
	func(ctx context.Context, i *model.User) (string, error) {
		return fmt.Sprintf("%s@carrier.go", i.Name), nil
	},
)
```
**> ent**

Ent is a little different because the struct is created after `Save`. And carrier call ent's `Set{Field}` method to set values.
So the input param here is not `*model.User`, but a temp containter struct created by carrier, hold all fields you can set.
```go
entUserMetaFactory.SetEmailLazy(
	func(ctx context.Context, i *factory.EntUserMutator) (string, error) {
		return fmt.Sprintf("%s@carrier.com", i.Name), nil
	},
)
```

#### Factory
If a field's value has related factory, use `relatedFactory.Create` method as param here. You can also set the function manually.
```go
// User struct has a Group field, type is Group
userMetaFactory.SetGroupFactory(groupFactory.Create)
```

#### AfterCreate
For struct factory, after create function is called after all lazy functions done. For ent factory, after create function is called next to ent's `Save` method.
```go
userMetaFactory.SetAfterCreateFunc(func(ctx context.Context, i *model.User) error {
	fmt.Printf("user: %d saved", i.Name)
	return nil
})
```

#### Post
Post functions will run once `AfterCreate` step done.
```go
// user MetaFactory
userMetaFactory.SetWelcomePostFunc(
	func(ctx context.Context, set bool, obj *model.User, i string) error {
		if set {
			message.SendTo(obj, i)
		}
		return nil
	},
)
// user Factory, send welcome message
userFactory.SetWelcomePost("welcome to carrier").Create(context.TODO())
// user Factory, no welcome message
userFactory.Create(context.TODO())
```

#### Trait
Trait is used to override some fields at once, activated by `With{Name}Trait` method.
```go
// override name
userMetaFactory.SetGopherTrait(factory.UserTrait().SetNameDefault("gopher"))
// user Factory
userFactory.WithGopherTrait().Create(context.TODO())
```
The `Trait` struct share same API with `MetaFactory` except `Set{Name}Trait` one, that means you can override 6 methods within a trait.
`Trait` only override methods you explicitly set, the exampe above will only override name field. So you can combine multiple traits together,
each change some parts of the struct. If multiple traits override same field, the last one will win:
```go
userMetaFactory.SetGopherTrait(factory.UserTrait().SetNameDefault("gopher")).
SetFooTrait(factory.UserTrait().SetNameDefault("foo"))
// user name is foo
userFactory.WithGopherTrait().WithFooTrait().Create(context.TODO())
```
#### Build
This is the final step for `MetaFactory` definition, call this method will return a `Factory` which you can use to create structs.

## Factory API
Factory API provide 3 types of method, `Set{Field}` to override some field, `Set{Field}Post` to call post function and `With{Name}Trait` to enable trait.
#### Set
Override field value. This method has the highest priority and will override your field method in `MetaFactory`.
```go
userFactory.SetName("foo").Create(context.TODO())
```
#### SetPost
Call post function defined in `MetaFactory` with param.
```go
// create a user with 3 friends
userFactory.SetFriendsPost(3).Create(context.TODO())
```
#### WithTrait
Enable a named trait. If you enable multi traits, and traits have overlapping, the latter one will override the former.
```go
userFactory.WithFooTrait().WithBarTrait().Create(context.TODO())
```
#### Create
Create pointer of struct.
#### CreateV
Create struct.
#### CreateBatch
Create slice of struct pointer.
#### CreateBatchV
Create slice of struct.

## Common Recipes

