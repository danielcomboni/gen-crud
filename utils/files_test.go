package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDoesDirectoryExist(t *testing.T) {
	println(DoesDirectoryExist("../models"))
}

var (
	targetFolder string
	targetFile   string
	searchResult []string
)

func findFile(path string, fileInfo os.FileInfo, err error) error {

	if err != nil {
		fmt.Println(err)
		return nil
	}

	// get absolute path of the folder that we are searching
	absolute, err := filepath.Abs(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if fileInfo.IsDir() {
		fmt.Println("Searching directory ... ", absolute)

		// correct permission to scan folder?
		testDir, err := os.Open(absolute)

		if err != nil {
			if os.IsPermission(err) {
				fmt.Println("No permission to scan ... ", absolute)
				fmt.Println(err)
			}
		}
		testDir.Close()
		return nil
	} else {
		// ok, we are dealing with a file
		// is this the target file?

		// yes, need to support wildcard search as well
		// https://www.socketloop.com/tutorials/golang-match-strings-by-wildcard-patterns-with-filepath-match-function

		matched, err := filepath.Match(targetFile, fileInfo.Name())
		if err != nil {
			fmt.Println(err)
		}

		if matched {
			// yes, add into our search result
			add := "Found : " + absolute
			searchResult = append(searchResult, add)
		}
	}

	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

}

func TestEnsure(t *testing.T) {
	ensureDir("controllers2")
}
