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
		`
add an $option to generate the files. 
examples:
        >> tgen python 
             (template of python)
        >> tgen cpp 
             (template of cpp)
        >> tgen cppz 
             (template of cpp with zig build)
        >> tgen cmake 
             (template of cmake)
        >> tgen cpptask 
             (template of cpptask)
        >> tgen go 
             (template of golang)
        >> tgen react 
             (template of react app with vite)
        >> tgen task
             (template of Taskfile.yml)
`)

	flag.Parse()

	*option = os.Args[1]

	switch *option {
	case "python", "py":
		gen.GenPython()
	case "cpp":
		gen.GenCpp()
	case "cpplib":
		gen.GenCppLib()
	case "cppz":
		gen.GenCppZig()
	case "cmake":
		gen.GenCMake()
	case "cpptask":
		gen.GenCppTask()
	case "go":
		gen.GenGolang()
	case "react":
		gen.GenReact()
	case "task":
		gen.GenTask()
	case "static":
		gen.GenStaticServerBun()
	default:
		fmt.Println("name not known!")
		flag.Usage()
	}
}
