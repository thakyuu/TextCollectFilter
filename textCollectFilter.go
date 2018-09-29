package main

import (
	"os"
	"flag"
	"strings"
	"io/ioutil"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/japanese"
)

const (
	Space   = " "
	NewLine = "\n"
)



func main() {
	flag.Parse()
	var s string = ""
	
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}

	if(strings.Contains(s, `誤認識 登録`) || strings.Contains(s, `ご認識 登録`) || strings.Contains(s, `小錦 登録`)) {
		os.Stdout.WriteString(ConvertShiftJIS(""))
		writeNeedCollectWords()
	} else {
		os.Stdout.WriteString(ConvertShiftJIS(s))
		writeBeforeWords(s)
	}
}

func writeBeforeWords(s string) {
	file, err := os.Create("./tmp_beforewords.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	_, err = file.WriteString(s)
	if err != nil {
		panic(err)
	}
}

func writeNeedCollectWords() {
	file, err := os.OpenFile("./000_誤認識リスト.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 600)
	
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	_, err = file.WriteString(readBeforeWords() + NewLine)
	if err != nil {
		panic(err)
	}
}

func readBeforeWords() (string) {
	data, err := ioutil.ReadFile("./tmp_beforewords.txt")
	if err != nil {
		panic(err)
	}
	return(string(data))
}

func ConvertShiftJIS(str string) (string) {
	text, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), str)
	if err != nil {
		panic(err)
	}
	return text
}