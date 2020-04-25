package stacktables

import (
	"strconv"
	"strings"
)

type StackCommander struct {
	arithmeticCommands   []string
	logicalCommands      []string
	memoryAccessCommands []string
	branchingCommands    []string
	functionCommands     []string
	registers            []string
	symbolsTable         map[string]string
	FileName             string
	jumpSequence         int
}

func New(fileName string) *StackCommander {
	return &StackCommander{
		getArithmeticCommands(),
		getLogicalCommands(),
		getMemoryAccessCommands(),
		getBranchingCommands(),
		getFunctionCommands(),
		getRegisters(),
		getSymbolsTable(),
		fileName,
		0,
	}
}

func (s *StackCommander) GetJumpSequenceAndInc() int {
	nextIndex := s.jumpSequence
	s.jumpSequence++
	return nextIndex
}

func getSymbolsTable() map[string]string {
	return map[string]string{
		"local":    "LCL",
		"constant": "constant",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
		"temp":     "R5",
		"pointer":  "R3",
	}
}

func getArithmeticCommands() []string {
	return []string{
		"add",
		"sub",
		"neg",
	}
}

func getLogicalCommands() []string {
	return []string{
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

func (s *StackCommander) StartsWithCommand(line string) bool {
	var allCommands []string
	for _, command := range s.arithmeticCommands {
		allCommands = append(allCommands, command)
	}
	for _, command := range s.logicalCommands {
		allCommands = append(allCommands, command)
	}
	for _, command := range s.memoryAccessCommands {
		allCommands = append(allCommands, command)
	}
	for _, command := range s.branchingCommands {
		allCommands = append(allCommands, command)
	}
	for _, command := range s.functionCommands {
		allCommands = append(allCommands, command)
	}
	for _, command := range allCommands {
		if strings.HasPrefix(line, command) {
			return true
		}
	}
	return false
}

func (s *StackCommander) GetCommandSymbol(symbol string, index string) string {
	if symbol == "static" {
		return s.FileName + "." + index
	}
	return s.symbolsTable[symbol]
}

func (s *StackCommander) IsRegister(segment string) bool {
	return contains(s.registers, segment)
}

func (s *StackCommander) IsMemoryAccessCommand(command string) bool {
	return contains(s.memoryAccessCommands, command)
}

func (s *StackCommander) IsArithmeticCommand(command string) bool {
	return contains(s.arithmeticCommands, command)
}

func (s *StackCommander) IsLogicalCommand(command string) bool {
	return contains(s.logicalCommands, command)
}

func contains(elements []string, elem string) bool {
	for _, listElem := range elements {
		if elem == listElem {
			return true
		}
	}
	return false
}
