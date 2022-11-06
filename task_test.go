package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleBuildSystem(t *testing.T) {
	b := newBuildSystem()
	b.addTask(newTask("step", nil, []action{"print foo"}))

	cmds, err := b.run()
	assert.NoError(t, err)
	assert.Equal(t, []action{"print foo"}, cmds)
}