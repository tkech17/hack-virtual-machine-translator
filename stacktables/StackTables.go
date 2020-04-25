package stacktables

import (
	"strconv"
	"strings"
)

type StackCommander struct {
	arithmeticLogicalCommands []string
	memoryAccessCommands      []string
	branchingCommands         []string
	functionCommands          []string
	registers                 []string
	fileName                  string
}

func New(fileName string) *StackCommander {
	return &StackCommander{
		getArithmeticLogicalCommands(),
		getMemoryAccessCommands(),
		getBranchingCommands(),
		getFunctionCommands(),
		getRegisters(),
		fileName,
	}
}

func getArithmeticLogicalCommands() []string {
	return []string{
		"add",
		"sub",
		"neg",
		"eq",
		"gt",
		"lt",
		"and",
		"or",
		"not",
	}
}

func getMemoryAccessCommands() []string {
	return []string{
		"pop",
		"push",
	}
}

func getBranchingCommands() []string {
	return []string{
		"label",
		"goto",
		"if-goto",
	}
}

func getFunctionCommands() []string {
	return []string{
		"function",
		"call",
		"return",
	}
}

func getRegisters() []string {
	var result []string
	for i := 0; i <= 15; i++ {
		register := "R" + strconv.Itoa(i)
		result = append(result, register)
	}
	return result
}

func (t *StackCommander) StartsWithCommand(line string) bool {
	var allCommands []string
	for _, command := range t.arithmeticLogicalCommands {
		allCommands = append(allCommands, command)
	}
	for _, command := range t.memoryAccessCommands {
		allCommands = append(allCommands, command)
	}
	for _, command := range t.branchingCommands {
		allCommands = append(allCommands, command)
	}
	for _, command := range t.functionCommands {
		allCommands = append(allCommands, command)
	}
	for _, command := range allCommands {
		if strings.HasPrefix(line, command) {
			return true
		}
	}
	return false
}

func (t *StackCommander) GetCommandSymbol(symbol string, index string) string {
	if symbol == "static" {
		return t.fileName + "." + index
	}
	return symbol
}

func (t *StackCommander) IsRegister(segment string) bool {
	for _, register := range t.registers {
		if segment == register {
			return true
		}
	}
	return false
}
