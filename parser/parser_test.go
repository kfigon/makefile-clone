package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	input := `default: program

program.o: program.c
	gcc -c program.c -o program.o

program: program.o
	gcc program.o -o program
clean:
	rm -f program.o
	rm -f program`

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