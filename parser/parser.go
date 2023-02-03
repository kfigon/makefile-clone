package parser

import (
	"fmt"
	. "makefile-clone/buildsystem"
	"unicode"
)

type tokenType int

const (
	identifier tokenType = iota
	colon
	tabedCmd
	doubleNewline
	newLine
)

type token struct {
	tok    tokenType
	lexeme string
}

func (t token) String() string {
	tokType := [...]string {
		"identifier",
		"colon",
		"tabedCmd",
		"doubleNewline",
		"newLine",
	}[t.tok]
	return fmt.Sprintf("(%v, %q)", tokType, t.lexeme)
}

func lex(input string) []token {
	out := []token{}
	i := 0
	peek := func() (rune, bool) {
		if i+1 >= len(input) {
			return 0, false
		}
		return rune(input[i+1]), true
	}

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
			if c == '\n' {
				if next, nextOk := peek(); nextOk && next == '\n' {
					i++
					out = append(out, token{doubleNewline, "\n\n"})
				} else {
					out = append(out, token{newLine, "\n"})
				}
			}
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

	for current, ok := p.current(); ok; current, ok = p.current(){
		next, nextOk := p.peek()

		if current.tok == identifier && nextOk && next.tok == colon {
			taskName := current.lexeme
			p.consume()
			p.consume()
			var tasks []TaskName
			var actions []Action

			for current, ok = p.current(); ok; current, ok = p.current() {
				if current.tok == doubleNewline {
					b.AddTask(NewTask(TaskName(taskName), tasks, actions))
					break
				} else if current.tok == identifier {
					tasks = append(tasks, TaskName(current.lexeme))
				} else if current.tok == tabedCmd {
					actions = append(actions, Action(current.lexeme))
				} else if current.tok != newLine {
					return nil, fmt.Errorf("Invalid token %v", current)
				}
				p.consume()
			}
			if p.idx >= len(p.tokens) && taskName != "" {
				b.AddTask(NewTask(TaskName(taskName), tasks, actions))
			}
		}
		p.consume()
	}

	return b, nil
}

func (p *parser) consume() {
	p.idx++
}

func (p *parser) current() (token, bool) {
	if p.idx >= len(p.tokens) {
		return token{}, false
	}
	return p.tokens[p.idx], true
}

func (p *parser) peek() (token, bool) {
	if p.idx+1 >= len(p.tokens) {
		return token{}, false
	}
	return p.tokens[p.idx+1], true
}