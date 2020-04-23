package main_test

import (
	"github.com/tkech17/hach_virtual_machine_translator/utils/functional"
	"github.com/tkech17/hach_virtual_machine_translator/utils/tests"
	"testing"
)


func TestFilter(t *testing.T) {
	var slice = []string{"1", "2", "22"}
	var filtered = functional.Filter(slice, func(elem string) bool {
		return len(elem)%2 == 0
	})
	tests.AssertEqualsInt(t, 1, len(filtered))
	tests.AssertEqualsInt(t, 1, len(filtered))
	tests.AssertEqualsString(t, "22", filtered[0])
}

