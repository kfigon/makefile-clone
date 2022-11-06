package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSystem(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		b := newBuildSystem()
		b.addTask(newTask("step", nil, []action{"print foo"}))

		cmds, err := b.run("step")
		assert.NoError(t, err)
		assert.Equal(t, []action{"print foo"}, cmds)
	})

	t.Run("more tasks", func(t *testing.T) {
		b := newBuildSystem()
		b.addTask(newTask("stepA", nil, []action{"print foo"}))
		b.addTask(newTask("stepB", nil, []action{"print bar", "print bar again"}))
		b.addTask(newTask("run", []taskName{"stepA", "stepB"}, []action{"execute main"}))

		cmds, err := b.run("run")
		assert.NoError(t, err)
		assert.Equal(t, []action{"print foo", "print bar", "print bar again", "execute main"}, cmds)
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