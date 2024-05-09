package band

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/band"
	"github.com/zimeg/instant-band-night/pkg/event"
)

// bandCommandCreateFlagSet contains the flags for this command
type bandCommandCreateFlagSet struct {
	artFlag      int
	confirmFlag  bool
	cooldownFlag int
	bassFlag     int
	drumsFlag    int
	guitarFlag   int
	nameFlag     string
	otherFlag    int
	vocalsFlag   int
}

// bandCommandCreateFlags implements the flags for this command
var bandCommandCreateFlags bandCommandCreateFlagSet

// BandCommandCreate forms new groups of somewhat random musicians
func BandCommandCreateNew(event *event.Event) *cobra.Command {
	bandCommandCreate := &cobra.Command{
		Use:   "draw",
		Short: "Draw musicians to form a band",
		Aliases: []string{
			"create",
		},
		Long: strings.Join([]string{
			"Collections of commands and individuals performers making a collective.",
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			return bandCommandCreateRunE(cmd, event)
		},
	}
	bandCommandCreate.Flags().IntVar(&bandCommandCreateFlags.artFlag, "artist", 1, "the number drawing band posters")
	bandCommandCreate.Flags().IntVar(&bandCommandCreateFlags.bassFlag, "bass", 1, "the number making lower sounds")
	bandCommandCreate.Flags().BoolVar(&bandCommandCreateFlags.confirmFlag, "confirm", false, "for provided values and the set defaults")
	bandCommandCreate.Flags().IntVar(&bandCommandCreateFlags.cooldownFlag, "cooldown", 1, "the number of performances between")
	bandCommandCreate.Flags().IntVar(&bandCommandCreateFlags.drumsFlag, "drums", 1, "the number often with the beat")
	bandCommandCreate.Flags().IntVar(&bandCommandCreateFlags.guitarFlag, "guitar", 1, "the number playing the six strings")
	bandCommandCreate.Flags().StringVar(&bandCommandCreateFlags.nameFlag, "name", "", "the title of the group performing")
	bandCommandCreate.Flags().IntVar(&bandCommandCreateFlags.otherFlag, "other", 0, "the number with another instrument")
	bandCommandCreate.Flags().IntVar(&bandCommandCreateFlags.vocalsFlag, "vocals", 1, "the number making person noises")
	return bandCommandCreate
}

// bandCommandCreateRunE forms a new band with random and unique musicians
func bandCommandCreateRunE(cmd *cobra.Command, event *event.Event) error {
	arrangement, cooldown, err := bandCommandCreatePrompt(cmd)
	if err != nil {
		return err
	}
	event.FilterCooldown(cooldown)
	band, err := band.NewBand(arrangement, *event.Buckets)
	if err != nil {
		return err
	}
	err = event.SaveBand(band)
	if err != nil {
		return err
	}
	_, musicians := event.BandMusiciansF(band)
	terminal.PrintInfo(display.Section(display.SectionF{
		Icon:   "star",
		Header: "The next band can take the stage",
		Body:   musicians,
	}))
	return nil
}
