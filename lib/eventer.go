package lib

import (
	"fmt"
	"project/lib/interfaces"
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
	Fires map[string]interfaces.EventFunc
}

// Events struct - collection of events
type Events struct {
	Collection map[string]event
}

// Add event to the collection
func (r *Events) Add(name string, f interfaces.EventFunc, call string) {
	fmt.Println("Adding event " + name + " to call " + call)
	if val, ok := r.Collection[name]; ok {
		val.Fires[call] = f
		r.Collection[name] = val
	} else {
		fires := make(map[string]interfaces.EventFunc)
		fires[call] = f
		e := event{Name: name, Fires: fires}
		r.Collection[name] = e
	}
}

// Emmit the event from the collection
func (r *Events) Emmit(name string, data ...interface{}) {
	if val, ok := r.Collection[name]; ok {
		for _, fire := range val.Fires {
			go fire(data)
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
