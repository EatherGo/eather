package eather

import (
	"sync"
)

var (
	instance *Registry
	once     sync.Once
)

// RegistryInterface - interface for a registry
type RegistryInterface interface {
	Get(name string) RegistryModuleInterface
	GetCollection() RegistryCollection
	Add(object Module, name string)
	Contains(name string) bool
	Remove(name string)
}

// Registry struct - collection for registry
type Registry struct {
	collection RegistryCollection
}

// RegistryModuleInterface interface for registryModule
type RegistryModuleInterface interface {
	GetCallable() Callable
}

// RegistryModule structure
type RegistryModule struct {
	Module Module
}

// GetCallable will return nil or callable interface
func (rm RegistryModule) GetCallable() Callable {
	if callable, isCallable := rm.Module.(Callable); isCallable {
		return callable
	}

	return nil
}

// RegistryCollection map of all modules in registry
type RegistryCollection map[string]Module

// Get module from registry by name
func (r *Registry) Get(name string) RegistryModuleInterface {
	if val, ok := r.collection[name]; ok {
		return RegistryModule{Module: val}
	}

	return nil
}

// GetCollection returns collection of all modules
func (r *Registry) GetCollection() RegistryCollection {
	return r.collection
}

// Add module to the registry object
func (r *Registry) Add(object Module, name string) {
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

		instance = &Registry{make(map[string]Module)}

	})

	return instance
}
