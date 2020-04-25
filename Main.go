package main

import (
	"github.com/tkech17/hach_virtual_machine_translator/utils/files"
	"github.com/tkech17/hach_virtual_machine_translator/vm"
	"log"
	"os"
	"strings"
)

type VirtualMachine interface {
	GenerateAssembly(fileName string, content string) string
}

func main() {
	var translator VirtualMachine = vm.GetVirtualMachine()
	var fileNameContents = getFileNameAndContents()
	for fullFileName, content := range fileNameContents {
		targetFileName := getTargetFileName(fullFileName)
		targetContent := translator.GenerateAssembly(files.GetFileName(fullFileName), content)
		files.SaveContent(targetFileName, targetContent)
	}
}

func getTargetFileName(fileName string) string {
	return strings.Replace(fileName, ".vm", ".asm", 1)
}

func getFileNameAndContents() map[string]string {
	if len(os.Args) != 2 {
		log.Fatal("arguments len must be 2")
	}
	var path = os.Args[1]
	var filesArray map[string]string
	if files.IsDir(path) {
		filesArray = files.GetFilesFromFolderWithSuffix(path, ".vm")
	} else {
		if !strings.HasSuffix(path, ".vm") {
			log.Fatal("file suffix must be .vm")
		}
		content := files.ReadFile(path)
		filesArray = map[string]string{path: content}
	}
	return filesArray
}
