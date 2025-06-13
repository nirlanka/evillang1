package main

import (
	"evil/lang1/source"
	"fmt"
)

func main() {
	fmt.Println("lang1 compiler")

	Init()

	fmt.Println("end.")
}

func Init() {
	source.GetChar()
}
