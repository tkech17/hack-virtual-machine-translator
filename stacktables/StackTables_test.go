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
	tests.AssertEqualsString(t, "bla.1", stackCommander.GetCommandSymbol("static", "1"))
}

func TestStackCommander_IsRegister(t *testing.T) {
	var stackCommander StackCommander = stacktables.New("bla")

	tests.AssertTrue(t, stackCommander.IsRegister("R0"), "R0 is register")
	tests.AssertFalse(t, stackCommander.IsRegister("0"), "0 is not register")
}
