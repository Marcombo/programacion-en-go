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

	err = ioutil.WriteFile(fileName, []byte("algo de texto\naqui"), 0644)
	exitIfErr(err)

	contents, err := ioutil.ReadFile(fileName)
	exitIfErr(err)

	fmt.Println("leido desde ", fileName, ":")
	fmt.Println(string(contents))
}

func exitIfErr(err error) {
	if err != nil {
		fmt.Println("Fallo: ", err.Error())
		os.Exit(-1)
	}
}
