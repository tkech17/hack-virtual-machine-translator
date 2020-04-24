package vm

type VirtualMachine struct {
}

func GetVirtualMachine() *VirtualMachine {
	return &VirtualMachine{}
}

func (t *VirtualMachine) GenerateAssembly(vmContent string) string   {

	return ""
}