package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := strings.NewReader("This is a test input string.")
	expected := 6
	actual := countWords(input)
	if actual != expected {
		t.Errorf("countWords returned %d, expected %d", actual, expected)
	}
}

func TestCountLines(t *testing.T) {
	input := strings.NewReader("Line 1\nLine 2\nLine 3")
	expected := 3
	actual := countLines(input)
	if actual != expected {
		t.Errorf("countLines returned %d, expected %d", actual, expected)
	}
}

func TestCountBytes(t *testing.T) {
	input := bytes.NewBufferString("This is a test input string.")
	expected := 28
	actual := countBytes(input)
	if actual != expected {
		t.Errorf("countBytes returned %d, expected %d", actual, expected)
	}
}

func TestMainFunction(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Count Words",
			args:     []string{"-w", "test.txt"},
			expected: "\t 58164 test.txt\n",
		},
		{
			name:     "Count Lines",
			args:     []string{"-l", "test.txt"},
			expected: "\t 7143 test.txt\n",
		},
		{
			name:     "Count Bytes",
			args:     []string{"-c", "test.txt"},
			expected: "\t 335039 test.txt\n",
		},
		// {
		// 	name:     "Default Case",
		// 	args:     []string{"test.txt"},
		// 	expected: "\t 3 6 33 test.txt\n",
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "main.go")
			cmd.Args = append(cmd.Args, tc.args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				t.Fatalf("execution error: %v", err)
			}
			actual := out.String()
			if actual != tc.expected {
				t.Errorf("got %q, want %q", actual, tc.expected)
			}
		})
	}
}
