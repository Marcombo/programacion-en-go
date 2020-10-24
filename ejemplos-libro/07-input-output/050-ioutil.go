package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	tmpDir, err := ioutil.TempDir("", "ioutil")
	exitIfErr(err)

	fileName := path.Join(tmpDir, "ioutil.txt")

	err = ioutil.WriteFile(fileName, []byte("some text\ngoes here."), 0644)
	exitIfErr(err)

	contents, err := ioutil.ReadFile(fileName)
	exitIfErr(err)

	fmt.Println("read from ", fileName, ":")
	fmt.Println(string(contents))
	/*
		Other ioutil functions:
			ReadAll,
			ReadDir,
			ReadFile,
			TempDir,
			TempFile,
			WriteFile
	*/
}

// maybe a convention to use across the book for the sake of simplicity?
func exitIfErr(err error) {
	if err != nil {
		fmt.Println("exiting. Cause: ", err.Error())
		os.Exit(-1)
	}
}
