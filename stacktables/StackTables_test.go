package stacktables_test

import (
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
	"github.com/tkech17/hach_virtual_machine_translator/utils/tests"
	"testing"
)

type StackCommander interface {
	StartsWithCommand(vmContent string) bool
	GetCommandSymbol(symbol string, index string) string
	IsRegister(segment string) bool
	IsMemoryAccessCommand(command string) bool
	IsArithmeticCommand(command string) bool
	IsLogicalCommand(command string) bool
	IsBranchingCommand(command string) bool
	GetJumpSequenceAndInc() int
}

func TestNew(t *testing.T) {
	tests.AssertNotNull(t, stacktables.New("bla"))
}

func TestStackCommander_StartsWithCommand(t *testing.T) {
	var stackCommander StackCommander = stacktables.New("")

	tests.AssertTrue(t, stackCommander.StartsWithCommand("push bla"), "push is memory access command")
	tests.AssertTrue(t, stackCommander.StartsWithCommand("add 1 2"), "add is arithmetic command")
	tests.AssertTrue(t, stackCommander.StartsWithCommand("not and"), "not is logical command")
	tests.AssertTrue(t, stackCommander.StartsWithCommand("goto  adas "), "goto is branching command")
	tests.AssertTrue(t, stackCommander.StartsWithCommand("call  adas "), "call is function command")
	tests.AssertFalse(t, stackCommander.StartsWithCommand("bla"), "bla is not command")
}

func TestStackCommander_GetCommandSymbol(t *testing.T) {
	var stackCommander StackCommander = stacktables.New("bla")

	tests.AssertEqualsString(t, "constant", stackCommander.GetCommandSymbol("constant", "1"))
	tests.AssertEqualsString(t, "ARG", stackCommander.GetCommandSymbol("argument", "1"))
	tests.AssertEqualsString(t, "bla.1", stackCommander.GetCommandSymbol("static", "1"))
}

func TestStackCommander_IsRegister(t *testing.T) {
	var stackCommander StackCommander = stacktables.New("bla")

	tests.AssertTrue(t, stackCommander.IsRegister("R0"), "R0 is register")
	tests.AssertFalse(t, stackCommander.IsRegister("0"), "0 is not register")
}

func TestStackCommander_IsMemoryAccessCommand(t *testing.T) {
	var stackCommander StackCommander = stacktables.New("bla")

	tests.AssertTrue(t, stackCommander.IsMemoryAccessCommand("pop"), "pop is memory access command")
	tests.AssertFalse(t, stackCommander.IsMemoryAccessCommand("bla"), "bla is not memory access command")
}

func TestStackCommander_IsArithmeticCommand(t *testing.T) {
	var stackCommander StackCommander = stacktables.New("bla")

	tests.AssertTrue(t, stackCommander.IsArithmeticCommand("add"), "add is arithmetic command")
	tests.AssertFalse(t, stackCommander.IsArithmeticCommand("not"), "not is not arithmetic command")
}

func TestStackCommander_IsLogicalCommand(t *testing.T) {
	var stackCommander StackCommander = stacktables.New("bla")

	tests.AssertTrue(t, stackCommander.IsLogicalCommand("not"), "not is logical command")
	tests.AssertFalse(t, stackCommander.IsLogicalCommand("add"), "add is not logical command")
}

func TestStackCommander_IsBranchingCommand(t *testing.T) {
	var stackCommander StackCommander = stacktables.New("bla")

	tests.AssertTrue(t, stackCommander.IsBranchingCommand("goto"), "goto is branching command")
	tests.AssertFalse(t, stackCommander.IsBranchingCommand("add"), "add is not branching command")
}

func TestStackCommander_GetJumpSequenceAndInc(t *testing.T) {
	var stackCommander StackCommander = stacktables.New("bla")

	tests.AssertEqualsInt(t, 0, int(stackCommander.GetJumpSequenceAndInc()))
	tests.AssertEqualsInt(t, 1, int(stackCommander.GetJumpSequenceAndInc()))
}

func TestStackCommander_FileName(t *testing.T) {
	tests.AssertEqualsString(t, "bla", stacktables.New("bla").FileName)
}
