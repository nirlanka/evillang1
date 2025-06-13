package main

import (
	"evil/lang1/ast"
	"evil/lang1/source"
	"fmt"
)

func main() {
	fmt.Println("lang1>>")

	Init()
	ast.Expression()

	fmt.Println()
}

func Init() {
	source.GetChar()
}
