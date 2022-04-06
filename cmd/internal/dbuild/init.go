package dbuild

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed assets/schema.yaml
var schema []byte

//go:embed assets/input.yaml
var input []byte

//go:embed assets/Dockerfile
var dockerfile []byte

//go:embed assets/generate.go
var generate []byte

const (
	defaultDir = "direktiv"
)

func InitCmd() *cobra.Command {
	var target string
	cmd := &cobra.Command{
		Use:   "init",
		Short: "initialize a direktiv container app",
		Run: func(cmd *cobra.Command, names []string) {
			if err := build(target); err != nil {
				log.Fatalln(fmt.Errorf("dbuild init: %v", err))
			}
		},
	}
	cmd.Flags().StringVar(&target, "target", defaultDir, "target directory for application")
	return cmd
}

func build(dir string) error {

	appDir := filepath.Join(dir, "app")

	if err := os.MkdirAll(appDir, os.ModePerm); err != nil {
		return fmt.Errorf("creating application directory: %w", err)
	}

	files := []struct {
		name string
		data []byte
	}{
		{
			"schema.yaml",
			schema,
		},
		{
			"input.yaml",
			input,
		},
		{
			"Dockerfile",
			dockerfile,
		},
	}

	for a := range files {

		f := files[a]
		if err := os.WriteFile(filepath.Join(dir, f.name),
			[]byte(f.data), 0644); err != nil {
			return fmt.Errorf("creating %s file: %w", f.name, err)
		}

	}

	if err := os.WriteFile(filepath.Join(appDir, "generate.go"),
		[]byte(generate), 0644); err != nil {
		return fmt.Errorf("creating generate.go file: %w", err)
	}

	return nil

}
