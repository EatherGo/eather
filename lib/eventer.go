package lib

import (
	"eather/lib/types"
	"fmt"
	"sync"
)

var (
	eventInstance *Events
	onceEvent     sync.Once
)

// EventsInterface interface of events
type EventsInterface interface {
	Emmit(name string, data ...interface{})
	Add(name string, f types.EventFunc, call string)
	Remove(name string)
	GetCollection() EventCollection
}

type event struct {
	Name  string
	Fires FiresCollection
}

// Events struct - collection of events
type Events struct {
	Collection EventCollection
}

// EventCollection is definition of events collection
type EventCollection map[string]event

// FiresCollection is definition of fires collection
type FiresCollection map[string]types.EventFunc

// Add event to the collection
func (r *Events) Add(name string, f types.EventFunc, call string) {
	fmt.Println("Adding event " + name + " to call " + call)
	if val, ok := r.Collection[name]; ok {
		val.Fires[call] = f
		r.Collection[name] = val
	} else {
		fires := make(FiresCollection)
		fires[call] = f
		e := event{Name: name, Fires: fires}
		r.Collection[name] = e
	}
}

// Emmit the event from the collection
// data will be passed to the event func
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

// GetCollection will return collection of events
func (r *Events) GetCollection() EventCollection {
	return r.Collection
}

// GetEvents - get collection of all registered events
func GetEvents() EventsInterface {

	onceEvent.Do(func() {
		eventInstance = &Events{make(map[string]event)}
	})

	return eventInstance
}
