package main

import (
	"fmt"
	"github.com/tkech17/hach_virtual_machine_translator/vm"
	"github.com/tkech17/hach_virtual_machine_translator/utils/files"
	"log"
	"os"
	"strings"
)

type VirtualMachine interface {
	GenerateAssembly(vmContent string) string
}

func main() {
	var translator VirtualMachine = vm.GetVirtualMachine()
	var fileNameContents map[string]string = getFileNameAndContents()
	for fileName, content := range fileNameContents {
		targetFileName := getTargetFileName(fileName)
		targetContent := translator.GenerateAssembly(content)
		files.SaveContent(targetFileName, targetContent)
	}
	fmt.Println(fileNameContents)
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
		if !strings.HasSuffix(path,".vm") {
			log.Fatal("file suffix must be .vm")
		}
		content := files.ReadFile(path)
		filesArray = map[string]string{path: content}
	}
	return filesArray
}