package schema

import (
	"github.com/Yiling-J/carrier"
	"github.com/Yiling-J/carrier/integration/ent"
	"github.com/Yiling-J/carrier/integration/model"
)

var (
	Schemas = []carrier.Schema{
		&carrier.StructSchema{
			To: model.Group{},
		},
		&carrier.StructSchema{
			To:     model.User{},
			Traits: []string{"default", "lazy", "sequence", "factory", "anonymous"},
			Posts:  []carrier.PostField{{Name: "foo", Input: ""}},
		},
		&carrier.EntSchema{
			To: &ent.UserCreate{},
		},
		&carrier.EntSchema{
			To: &ent.CarCreate{},
		},
		&carrier.EntSchema{
			To: &ent.GroupCreate{},
		},
	}
)
