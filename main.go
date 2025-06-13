package main

import (
	"evil/lang1/ast"
	"evil/lang1/input"
	"evil/lang1/source"
	"fmt"
)

func main() {
	fmt.Println("lang1>>")

	input.Setup()

	Init()
	ast.Expression()

	input.Close()

	fmt.Println()
}

func Init() {
	source.GetChar()
}
