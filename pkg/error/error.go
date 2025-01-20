package interr

import "fmt"

func Error(line, col int, message string) {
	Report(line, col, "", message)
}

func Report(line, col int, where, message string) {
	fmt.Printf("[line: %d:%d] Error %s: %s\n", line, col, where, message)
}
