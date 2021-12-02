package integration

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Yiling-J/carrier/integration/carrier"
	"github.com/Yiling-J/carrier/integration/carrier/factory"
	"github.com/Yiling-J/carrier/integration/ent"
	"github.com/Yiling-J/carrier/integration/model"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func getStructFactory() *carrier.Factory {
	groupCatMetaFactory := carrier.GroupCategoryMetaFactory()
	groupCatFactory := groupCatMetaFactory.SetNameDefault("cat").Build()
	groupMetaFactory := carrier.GroupMetaFactory()
	_ = groupMetaFactory.SetNameSequence(
		func(ctx context.Context, i int) (string, error) { return fmt.Sprintf("group-%d", i), nil },
	).SetCategoryFactory(groupCatFactory.CreateV)
	groupFactory := groupMetaFactory.Build()
	userMetaFactory := carrier.UserMetaFactory()
	_ = userMetaFactory.SetNameSequence(
		func(ctx context.Context, i int) (string, error) { return fmt.Sprintf("user-%d", i), nil },
	)
	_ = userMetaFactory.SetEmailLazy(
		func(ctx context.Context, i *model.User) (string, error) {
			return fmt.Sprintf("%s@test", i.Name), nil
		},
	)
	_ = userMetaFactory.SetFooPostFunc(
		func(ctx context.Context, set bool, obj *model.User, i string) error {
			if set {
				obj.Name += i
			}
			return nil
		},
	)
	_ = userMetaFactory.
		SetDefaultTrait(factory.UserTrait().SetNameDefault("default_user")).
		SetSequenceTrait(factory.UserTrait().SetNameSequence(
			func(ctx context.Context, i int) (string, error) {
				return fmt.Sprintf("%d_user", i), nil
			},
		),
		).
		SetLazyTrait(factory.UserTrait().SetNameLazy(
			func(ctx context.Context, i *model.User) (string, error) { return "lazy_user", nil }),
		).
		SetFactoryTrait(factory.UserTrait().SetNameFactory(
			func(ctx context.Context) (string, error) { return "factory_user", nil }),
		).
		SetAnonymousTrait(factory.UserTrait().SetNameDefault("anonymous").SetGroupFactory(nil)).
		SetNilTrait(
			factory.UserTrait().SetNameFactory(nil).SetEmailLazy(nil).SetGroupSequence(nil).SetAfterCreateFunc(nil),
		).
		SetMixnameTrait(factory.UserTrait().SetNameDefault("mix_name")).
		SetMixemailTrait(factory.UserTrait().SetEmailDefault("mix_email").SetAfterCreateFunc(nil)).
		SetMixtitleTrait(factory.UserTrait().SetTitleDefault("mix_title")).
		SetAfterCreateFunc(func(ctx context.Context, i *model.User) error {
			i.Email = i.Email + ".com"
			return nil
		})
	_ = userMetaFactory.SetGroupFactory(groupFactory.Create)
	userFactory := userMetaFactory.Build()
	barFactory := carrier.BarMetaFactory().SetNameDefault("foo").Build()
	foodFactory := carrier.FoodMetaFactory().SetFooFactory(barFactory.CreateV).Build()
	factory := &carrier.Factory{}
	factory.SetGroupFactory(groupFactory)
	factory.SetUserFactory(userFactory)
	factory.SetBarFactory(barFactory)
	factory.SetFoodFactory(foodFactory)
	return factory
}

func getEntFactory() (*carrier.EntFactory, error) {
	entFactory := &carrier.EntFactory{}
	groupFactory := carrier.EntGroupMetaFactory().SetNameSequence(
		func(ctx context.Context, i int) (string, error) {
			return fmt.Sprintf("group%d", i), nil
		},
	).
		SetAfterCreateFunc(func(ctx context.Context, i *ent.Group) error {
			user, err := entFactory.UserFactory().SetGroupsPost(0).Create(ctx)
			if err != nil {
				return err
			}
			_, err = i.Update().AddUsers(user).Save(ctx)
			return err
		}).
		SetNouserTrait(factory.EntGroupTrait().SetAfterCreateFunc(nil)).
		Build()
	userFactory := carrier.EntUserMetaFactory().
		SetNameSequence(
			func(ctx context.Context, i int) (string, error) {
				return fmt.Sprintf("user-%d", i), nil
			},
		).SetAgeDefault(20).
		SetEmailLazy(func(ctx context.Context, i *factory.EntUserMutator) (string, error) {
			return fmt.Sprintf("%s@test.com", i.Name), nil
		}).
		SetGroupsPostFunc(func(ctx context.Context, set bool, obj *ent.User, i int) error {
			if !set {
				group, err := groupFactory.WithNouserTrait().Create(ctx)
				if err != nil {
					return err
				}
				_, err = obj.Update().AddGroups(group).Save(ctx)
				return err
			}
			if i == 0 {
				return nil
			}
			groups, err := groupFactory.WithNouserTrait().CreateBatch(ctx, i)
			if err != nil {
				return err
			}
			_, err = obj.Update().AddGroups(groups...).Save(ctx)
			return err
		}).
		SetAfterCreateFunc(nil).
		Build()
	carFactory := carrier.EntCarMetaFactory().
		SetModelDefault("Tesla").
		SetOwnerFactory(userFactory.Create).
		Build()

	client, err := ent.Open("sqlite3", ":memory:?_fk=1")
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}
	entFactory.SetUserFactory(userFactory).
		SetCarFactory(carFactory).
		SetGroupFactory(groupFactory).
		SetClient(client)
	return entFactory, nil
}

func TestBasicWithTraits(t *testing.T) {
	testCases := []struct {
		Name     string
		Type     int
		Expected string
	}{
		{"default test", factory.TypeDefault, "default_user"},
		{"sequence test", factory.TypeSequence, "1_user"},
		{"lazy test", factory.TypeLazy, "lazy_user"},
		{"factory test", factory.TypeFactory, "factory_user"},
		{"trait override", -1, "lazy_user"},
		{"set override", -2, "over_user"},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			f := getStructFactory()
			uf := f.UserFactory()
			var ub *factory.UserBuilder
			switch tc.Type {
			case factory.TypeDefault:
				ub = uf.WithDefaultTrait()
			case factory.TypeSequence:
				ub = uf.WithSequenceTrait()
			case factory.TypeLazy:
				ub = uf.WithLazyTrait()
			case factory.TypeFactory:
				ub = uf.WithFactoryTrait()
			case -1:
				ub = uf.WithDefaultTrait().WithFactoryTrait().WithLazyTrait()
			case -2:
				ub = uf.WithSequenceTrait().SetName("over_user")
			default:
				t.FailNow()
			}
			user, err := ub.Create(context.TODO())
			require.Nil(t, err)
			require.Equal(t, tc.Expected, user.Name)
			require.Equal(t, fmt.Sprintf("%s@test.com", tc.Expected), user.Email)
		})
	}
}

func TestNilFunc(t *testing.T) {
	f := getStructFactory()
	user, err := f.UserFactory().WithNilTrait().Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, user.Name, "")
	require.Equal(t, user.Email, "")
	require.True(t, user.Group == nil)
}

func TestSequenceCounter(t *testing.T) {
	f := getStructFactory()
	var wg sync.WaitGroup
	mu := &sync.Mutex{}
	var names []string
	var groups []string
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			user, err := f.UserFactory().Create(context.TODO())
			require.Nil(t, err)
			mu.Lock()
			names = append(names, user.Name)
			groups = append(groups, user.Group.Name)
			mu.Unlock()

		}()
	}
	wg.Wait()
	var expected []string
	var expectedGroups []string
	for i := 1; i <= 20; i++ {
		expected = append(expected, fmt.Sprintf("user-%d", i))
		expectedGroups = append(expectedGroups, fmt.Sprintf("group-%d", i))
	}
	require.ElementsMatch(t, names, expected)
	require.ElementsMatch(t, groups, expectedGroups)
}

func TestPost(t *testing.T) {
	f := getStructFactory()
	user, err := f.UserFactory().SetFooPost("foo").Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "user-1foo", user.Name)

	user, err = f.UserFactory().Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "user-2", user.Name)
}

func TestTrait(t *testing.T) {
	f := getStructFactory()
	user, err := f.UserFactory().Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "user-1", user.Name)
	require.Equal(t, "group-1", user.Group.Name)
	require.Equal(t, "cat", user.Group.Category.Name)

	user, err = f.UserFactory().WithAnonymousTrait().Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "anonymous", user.Name)
	require.Nil(t, user.Group)
}

func TestCreateV(t *testing.T) {
	f := getStructFactory()
	user, err := f.UserFactory().CreateV(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "user-1", user.Name)
	require.Equal(t, "group-1", user.Group.Name)
}

func TestCreateBatch(t *testing.T) {
	f := getStructFactory()
	users, err := f.UserFactory().CreateBatch(context.TODO(), 3)
	require.Nil(t, err)
	require.Equal(t, 3, len(users))
	var names []string
	for _, user := range users {
		names = append(names, user.Name)
	}
	require.Equal(t, []string{"user-1", "user-2", "user-3"}, names)
}

func TestCreateEntBatch(t *testing.T) {
	f, err := getEntFactory()
	require.Nil(t, err)
	users, err := f.UserFactory().CreateBatch(context.TODO(), 5)
	require.Nil(t, err)
	require.Equal(t, 5, len(users))
	var names []string
	users, err = f.Client().User.Query().Order(ent.Asc("id")).All(context.TODO())
	require.Nil(t, err)
	for _, user := range users {
		names = append(names, user.Name)
	}
	require.Equal(t, []string{"user-1", "user-2", "user-3", "user-4", "user-5"}, names)
}

func TestCreateBatchV(t *testing.T) {
	f := getStructFactory()
	users, err := f.UserFactory().CreateBatchV(context.TODO(), 3)
	require.Nil(t, err)
	require.Equal(t, 3, len(users))
	var names []string
	for _, user := range users {
		names = append(names, user.Name)
	}
	require.Equal(t, []string{"user-1", "user-2", "user-3"}, names)
}

func TestMixTrait(t *testing.T) {
	f := getStructFactory()
	user, err := f.UserFactory().
		WithMixnameTrait().
		WithMixemailTrait().
		WithMixtitleTrait().
		Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "mix_name", user.Name)
	require.Equal(t, "mix_email", user.Email)
	require.Equal(t, "mix_title", user.Title)
}

func TestAlias(t *testing.T) {
	f := getStructFactory()
	foo, err := f.BarFactory().SetName("foo").Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "foo", foo.Name)
}

func TestEmbed(t *testing.T) {
	f := getStructFactory()
	food, err := f.FoodFactory().SetCategory("apple").Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "foo", food.Name)
}

func TestEntBasic(t *testing.T) {
	ctx := context.TODO()
	f, err := getEntFactory()
	require.Nil(t, err)
	user, err := f.UserFactory().Create(ctx)
	require.Nil(t, err)
	require.Equal(t, user.Name, "user-1")
	require.Equal(t, user.ID, 1)
	require.Equal(t, *user.Email, "user-1@test.com")
	// sub factory
	car, err := f.CarFactory().SetRegisteredAt(time.Now()).Create(ctx)
	require.Nil(t, err)
	require.Equal(t, "Tesla", car.Model)
	require.Equal(t, car.ID, 1)
	owner, err := car.QueryOwner().First(ctx)
	require.Nil(t, err)
	require.Equal(t, "user-2", owner.Name)
	require.Equal(t, owner.ID, 2)
	// post default
	groups, err := user.QueryGroups().All(ctx)
	require.Nil(t, err)
	require.Equal(t, 1, len(groups))
	for _, g := range groups {
		c, err := g.QueryUsers().Count(ctx)
		require.Nil(t, err)
		require.Equal(t, 1, c)
	}
	// post set
	user2, err := f.UserFactory().SetGroupsPost(3).Create(ctx)
	require.Nil(t, err)
	groups, err = user2.QueryGroups().All(ctx)
	require.Nil(t, err)
	require.Equal(t, 3, len(groups))
	for _, g := range groups {
		c, err := g.QueryUsers().Count(ctx)
		require.Nil(t, err)
		require.Equal(t, 1, c)
	}
	// group default user
	group, err := f.GroupFactory().Create(ctx)
	require.Nil(t, err)
	c, err := group.QueryUsers().Count(ctx)
	require.Nil(t, err)
	require.Equal(t, 1, c)
	// total group count 6, user:1 + car.owner:1 + user2:3 + group:1
	total, err := f.Client().Group.Query().Count(ctx)
	require.Nil(t, err)
	require.Equal(t, 6, total)
}

func TestEntClient(t *testing.T) {
	carFactory := carrier.EntCarMetaFactory().
		SetModelDefault("Tesla").
		Build()

	client, err := ent.Open("sqlite3", ":memory:?_fk=1")
	require.Nil(t, err)
	err = client.Schema.Create(context.Background())
	require.Nil(t, err)
	car, err := carFactory.Client(client).SetRegisteredAt(time.Now()).Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "Tesla", car.Model)
}
