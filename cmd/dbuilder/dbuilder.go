package main

import (
	"log"

	"github.com/jensg-st/dbuilder/cmd/internal/dbuilder"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "dbuilder"}
	cmd.AddCommand(
		dbuilder.InitCmd(),
		dbuilder.GenerateCmd(),
	)
	_ = cmd.Execute()
}
