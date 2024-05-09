package band

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/band"
	"github.com/zimeg/instant-band-night/pkg/event"
	"github.com/zimeg/instant-band-night/pkg/instrument"
	"github.com/zimeg/instant-band-night/pkg/musician"
)

// bandCommandEditFlagSet contains the flags for this command
type bandCommandEditFlagSet struct {
	addFlag        bool
	bandFlag       int
	instrumentFlag string
	monikerFlag    string
	musicianFlag   string
	removeFlag     bool
}

// bandCommandEditFlags implements the flags for this command
var bandCommandEditFlags bandCommandEditFlagSet

// bandCommandEditActionOptions represents a single edit option
type bandCommandEditActionOptions int

const (
	bandCommandEditActionUndefined bandCommandEditActionOptions = iota
	bandCommandEditActionAdd
	bandCommandEditActionRemove
	bandCommandEditActionSave
)

// IsAdd returns if the add action is selected
func (opt bandCommandEditActionOptions) IsAdd() bool {
	return opt == bandCommandEditActionAdd
}

// IsRemove returns if the remove action is selected
func (opt bandCommandEditActionOptions) IsRemove() bool {
	return opt == bandCommandEditActionRemove
}

// BandCommandEditNew adjusts details about a performing group of musicians
func BandCommandEditNew(event *event.Event) *cobra.Command {
	bandCommandEdit := &cobra.Command{
		Use:   "edit",
		Short: "Edit information about a band",
		Long: strings.Join([]string{
			"Edit which musicians compose a band or details about the actual band.",
			"",
			"All attempts will be made to make the edit happen. Musicians can perform",
			"multiple instruments at the same time if such enthusiasm exists.",
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			return bandCommandEditRunE(cmd, event)
		},
	}
	bandCommandEdit.Flags().BoolVar(&bandCommandEditFlags.addFlag, "add", false, "if adding a musician to the band")
	bandCommandEdit.Flags().IntVar(&bandCommandEditFlags.bandFlag, "band", 0, "the order of the performing group")
	bandCommandEdit.Flags().StringVar(&bandCommandEditFlags.instrumentFlag, "instrument", "", "the instrument being updated")
	bandCommandEdit.Flags().StringVar(&bandCommandEditFlags.monikerFlag, "moniker", "", "the formal title of a band")
	bandCommandEdit.Flags().StringVar(&bandCommandEditFlags.musicianFlag, "musician", "", "the musician ID being changed")
	bandCommandEdit.Flags().BoolVar(&bandCommandEditFlags.removeFlag, "remove", false, "if removing a musician from the band")
	return bandCommandEdit
}

// bandCommandEditRunE makes edits to a band at the event
func bandCommandEditRunE(cmd *cobra.Command, event *event.Event) error {
	band, musician, instrument, action, err := bandCommandEditPrompt(cmd, event)
	if err != nil {
		return err
	}
	switch action {
	case bandCommandEditActionAdd:
		band.Instruments.GetInstrument(instrument).AddMusician(musician.GetID())
		bandCommandEditRunEPrintAdd(band, instrument, musician)
	case bandCommandEditActionRemove:
		band.Instruments.GetInstrument(instrument).RemoveMusician(musician.GetID())
		bandCommandEditRunEPrintRemove(band, instrument, musician)
	case bandCommandEditActionSave:
		bandCommandEditRunEPrintSave(event, band)
	}
	return nil
}

// bandCommandEditRunEPrintAdd outputs a confirmation of addition to the band
func bandCommandEditRunEPrintAdd(
	band band.Band,
	instrument instrument.Instrument,
	musician musician.Musician,
) {
	terminal.PrintInfo(display.Section(display.SectionF{
		Icon:   "star",
		Header: "A musician has joined the band",
		Body: []string{
			fmt.Sprintf(
				"Band #%d has someone new on '%s': %s",
				band.GetID(),
				instrument,
				musician.Name,
			),
		},
	}))
}

// bandCommandEditRunEPrintAdd outputs a confirmation of removal to the band
func bandCommandEditRunEPrintRemove(
	band band.Band,
	instrument instrument.Instrument,
	musician musician.Musician,
) {
	terminal.PrintInfo(display.Section(display.SectionF{
		Icon:   "star",
		Header: "A musician has left the band",
		Body: []string{
			fmt.Sprintf(
				"Band #%d has one fewer performers on '%s': %s",
				band.GetID(),
				instrument,
				musician.Name,
			),
		},
	}))
}

// bandCommandEditRunEPrintSave outputs a message that the band was saved
func bandCommandEditRunEPrintSave(event *event.Event, band band.Band) {
	_, musicians := event.BandMusiciansF(band)
	terminal.PrintInfo(display.Section(display.SectionF{
		Icon: "star",
		Header: fmt.Sprintf(
			"Details have been saved for Band #%d: '%s'",
			band.GetID(),
			band.Moniker,
		),
		Body: musicians,
	}))
}
