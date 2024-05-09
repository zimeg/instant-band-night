package band

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/band"
	ibnband "github.com/zimeg/instant-band-night/pkg/band"
	"github.com/zimeg/instant-band-night/pkg/event"
	ibninstrument "github.com/zimeg/instant-band-night/pkg/instrument"
	ibnmusician "github.com/zimeg/instant-band-night/pkg/musician"
)

// bandCommandEditPrompt collets band and musician information to change
func bandCommandEditPrompt(
	cmd *cobra.Command,
	event *event.Event,
) (
	ibnband.Band,
	ibnmusician.Musician,
	ibninstrument.Instrument,
	bandCommandEditActionOptions,
	error,
) {
	bandID, moniker, musicianID, instrument, action, err := bandCommandEditPromptFlags(cmd)
	if err != nil {
		return ibnband.Band{}, ibnmusician.Musician{}, "", 0, err
	}
	if !cmd.Flag("band").Changed {
		bandID, err = bandCommandEditPromptBand(event.Bands)
		if err != nil {
			return ibnband.Band{}, ibnmusician.Musician{}, "", 0, err
		}
	}
	band, err := event.Bands.GetBand(bandID)
	if err != nil {
		return ibnband.Band{}, ibnmusician.Musician{}, "", 0, err
	}
	if cmd.Flag("moniker").Changed {
		band.Moniker = moniker
	}
	if !cmd.Flag("musician").Changed && !cmd.Flag("moniker").Changed {
		switch action {
		case bandCommandEditActionAdd:
		default:
			musicianID, instrument, err = bandCommandEditPromptMusicianIDBand(
				event,
				band,
				&action,
			)
			if err != nil {
				return ibnband.Band{}, ibnmusician.Musician{}, "", 0, err
			}
		}
		switch action {
		case bandCommandEditActionAdd:
			musicians := event.Musicians.GetMusicians()
			musicianID, action, err = bandCommandEditPromptMusicianIDEvent(
				musicians,
			)
			if err != nil {
				return ibnband.Band{}, ibnmusician.Musician{}, "", 0, err
			}
		case bandCommandEditActionSave:
			return *band, ibnmusician.Musician{}, "", bandCommandEditActionSave, nil
		}
	}
	musician, err := event.Musicians.GetMusician(musicianID)
	if err != nil {
		return ibnband.Band{}, ibnmusician.Musician{}, "", 0, err
	}
	switch action {
	case bandCommandEditActionAdd:
		if !instrument.Exists() {
			instrument, err = bandCommandEditPromptInstrument(musician.Instruments)
			if err != nil {
				return ibnband.Band{}, ibnmusician.Musician{}, "", 0, err
			}
		}
	}
	if !cmd.Flag("add").Changed && !cmd.Flag("remove").Changed {
		confirm, err := bandCommandEditPromptConfirm(band, action, musician, instrument)
		if err != nil {
			return ibnband.Band{}, ibnmusician.Musician{}, "", 0, err
		}
		if !confirm {
			return ibnband.Band{}, ibnmusician.Musician{}, "", 0,
				errors.ErrPromptInterrupt.WithMessage("No performers were changed")
		}
	}
	return *band, musician, instrument, action, nil
}

// bandCommandEditPromptFlags finds the band and musician to change from flags
func bandCommandEditPromptFlags(cmd *cobra.Command) (
	bandID int,
	moniker string,
	musicianID string,
	instrument ibninstrument.Instrument,
	action bandCommandEditActionOptions,
	err error,
) {
	bandID = terminal.FlagToInt(cmd.Flag("band"))
	moniker = terminal.FlagToString(cmd.Flag("moniker"))
	musicianID = terminal.FlagToString(cmd.Flag("musician"))
	instrumentFlag := terminal.FlagToString(cmd.Flag("instrument"))
	instrument = ibninstrument.NewInstrument(instrumentFlag)
	add := terminal.FlagToBool(cmd.Flag("add"))
	remove := terminal.FlagToBool(cmd.Flag("remove"))
	switch {
	case add && remove:
		return 0, "", "", "", bandCommandEditActionUndefined, errors.ErrConfusedFlags
	case add:
		action = bandCommandEditActionAdd
	case remove:
		action = bandCommandEditActionRemove
	default:
		action = bandCommandEditActionUndefined
	}
	return
}

// bandCommandEditPromptBand determines the band ID to update
func bandCommandEditPromptBand(bands band.Bands) (bandID int, err error) {
	var bandOptions []huh.Option[int]
	bandID = len(bands) - 1
	sorted := slices.Clone(bands)
	slices.Reverse(sorted)
	for _, band := range sorted {
		bandLabel := display.Secondary("Band #%d", band.GetID())
		bandOptions = append(bandOptions, huh.NewOption(bandLabel, band.GetID()))
	}
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Rearrange a band").
				Options(bandOptions...).
				Value(&bandID),
		),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	if err != nil {
		return 0, errors.ToIBNError(err)
	}
	return
}

// bandCommandEditPromptMusicianIDBand gathers a musician ID from a band and
// signals if a new musician should be added or the selected one removed
func bandCommandEditPromptMusicianIDBand(
	event *event.Event,
	band *band.Band,
	action *bandCommandEditActionOptions,
) (
	musicianID string,
	instrument ibninstrument.Instrument,
	err error,
) {
	type musicianActioner struct {
		musicianID string
		instrument ibninstrument.Instrument
		action     bandCommandEditActionOptions
	}
	var musicianAction musicianActioner
	var musicianOptions []huh.Option[musicianActioner]
	for _, musicianID := range band.GetMusicianIDs() {
		musician, instrument, err := event.BandMusician(*band, musicianID)
		if err != nil {
			return "", ibninstrument.NewInstrument(""), err
		}
		musicianOption := musicianActioner{
			musicianID: musicianID,
			instrument: instrument,
			action:     bandCommandEditActionRemove,
		}
		musicianLabel := instrument.InstrumentF(
			"%s %s",
			display.Secondary(musicianID),
			musician.Name,
		)
		musicianOptions = append(
			musicianOptions,
			huh.NewOption(musicianLabel, musicianOption),
		)
	}
	var formFields []huh.Field
	switch *action {
	case bandCommandEditActionUndefined:
		musicianOptions = append(
			musicianOptions,
			huh.NewOption(
				display.Secondary("Add a new musician"),
				musicianActioner{
					action: bandCommandEditActionAdd,
				},
			),
		)
		musicianOptions = append(
			musicianOptions,
			huh.NewOption(
				display.Secondary("Save this setup"),
				musicianActioner{
					action: bandCommandEditActionSave,
				},
			),
		)
		formFields = append(
			formFields,
			huh.NewInput().
				Title("Change a moniker").
				Placeholder(fmt.Sprintf("Band #%d", band.GetID())).
				Value(&band.Moniker),
		)
	}
	formFields = append(
		formFields,
		huh.NewSelect[musicianActioner]().
			Title("Orchestrate an arrangement").
			Options(musicianOptions...).
			Value(&musicianAction),
	)
	err = huh.NewForm(
		huh.NewGroup(formFields...),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	if err != nil {
		return "", ibninstrument.NewInstrument(""), errors.ToIBNError(err)
	}
	*action = musicianAction.action
	return musicianAction.musicianID, musicianAction.instrument, nil
}

// bandCommandEditPromptMusicianIDEvent gathers a musician ID from the event
func bandCommandEditPromptMusicianIDEvent(
	musicians []ibnmusician.Musician,
) (
	musicianID string,
	action bandCommandEditActionOptions,
	err error,
) {
	var musicianOptions []huh.Option[string]
	for _, musician := range musicians {
		label := strings.TrimLeft(
			musician.MusicianF(),
			fmt.Sprintf("%s ", display.Emoji("star")),
		)
		musicianLabel := huh.NewOption(label, musician.GetID())
		musicianOptions = append(musicianOptions, musicianLabel)
	}
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Join the stage").
				Options(musicianOptions...).
				Value(&musicianID),
		),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	if err != nil {
		return "", bandCommandEditActionUndefined, errors.ToIBNError(err)
	}
	return musicianID, bandCommandEditActionAdd, nil
}

// bandCommandEditPromptInstrument gathers an instrument to perform with
func bandCommandEditPromptInstrument(
	instruments ibninstrument.Instruments,
) (
	instrument ibninstrument.Instrument,
	err error,
) {
	var instrumentOptions []huh.Option[ibninstrument.Instrument]
	for _, instrument := range instruments {
		instrumentOptions = append(
			instrumentOptions,
			huh.NewOption(instrument.InstrumentF(""), instrument),
		)
	}
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[ibninstrument.Instrument]().
				Title("Pick an instrument").
				Options(instrumentOptions...).
				Value(&instrument),
		),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	if err != nil {
		return ibninstrument.NewInstrument(""), errors.ToIBNError(err)
	}
	return
}

// bandCommandEditPromptConfirm checks that this selected action is correct
func bandCommandEditPromptConfirm(
	band *band.Band,
	action bandCommandEditActionOptions,
	musician ibnmusician.Musician,
	instrument ibninstrument.Instrument,
) (
	confirmed bool,
	err error,
) {
	var title string
	var description string
	switch action {
	case bandCommandEditActionAdd:
		title = fmt.Sprintf("Is it true that '%s' is joining the group?", musician.Name)
		description = fmt.Sprintf(
			"Band #%d will have a one more performer of %s. The crowd is excited.",
			band.GetID(),
			instrument,
		)
	case bandCommandEditActionRemove:
		title = fmt.Sprintf("Is it true that '%s' is leaving the group?", musician.Name)
		description = fmt.Sprintf(
			"Band #%d will have one fewer performers of %s. The crowd is curious.",
			band.GetID(),
			instrument,
		)
	}
	confirmed = true
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Value(&confirmed).
				Description(description),
		),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	if err != nil {
		return false, errors.ToIBNError(err)
	}
	return
}
