package files

import (
	"io/ioutil"
	"log"
	"strings"
)

func IsDir(path string) bool {
	if path[len(path)-1] == '/' || strings.LastIndex(path, ".") == -1 {
		return true
	}
	return false
}

func CountOfFilesWithSuffix(folder string, suffix string) int {
	filesArray, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	var count = 0
	for _, file := range filesArray {
		fileName := file.Name()
		if strings.HasSuffix(fileName, suffix) {
			count++
		}
	}
	return count
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

func SaveContent(fileName string, content string) {
	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func GetFileName(fullFileName string) string {
	indexOfLastSlash := strings.LastIndex(fullFileName, "/")
	if indexOfLastSlash == -1 {
		return fullFileName
	} else {
		return fullFileName[indexOfLastSlash+1:]
	}
}
