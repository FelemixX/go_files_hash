package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
)

var fileHashes = make(map[string][]string)

func hashFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash := sha1.Sum(data)
	hashKey := fmt.Sprintf("%x", hash)
	fileHashes[hashKey] = append(fileHashes[hashKey], path)

	fmt.Println(hashKey)
}

func ifNotIsDir(path string, f os.FileInfo, _ error) error {
	if !f.IsDir() {
		hashFile(path)
	}

	return nil
}

func searchFile(hashKey string) {
	paths, found := fileHashes[hashKey]
	if found {
		fmt.Println("Files found:")
		for _, path := range paths {
			fmt.Println(path)
		}

		return
	}

	fmt.Println("No files found with the given hash")
}

func main() {
	root := "." // заменить на нужный каталог
	err := filepath.Walk(root, ifNotIsDir)
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
		return
	}

	var hashKey string
	fmt.Print("Enter file hash to search: ")
	_, err = fmt.Scanln(&hashKey)
	if err != nil {
		panic(err)
	}

	searchFile(hashKey)
}
