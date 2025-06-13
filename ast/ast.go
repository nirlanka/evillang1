package ast

import "evil/lang1/source"

func Expression() {
	source.EmitLn("MOVE #" + string(source.GetNum()) + ",D0")
}
