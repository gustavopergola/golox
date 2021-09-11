package compiler

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavopergola/golox/src/scanner"
)

func Execute() {
	var err error
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		err = runFile(args[0])
	} else {
		err = runPrompt()
		if err != nil {
			os.Exit(65)
		}
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func runFile(path string) error {
	inputFile, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Could not read file %s", path)
	}
	fmt.Print(string(inputFile))
	return nil
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		bytesLine, _, err := reader.ReadLine()
		line := string(bytesLine)
		if err != nil || line == "exit" {
			return err
		}
		run(line)
	}
}

func run(source string) {
	scanner.ScanFile(source)
}

func errorHandling(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where string, msg string) {
	fmt.Printf("[line %d] Error %s: %s", line, where, msg)
}
