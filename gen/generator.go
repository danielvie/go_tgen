// generator.go
package gen

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
//go:embed templates/python/.python-version
var templates embed.FS

func GetFilesInPath(templates embed.FS, root string) []string {
	paths := make([]string, 0)

	err := fs.WalkDir(templates, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !d.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path:  %v", err)
		os.Exit(1)
	}

	return paths
}

func GenFile(filename, templatePath string) error {
	// ensure the directory exists
	filePath := filepath.Dir(filename)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}

	// read the template file content from disk
	t, err := template.ParseFS(templates, templatePath)
	if err != nil {
		return err
	}

	// open or create the file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// execute the template and write the output to the file
	return t.Execute(file, nil)
}

func GenPython() error {

	if err := GenFile("./main.py", "templates/python/main.py"); err != nil {
		fmt.Println(err)
	}

	if err := GenFile("./.python-version", "templates/python/.python-version"); err != nil {
		fmt.Println(err)
	}

	if err := GenFile("./pyproject.toml", "templates/python/pyproject.toml"); err != nil {
		fmt.Println(err)
	}

	if err := GenFile("./Taskfile.yml", "templates/python/Taskfile.yml"); err != nil {
		fmt.Println(err)
	}

	return nil
}

func GenCpp() error {
	err := GenWalkPath("templates/cpp")
	if err != nil {
		return fmt.Errorf("error walking template: %v", err)
	}

	return nil
}

func GenCppLib() error {
	err := GenWalkPath("templates/cpplib")
	if err != nil {
		return fmt.Errorf("error walking template: %v", err)
	}

	return nil
}

func GenCppZig() error {

	err := GenWalkPath("templates/cppzig")
	if err != nil {
		return fmt.Errorf("error walking template: %v", err)
	}

	return nil
}

func GenCMake() error {
	if err := GenFile("./CMakeLists.txt", "templates/cmake/CMakeLists.txt"); err != nil {
		fmt.Println(err)
	}

	return nil
}

func GenCppTask() error {
	if err := GenFile("./Taskfile.yml", "templates/cpp/Taskfile.yml"); err != nil {
		fmt.Println(err)
	}

	return nil
}

func GenGolang() error {
	if err := GenFile("./Taskfile.yml", "templates/go/Taskfile.yml"); err != nil {
		fmt.Println(err)
	}

	if err := GenFile("./main.go", "templates/go/main.go"); err != nil {
		fmt.Println(err)
	}

	return nil
}

func GenWalkPath(root string) error {
	// Walk through the embedded templates
	err := fs.WalkDir(templates, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Create destination path by removing the templates/react_ts prefix
		destPath := strings.TrimPrefix(path, root)
		destPath = "./" + destPath

		fmt.Printf("Generating: %s -> %s\n", path, destPath)

		if err := GenFile(destPath, path); err != nil {
			fmt.Printf("Error generating %s: %v\n", destPath, err)
			return err
		}

		return nil
	})

	return err
}

func GenReact() error {

	// create all files from template
	err := GenWalkPath("templates/react_ts")
	if err != nil {
		return fmt.Errorf("error walking templates: %v", err)
	}

	// create `public` folder
	err = os.Mkdir("public", 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	return nil
}
