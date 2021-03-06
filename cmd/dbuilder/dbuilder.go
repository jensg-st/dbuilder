package main

import (
	"log"

	dbuild "github.com/jensg-st/dbuilder/cmd/internal/dbuilder"

	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "dbuilder"}
	cmd.AddCommand(
		dbuild.InitCmd(),
		dbuild.GenerateCmd(),
	)
	_ = cmd.Execute()
}
