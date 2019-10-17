package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var e = GetEvents()

func TestEventsCanBeAdded(t *testing.T) {
	assert.Equal(t, 0, len(e.Collection), "Event collection should be empty")

	e.Add("test", func() {}, "testing")

	assert.Equal(t, 1, len(e.Collection), "Event collection should have only 1 event")
}

func TestEventsCanBeAddedOnlyOnce(t *testing.T) {

	e.Add("test", func() {}, "testing")
	e.Add("test", func() {}, "testing")

	assert.Equal(t, 1, len(e.Collection), "Event collection should have only 1 event")
}

func TestEventsCanBeRemoved(t *testing.T) {
	e.Add("test", func() {}, "testing")

	e.Remove("test")

	assert.Equal(t, 0, len(e.Collection), "Event collection be empty")
}
