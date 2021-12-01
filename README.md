# carrier - A Test Fixture Generator for Go

- **Statically Typed** - 100% statically typed using code generation.
- **Developer Friendly API** - factory_bot/factory_boy style API.
- **Ent Support** - [ent: An Entity Framework For Go](https://github.com/ent/ent)

A snippet how carrier works:

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
