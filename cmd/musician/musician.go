package musician

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/pkg/event"
	"github.com/zimeg/instant-band-night/pkg/musician"
)

// Musician contains information about the performer
type Musician = musician.Musician

// MusicianCommandNew contains commands centered around the musician
func MusicianCommandNew(event *event.Event) *cobra.Command {
	musicianCommand := &cobra.Command{
		Use:     "musician",
		Aliases: []string{"bucket"},
		Short:   "Makers of music",
		Long: strings.Join([]string{
			"Individual instrumentation for the makers of music",
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	musicianCommand.AddCommand(MusicianCommandJoinNew(event))
	return musicianCommand
}
