// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = []struct {
		name string
		input string
		want int
	} {
		{"zero", "#00", 0},
		{"one", "a2_32_^&/)", 1},
		{"two", "812_%6ab//", 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := letterCounter.count(letterCounter{identifier: tc.name}, tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
  var tests = []struct {
		name string
		input string
		want int
	} {
		{"#00", "#00", 2},
		{"abc_1,?/", "abc_1,?/", 1},
		{"abc_12&8_^", "abc_12&8_^", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := numberCounter.count(numberCounter{designation: tc.name}, tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}

}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
  var tests = []struct {
		name string
		input string
		want int
	} {
		{"#00", "#00", 1},
		{"abc_1,?/", "abc_1,?/", 4},
		{"abc_12&8_^", "abc_12&8_^", 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := symbolCounter.count(symbolCounter{label: tc.name}, tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
