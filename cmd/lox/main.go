package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brewinski/crafting-interpreters/pkg/scanner"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(63)
	} else if len(args) == 2 {
		err := runFile(args[1])
		if err != nil {
			fmt.Println("error runFile: ", err)
			os.Exit(1)
		}
	} else {
		runPrompt(os.Stdin, os.Stdout)
	}
}

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		// TODO: handle error properly.
		fmt.Println("error in runFile: ", err)
		return err
	}

	run(string(bytes))

	return nil
}

func runPrompt(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewReader(reader)
	for {
		fmt.Printf("> ")
		line, _, err := scanner.ReadLine()
		if err == io.EOF {
			fmt.Println()
			os.Exit(0)
		}
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("failed to read line %v.", err)))
		}

		run(string(line))
	}
}

func run(sourceFile string) error {
	scanner := scanner.NewScanner(sourceFile)

	for _, token := range scanner.ScanTokens() {
		fmt.Printf("[RUN] token: '%s', line: '%d', lexeme: '%s', literal: '%s' \n", token.TokenType.String(), token.Line, token.Lexeme, token.Literal)
	}

	return nil
}
