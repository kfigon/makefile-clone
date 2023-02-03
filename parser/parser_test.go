package parser

import (
	. "makefile-clone/buildsystem"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const input = `default: program

program.o: program.c
	gcc -c program.c -o program.o

program: program.o
	gcc program.o -o program
clean:
	rm -f program.o
	rm -f program`

func TestLexer(t *testing.T) {
	expected := []token{
		{identifier, "default"},
		{colon, ":"},
		{identifier, "program"},

		{identifier, "program.o"},
		{colon, ":"},
		{identifier, "program.c"},
		{tabedCmd, "gcc -c program.c -o program.o"},

		{identifier, "program"},
		{colon, ":"},
		{identifier, "program.o"},
		{tabedCmd, "gcc program.o -o program"},

		{identifier, "clean"},
		{colon, ":"},
		{tabedCmd, "rm -f program.o"},
		{tabedCmd, "rm -f program"},

	}
	
	got := lex(input)
	assert.Equal(t, expected, got)
}

func TestParse(t *testing.T) {
	b, err := ParseInput(input)

	exp := NewBuildSystem()
	exp.AddTask(NewTask("default", []TaskName{"program"}, nil))
	exp.AddTask(NewTask("program.o", []TaskName{"program.c"}, []Action{"gcc -c program.c -o program.o"}))
	exp.AddTask(NewTask("program", []TaskName{"program.o"}, []Action{"gcc program.o -o program"}))
	exp.AddTask(NewTask("clean", nil, []Action{"rm -f program.o", "rm -f program"}))

	require.NoError(t, err)
	assert.Equal(t, *exp, *b)
}