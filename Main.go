package main

import (
	"fmt"
	"github.com/tkech17/hach_virtual_machine_translator/utils/files"
	"github.com/tkech17/hach_virtual_machine_translator/vm"
	"log"
	"os"
	"strings"
)

type VirtualMachine interface {
	GenerateAssembly(fileName string, content string, dir bool) string
}

func main() {
	checkArgs()
	var path = os.Args[1]
	var translator VirtualMachine = vm.GetVirtualMachine()
	var fileNameContents = getFileNameAndContents(path)
	init := files.IsDir(path)
	init2 := files.IsDir(path)
	var result string

	//var keys []string
	//for key := range fileNameContents {
	//	keys = append(keys, key)
	//}
	//
	//for i := len(keys) - 1; i >= 0; i-- {
	//	fullFileName := keys[i]
	//	content := fileNameContents[fullFileName]
	//	targetFileName := getTargetFileName(fullFileName)
	//	targetContent := translator.GenerateAssembly(files.GetFileName(fullFileName), content, init)
	//	result += targetContent
	//	files.SaveContent(targetFileName, targetContent)
	//	init = false
	//}

	for fullFileName, content := range fileNameContents {
		targetFileName := getTargetFileName(fullFileName)
		targetContent := translator.GenerateAssembly(files.GetFileName(fullFileName), content, init)
		result += targetContent
		files.SaveContent(targetFileName, targetContent)
		init = false
	}
	if init2 {
		folderName := getFolderName(path)
		targetFile := path + "/" + folderName + ".asm"
		files.SaveContent(targetFile, result)
	}
	fmt.Println("b")
}

func getFolderName(path string) string {
	lastIndexOfSlash := strings.LastIndex(path, "/")
	return path[lastIndexOfSlash+1:]
}

func checkArgs() {
	if len(os.Args) != 2 {
		log.Fatal("arguments len must be 2")
	}
}

func getTargetFileName(fileName string) string {
	lastSlashIndex := strings.LastIndex(fileName, "/")
	folder := fileName[:lastSlashIndex]
	lastSlashIndex = strings.LastIndex(folder, "/")
	folderName := folder[lastSlashIndex+1:]
	return folder + "/" + folderName + ".asm"
}

func getFileNameAndContents(path string) map[string]string {
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
