package parser

import (
	"github.com/tkech17/hach_virtual_machine_translator/utils/functional"
	"strings"
)

type ParseResult struct {
	AssemblyLines []string
}

func New() *ParseResult {
	return &ParseResult{nil}
}

func (p *ParseResult) Parse(content string) {
	p.AssemblyLines = strings.Split(content, "\n")
	p.AssemblyLines = functional.Filter(p.AssemblyLines, isNotEmptyLine)
	p.AssemblyLines = functional.Filter(p.AssemblyLines, isNotComment)
	p.AssemblyLines = functional.MapString(p.AssemblyLines, removeCommentLine)
	p.AssemblyLines = functional.MapString(p.AssemblyLines, strings.TrimSpace)
	p.AssemblyLines = functional.MapString(p.AssemblyLines, removeSpaces)
}

func removeSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func removeCommentLine(str string) string {
	index := strings.Index(str, "//")
	if index == -1 {
		return str
	}
	return str[0:index]
}

func isNotEmptyLine(line string) bool {
	return len(strings.TrimSpace(line)) != 0
}

func isNotComment(line string) bool {
	trimmed := strings.TrimSpace(line)
	isComment := strings.HasPrefix(trimmed, "//")
	return !isComment
}
