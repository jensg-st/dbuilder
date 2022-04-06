package dbuild

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed assets/schema.json
var schema []byte

//go:embed assets/input.yaml
var input []byte

//go:embed assets/Dockerfile
var dockerFile []byte

//go:embed assets/generate.go
var genFile []byte

//go:embed assets/main.templ.go
var mainFile []byte

const (
	defaultDir = "app"
)

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initialize a direktiv container app",
		Run: func(cmd *cobra.Command, names []string) {
			if err := build(); err != nil {
				log.Fatalln(fmt.Errorf("dbuild init: %v", err))
			}
		},
	}
	return cmd
}

func build() error {

	files := []struct {
		name string
		data []byte
	}{
		{
			"schema.json",
			schema,
		},
		{
			"input.yaml",
			input,
		},
		{
			"Dockerfile",
			dockerFile,
		},
	}

	for a := range files {

		f := files[a]
		if err := os.WriteFile(f.name,
			[]byte(f.data), 0644); err != nil {
			return fmt.Errorf("creating %s file: %w", f.name, err)
		}

	}

	if err := os.MkdirAll(defaultDir, os.ModePerm); err != nil {
		return fmt.Errorf("creating application directory: %w", err)
	}

	if err := os.WriteFile(filepath.Join(defaultDir, "generate.go"),
		[]byte(genFile), 0644); err != nil {
		return fmt.Errorf("creating generate.go file: %w", err)
	}

	return nil

}
