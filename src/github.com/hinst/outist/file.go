package outist

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var AppDirectory, _ = filepath.Abs(filepath.Dir(os.Args[0]))

func WriteStringToFile(filePath, text string) error {
	var data = []byte(text)
	var result = ioutil.WriteFile(filePath, data, os.ModePerm)
	return result
}

func ReadStringFromFile(filePath string) string {
	var data, _ = ioutil.ReadFile(filePath)
	var text = string(data)
	return text
}
