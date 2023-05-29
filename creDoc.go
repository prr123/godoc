// creDoc.go
// program that creates an md file from
// author: prr azul software
// date: 29 May 2023
// copyright 2023 prr, azul software
//

package main

import (
	"os"
	"fmt"
	"log"
)

func main() {

	numArgs := len(os.Args)

	useStr:= "creDirList [-help]"
	if numArgs > 2 {
		fmt.Printf("usage: %s\n", useStr)
		log.Fatalf("too many args!")
	}

	if numArgs == 2 {
		if os.Args[1]== "-help" {
			fmt.Printf("usage: %s\n", useStr)
			os.Exit(1)
		}
	}

	log.Printf("creDocList start\n")

}
