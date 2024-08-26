package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"golang.org/x/tools/imports"
)

func printFiles(dir string, files map[string]string) error {
	for file, content := range files {
		if !strings.HasPrefix(file, dir) {
			file = dir + "/" + file
		}

		if strings.HasSuffix(file, ".go") {
			imports.LocalPrefix = "github.com/smartcontractkit"
			rawContent, err := imports.Process(file, []byte(content), nil)
			if err != nil {
				// print an error, but also write the file so debugging the generator isn't a pain.
				fmt.Printf("Error formatting file %s: %s\n", file, err)
			} else {
				content = string(rawContent)
			}
		}

		if err := os.MkdirAll(path.Dir(file), 0600); err != nil {
			return err
		}

		if err := os.WriteFile(file, []byte(content), 0600); err != nil {
			return err
		}
	}

	return nil
}

func capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}
