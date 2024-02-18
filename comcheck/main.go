package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for _, str := range arg {
		if str == "01" || str == "galaxy" || str == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
