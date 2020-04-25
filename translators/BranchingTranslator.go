package translators

import (
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
)

type BranchingTranslator struct {
}

func (b *BranchingTranslator) GetBranchingAssemblyCode(args []string, commander *stacktables.StackCommander) string {
	var result string
	command, lbl := args[0], commander.FileName+"-_-"+args[1]
	switch command {
	case "label":
		result = label(lbl)
	case "goto":
		result = gotoF(lbl)
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

func gotoF(lbl string) string {
	return "@" + lbl + "\n" +
		"0;JMP\n"
}

func ifGoto(lbl string) string {
	return "@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"@" + lbl + "\n" +
		"D;JNE\n"
}
