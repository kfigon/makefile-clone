package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSystem(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		b := newBuildSystem()
		b.addTask(newTask("step", nil, []action{"print foo"}))

		cmds, err := b.run()
		assert.NoError(t, err)
		assert.Equal(t, []action{"print foo"}, cmds)
	})
}

func TestInvalidCases(t *testing.T) {
	t.Run("empty system", func(t *testing.T) {
		t.Fatal("todo")
	})

	t.Run("dependency not provided", func(t *testing.T) {
		t.Fatal("todo")
	})

	t.Run("circular dependency", func(t *testing.T) {
		t.Fatal("todo")
	})
}