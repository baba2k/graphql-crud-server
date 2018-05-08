package main

import (
	"os/exec"
	"path/filepath"
	"os"
	"fmt"
	"bytes"
)

const (
	antlrPath   = "lib/antlr-4.7.1-complete.jar"
	grammarPath = "parser/grammar/GraphQL.g4"
	outputPath  = "parser/antlr"
)

func main() {
	generate()
}

func generate() {
	if _, err := os.Stat(antlrPath); os.IsNotExist(err) {
		panic("antlr jar not found")
	}

	if _, err := os.Stat(antlrPath); os.IsNotExist(err) {
		panic("antlr jar not found")
	}

	antlr, err := filepath.Abs(antlrPath)
	if err != nil {
		panic("could not get abs path of antlr jar")
	}
	grammar, err := filepath.Abs(grammarPath)
	if err != nil {
		panic("could not get abs path of grammar")
	}

	cmd := exec.Command("java", "-jar", antlr, "-Dlanguage=Go", grammar, "-o", outputPath)
	if err != nil {
		panic(err)
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		panic(fmt.Sprint(err) + "\n" + stderr.String())
	}

	fmt.Println("ANTLR files generated successfully in dir \"" + outputPath + "\"")
}
