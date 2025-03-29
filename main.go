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
	tgen python
	tgen cpp
	tgen cppzig
	tgen cmake
	tgen cpptask
	tgen go
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

	switch *option {
	case "python", "py":
		gen.GenPython()
	case "cpp":
		gen.GenCpp()
	case "cppzig":
		gen.GenCppZig()
	case "cmake":
		gen.GenCMake()
	case "cpptask":
		gen.GenCppTask()
	default:
		fmt.Println("name not known!")
		flag.Usage()
	}
}
