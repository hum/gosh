package env

import (
	"os"
	"testing"
)

func TestGetEnvironmentalValues(t *testing.T) {
	if len(env) > 0 {
		t.Fatalf("map env is not empty, got=%v", env)
	}

	key, val := "GOSH_TEST_ENV_VAL", "gosh"
	os.Setenv(key, val)

	LoadEnv()

	v := Get(key)
	if v != val {
		t.Fatalf("loaded env val %s does not match expected value, wanted=%s, got=%s", key, val, v)
	}
}

func TestSetEnvironmentalValues(t *testing.T) {
	key, val := "GOSH_TEST_ENV_VAL", "gosh"
	Set(key, val)

	v := Get(key)

	if v != val {
		t.Fatalf("could not set and get env value, set=%s, got=%s", val, v)
	}
}
