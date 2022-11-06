package main

type action string
type taskName string

type task struct{
	name taskName
	depedencies []taskName
	actions []action
}

func newTask(name taskName, deps []taskName, actions []action) task {
	return task {
		name: name, 
		depedencies: deps,
		actions: actions,
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
	
	out := []action{}
	seen := map[taskName]struct{}{}
	
	var foo func(taskName)
	foo = func(currentTask taskName) {
		
	}
	
	foo(t)
	return out, nil
}

func (b *buildSystem) buildGraph() map[taskName][]taskName {
	out := map[taskName][]taskName{}
	for name,t := range b.tasks {
		tasks, ok := out[name]
		if !ok {
			out[name] = t.depedencies
		} else {
			tasks = append(tasks, t.depedencies...) 
			out[name] = tasks
		}
	}
	return out
}