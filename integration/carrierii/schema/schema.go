package schema

import (
	"github.com/Yiling-J/carrier"
	"github.com/Yiling-J/carrier/integration/ten"
)

var (
	Schemas = []carrier.Schema{
		&carrier.EntSchema{
			To:    &ten.UserCreate{},
			Posts: []carrier.PostField{{Name: "groups", Input: 2}},
		},
		&carrier.EntSchema{
			To: &ten.CarCreate{},
		},
		&carrier.EntSchema{
			To:     &ten.GroupCreate{},
			Traits: []string{"nouser"},
		},
	}
)
