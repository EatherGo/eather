package eather

import (
	"sync"

	"github.com/EatherGo/eather/types"
)

var (
	instance *Registry
	once     sync.Once
)

// RegistryInterface - interface for a registry
type RegistryInterface interface {
	Get(name string) types.Module
	GetCollection() RegistryCollection
	Add(object types.Module, name string)
	Contains(name string) bool
	Remove(name string)
}

// Registry struct - collection for registry
type Registry struct {
	collection RegistryCollection
}

// RegistryCollection map of all modules in registry
type RegistryCollection map[string]types.Module

// Get module from registry by name
func (r *Registry) Get(name string) types.Module {
	if val, ok := r.collection[name]; ok {
		return val
	}

	return nil
}

// GetCollection returns collection of all modules
func (r *Registry) GetCollection() RegistryCollection {
	return r.collection
}

// Add module to the registry object
func (r *Registry) Add(object types.Module, name string) {
	r.collection[name] = object
}

// Contains check if module is already in registry
func (r *Registry) Contains(name string) bool {
	_, ok := r.collection[name]

	return ok
}

// Remove module from collection
func (r *Registry) Remove(name string) {
	delete(r.collection, name)
}

// GetRegistry load registry collection
func GetRegistry() RegistryInterface {

	once.Do(func() {

		instance = &Registry{make(map[string]types.Module)}

	})

	return instance
}
