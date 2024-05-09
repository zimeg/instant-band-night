package musician

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/event"
)

// MusicianCommandListNew outputs all musicians and instruments
func MusicianCommandListNew(event *event.Event) *cobra.Command {
	musicianCommandList := &cobra.Command{
		Use:   "list",
		Short: "List instruments of musicians",
		Long: strings.Join([]string{
			"Join different instrument buckets as a participating band member.",
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			return musicianCommandListRunE(event)
		},
	}
	return musicianCommandList
}

// musicianCommandListRunE outputs information about all musicians
func musicianCommandListRunE(event *event.Event) error {
	musicians := event.Musicians.GetMusicians()
	if len(musicians) <= 0 {
		terminal.PrintInfo(display.Section(display.SectionF{
			Icon:   "star",
			Header: "No musicicans are performing",
			Body: []string{
				"Join as a musician with 'musician join'",
			},
		}))
	}
	for _, musician := range musicians {
		terminal.PrintInfo(musician.MusicianF())
	}
	return nil
}
