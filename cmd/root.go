package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/cmd/band"
	"github.com/zimeg/instant-band-night/cmd/musician"
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
	ibnevent "github.com/zimeg/instant-band-night/pkg/event"
)

// rootCommandFlagSet contains global flags for this program
type rootCommandFlagSet struct {
	configFlag string
	dateFlag   string
}

// rootCommandFlags implements the global flags for this prompts
var rootCommandFlags rootCommandFlagSet

// rootCommandOptions contains values to persist for the command
type rootCommandOptionSet struct {
	configDir string
	eventDate string
}

// rootCommandNew creates the top level command
func rootCommandNew() *cobra.Command {
	rootCommandOptions := rootCommandOptionSet{}
	event := ibnevent.Event{}
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
	rootCommand.AddCommand(band.BandCommandNew(&event))
	rootCommand.AddCommand(musician.MusicianCommandNew(&event))
	aliases := []*cobra.Command{
		band.BandCommandCreateNew(&event),
		musician.MusicianCommandJoinNew(&event),
	}
	for _, alias := range aliases {
		alias.Hidden = true
		rootCommand.AddCommand(alias)
	}
	rootCommand.PersistentFlags().StringVarP(&rootCommandFlags.configFlag, "config", "c", "~/.config/ibn", "path to save data")
	rootCommand.PersistentFlags().StringVarP(&rootCommandFlags.dateFlag, "date", "d", time.Now().Format("2006-01-02"), "date of the event")
	cobra.OnInitialize(func() {
		err := setRootCommandOptions(rootCommand, &rootCommandOptions)
		if err != nil {
			terminal.PrintError(err)
			os.Exit(1)
		}
		tonight, err := ibnevent.LoadEvent(
			rootCommandOptions.configDir,
			rootCommandOptions.eventDate,
		)
		if err != nil {
			terminal.PrintError(err)
			os.Exit(1)
		} else {
			event = tonight
		}
	})
	cobra.OnFinalize(func() {
		err := ibnevent.SaveEvent(event)
		if err != nil {
			terminal.PrintError(err)
			os.Exit(1)
		}
	})
	return rootCommand
}

// setRootCommandOptions sets the base configurations for the root command setup
func setRootCommandOptions(
	rootCommand *cobra.Command,
	rootCommandOptions *rootCommandOptionSet,
) (
	err error,
) {
	if !rootCommand.Flag("config").Changed {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		} else {
			rootCommandOptions.configDir = fmt.Sprintf("%s/.config/ibn", homeDir)
		}
	} else {
		rootCommandOptions.configDir = rootCommand.Flag("config").Value.String()
	}
	rootCommandOptions.eventDate = rootCommand.Flag("date").Value.String()
	return
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
