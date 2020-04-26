package translators

import (
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
	"strings"
)

type Translator struct {
	ArithmeticTranslator
	LogicalTranslator
	BranchingTranslator
	MemoryAccessTranslator
	FunctionTranslator
}

func New() *Translator {
	return &Translator{}
}

func (t *Translator) TranslateVMCodeIntoAssembly(lines []string, commander *stacktables.StackCommander, init bool) string {
	var result = ""
	if init {
		result = initAssembly(t, commander)
	}
	for _, line := range lines {
		args := strings.Split(line, " ")
		result += translateVMLineIntoAssembly(t, commander, args)
	}
	return result
}

func initAssembly(t *Translator, commander *stacktables.StackCommander) string {
	var result string
	result += "@256\n" +
		"D=A\n" +
		"@SP\n" +
		"M=D\n"
	fileName := commander.FileName
	commander.FileName = "Sys"
	result += t.GetFunctionAssemblyCode([]string{"call", "Sys.init", "0"}, commander)
	commander.FileName = fileName
	return result
}

func translateVMLineIntoAssembly(t *Translator, commander *stacktables.StackCommander, args []string) string {
	command := args[0]
	var result string
	if commander.IsArithmeticCommand(command) {
		result = t.GetArithmeticAssemblyCode(args, commander)
	} else if commander.IsLogicalCommand(command) {
		result = t.GetLogicalAssemblyCode(args, commander)
	} else if commander.IsMemoryAccessCommand(command) {
		result = t.GetMemoryAccessAssemblyCode(args, commander)
	} else if commander.IsBranchingCommand(command) {
		result = t.GetBranchingAssemblyCode(args, commander)
	} else if commander.IsFunctionCommand(command) {
		result = t.GetFunctionAssemblyCode(args, commander)
	} else {
		panic("not defined command")
	}
	return result
}
