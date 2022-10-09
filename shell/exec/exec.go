package exec

import (
	"fmt"
	"os"
	"strings"

	"github.com/hum/gosh/shell/env"
)

var (
	Executables map[string]string = map[string]string{}
)

func LoadExecutablesFromPath() error {
	for _, dir := range strings.Split(env.Get("PATH"), ":") {
		de, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		for _, entry := range de {
			if entry.IsDir() {
				// TODO: handle recursively later
				continue
			}
			name := entry.Name()
			path := fmt.Sprintf("%s/%s", dir, name)
			Executables[name] = path
		}
	}
	return nil
}
