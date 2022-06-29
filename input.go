package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func takeInputWithSpaces() string {

	fmt.Println("Enter your name")

	stdInobj := bufio.NewReader(os.Stdin)

	input, err := stdInobj.ReadString('\n')

	if err != nil {
		log.Fatal(err)

	}

	return input

}
