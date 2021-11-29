package schema

import (
	"github.com/Yiling-J/carrier"
	"github.com/Yiling-J/carrier/integration/model"
)

var (
	Schemas = []carrier.Schema{
		&carrier.StructSchema{
			To: model.Group{},
		},
		&carrier.StructSchema{
			To:     model.User{},
			Traits: []string{"default", "lazy", "sequence", "factory"},
		},
	}
)
