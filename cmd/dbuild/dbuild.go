package main

import (
	"log"
	"tetstme/cmd/internal/dbuild"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "dbuild"}
	cmd.AddCommand(
		dbuild.InitCmd(),
		dbuild.GenerateCmd(),
	)
	_ = cmd.Execute()
}
