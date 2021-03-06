package main

import (
	"testing"

	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCheckArgs(t *testing.T) {
	assert := assert.New(t)
	event := types.FixtureEvent("entity1", "check1")
	assert.NoError(checkArgs(event))
}

func TestExecuteMutator(t *testing.T) {
	assert := assert.New(t)
	event := types.FixtureEvent("entity1", "registration")
	event.Check.Status = 1
	ev, err := executeMutator(event)
	assert.NoError(err)
	assert.EqualValues(ev.Check.Status, 0)
}
