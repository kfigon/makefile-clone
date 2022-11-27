package main

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

type action string
type taskName string

type task struct {
	name        taskName
	depedencies []taskName
	actions     []action
}

func newTask(name taskName, deps []taskName, actions []action) task {
	return task{
		name:        name,
		depedencies: deps,
		actions:     actions,
	}
}

type buildSystem struct {
	tasks map[taskName]task
}

func newBuildSystem() *buildSystem {
	return &buildSystem{
		tasks: map[taskName]task{},
	}
}

func (b *buildSystem) addTask(t task) {
	b.tasks[t.name] = t
}

func (b *buildSystem) run(t taskName) ([]action, error) {
	dependencyGraph := b.buildGraph()
	if cycle := hasCycles(t, dependencyGraph); len(cycle) != 0 {
		var vals []string
		for _, v := range cycle {
			vals = append(vals, string(v))
		}
		return nil, fmt.Errorf("Cycle detected %s", strings.Join(vals, "->"))
	}

	top := topology(t, dependencyGraph)
	out := []action{}
	for _,v := range top {
		task := b.tasks[v]
		out = append(out, task.actions...)
	}
	return out, nil
}

type graph map[taskName][]taskName

func (b *buildSystem) buildGraph() map[taskName][]taskName {
	out := graph{}
	for name, t := range b.tasks {
		tasks := out[name]
		tasks = append(tasks, t.depedencies...)
		out[name] = tasks
	}
	return out
}

func hasCycles(start taskName, g graph) []taskName {
	visited := set{}
	pathTo := map[taskName]taskName{}
	out := []taskName{}
	onStack := set{}

	var dfs func(taskName)
	dfs = func(t taskName) {
		if visited.present(string(t)) {
			return
		}
		onStack.add(string(t))
		visited.add(string(t))
		for _,children := range g[t] {
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

func topology(start taskName, g graph) []taskName {
	visited := set{}
	out := []taskName{}
	var dfs func(t taskName)
	dfs = func(t taskName) {
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