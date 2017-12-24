package call911

import (
	"github.com/maruel/panicparse/stack"
	"github.com/mgutz/ansi"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

const resetFG = ansi.DefaultFG + "\033[m"

var errorColor = ansi.ColorCode("red+b")

var defaultPalette = stack.Palette{
	EOLReset:               resetFG,
	RoutineFirst:           ansi.ColorCode("magenta+b"),
	CreatedBy:              ansi.LightBlack,
	Package:                ansi.ColorCode("default+b"),
	SourceFile:             resetFG,
	FunctionStdLib:         ansi.Green,
	FunctionStdLibExported: ansi.ColorCode("green+b"),
	FunctionMain:           ansi.ColorCode("yellow+b"),
	FunctionOther:          ansi.Red,
	FunctionOtherExported:  ansi.ColorCode("red+b"),
	Arguments:              resetFG,
}
