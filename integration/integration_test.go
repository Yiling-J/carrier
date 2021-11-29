package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/Yiling-J/carrier/integration/carrier"
	"github.com/Yiling-J/carrier/integration/carrier/factory"
	"github.com/Yiling-J/carrier/integration/model"
	"github.com/stretchr/testify/require"
)

func getFactory() *carrier.Factory {
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
		)
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
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			f := getFactory()
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
			default:
				t.FailNow()
			}
			user, err := ub.Create(context.TODO())
			require.Nil(t, err)
			require.Equal(t, tc.Expected, user.Name)
		})
	}
}
