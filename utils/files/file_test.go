package files_test

import (
	"github.com/tkech17/hach_virtual_machine_translator/utils/files"
	"github.com/tkech17/hach_virtual_machine_translator/utils/tests"
	"testing"
)

func TestIsDir_when_IsDirectory(t *testing.T) {
	var folders = [2]string{"folder", "folder/"}

	for _, path := range folders {
		tests.AssertTrue(t, files.IsDir(path), path + " Is Not Folder")
	}
}
func TestIsDir_when_IsnOTDirectory(t *testing.T) {
	var filesArray = [2]string{"folder.VM", "folder/BLA.JPG"}

	for _, path := range filesArray {
		tests.AssertFalse(t, files.IsDir(path), path + " Is Folder")
	}
}
