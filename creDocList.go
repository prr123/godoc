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
	"strings"
	"bytes"
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

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf("ReadDir: %v\n", err)
	}

	outfil,err := os.Create("GoDir.md")
	if err != nil {log.Fatalf("Create GoDir.md: %v\n", err)}
	defer outfil.Close()

	outfil.WriteString("# GoDir\n\n")

	for _, fil := range files {
		nam := fil.Name()
		idx := strings.Index(nam, ".go")
		if idx > 0 {
			fmt.Println(nam)
			namStr := string(nam[:idx])
			outfil.WriteString("## " + namStr +"\n\n")
			err = GetFileHeader(outfil, nam)
			if err != nil {log.Fatalf("GetFileHeader: %v\n", err)}
		}
	}
	log.Printf("*** End CreDocList ***\n")
}

func GetFileHeader(outfil *os.File, filnam string) (err error) {

	log.Printf("file name: %s\n", filnam)
	offset := int64(0)
	lin := make([]byte, 200)
	infil, err := os.Open(filnam)
	if err != nil {return fmt.Errorf("os.Open:%v\n", err)}
	defer infil.Close()

	for i:=0; i< 10; i++ {
		_, err =infil.ReadAt(lin,offset)
//		fmt.Printf("eol: %d, lin: %s\n", offset, string(lin))
		if err != nil {return fmt.Errorf("os.Read:%v\n", err)}
		eol := -1
		ist := 0
		istate := 1
		for j:=0; j< 200; j++ {
			switch lin[j] {
			case '\n':
				eol = j
			case '/', ' ':
				if istate == 1 {ist++}
			default:
				istate = 2
			}
			if eol > -1 {break}
		}
		if eol<0 {return fmt.Errorf("EOL not found!\n")}

		pidx := bytes.Index(lin[:eol+1],[]byte("package"))
//		fmt.Printf("pidx: %d, %s\n", pidx, string(lin[:eol+1]))
		if pidx > -1 {break}

//		fmt.Printf("eol: %d, lin: %s\n", eol, string(lin[:eol+1]))
		outfil.WriteString(string(lin[ist:eol+1]))
		offset += int64(eol+1)
	}

	return nil
}
