// creDoc.go
// program that creates an md file describing each type, function and method of a go file
// author: prr azul software
// date: 29 May 2023
// copyright 2023 prr, azul software
//

package main

import (
	"os"
	"fmt"
	"log"
	"bytes"
)

func main() {

	var inFilnam string
	numArgs := len(os.Args)

	useStr:= "creDoc [-help/ file.go]"
	if numArgs > 2 {
		fmt.Printf("usage: %s\n", useStr)
		log.Fatalf("too many args!")
	}

	if numArgs < 2 {
		fmt.Printf("usage: %s\n", useStr)
		log.Fatalf("insufficient args!")
	}

	if numArgs == 2 {
		if os.Args[1]== "-help" {
			fmt.Printf("usage: %s\n", useStr)
			os.Exit(1)
		}
		inFilnam = os.Args[1]
	}

	// check file name
	inFilByt := []byte(inFilnam)
	idx := bytes.Index(inFilByt, []byte(".go"))
	if idx == -1 {
		fmt.Printf("usage: %s\n", useStr)
		log.Fatalf("cli file %s not a go file!", inFilnam)
	}


	data, err := os.ReadFile(inFilnam)
	if err != nil {log.Fatalf("os.ReadFile: %v", err)}
	log.Printf("read file: %s size: %d\n", inFilnam, len(data))

	log.Printf("creDoc start parsing\n")

	log.Printf("creDoc end parsing\n")
}
