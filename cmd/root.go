package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/cmd/musician"
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/event"
)

// rootCommandFlagSet contains global flags for this program
type rootCommandFlagSet struct {
	configFlag string
	dateFlag   string
}

// rootCommandFlags implements the global flags for this prompts
var rootCommandFlags rootCommandFlagSet

// rootCommandNew creates the top level command
func rootCommandNew() *cobra.Command {
	tonight := event.Event{}
	rootCommand := &cobra.Command{
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
	rootCommand.CompletionOptions.DisableDefaultCmd = true
	rootCommand.SilenceErrors = true
	rootCommand.SilenceUsage = true
	rootCommand.AddCommand(musician.MusicianCommandNew(&tonight))
	rootCommand.PersistentFlags().StringVarP(&rootCommandFlags.configFlag, "config", "c", "~/.config/ibn", "path to save data")
	rootCommand.PersistentFlags().StringVarP(&rootCommandFlags.dateFlag, "date", "d", time.Now().Format("2006-01-02"), "date of the event")
	cobra.OnInitialize(func() {
		err := loadConfiguration(rootCommand, &tonight)
		if err != nil {
			terminal.PrintError(err)
			os.Exit(1)
		}
	})
	return rootCommand
}

// loadConfiguration updates the event of tonight with saved values and settings
func loadConfiguration(rootCommand *cobra.Command, tonight *event.Event) error {
	var config string
	if !rootCommand.Flag("config").Changed {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		} else {
			config = fmt.Sprintf("%s/.config/ibn", homedir)
		}
	} else {
		config = rootCommand.Flag("config").Value.String()
	}
	date := rootCommand.Flag("date").Value.String()
	loaded, err := event.LoadEvent(config, date)
	if err != nil {
		return err
	} else {
		*tonight = loaded
	}
	return nil
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
