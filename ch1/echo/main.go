package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(naiveEcho(os.Args[1:]))
	fmt.Println(libEcho(os.Args[1:]))
}

func naiveEcho(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}


func libEcho(args []string) string {
	return strings.Join(args, " ")
}
