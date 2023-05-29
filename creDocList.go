// creDoc.go
// program that creates an md file containing the header description of each go file in the working directory.
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
	"time"
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
	timStr := fmt.Sprintf("created on %s\n\n", time.Now().Format(time.RFC1123))

	outfil.WriteString(timStr)


	for _, fil := range files {
		nam := fil.Name()
		idx := strings.Index(nam, ".go")
		if idx > 0 {
//			fmt.Println(nam)
			namStr := string(nam[:idx])
			outfil.WriteString("## " + namStr +"\n\n")
			err = WriteFileHeader(outfil, nam)
			if err != nil {log.Fatalf("WriteFileHeader: %v\n", err)}
		}
	}
	log.Printf("*** End CreDocList ***\n")
}

func WriteFileHeader(outfil *os.File, filnam string) (err error) {

	log.Printf("file name: %s\n", filnam)
	offset := int64(0)
	lin := make([]byte, 200)
	infil, err := os.Open(filnam)
	if err != nil {return fmt.Errorf("os.Open:%v\n", err)}
	defer infil.Close()
	outStr :=""
	pidx := -1
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

		pidx = bytes.Index(lin[:eol+1],[]byte("package"))
//		fmt.Printf("pidx: %d, %s\n", pidx, string(lin[:eol+1]))

		if pidx > -1 {break}

//		fmt.Printf("eol: %d, lin: %s\n", eol, string(lin[:eol+1]))
		outStr += string(lin[ist:eol+1]) +"\n"
		offset += int64(eol+1)
	}

	useStr :=""
	for i:=0; i< 20; i++ {
		n, err :=infil.ReadAt(lin,offset)
//		fmt.Printf("eol: %d, lin: %s\n", offset, string(lin))
//		if err != nil {return fmt.Errorf("os.Read:%v\n", err)}
		if err != nil {break}
		eol := -1
		for j:=0; j< n; j++ {
			if lin[j] == '\n' {
				eol = j
				break
			}
		}
		if eol<0 {return fmt.Errorf("EOL not found!\n")}

		idx := bytes.Index(lin[:eol+1], []byte("useStr"))
		if idx > -1 {
			endidx := eol
			stidx := idx+7
			state :=1

			for k:=idx+7; k<eol+1; k++ {
				switch lin[k] {
				case '"':
					switch state {
					case 1:
						stidx = k+1
						state = 2
					case 2:
						endidx = k
						state = 0
					default:
						return fmt.Errorf("parsing useStr: unknown state")
				}
				default:
				}
				if state == 0 {break}
			}

//fmt.Printf("useStr: %s\n", string(lin[idx:eol+1]))

//fmt.Printf("idx: %d stidx: %d endidx: %d eol: %d str: %s\n", idx, stidx, endidx, eol, string(lin[stidx: endidx]))

			useStr = string(lin[stidx: endidx]) +"\n\n"
			break
		}
		offset += int64(eol+1)
	}

	if len(useStr) > 0 {outfil.WriteString(useStr)}
	

	outfil.WriteString(outStr)

	return nil
}
