package main_test

import (
	"fmt"
	"github.com/tkech17/hach_virtual_machine_translator/utils/files"
	"github.com/tkech17/hach_virtual_machine_translator/vm"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"testing"
)

const resources = "resources"

type VirtualMachine interface {
	GenerateAssembly(fileName string, content string) string
}

func Test_main(t *testing.T) {
	var translator VirtualMachine = vm.GetVirtualMachine()
	var fileNameContents = getFileNameAndContents()
	for fullFileName, content := range fileNameContents {
		if !strings.HasSuffix(fullFileName, "FibonacciSeries.vm") {
			continue
		}
		targetFileName := getTargetFileName(fullFileName)
		targetContent := translator.GenerateAssembly(files.GetFileName(fullFileName), content)
		files.SaveContent(targetFileName, targetContent)
	}
	fmt.Println("finish")
}

func getTargetFileName(fileName string) string {
	return strings.Replace(fileName, ".vm", ".asm", 1)
}

func getFileNameAndContents() map[string]string {
	filesArray, err := ioutil.ReadDir(resources)
	if err != nil {
		log.Fatal(err)
	}
	var result = make(map[string]string)
	for _, dir := range filesArray {
		if files.IsDir(dir.Name()) {
			res := files.GetFilesFromFolderWithSuffix(resources+"/"+dir.Name(), ".vm")
			MapCopy(result, res)
		}
	}
	return result
}
func MapCopy(dst, src interface{}) {
	dv, sv := reflect.ValueOf(dst), reflect.ValueOf(src)

	for _, k := range sv.MapKeys() {
		dv.SetMapIndex(k, sv.MapIndex(k))
	}
}
