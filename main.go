package main

import (
	"fmt"
	. "makefile-clone/buildsystem"
)

// https://xmonader.github.io/nimdays/day11_buildsystem.html
func main() {
	b := NewBuildSystem()
	b.AddTask(NewTask("clean", nil, []Action{"cleaning"}))
	b.AddTask(NewTask("stepA", nil, []Action{"print foo"}))
	b.AddTask(NewTask("stepB", nil, []Action{"print bar", "print bar again"}))
	b.AddTask(NewTask("run", []TaskName{"stepA", "stepB"}, []Action{"execute main"}))

	cmds, err := b.Run("run")
	if err != nil {
		fmt.Println("got error:", err)
		return
	}
	for _, v := range cmds {
		fmt.Println(v)
	}
}
