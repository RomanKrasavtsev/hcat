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
)

func main() {
	color := make(map[string]string)

	color["package"] = Purple
	color["import"] = Purple
	color["var"] = Purple
	color["func"] = Purple
	color["main"] = Blue
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
