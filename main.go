package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
)

var fileHashes = make(map[string]string)

func hashFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash := sha1.Sum(data)
	fileHashes[path] = fmt.Sprintf("%x", hash)
}

func ifNotIsDir(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {
		hashFile(path)
	}

	return nil
}

func searchFileByName(fileName string) {
	for path, _ := range fileHashes {
		if filepath.Base(path) == fileName {
			fmt.Println("File found:", path)
			return
		}
	}
	fmt.Println("File not found")
}

func main() {
	root := "../../" //заменить на нужную папку
	err := filepath.Walk(root, ifNotIsDir)
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
		return
	}

	var fileName string
	fmt.Print("Enter file_name.extension to search: ")
	_, err = fmt.Scanln(&fileName)
	if err != nil {
		panic(err)
	}

	searchFileByName(fileName)
}
