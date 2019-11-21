package eather

import (
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
	Add(name string, f EventFunc, call string)
	Remove(name string)
	GetCollection() EventCollection
}

// Events struct - collection of events
type Events struct {
	Collection EventCollection
}

// Event structure
type Event struct {
	Name  string `json:"name"`
	Fires []Fire `json:"-"`
}

// EventCollection is definition of events collection
type EventCollection map[string]Event

// EventFunc type of events func
type EventFunc func(data ...interface{})

// Fire struct of Fires
type Fire struct {
	Call string
	Func EventFunc
}

// Add event to the collection
func (r *Events) Add(name string, f EventFunc, call string) {
	fmt.Println("Adding event " + name + " to call " + call)
	if val, ok := r.Collection[name]; ok {
		val.Fires = append(val.Fires, Fire{Call: call, Func: f})
		// val.Fires[call] = f
		// r.Collection[name] = val
	} else {
		// fires := make(Fire)
		fire := Fire{Call: call, Func: f}
		e := Event{Name: name, Fires: []Fire{fire}}
		r.Collection[name] = e
	}
}

// Emmit the event from the collection
// data will be passed to the event func
func (r *Events) Emmit(name string, data ...interface{}) {
	if val, ok := r.Collection[name]; ok {
		for _, fire := range val.Fires {
			go fire.Func(data)
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
		eventInstance = &Events{make(map[string]Event)}
	})

	return eventInstance
}
