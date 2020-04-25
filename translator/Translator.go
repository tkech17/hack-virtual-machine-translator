package translator

import (
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
	"strings"
)

type Translator struct {
	ArithmeticLogicalTranslator
	BranchingTranslator
	MemoryAccessTranslator
	FunctionTranslator
}

func New() *Translator {
	return &Translator{}
}

func (t *Translator) TranslateVMCodeIntoAssembly(lines []string, commander *stacktables.StackCommander) string {
	var result = ""
	for _, line := range lines {
		args := strings.Split(line, " ")
		result += translateVMLineIntoAssembly(t, commander, args)
	}
	return result
}

func translateVMLineIntoAssembly(t *Translator, commander *stacktables.StackCommander, args []string) string {
	command := args[0]
	var result string
	switch command {
	case "Push":
		result = t.Push(args, commander)
	case "pop":
		result = t.Pop(args, commander)
	default:
		panic("bla")
	}
	return result
}
