package vm

import (
	"github.com/tkech17/hach_virtual_machine_translator/parsers"
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
	"github.com/tkech17/hach_virtual_machine_translator/translator"
)

type VirtualMachine struct {
}

func GetVirtualMachine() *VirtualMachine {
	return &VirtualMachine{}
}

func (t *VirtualMachine) GenerateAssembly(fileName string, content string) string {
	var stackCommander = stacktables.New(fileName)
	var parser = parsers.New()
	var trans = translator.New()
	vmLines := parser.Parse(content, stackCommander)
	return trans.TranslateVMCodeIntoAssembly(vmLines, stackCommander)
}
