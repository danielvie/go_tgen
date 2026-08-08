package main

import (
	"flag"
	"fmt"
	"os"

	"tgen/gen"

	_ "golang.org/x/mod/modfile"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <option>\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Available options:")
		fmt.Fprintln(os.Stderr, "  agents   template for AGENTS.md")
		fmt.Fprintln(os.Stderr, "  python   template for Python")
		fmt.Fprintln(os.Stderr, "  py       alias for python")
		fmt.Fprintln(os.Stderr, "  c        template for C")
		fmt.Fprintln(os.Stderr, "  cpp      template for C++")
		fmt.Fprintln(os.Stderr, "  cpplib   template for C++ library")
		fmt.Fprintln(os.Stderr, "  cppz     template for C++ with Zig build")
		fmt.Fprintln(os.Stderr, "  cmake    template for CMake")
		fmt.Fprintln(os.Stderr, "  cpptask  template for C++ task setup")
		fmt.Fprintln(os.Stderr, "  go       template for Go")
		fmt.Fprintln(os.Stderr, "  react    template for React app with Vite")
		fmt.Fprintln(os.Stderr, "  task     template for Taskfile.yml")
		fmt.Fprintln(os.Stderr, "  static   template for Bun static server")
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	option := flag.Arg(0)

	switch option {
	case "agents":
		gen.GenAgents()
	case "python", "py":
		gen.GenPython()
	case "c":
		gen.GenC()
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
