package source

import (
	"evil/lang1/consts"
	"evil/lang1/input"
	"evil/lang1/state"
	"fmt"
	"strings"
	"unicode"
)

func GetChar() {
	var err error
	state.Look, err = input.ReadRune()
	if err != nil {
		AbortError(err)
	}
}

func AbortCustom(s string) {
	write("\n")
	panic("Error: " + s + ".")
}

func AbortError(err error) {
	write("\n")
	panic(err)
}

func Expected(s string) {
	AbortCustom(s + " expected")
}

func Match(r rune) {
	if state.Look == r {
		GetChar()
	} else {
		Expected("\"" + string(r) + "\"")
	}
}

func isAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func GetName() rune {
	if !isAlpha(state.Look) {
		Expected("Name")
	}

	s := strings.ToUpper(string(state.Look))[0]
	res := rune(s)
	GetChar()

	return res
}

func GetNum() rune {
	if !isDigit(state.Look) {
		Expected("Integer")
	}

	res := state.Look
	GetChar()

	return res
}

func write(a ...any) {
	fmt.Print(a...)
}

func Emit(s string) {
	write(string(consts.Tab), s)
}

func EmitLn(s string) {
	Emit(s)
	write("\n")
}
