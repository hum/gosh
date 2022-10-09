package lexer

import (
	"testing"
)

func TestLexerParseLine(t *testing.T) {
	tests := []struct {
		input        string
		expectedCmd  string
		expectedArgv []string
	}{
		{
			input:        "echo hello world",
			expectedCmd:  "echo",
			expectedArgv: []string{"hello", "world"},
		},
		{
			input:        "echo hello world\n",
			expectedCmd:  "echo",
			expectedArgv: []string{"hello", "world"},
		},
		{
			input:        "echo \"hello world\"",
			expectedCmd:  "echo",
			expectedArgv: []string{"\"hello", "world\""},
		},
		{
			input:        "",
			expectedCmd:  "",
			expectedArgv: []string{},
		},
	}

	for i, tt := range tests {
		cmd, argv, err := Process(tt.input)
		if err != nil {
			t.Fatalf("tests[%d]: processing error: %s\n", i, err)
		}
		if cmd != tt.expectedCmd {
			t.Fatalf("tests[%d]: cmd value mismatch, wanted: %s, received %s\n", i, tt.expectedCmd, cmd)
		}

		if len(argv) != len(tt.expectedArgv) {
			t.Fatalf("tests[%d]: argv length mismatch, wanted: %s - len(%d), received %s - len(%d)\n", i, tt.expectedArgv, len(tt.expectedArgv), argv, len(argv))
		}

		for i, arg := range argv {
			if arg != tt.expectedArgv[i] {
				t.Fatalf("tests[%d]: arg mismatch, wanted: %s, got: %s", i, tt.expectedArgv[i], arg)
			}
		}
	}
}
