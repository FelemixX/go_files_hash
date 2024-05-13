package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var fileHashes = make(map[string]string)

func hashFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash := sha1.Sum(data)
	fileHashes[path] = fmt.Sprintf("%x", hash)
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		hashFile(path)
	}
	return nil
}

func searchFile(fileName string) {
	for path, _ := range fileHashes {
		if filepath.Base(path) == fileName {
			fmt.Println("File found:", path)
			return
		}
	}
	fmt.Println("File not found")
}

func main() {
	root := "C:\\Users\\bythe\\GolandProjects\\awesomeProject1" //заменить на нужную папку
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
		return
	}

	var fileName string
	fmt.Print("Enter file name to search: ")
	fmt.Scanln(&fileName)

	searchFile(fileName)
}
