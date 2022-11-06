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
	// validate and run
	dependencyGraph := map[taskName][]taskName{}


	return nil, nil
}

