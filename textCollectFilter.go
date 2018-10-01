package main

import (
	"os"
	"flag"
	"time"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/japanese"
)

const (
	Space   = " "
	NewLine = "\n"
	Tab = "\t"
)

func main() {
	flag.Parse()
	s := ""
	
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}

	writeWords(s)
	os.Stdout.WriteString(ConvertShiftJIS(s))
}

func writeWords(s string) {
	file, err := os.OpenFile("./000_talk.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 600)
	
	if err != nil {
		panic(err)
	}
	
	defer file.Close()
	
	_, err = file.WriteString(time.Now().String() + Tab + s + NewLine)
	if err != nil {
		panic(err)
	}
}

func ConvertShiftJIS(str string) (string) {
	text, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), str)
	if err != nil {
		panic(err)
	}
	return text
}