package integration

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/Yiling-J/carrier/integration/carrier"
	"github.com/Yiling-J/carrier/integration/carrier/factory"
	"github.com/Yiling-J/carrier/integration/model"
	"github.com/stretchr/testify/require"
)

func getStructFactory() *carrier.Factory {
	groupMetaFactory := carrier.GroupMetaFactory()
	_ = groupMetaFactory.SetNameSequence(
		func(ctx context.Context, i int) (string, error) { return fmt.Sprintf("group-%d", i), nil },
	)
	groupFactory := groupMetaFactory.Build()
	userMetaFactory := carrier.UserMetaFactory()
	_ = userMetaFactory.SetNameSequence(
		func(ctx context.Context, i int) (string, error) { return fmt.Sprintf("user-%d", i), nil },
	)
	_ = userMetaFactory.SetEmailLazy(
		func(ctx context.Context, i *model.User) (string, error) {
			return fmt.Sprintf("%s@test.com", i.Name), nil
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
		SetAnonymousTrait(factory.UserTrait().SetNameDefault("anonymous").SetGroupFactory(nil))
	_ = userMetaFactory.SetGroupFactory(groupFactory.Create)
	userFactory := userMetaFactory.Build()
	factory := &carrier.Factory{}
	factory.SetGroupFactory(groupFactory)
	factory.SetUserFactory(userFactory)
	return factory
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

	user, err = f.UserFactory().WithAnonymousTrait().Create(context.TODO())
	require.Nil(t, err)
	require.Equal(t, "anonymous", user.Name)
	require.Nil(t, user.Group)
}
