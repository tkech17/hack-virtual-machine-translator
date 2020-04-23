package parser_test

import (
	"github.com/tkech17/hach_virtual_machine_translator/parser"
	"github.com/tkech17/hach_virtual_machine_translator/utils/tests"
	"testing"
)

func TestParse(t *testing.T) {
	var str = "" +
		"// This file is part of www.nand2tetris.org\n" +
		"// and the book \"The Elements of Computing Systems\"\n" +
		"// by Nisan and Schocken, MIT Press.\n" +
		"       // File name: projects/06/add/Add.asm\n" +
		"  \n" +
		"  // Computes R0 = 2 + 3  (R0 refers to RAM[0])\n" +
		"\n" +
		"@2 //ბლა\n" +
		"D =   A \n" +
		"  @3  \n" +
		"    D=D+A\n" +
		"@temp\n" +
		"M   =  D\n"

	var expectedLines = []string{
		"@2",
		"D=A",
		"@3",
		"D=D+A",
		"@temp",
		"M=D",
	}

	result := parser.New()
	result.Parse(str)

	tests.AssertEqualsInt(t, len(result.AssemblyLines), len(expectedLines))
	for i := range expectedLines {
		expected := expectedLines[i]
		actual := result.AssemblyLines[i]
		tests.AssertEqualsString(t, expected, actual)
	}
}
