package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var e = GetEvents()

func TestEventsCanBeAdded(t *testing.T) {
	assert.Equal(t, 0, len(e.GetCollection()), "Event collection should be empty")

	e.Add("test", func(data ...interface{}) {}, "testing")

	assert.Equal(t, 1, len(e.GetCollection()), "Event collection should have only 1 event")
}

func TestEventsCanBeAddedOnlyOnce(t *testing.T) {

	e.Add("test", func(data ...interface{}) {}, "testing")
	e.Add("test", func(data ...interface{}) {}, "testing")

	assert.Equal(t, 1, len(e.GetCollection()), "Event collection should have only 1 event")
}

func TestEventsCanBeRemoved(t *testing.T) {
	e.Add("test", func(data ...interface{}) {}, "testing")

	e.Remove("test")

	assert.Equal(t, 0, len(e.GetCollection()), "Event collection be empty")
}
