package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Привіт, світе!")
	fmt.Println(12)
	
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
