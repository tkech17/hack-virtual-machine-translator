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
	nonMultiLinedContent := multiLineCommentsRemovedContent(content)
	result := strings.Split(nonMultiLinedContent, "\n")
	result = functional.MapString(result, removeCommentLine)
	result = functional.MapString(result, strings.TrimSpace)
	return functional.Filter(result, func(elem string) bool {
		return commander.StartsWithCommand(elem)
	})
}

func multiLineCommentsRemovedContent(content string) string {
	for true {
		start := strings.Index(content, "{")
		end := strings.Index(content, "}")
		if start == -1 {
			break
		}
		content = content[:start] + content[end+1:]
	}
	return content
}

func removeCommentLine(str string) string {
	index := strings.Index(str, "//")
	if index == -1 {
		return str
	}
	return str[0:index]
}
