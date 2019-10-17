package lib

import (
	"fmt"
	"sync"
)

var (
	eventInstance *Events
	onceEvent     sync.Once
)

type events interface {
	Emmit(name string)
	Add(name string)
	Remove(name string)
}

type event struct {
	Name  string
	Fires map[string]func()
}

// Events struct - collection of events
type Events struct {
	Collection map[string]event
}

// Add event to the collection
func (r *Events) Add(name string, f func(), call string) {
	fmt.Println("Adding event " + name + " to call " + call)
	if val, ok := r.Collection[name]; ok {
		val.Fires[call] = f
		r.Collection[name] = val
	} else {
		fires := make(map[string]func())
		fires[call] = f
		e := event{Name: name, Fires: fires}
		r.Collection[name] = e
	}
}

// Emmit the event from the collection
func (r *Events) Emmit(name string) {
	if val, ok := r.Collection[name]; ok {
		for _, fire := range val.Fires {
			go fire()
		}
	}
}

// Remove the event from the collection
func (r *Events) Remove(name string) {
	delete(r.Collection, name)
}

// GetEvents - get collection of all registered events
func GetEvents() *Events {

	onceEvent.Do(func() {

		eventInstance = &Events{make(map[string]event)}

	})

	return eventInstance
}
