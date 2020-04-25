package translators

import (
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
)

type MemoryAccessTranslator struct {
}

func (m *MemoryAccessTranslator) GetMemoryAccessAssemblyCode(args []string, commander *stacktables.StackCommander) string {
	var result string
	command := args[0]
	switch command {
	case "push":
		result = push(args, commander)
	case "pop":
		result = pop(args, commander)
	default:
		panic(command + " is not memory access command")
	}
	return result
}

func pop(args []string, commander *stacktables.StackCommander) string {
	var result string
	seg, index := args[1], args[2]
	symbol := commander.GetCommandSymbol(seg, index)
	if seg == "static" {
		result += "@" + symbol + "\n" +
			"D=A\n"
	} else {
		result += "@" + index + "\n" +
			"D=A\n"
		if commander.IsRegister(symbol) {
			result += "@" + symbol + "\n" +
				"D=D+A\n"
		} else {
			result += "@" + symbol + "\n" +
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

func push(args []string, commander *stacktables.StackCommander) string {
	var result string
	seg, index := args[1], args[2]
	symbol := commander.GetCommandSymbol(seg, index)
	if seg == "static" {
		result += "@" + symbol + "\n" +
			"D=M\n"
	} else {
		result += "@" + index + "\n" +
			"D=A\n"
		if seg != "constant" {
			if commander.IsRegister(symbol) {
				result += "@" + symbol + "\n" +
					"A=D+A\n" +
					"D=M\n"
			} else {
				result += "@" + symbol + "\n" +
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
