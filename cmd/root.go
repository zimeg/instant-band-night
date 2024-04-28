package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/cmd/musician"
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
	rootCmd.AddCommand(musician.MusicianCommandNew())
	return rootCmd
}

// Execute runs the root command of the program
func Execute() {
	if err := rootCommandNew().Execute(); err != nil {
		ibnerr := errors.ToIBNError(err)
		switch ibnerr.Code {
		case errors.ErrPromptInterrupt.Code:
			if ibnerr.Message != "" {
				terminal.PrintError(err)
			}
		default:
			terminal.PrintError(err)
		}
	}
}
