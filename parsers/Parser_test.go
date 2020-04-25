package parsers_test

import (
	"github.com/tkech17/hach_virtual_machine_translator/parsers"
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
	"github.com/tkech17/hach_virtual_machine_translator/utils/tests"
	"testing"
)

func TestParse(t *testing.T) {
	var str = "" +
		"// This files is part of www.nand2tetris.org\n" +
		"// and the book \"The Elements of Computing Systems\"\n" +
		"// by Nisan and Schocken, MIT Press.\n" +
		"       // File name: projects/06/add/Add.asm\n" +
		"  \n" +
		"  // Computes R0 = 2 + 3  (R0 refers to RAM[0])\n" +
		"\n" +
		"push constant 7\n" +
		"push constant 8\n" +
		"add\n"

	var expectedLines = []string{
		"push constant 7",
		"push constant 8",
		"add",
	}

	parser := parsers.New()
	result := parser.Parse(str, stacktables.New(""))

	tests.AssertEqualsInt(t, len(result), len(expectedLines))
	for i := range expectedLines {
		expected := expectedLines[i]
		actual := result[i]
		tests.AssertEqualsString(t, expected, actual)
	}
}
