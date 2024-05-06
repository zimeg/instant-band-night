package band

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/pkg/event"
)

// BandCommandNew contains commands centered around the band
func BandCommandNew(event *event.Event) *cobra.Command {
	bandCommand := &cobra.Command{
		Use:   "band",
		Short: "Ensembles of sound",
		Long: strings.Join([]string{
			"Collections of commands and individuals performers making a collective.",
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	bandCommand.AddCommand(BandCommandCreateNew(event))
	return bandCommand
}
