package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	password, input string
	tip             string
}{
	{password: "", input: "", tip: ""},
	{password: "secret", input: "secret", tip: ""},
	{password: "a", input: "b", tip: "a"},
	{password: "foobar", input: "fo", tip: "foo."},
	{password: "foobar", input: "o", tip: "fo."},
	{password: "bar", input: "bax", tip: ".ar"},
	{password: "bar", input: "bbr", tip: "bar"},
	{password: "foobar", input: "fooaar", tip: ".oba."},
	{password: "foobär", input: "fooöar", tip: ".obä."},
}

func TestHint(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprintf("'%s'+'%s'", tc.password, tc.input), func(t *testing.T) {
			tip := hint([]rune(tc.password), []rune(tc.input))
			if tip != tc.tip {
				t.Errorf("want '%s'; got '%s'", tc.tip, tip)
			}
		})
	}
}

func FuzzHint(f *testing.F) {
	for _, tc := range tests {
		f.Add(tc.password, tc.input)
	}

	f.Fuzz(func(t *testing.T, password, input string) {
		_ = hint([]rune(password), []rune(input))
	})
}
