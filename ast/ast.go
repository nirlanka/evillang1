package ast

import (
	"evil/lang1/source"
	"evil/lang1/state"
)

func Term() {
	num := string(source.GetNum())
	source.EmitLn("MOVE #" + num + ",D0")
}

func Expression() {
	Term()

	source.EmitLn("MOVE D0,D1")

	switch state.Look {
	case '+':
		Add()
	case '-':
		Subtract()
	default:
		source.Expected("Addop")
	}
}

func Add() {
	source.Match('+')
	Term()
	source.EmitLn("ADD D1,D0")
}

func Subtract() {
	source.Match('-')
	Term()
	source.EmitLn("SUB D1,D0")
}
