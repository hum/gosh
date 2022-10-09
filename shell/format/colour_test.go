package format

import "testing"

func TestSetColourOnText(t *testing.T) {
	expectedOutput := "\x1b[31mtext\x1b[0m"
	output := SetColour(ColorRed, "text")
	if output != expectedOutput {
		t.Fatalf("color set did not match expected output, wanted=%s, got=%s", expectedOutput, output)
	}
}
