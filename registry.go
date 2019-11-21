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
	GetCronable() Cronable
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

// GetCronable will return nil or cronable interface
func (rm RegistryModule) GetCronable() Cronable {
	if cronable, isCronable := rm.Module.(Cronable); isCronable {
		return cronable
	}

	return nil
}

// RegistryCollection map of all modules in registry
type RegistryCollection map[string]RegistryModule

// Get module from registry by name
func (r *Registry) Get(name string) RegistryModuleInterface {
	if reg, ok := r.collection[name]; ok {
		return reg
	}

	return nil
}

// GetCollection returns collection of all modules
func (r *Registry) GetCollection() RegistryCollection {
	return r.collection
}

// Add module to the registry object
func (r *Registry) Add(object Module, name string) {
	r.collection[name] = RegistryModule{Module: object}

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

		instance = &Registry{make(map[string]RegistryModule)}

	})

	return instance
}
