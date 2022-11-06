package main

type action string

type task struct{
	name string
	depedencies []task
	actions []action
}

func newTask(name string, deps []task, actions []action) task{
	return task{
		name: name, 
		depedencies: deps,
		actions: actions,
	}
}


type buildSystem struct {
	tasks []task
}

func newBuildSystem() *buildSystem {
	return &buildSystem{}
}

func (b *buildSystem) addTask(t task) {
	b.tasks = append(b.tasks, t)
}

func (b *buildSystem) run() ([]action, error) {
	// validate and run
	return nil, nil
}

