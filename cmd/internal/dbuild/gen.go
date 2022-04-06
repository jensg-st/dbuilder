package dbuild

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func GenerateCmd() *cobra.Command {
	var target string
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
			// if err := initEnv(target, names); err != nil {
			// 	log.Fatalln(fmt.Errorf("ent/init: %w", err))
			// }
			if err := buildApp(); err != nil {
				log.Fatalln(fmt.Errorf("ent/init: %w", err))
			}
		},
	}
	cmd.Flags().StringVar(&target, "target", defaultDir, "application directory")
	return cmd
}

func buildApp() error {
	fmt.Println("kkk")
	return nil
}
