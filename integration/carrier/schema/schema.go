package schema

import (
	"github.com/Yiling-J/carrier"
	"github.com/Yiling-J/carrier/integration/ent"
	"github.com/Yiling-J/carrier/integration/model"
)

var (
	Schemas = []carrier.Schema{
		&carrier.StructSchema{
			To: &model.GroupCategory{},
		},
		&carrier.StructSchema{
			To: model.Group{},
		},
		&carrier.StructSchema{
			To:    model.Foo{},
			Alias: "Bar",
		},
		&carrier.StructSchema{
			To: model.User{},
			Traits: []string{
				"default", "lazy", "sequence", "factory", "anonymous", "nil",
				"mixname", "mixemail", "mixtitle",
			},
			Posts: []carrier.PostField{{Name: "foo", Input: ""}},
		},
		&carrier.EntSchema{
			To:    &ent.UserCreate{},
			Posts: []carrier.PostField{{Name: "groups", Input: 2}},
		},
		&carrier.EntSchema{
			To: &ent.CarCreate{},
		},
		&carrier.EntSchema{
			To:     &ent.GroupCreate{},
			Traits: []string{"nouser"},
		},
	}
)
