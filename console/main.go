package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"project/lib"
	"project/lib/interfaces"
)

var (
	generate      = flag.Bool("generate", false, "Generate structs")
	generateClean = flag.Bool("cleanGen", false, "Clean generated structs")
)

func main() {
	flag.Parse()

	if *generate {
		generator()
		return
	}

	if *generateClean {
		generatorClean()
		return
	}

	fmt.Println("Nothing to do")
}

const genFolder = "./gen/"

func generator() {
	files := lib.GetListOfModuleFolders()

	for _, f := range files {
		name := f.Name()

		modelsDir := fmt.Sprintf("%s/%s/models/", interfaces.ModulesDir, name)

		filepath.Walk(modelsDir, copyToGen)

	}
}

func generatorClean() {
	filepath.Walk("./"+genFolder+"models/", rmGenModels)
}

func copyToGen(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		rmcmd := exec.Command("cp", "./"+path, genFolder+"models/.")

		err := rmcmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	return err
}

func rmGenModels(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		rmcmd := exec.Command("rm", "./"+path)

		err := rmcmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	return err
}
