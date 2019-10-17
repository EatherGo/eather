package lib

import (
	"project/lib/interfaces"
	"sync"
)

var (
	instance *Registry
	once     sync.Once
)

// RegistryInterface - interface for a registry
type RegistryInterface interface {
	Get(name string) interfaces.Module
	Add(object interfaces.Module, name string)
	Contains(name string) bool
	Remove(name string)
}

// Registry struct - collection for registry
type Registry struct {
	Collection map[string]interfaces.Module
}

// Get module from registry by name
func (r *Registry) Get(name string) interfaces.Module {
	if val, ok := r.Collection[name]; ok {
		return val
	}

	return nil
}

// Add module to the registry object
func (r *Registry) Add(object interfaces.Module, name string) {
	r.Collection[name] = object
}

// Contains check if module is already in registry
func (r *Registry) Contains(name string) bool {
	_, ok := r.Collection[name]

	return ok
}

// Remove module from collection
func (r *Registry) Remove(name string) {
	delete(r.Collection, name)
}

// GetRegistry load registry collection
func GetRegistry() *Registry {

	once.Do(func() {

		instance = &Registry{make(map[string]interfaces.Module)}

	})

	return instance
}
