package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCommandNew creates the top level command
func rootCommandNew() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ibn",
		Short: "🎶 Instant Band Night CLI",
		Long: strings.Join([]string{
			"🎶 Instant Band Night CLI",
			"",
			"Pairing random musicians to buckets of instrument for impromptu playing.",
		}, "\n"),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("🎺 doot")
		},
	}
	return rootCmd
}

// Execute runs the root command of the program
func Execute() {
	if err := rootCommandNew().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
