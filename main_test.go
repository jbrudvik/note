package main_test

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const executableName string = "note"
const helpOutput string = `NAME:
   note - Append to latest Apple Notes note

USAGE:
   note [global options] "Text to append"

VERSION:
   (devel)

DESCRIPTION:
   Ignores shared notes. Formats as new line by default.

GLOBAL OPTIONS:
   --bulleted, -b  Format text as bulleted list item (default: false)
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)
`

type executableTestCase struct {
	args             []string
	expectedExitCode int
	expectedStdout   string
	expectedStderr   string
}

func TestExecutable(t *testing.T) {
	// Test cases
	tests := []executableTestCase{
		{
			args:             nil,
			expectedExitCode: 1,
			expectedStdout:   helpOutput,
			expectedStderr:   "",
		},
		{
			args:             []string{"-h"},
			expectedExitCode: 0,
			expectedStdout:   helpOutput,
			expectedStderr:   "",
		},
	}
	for _, tc := range tests {
		testName := strings.Join(tc.args, " ")
		t.Run(testName, func(t *testing.T) {
			runExecutableTestCase(t, tc)
		})
	}
}

func runExecutableTestCase(t *testing.T, tc executableTestCase) {
	// Create a temporary test dir (automatically cleaned up)
	tempTestDir := t.TempDir()

	// Build executable
	buildCmd := exec.Command("go", "build", "-o", tempTestDir)
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Unable to `go build` %s: $ %s\n", executableName, buildCmd)
	}
	executablePath := filepath.Join(tempTestDir, executableName)

	// Run executable and test outputs
	cmd := exec.Command(executablePath, tc.args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			assert.Equal(t, tc.expectedExitCode, exitError.ExitCode(), "exit code")
		} else {
			t.Fatalf("Unable to parse exit code from: %s", err.Error())
		}
	}
	assert.Equal(t, tc.expectedStdout, stdout.String(), "stdout")
	assert.Equal(t, tc.expectedStderr, stderr.String(), "stderr")
}
