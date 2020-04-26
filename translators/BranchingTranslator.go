package translators

import (
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
)

type BranchingTranslator struct {
}

func (b *BranchingTranslator) GetBranchingAssemblyCode(args []string, commander *stacktables.StackCommander) string {
	var result string
	command, lbl := args[0], commander.FunctionName+"-_-"+args[1]
	switch command {
	case "label":
		result = label(lbl)
	case "goto":
		result = gotoF(lbl, commander)
	case "if-goto":
		result = ifGoto(lbl)
	default:
		panic(command + " is not logical command")
	}
	return result + "\n"
}

func label(lbl string) string {
	return "(" + lbl + ")\n"
}

func gotoF(lbl string, commander *stacktables.StackCommander) string {
	var result string
	result += "@" + lbl + "\n"
	if commander.IsTemp(lbl) {
		result += "A=M\n"
	}
	result += "0;JMP\n"
	return result
}

func ifGoto(lbl string) string {
	return "@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"@" + lbl + "\n" +
		"D;JNE\n"
}
