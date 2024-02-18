package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := ""
	for i, v := range os.Args {
		if i == 1 {
			filename = v
		}
	}
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	strtxt, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(strtxt))
}
