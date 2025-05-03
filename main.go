package main

import (
	"flag"
	"fmt"
	"os"

	"tgen/gen"

	_ "golang.org/x/mod/modfile"
)

func main() {
	option := flag.String("option", "",
		`an option to generate, 
examples:
	tgen python -> template of python
	tgen cpp -> template of cpp
	tgen cppz -> template of cpp with zig build
	tgen cmake -> template of cmake
	tgen cpptask -> template of cpptask
	tgen go -> template of golang
`)

	flag.Parse()

	// Check if the option is provided via the flag
	if *option == "" {
		if len(os.Args) < 2 {
			fmt.Println("Error: options not provided")
			flag.Usage()
		} else {
			*option = os.Args[1] // Get the first argument after the program name
		}
	}
	// bla ble
	switch *option {
	case "python", "py":
		gen.GenPython()
	case "cpp":
		gen.GenCpp()
	case "cppz":
		gen.GenCppZig()
	case "cmake":
		gen.GenCMake()
	case "cpptask":
		gen.GenCppTask()
	case "go":
		gen.GenGolang()
	default:
		fmt.Println("name not known!")
		flag.Usage()
	}
}
