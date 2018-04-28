package lite

import (
	"fmt"
	"sync"
)

// Module represents single module with lite API (it means that all its routes
// can be generated only once at startup).
type Module interface {
	// Register should add Controller to module resources.
	Register(alias string, controller Controller) error
	// Unregister should remove the Controller from module resources.
	Unregister(alias string) error
	// Controllers should call the provided func sequentially for each available
	// resource. If func returns false the "for" loop should be interrupted.
	Controllers(func(alias string, controller Controller) bool)
}

// BaseModule contains a basic set of logic and provides basic operations on
// resources (like "Register", "Unregister" etc).
type BaseModule struct {
	sync.RWMutex
	resources map[string]Controller
}

// NewBaseModule is a constructor func for BaseModule.
func NewBaseModule() *BaseModule {
	return &BaseModule{resources: make(map[string]Controller)}
}

// Register makes resource available by provided alias.
func (m *BaseModule) Register(name string, resource Controller) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.resources[name]; ok {
		return fmt.Errorf("already registered: %q", name)
	}
	m.resources[name] = resource
	return nil
}

// Unregister removes resource from the list by alias.
func (m *BaseModule) Unregister(name string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.resources[name]; !ok {
		return fmt.Errorf("not registered: %q", name)
	}
	delete(m.resources, name)
	return nil
}

// Controllers calls the provided func sequentially for each available resource.
// If func returns false the "for" loop will be stopped.
func (m *BaseModule) Controllers(f func(string, Controller) bool) {
	m.Lock()
	defer m.Unlock()

	for alias, resource := range m.resources {
		if !f(alias, resource) {
			break
		}
	}
}
