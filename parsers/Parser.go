package parsers

import (
	"github.com/tkech17/hach_virtual_machine_translator/stacktables"
	"github.com/tkech17/hach_virtual_machine_translator/utils/functional"
	"strings"
)

type ParseResult struct {
}

func New() *ParseResult {
	return &ParseResult{}
}

func (p *ParseResult) Parse(content string, commander *stacktables.StackCommander) []string {
	result := strings.Split(content, "\n")
	result = functional.MapString(result, removeCommentLine)
	result = functional.MapString(result, strings.TrimSpace)
	return functional.Filter(result, func(elem string) bool {
		return commander.StartsWithCommand(elem)
	})
}

func removeCommentLine(str string) string {
	index := strings.Index(str, "//")
	if index == -1 {
		return str
	}
	return str[0:index]
}
