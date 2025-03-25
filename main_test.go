package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	password, input string
	tip             string
}{
	{
		password: "",
		input:    "",
		tip:      "",
	},
	{
		password: "foobar",
		input:    "fo",
		tip:      "oob",
	},
	{
		password: "foobar",
		input:    "o",
		tip:      "fo",
	},
	{
		password: "bar",
		input:    "bax",
		tip:      "ar",
	},
	{
		password: "bar",
		input:    "bbr",
		tip:      "bar",
	},
	{
		password: "foobar",
		input:    "fooaar",
		tip:      "oba",
	},
	{
		password: "foobär",
		input:    "fooöar",
		tip:      "obä",
	},
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
