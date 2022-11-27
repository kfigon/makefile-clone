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
		b.addTask(newTask("clean", nil, []action{"cleaning"}))
		b.addTask(newTask("stepA", nil, []action{"print foo"}))
		b.addTask(newTask("stepB", nil, []action{"print bar", "print bar again"}))
		b.addTask(newTask("run", []taskName{"stepA", "stepB"}, []action{"execute main"}))

		cmds, err := b.run("run")
		assert.NoError(t, err)
		assert.Equal(t, []action{"print foo", "print bar", "print bar again", "execute main"}, cmds)
	})

	t.Run("more tasks but run only intermediate cmd", func(t *testing.T) {
		b := newBuildSystem()
		b.addTask(newTask("clean", nil, []action{"cleaning"}))
		b.addTask(newTask("stepA", nil, []action{"print foo"}))
		b.addTask(newTask("stepB", nil, []action{"print bar", "print bar again"}))
		b.addTask(newTask("run", []taskName{"stepA", "stepB"}, []action{"execute main"}))

		cmds, err := b.run("stepA")
		assert.NoError(t, err)
		assert.Equal(t, []action{"print foo"}, cmds)
	})
}

func TestInvalidCases(t *testing.T) {
	t.Run("empty system", func(t *testing.T) {
		b := newBuildSystem()
		cmds, err := b.run("foo")
		assert.NoError(t, err)
		assert.Equal(t, []action{}, cmds)
	})

	t.Run("circular dependency", func(t *testing.T) {
		b := newBuildSystem()
		b.addTask(newTask("clean", nil, []action{"cleaning"}))
		b.addTask(newTask("stepA", []taskName{"stepB"}, []action{"print foo"}))
		b.addTask(newTask("stepB", []taskName{"stepA"}, []action{"print bar", "print bar again"}))
		b.addTask(newTask("run", []taskName{"stepA"}, []action{"execute main"}))

		_, err := b.run("run")
		assert.Error(t, err)
	})
}