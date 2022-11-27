package main

import "fmt"

// https://xmonader.github.io/nimdays/day11_buildsystem.html
func main() {
	b := newBuildSystem()
	b.addTask(newTask("clean", nil, []action{"cleaning"}))
	b.addTask(newTask("stepA", nil, []action{"print foo"}))
	b.addTask(newTask("stepB", nil, []action{"print bar", "print bar again"}))
	b.addTask(newTask("run", []taskName{"stepA", "stepB"}, []action{"execute main"}))

	cmds, err := b.run("run")
	if err != nil {
		fmt.Println("got error:", err)
		return
	}
	for _, v := range cmds {
		fmt.Println(v)
	}
}