// Binary pwtrainer is a tiny utility to help memorize passwords.
package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func main() {
	fmt.Println("Enter the correct password:")
	input, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fatalf("failed to read input: %v", err)
	}
	password := []rune(string(input))

	for {
		fmt.Println("Enter the password (or nothing to stop):")
		input, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fatalf("failed to read input: %v", err)
		}
		if len(input) == 0 {
			return
		}

		if tip := hint(password, []rune(string(input))); tip == "" {
			fmt.Println("Passwords match.")
		} else {
			fmt.Printf("Difference near ...%s...\n", tip)
		}
	}
}

// hint computes a hint how the entered input password differs from the correct
// password. If both slices are the same an empty string is returned.
func hint(password, input []rune) string {
	diff := -1
	for i, r := range input {
		if i >= len(password) || r != password[i] {
			diff = i
			break
		}
	}
	if diff == -1 && len(input) == len(password) {
		return ""
	}
	if diff == -1 && len(input) < len(password) {
		diff = len(input)
	}

	var delta []rune
	if diff-1 >= 0 {
		delta = append(delta, password[diff-1])
	}
	delta = append(delta, password[diff])
	if diff+1 < len(password) {
		delta = append(delta, password[diff+1])
	}
	return string(delta)
}
