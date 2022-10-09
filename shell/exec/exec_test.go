package exec

import (
	"testing"

	"github.com/hum/gosh/shell/env"
)

func TestExecutablesProcessPath(t *testing.T) {
	if len(Executables) > 0 {
		t.Fatalf("map Executables is not empty, got=%v", Executables)
	}
	env.LoadEnv()
	err := LoadExecutablesFromPath()
	if err != nil {
		t.Fatalf("did not load executables from the default $PATH, got err=%s", err)
	}
	if len(Executables) == 0 {
		t.Fatalf("map Executables is empty after load from $PATH")
	}
}
