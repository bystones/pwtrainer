// Binary pwtrainer is a tiny utility to help memorize passwords.
package main

import (
	"fmt"
	"os"
	"slices"
	"syscall"

	"golang.org/x/term"
)

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func main() {
	fmt.Println("Enter the correct password:")
	input, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fatalf("failed to read input: %v", err)
	}
	password := []rune(string(input))

	for {
		fmt.Println("Enter the password (or nothing to stop):")
		input, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fatalf("failed to read input: %v", err)
		}
		if len(input) == 0 {
			return
		}

		if tip := hint(password, []rune(string(input))); tip == "" {
			fmt.Println("Passwords match.")
		} else {
			fmt.Printf("Difference near %s\n", tip)
		}
	}
}

// hint computes a hint how the entered input password differs from the correct
// password. If both slices are the same an empty string is returned.
func hint(password, input []rune) string {
	if slices.Equal(password, input) {
		return ""
	}

	length := min(len(password), len(input))
	var diff int
	for diff = range length {
		if input[diff] != password[diff] {
			break
		}
	}
	if diff == length {
		diff--
	}

	atIndex := func(i int) string {
		if i < 0 || i >= len(password) {
			return ""
		}
		return string(password[i])
	}
	hasIndex := func(i int) bool {
		return i >= 0 && i < len(password)
	}

	delta := atIndex(diff-1) + atIndex(diff) + atIndex(diff+1)
	if hasIndex(diff - 2) {
		delta = "." + delta
	}
	if hasIndex(diff + 2) {
		delta = delta + "."
	}
	return delta
}
