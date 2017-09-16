package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	Purple = "\033[0;35m"
	Green  = "\033[0;32m"
	Red    = "\033[0;31m"
	Blue   = "\033[0;34m"
  Orange = "\033[0;33m"
)

func main() {
	color := make(map[string]string)

	color["main"] = Blue

	color["break"] = Purple
	color["default"] = Purple
	color["func"] = Purple
	color["interface"] = Purple
	color["select"] = Purple
	color["case"] = Purple
	color["defer"] = Purple
	color["go"] = Purple
	color["map"] = Purple
	color["struct"] = Purple
	color["chan"] = Purple
	color["else"] = Purple
	color["goto"] = Purple
	color["package"] = Purple
	color["switch"] = Purple
	color["const"] = Purple
	color["fallthrough"] = Purple
	color["if"] = Purple
	color["range"] = Purple
	color["type"] = Purple
	color["continue"] = Purple
	color["for"] = Purple
	color["import"] = Purple
	color["return"] = Purple
	color["var"] = Purple

	// types
	color["int"] = Orange
	color["int8"] = Orange
	color["int16"] = Orange
	color["int32"] = Orange
	color["int64"] = Orange
	color["uint"] = Orange
	color["unit8"] = Orange
	color["uint16"] = Orange
	color["uint32"] = Orange
	color["uint64"] = Orange
	color["uintptr"] = Orange
	color["float32"] = Orange
	color["float64"] = Orange
	color["complex128"] = Orange
	color["complex64"] = Orange
	color["bool"] = Orange
	color["byte"] = Orange
	color["rune"] = Orange
	color["string"] = Orange
	color["error"] = Orange

	// functions
	color["make"] = Blue
	color["len"] = Blue
	color["cap"] = Blue
	color["new"] = Blue
	color["append"] = Blue
	color["copy"] = Blue
	color["close"] = Blue
	color["delete"] = Blue
	color["complex"] = Blue
	color["real"] = Blue
	color["imag"] = Blue
	color["panic"] = Blue
	color["recover"] = Blue

	file := os.Args[1]
	dat, err := ioutil.ReadFile(file)
	check(err)
	words := strings.Split(string(dat), " ")
	fmt.Println(len(words))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
 - read file
 - split by word
 - identify word and color
 - concatenate color and word
*/
