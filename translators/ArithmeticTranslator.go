package translators

import "github.com/tkech17/hach_virtual_machine_translator/stacktables"

type ArithmeticTranslator struct {
}

func (a *ArithmeticTranslator) GetArithmeticAssemblyCode(args []string, commander *stacktables.StackCommander) string {
	var result string
	command := args[0]
	switch command {
	case "add":
		result = add(result)
	case "sub":
		result = sub(result)
	case "neg":
		result = neg(result)
	default:
		panic(command + " is not arithmetic command")
	}
	return result + "\n"
}

func neg(result string) string {
	result = "@SP\n" +
		"A=M-1\n" +
		"M=-M\n"
	return result
}

func sub(result string) string {
	result = "@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=M-D\n"
	return result
}

func add(result string) string {
	result = "@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=D+M\n"
	return result
}
