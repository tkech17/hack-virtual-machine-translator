package translators

import (
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
	"strconv"
)

type LogicalTranslator struct {
}

func (l *LogicalTranslator) GetLogicalAssemblyCode(args []string, commander *stacktables.StackCommander) string {
	var result string
	command := args[0]
	switch command {
	case "and":
		result = and()
	case "or":
		result = or()
	case "not":
		result = not()
	case "eq":
		result = eq(commander)
	case "gt":
		result = gt(commander)
	case "lt":
		result = lt(commander)
	default:
		panic(command + " is not logical command")
	}
	return result + "\n"
}

func lt(commander *stacktables.StackCommander) string {
	jumpIndex := commander.GetJumpSequenceAndInc()
	label := "FOR_JUMP_" + commander.FileName + strconv.Itoa(jumpIndex)
	return "@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"M=-1\n" +
		"@" + label + "\n" +
		"D;JLT\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"(" + label + ")\n"
}

func gt(commander *stacktables.StackCommander) string {
	jumpIndex := commander.GetJumpSequenceAndInc()
	label := "FOR_JUMP_" + commander.FileName + strconv.Itoa(jumpIndex)
	return "@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"M=-1\n" +
		"@" + label + "\n" +
		"D;JGT\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"(" + label + ")\n"
}

func eq(commander *stacktables.StackCommander) string {
	jumpIndex := commander.GetJumpSequenceAndInc()
	label := "FOR_JUMP_" + commander.FileName + strconv.Itoa(jumpIndex)
	return "@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"A=A-1\n" +
		"D=M-D\n" +
		"M=-1\n" +
		"@" + label + "\n" +
		"D;JEQ\n" +
		"@SP\n" +
		"A=M-1\n" +
		"M=0\n" +
		"(" + label + ")\n"
}

func not() string {
	return "@SP\n" +
		"A=M-1\n" +
		"M=!M\n"
}

func or() string {
	return "@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=M|D\n"
}

func and() string {
	return "@SP\n" +
		"M=M-1\n" +
		"A=M\n" +
		"D=M\n" +
		"A=A-1\n" +
		"M=D&M\n"
}
