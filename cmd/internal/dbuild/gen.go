package dbuild

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/a-h/generate"
	"github.com/spf13/cobra"
)

func GenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate [flags]",
		Short: "generates direktiv application",
		// Example: examples(
		// 	"ent init Example",
		// 	"ent init --target entv1/schema User Group",
		// ),
		// Args: func(_ *cobra.Command, names []string) error {
		// 	for _, name := range names {
		// 		if !unicode.IsUpper(rune(name[0])) {
		// 			return errors.New("schema names must begin with uppercase")
		// 		}
		// 	}
		// 	return nil
		// },
		Run: func(cmd *cobra.Command, names []string) {
			if err := buildApp(); err != nil {
				log.Fatalln(fmt.Errorf("direktiv generate: %w", err))
			}
		},
	}
	return cmd
}

func buildApp() error {

	genStructs()

	if err := os.WriteFile("main.go",
		[]byte(mainFile), 0644); err != nil {
		return fmt.Errorf("creating main.go file: %w", err)
	}

	return nil
}

func genStructs() error {

	out, err := os.OpenFile("structs.go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	schemas, err := generate.ReadInputFiles([]string{"../schema.json"}, false)
	if err != nil {
		return err
	}

	g := generate.New(schemas...)

	err = g.CreateTypes()
	if err != nil {
		return err
	}

	var w io.Writer = out
	generate.Output(w, g, "main")

	return nil
}
