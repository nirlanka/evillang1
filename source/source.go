package source

import (
	"evil/lang1/consts"
	"evil/lang1/state"
	"fmt"
	"strings"
	"unicode"
)

func GetChar() {
	var err error
	state.Look, _, err = state.Reader.ReadRune()
	if err != nil {
		AbortError(err)
	}
}

func AbortCustom(s string) {
	panic("Error: " + s + ".")
}

func AbortError(err error) {
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

func IsAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

func IsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func GetName() rune {
	if !IsAlpha(state.Look) {
		Expected("Name")
	}

	res := []rune(strings.ToUpper(string(state.Look)))
	GetChar()

	return res[0]
}

func GetNum() rune {
	if !IsDigit(state.Look) {
		Expected("Intgeger")
	}

	res := state.Look
	GetChar()

	return res
}

func Write(a ...any) {
	fmt.Print(a...)
}

func Emit(s string) {
	Write(string(consts.Tab), s)
}

func EmitLn(s string) {
	Emit(s)
	Write('\n')
}
