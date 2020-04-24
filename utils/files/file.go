package files

import (
	"io/ioutil"
	"log"
	"strings"
)

func IsDir(path string) bool {
	if path[len(path) - 1] == '/' || strings.LastIndex(path, ".") == -1  {
		return true
	}
	return false
}

func GetFilesFromFolderWithSuffix(folder string, suffix string) map[string]string {
	filesArray, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	var result = make(map[string]string)
	for _, file := range filesArray {
		fileName := file.Name()
		if !strings.HasSuffix(fileName, suffix) {
			continue
		}
		fileFullName := folder + "/" + fileName
		content := ReadFile(fileFullName)
		result[fileFullName] = content
	}
	return result
}

func ReadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func SaveContent(fileName string, content string)  {
	//TODO save
}