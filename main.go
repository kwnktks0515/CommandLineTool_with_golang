package main

import (
	"fmt"
	"os"
)

var option, parameter = []string{}, []string{}
var version = "0.1"

func main() {
	if 1 < len(os.Args) {
		for _, i := range os.Args[1:] {
			if '-' == i[0] {
				option = append(option, i)
			} else {
				parameter = append(parameter, i)
			}
		}
	}
	if 0 == len(parameter) {
		parameter = append(parameter, "help")
	}
	pexe()
}

func pexe() {
	switch parameter[0] {
	case "version":
		Cversion(version)
	case "help":
		Chelp(parameter)
	case "hello":
		Chello(parameter[1:])
	case "build":
		Cbuild(parameter[1:])
	default:
		error("run", "unknown command "+parameter[0], "Run 'run help'")
	}
}

func error(source string, s ...string) {
	for i := range s {
		fmt.Printf("%s: %s\n", source, s[i])
	}
}
