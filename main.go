package main

import (
	"fmt"
	. "makefile-clone/buildsystem"
	"makefile-clone/parser"
)

// https://xmonader.github.io/nimdays/day11_buildsystem.html
func main() {
	byHand()

	fmt.Println("")
	fmt.Println("now same one, but parsed")
	fmt.Println("")

	parsed()
}

func execute(b *BuildSystem) {
	cmds, err := b.Run("run")
	if err != nil {
		fmt.Println("got error:", err)
		return
	}
	for _, v := range cmds {
		fmt.Println(v)
	}
}

func byHand() {
	b := NewBuildSystem()
	b.AddTask(NewTask("clean", nil, []Action{"cleaning"}))
	b.AddTask(NewTask("stepA", nil, []Action{"print foo"}))
	b.AddTask(NewTask("stepB", nil, []Action{"print bar", "print bar again"}))
	b.AddTask(NewTask("run", []TaskName{"stepA", "stepB"}, []Action{"execute main"}))

	execute(b)
}

func parsed() {
	input := `clean:
	cleaning

stepA:
	print foo

stepB:
	print bar
	print bar again

run: stepA stepB
	execute main`

	b2, err := parser.ParseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	execute(b2)
}