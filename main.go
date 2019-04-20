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
		pw := []rune(string(input))

		diff := -1
		for i, r := range pw {
			if i >= len(password) || r != password[i] {
				diff = i
				break
			}
		}
		if diff == -1 && len(pw) < len(password) {
			diff = len(pw)
		}

		if diff == -1 {
			fmt.Println("passwords match")
			continue
		}
		var delta []rune
		if diff-1 > 0 {
			delta = append(delta, password[diff-1])
		}
		delta = append(delta, password[diff])
		if diff+1 < len(password) {
			delta = append(delta, password[diff+1])
		}
		fmt.Printf("difference near ...%s...\n", string(delta))
	}
}
