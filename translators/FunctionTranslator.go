package translators

import (
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
	"log"
	"strconv"
)

type FunctionTranslator struct {
}

func (f *FunctionTranslator) GetFunctionAssemblyCode(args []string, commander *stacktables.StackCommander) string {
	var result string
	command := args[0]
	switch command {
	case "function":
		result = function(args, commander)
	case "call":
		result = call(args, commander)
	case "return":
		result = returnF(commander)
	default:
		panic(command + " is not logical command")
	}
	return result + "\n"
}

func returnF(commander *stacktables.StackCommander) string {
	var result string
	result += copyValue("R14", "LCL", 666)
	result += copyPointer("R15", "R14", -5)
	result += pop([]string{"pop", "argument", "0"}, commander)
	result += copyValue("SP", "ARG", 1)
	result += copyPointer("THAT", "R14", -1)
	result += copyPointer("THIS", "R14", -2)
	result += copyPointer("ARG", "R14", -3)
	result += copyPointer("LCL", "R14", -4)
	result += gotoF("R15", commander)
	return result
}

func copyPointer(dest string, src string, steps int) string {
	var result string
	result += "@" + src + "\n" +
		"D=M\n"
	var op string
	if steps > 0 {
		op = "+"
	} else {
		op = "-"
		steps *= -1
	}
	result += "D=M\n"
	result += "@" + strconv.Itoa(steps) + "\n"
	result += "A=D" + op + "A\n"
	result += "D=M\n"

	result += "@" + dest + "\n" +
		"M=D\n"
	return result + "\n"
}

func copyValue(dest string, src string, steps int) string {
	var result string
	result += "@" + src + "\n" +
		"D=M\n"
	if steps != 666 {
		var op string
		if steps > 0 {
			op = "+"
		} else {
			op = "-"
			steps *= -1
		}
		result += "@" + strconv.Itoa(steps) + "\n"
		result += "D=D" + op + "A\n"
	}
	result += "@" + dest + "\n" +
		"M=D\n"
	return result
}

func call(args []string, commander *stacktables.StackCommander) string {
	var result string
	functionName, argc := args[1], args[2]
	argcInt, err := strconv.Atoi(argc)
	if err != nil {
		log.Fatal(argc + " must be int")
	}
	returnIndex := commander.GetReturnSequenceAndInc()
	returnLabel := commander.FileName + "_-_return" + strconv.Itoa(returnIndex)
	result += push([]string{"push", "constant", returnLabel}, commander)
	result += push([]string{"push", "R1", "0"}, commander)
	result += push([]string{"push", "R2", "0"}, commander)
	result += push([]string{"push", "R3", "0"}, commander)
	result += push([]string{"push", "R4", "0"}, commander)

	backSteps := 0 - 5 - argcInt
	result += copyValue("ARG", "SP", backSteps)
	result += copyValue("LCL", "SP", 0)
	result += gotoF(functionName, commander)
	result += label(returnLabel)

	return result
}

func function(args []string, commander *stacktables.StackCommander) string {
	var result string
	functionName, argc := args[1], args[2] //todo
	commander.FileName = functionName
	result += label(commander.FileName)
	argcInt, err := strconv.Atoi(argc)
	if err != nil {
		log.Fatal(argc + " must be int")
	}
	for i := 0; i < argcInt; i++ {
		result += push([]string{"push", "constant", "0"}, commander)
	}
	return result
}
