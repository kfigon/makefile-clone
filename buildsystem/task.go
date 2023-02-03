package buildsystem

import (
	"fmt"
	"strings"
)

type set map[string]struct{}

func (s set) add(v string) {
	s[v] = struct{}{}
}
func (s set) present(v string) bool {
	_, ok := s[v]
	return ok
}

type Action string
type TaskName string

type Task struct {
	name        TaskName
	depedencies []TaskName
	actions     []Action
}

func NewTask(name TaskName, deps []TaskName, actions []Action) Task {
	return Task{
		name:        name,
		depedencies: deps,
		actions:     actions,
	}
}

type BuildSystem struct {
	tasks map[TaskName]Task
}

func NewBuildSystem() *BuildSystem {
	return &BuildSystem{
		tasks: map[TaskName]Task{},
	}
}

func (b *BuildSystem) AddTask(t Task) {
	b.tasks[t.name] = t
}

func (b *BuildSystem) Run(t TaskName) ([]Action, error) {
	dependencyGraph := b.buildGraph()
	if cycle := hasCycles(t, dependencyGraph); len(cycle) != 0 {
		var vals []string
		for _, v := range cycle {
			vals = append(vals, string(v))
		}
		return nil, fmt.Errorf("Cycle detected %s", strings.Join(vals, "->"))
	}

	top := topology(t, dependencyGraph)
	out := []Action{}
	for _, v := range top {
		task := b.tasks[v]
		out = append(out, task.actions...)
	}
	return out, nil
}

type graph map[TaskName][]TaskName

func (b *BuildSystem) buildGraph() map[TaskName][]TaskName {
	out := graph{}
	for name, t := range b.tasks {
		tasks := out[name]
		tasks = append(tasks, t.depedencies...)
		out[name] = tasks
	}
	return out
}

func hasCycles(start TaskName, g graph) []TaskName {
	visited := set{}
	pathTo := map[TaskName]TaskName{}
	out := []TaskName{}
	onStack := set{}

	var dfs func(TaskName)
	dfs = func(t TaskName) {
		if visited.present(string(t)) {
			return
		}
		onStack.add(string(t))
		visited.add(string(t))
		for _, children := range g[t] {
			if len(out) != 0 {
				return
			}
			if !visited.present(string(children)) {
				pathTo[children] = t
				dfs(children)
			} else if onStack.present(string(children)) {
				// cycle detected
				for next := t; next != children; next = pathTo[next] {
					out = append(out, next)
				}
				out = append(out, children)
				out = append(out, t)
			}
		}

		delete(onStack, string(t))
	}

	dfs(start)
	return out
}

func topology(start TaskName, g graph) []TaskName {
	visited := set{}
	out := []TaskName{}
	var dfs func(t TaskName)
	dfs = func(t TaskName) {
		if visited.present(string(t)) {
			return
		}
		visited.add(string(t))
		for _, children := range g[t] {
			dfs(children)
		}
		out = append(out, t)
	}

	dfs(start)
	return out
}
