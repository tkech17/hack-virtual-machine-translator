package translators_test

import (
	"github.com/tkech17/hach_virtual_machine_translator/translators"
	"github.com/tkech17/hach_virtual_machine_translator/utils/tests"
	"testing"
)

func TestNew(t *testing.T) {
	tests.AssertNotNull(t, translators.New())
}
