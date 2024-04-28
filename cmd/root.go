package cmd

import (
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/cmd/musician"
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
)

// rootCommandFlagSet contains global flags for this program
type rootCommandFlagSet struct {
	dateFlag string
}

// rootCommandFlags implements the global flags for this prompts
var rootCommandFlags rootCommandFlagSet

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
	rootCmd.PersistentFlags().StringVarP(&rootCommandFlags.dateFlag, "date", "d", time.Now().Format("2006-01-02"), "date of the event")
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
			os.Exit(130)
		default:
			terminal.PrintError(err)
			os.Exit(1)
		}
	}
}
