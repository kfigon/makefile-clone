package parser

import (
	"unicode"
	. "makefile-clone/buildsystem"
)

type tokenType int

const (
	identifier tokenType = iota
	colon
	tabedCmd
)

type token struct {
	tok    tokenType
	lexeme string
}

func lex(input string) []token {
	out := []token{}
	i := 0
	readUntil := func(predicate func(rune)bool) string {
		word := ""
		word += string(input[i])
		for i+1 < len(input) {
			if !predicate(rune(input[i+1])) {
				break
			}
			word += string(input[i+1])
			i++
		}
		return word
	}

	for i < len(input) {
		c := rune(input[i])
		if c == '\t' {
			i++
			word := readUntil(func(r rune) bool { return r != '\n'})
			out = append(out, token{tabedCmd, word})
		} else if unicode.IsSpace(c) {
			// skip newlines or spaces
		} else if c == ':'{
			out = append(out, token{colon, ":"})
		} else {
			word := readUntil(func(r rune) bool { return !unicode.IsSpace(r) && r != ':'})
			out = append(out, token{identifier, word})
		}
		i++
	}
	return out
}

func ParseInput(input string) (*BuildSystem, error) {
	tokens := lex(input)
	p := &parser{tokens: tokens}
	return p.parse()
}

type parser struct{
	tokens []token
	idx int
}

func (p *parser) parse() (*BuildSystem, error) {
	b := NewBuildSystem()
	// b.AddTask(NewTask("clean", nil, []Action{"cleaning"}))
	// b.AddTask(NewTask("stepA", nil, []Action{"print foo"}))
	// b.AddTask(NewTask("stepB", nil, []Action{"print bar", "print bar again"}))
	// b.AddTask(NewTask("run", []TaskName{"stepA", "stepB"}, []Action{"execute main"}))

	return b, nil
}