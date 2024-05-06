package band

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/event"
)

// BandCommandListNew outputs all musicians in performing bands
func BandCommandListNew(event *event.Event) *cobra.Command {
	bandCommandList := &cobra.Command{
		Use:   "list",
		Short: "List musicians that form bands",
		Long: strings.Join([]string{
			"Output information about who performed with who and in which band.",
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			return bandCommandListRunE(event)
		},
	}
	return bandCommandList
}

// bandCommandListRunE prints the band information
func bandCommandListRunE(event *event.Event) error {
	for ii, band := range event.Bands {
		terminal.PrintInfo(display.Section(display.SectionF{
			Icon:   "star",
			Header: fmt.Sprintf("Band #%d", ii),
			Body:   event.BandMusicianF(band),
		}))
	}
	return nil
}
