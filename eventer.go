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
	Add(eventName string, f EventFunc, call string, name string)
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
	Fires []Fire `json:"fires"`
}

// EventCollection is definition of events collection
type EventCollection map[string]Event

// EventFunc type of events func
type EventFunc func(data ...interface{})

// Fire struct of Fires
type Fire struct {
	Call string    `json:"call"`
	Name string    `json:"name"`
	Func EventFunc `json:"-"`
}

// Add event to the collection
func (r *Events) Add(eventName string, f EventFunc, call string, name string) {
	fmt.Println("Adding event " + eventName + " to call " + call + " with name " + name)

	fire := Fire{Call: call, Func: f, Name: name}

	if e, ok := r.Collection[eventName]; ok {
		e.Fires = append(e.Fires, fire)
		r.Collection[eventName] = e
	} else {
		e := Event{Name: eventName, Fires: []Fire{fire}}
		r.Collection[eventName] = e
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
