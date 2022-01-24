// Code generated by carrier, DO NOT EDIT.
package carrier

import (
	"github.com/Yiling-J/carrier/integration/carrierii/factory"

	"github.com/Yiling-J/carrier/integration/ten"
)

// Factory is struct factory wrapper
type Factory struct {
}

// NewFactory return a new struct factory wrapper
func NewFactory() *Factory {
	return &Factory{}
}

// EntFactory is ent factory wrapper
type EntFactory struct {
	userFactory *factory.EntUserFactory

	carFactory *factory.EntCarFactory

	groupFactory *factory.EntGroupFactory

	client *ten.Client
}

// NewEntFactory return a new ent factory wrapper
func NewEntFactory(client *ten.Client) *EntFactory {
	return &EntFactory{client: client}
}

// Client return wrappper's ent client
func (f *EntFactory) Client() *ten.Client {
	return f.client
}

// EntUserMetaFactory return a new meta factory with given ent client
func EntUserMetaFactory() *factory.EntUserMetaFactory {
	return &factory.EntUserMetaFactory{}
}

// SetUserFactory set a factory in wrapper
func (f *EntFactory) SetUserFactory(c *factory.EntUserFactory) *EntFactory {
	f.userFactory = c.Client(f.client)
	return f
}

// UserFactory return the EntUserFactory in wrapper
func (f *EntFactory) UserFactory() *factory.EntUserFactory {
	return f.userFactory
}

// EntCarMetaFactory return a new meta factory with given ent client
func EntCarMetaFactory() *factory.EntCarMetaFactory {
	return &factory.EntCarMetaFactory{}
}

// SetCarFactory set a factory in wrapper
func (f *EntFactory) SetCarFactory(c *factory.EntCarFactory) *EntFactory {
	f.carFactory = c.Client(f.client)
	return f
}

// CarFactory return the EntCarFactory in wrapper
func (f *EntFactory) CarFactory() *factory.EntCarFactory {
	return f.carFactory
}

// EntGroupMetaFactory return a new meta factory with given ent client
func EntGroupMetaFactory() *factory.EntGroupMetaFactory {
	return &factory.EntGroupMetaFactory{}
}

// SetGroupFactory set a factory in wrapper
func (f *EntFactory) SetGroupFactory(c *factory.EntGroupFactory) *EntFactory {
	f.groupFactory = c.Client(f.client)
	return f
}

// GroupFactory return the EntGroupFactory in wrapper
func (f *EntFactory) GroupFactory() *factory.EntGroupFactory {
	return f.groupFactory
}
