# carrier - A Test Fixture Generator for Go
![example workflow](https://github.com/Yiling-J/carrier/actions/workflows/go.yml/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/Yiling-J/carrier?style=flat-square)

- **Statically Typed** - 100% statically typed using code generation.
- **Developer Friendly API** - factory_bot/factory_boy style API
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
- *Run generator* 🎉
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
			To: model.User{},,
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

## Build Factory and Generate Fixtures
To construct a real factory for testing:

**Create MetaFactory struct**
```go
userMetaFactory := carrier.UserMetaFactory()
```
**Build factory from meta factory**
```go
userFactory := userMetaFactory.SetNameDefault("carrier").Build()
```
MetaFactory provide several methods to help you initial field values automatically, [MetaFactory API Reference](#metafactory-api)

**Create fixtures**

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
## MetaFactory API
## Factory API
## Common Recipes

