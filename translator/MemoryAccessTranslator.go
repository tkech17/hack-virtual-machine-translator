package translator

import "github.com/tkech17/hach_virtual_machine_translator/stacktables"

type MemoryAccessTranslator struct {
}

func (m *MemoryAccessTranslator) Pop(args []string, commander *stacktables.StackCommander) string {
	var result string
	segment, index := args[1], args[2]
	symbol := commander.GetCommandSymbol(segment, index)
	if segment == "static" {
		result += "@" + symbol + "\n" +
			"D=A\n"
	} else {
		result += "@" + symbol + "\n" +
			"D=A\n"
		if commander.IsRegister(segment) {
			result += "@" + segment + "\n" +
				"D=D+A\n"
		} else {
			result += "@" + segment + "\n" +
				"D=D+M\n"
		}
	}
	result += "@R13\n" +
		"M=D\n" +
		"@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"@R13\n" +
		"A=M\n" +
		"M=D\n"
	return result + "\n"
}

func (m *MemoryAccessTranslator) Push(args []string, commander *stacktables.StackCommander) string {
	var result string
	segment, index := args[1], args[2]
	symbol := commander.GetCommandSymbol(segment, index)
	if segment == "static" {
		result += "@" + symbol + "\n" +
			"D=M\n"
	} else {
		result += "@" + index + "\n" +
			"D=A\n"
		if segment != "constant" {
			if commander.IsRegister(segment) {
				result += "@" + segment + "\n" +
					"A=D+A\n" +
					"D=M\n"
			} else {
				result += "@" + segment + "\n" +
					"A=D+M\n" +
					"D=M\n"
			}
		}
	}
	result += "@SP\n" +
		"A=M\n" +
		"M=D\n" +
		"@SP\n" +
		"M=M+1\n"
	return result + "\n"
}
