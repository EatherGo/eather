package eather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var r = GetRegistry()

type mod struct{}

func TestItShouldAddModuleToRegistry(t *testing.T) {
	assert.Equal(t, 0, len(r.GetCollection()), "Registry collection should be empty")

	r.Add(mod{}, "test")

	assert.Equal(t, 1, len(r.GetCollection()), "Registry collection should have one module in registry")

	assert.Equal(t, mod{}, r.Get("test"), "Get function should return added module")

	assert.Equal(t, true, r.Contains("test"), "Contains should return true that test exists in collection")
}

func TestModulesWithSameNameCanBeInCollectionOnlyOnce(t *testing.T) {
	r.Add(mod{}, "test")
	r.Add(mod{}, "test")

	assert.Equal(t, 1, len(r.GetCollection()), "Registry collection should contains test only ones")
}

func TestModulesCanBeRemovedFromRegistry(t *testing.T) {
	r.Add(mod{}, "test")

	assert.Equal(t, 1, len(r.GetCollection()), "Registry collection should contains test")

	r.Remove("test")

	assert.Equal(t, 0, len(r.GetCollection()), "Registry collection should not contains test")
}
