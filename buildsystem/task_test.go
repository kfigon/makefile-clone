package buildsystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSystem(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		b := NewBuildSystem()
		b.AddTask(NewTask("step", nil, []Action{"print foo"}))

		cmds, err := b.Run("step")
		assert.NoError(t, err)
		assert.Equal(t, []Action{"print foo"}, cmds)
	})

	t.Run("more tasks", func(t *testing.T) {
		b := NewBuildSystem()
		b.AddTask(NewTask("clean", nil, []Action{"cleaning"}))
		b.AddTask(NewTask("stepA", nil, []Action{"print foo"}))
		b.AddTask(NewTask("stepB", nil, []Action{"print bar", "print bar again"}))
		b.AddTask(NewTask("run", []TaskName{"stepA", "stepB"}, []Action{"execute main"}))

		cmds, err := b.Run("run")
		assert.NoError(t, err)
		assert.Equal(t, []Action{"print foo", "print bar", "print bar again", "execute main"}, cmds)
	})

	t.Run("more tasks but run only intermediate cmd", func(t *testing.T) {
		b := NewBuildSystem()
		b.AddTask(NewTask("clean", nil, []Action{"cleaning"}))
		b.AddTask(NewTask("stepA", nil, []Action{"print foo"}))
		b.AddTask(NewTask("stepB", nil, []Action{"print bar", "print bar again"}))
		b.AddTask(NewTask("run", []TaskName{"stepA", "stepB"}, []Action{"execute main"}))

		cmds, err := b.Run("stepA")
		assert.NoError(t, err)
		assert.Equal(t, []Action{"print foo"}, cmds)
	})
}

func TestInvalidCases(t *testing.T) {
	t.Run("empty system", func(t *testing.T) {
		b := NewBuildSystem()
		cmds, err := b.Run("foo")
		assert.NoError(t, err)
		assert.Equal(t, []Action{}, cmds)
	})

	t.Run("circular dependency", func(t *testing.T) {
		b := NewBuildSystem()
		b.AddTask(NewTask("clean", nil, []Action{"cleaning"}))
		b.AddTask(NewTask("stepA", []TaskName{"stepB"}, []Action{"print foo"}))
		b.AddTask(NewTask("stepB", []TaskName{"stepA"}, []Action{"print bar", "print bar again"}))
		b.AddTask(NewTask("run", []TaskName{"stepA"}, []Action{"execute main"}))

		_, err := b.Run("run")
		assert.Error(t, err)
	})
}
