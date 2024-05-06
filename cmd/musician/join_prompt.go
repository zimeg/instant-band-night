package musician

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/instrument"
	"github.com/zimeg/instant-band-night/pkg/musician"
)

// musicianCommandJoinPrompt gathers information about the performing musician
// from flags and interactive input
func musicianCommandJoinPrompt(cmd *cobra.Command) (musician Musician, err error) {
	other, err := musicianCommandJoinPromptFlags(cmd, &musician)
	if err != nil {
		return Musician{}, err
	}
	if musician.Name != "" && len(musician.Instruments) > 0 {
		if ok, _ := musician.Instruments.HasInstrument(instrument.OTHER); ok {
			musician.Instruments.SetOther(instrument.NewInstrument(other))
		}
		return musician, nil
	}
	err = musicianCommandJoinPromptMusician(&musician)
	if err != nil {
		return Musician{}, err
	}
	if ok, _ := musician.Instruments.HasInstrument(instrument.OTHER); ok {
		err := musicianCommandJoinPromptOtherInstrument(&musician, other)
		if err != nil {
			return Musician{}, err
		}
	}
	confirmed, err := musicianCommandJoinPromptConfirm(&musician)
	if err != nil {
		return Musician{}, errors.ToIBNError(err)
	}
	if !confirmed {
		return Musician{}, errors.ErrPromptInterrupt.WithMessage("No musician created")
	}
	return
}

// musicianCommandJoinPromptFlags sets instrument values using command flags and
// returns the string value of the other flag
func musicianCommandJoinPromptFlags(cmd *cobra.Command, musician *musician.Musician) (string, error) {
	if terminal.FlagToBool(cmd.Flag("guitar")) {
		musician.Instruments.AddInstrument(instrument.GUITAR)
	}
	if terminal.FlagToBool(cmd.Flag("bass")) {
		musician.Instruments.AddInstrument(instrument.BASS)
	}
	if terminal.FlagToBool(cmd.Flag("drums")) {
		musician.Instruments.AddInstrument(instrument.DRUMS)
	}
	if terminal.FlagToBool(cmd.Flag("vocals")) {
		musician.Instruments.AddInstrument(instrument.VOCALS)
	}
	if terminal.FlagToBool(cmd.Flag("artist")) {
		musician.Instruments.AddInstrument(instrument.ART)
	}
	if cmd.Flag("other").Changed {
		musician.Instruments.AddInstrument(instrument.OTHER)
	}
	if cmd.Flag("name").Changed {
		if name := terminal.FlagToString(cmd.Flag("name")); name != "" {
			musician.Name = name
		} else {
			return "", errors.ErrMissingInputMusicianName
		}
	}
	return terminal.FlagToString(cmd.Flag("other")), nil
}

// musicianCommandJoinPromptMusician gathers input about a new musician
func musicianCommandJoinPromptMusician(musician *musician.Musician) error {
	instruments := musician.Instruments.GetInstruments()
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Sign a stage name").
				Value(&musician.Name).
				CharLimit(25).
				Validate(func(s string) error {
					if len(strings.TrimSpace(s)) == 0 {
						return errors.ErrMissingInputMusicianName
					}
					return nil
				}).
				Placeholder("Rockstar"),
			huh.NewMultiSelect[instrument.Instrument]().
				Title("Pick some instruments").
				Options(
					huh.NewOption("Guitar", instrument.GUITAR),
					huh.NewOption("Bass", instrument.BASS),
					huh.NewOption("Drums", instrument.DRUMS),
					huh.NewOption("Vocals", instrument.VOCALS),
					huh.NewOption("Art", instrument.ART),
					huh.NewOption("Other", instrument.OTHER),
				).
				Validate(func(s []instrument.Instrument) error {
					if len(s) == 0 {
						return errors.ErrMissingInputMusicianInstrument
					}
					return nil
				}).
				Filterable(false).
				Value(&instruments),
		),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	if err != nil {
		return errors.ToIBNError(err)
	}
	musician.Name = strings.TrimSpace(musician.Name)
	musician.Instruments.SetInstruments(instruments)
	return nil
}

// musicianCommandJoinPromptOtherInstrument gathers input for  other instrument
func musicianCommandJoinPromptOtherInstrument(musician *musician.Musician, other string) error {
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is the other instrument?").
				Value(&other).
				Placeholder("Bassoon").
				Validate(func(s string) error {
					if len(strings.TrimSpace(s)) <= 0 {
						return errors.ErrMissingInputMusicianInstrument
					}
					return nil
				}),
		),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	if err != nil {
		return errors.ToIBNError(err)
	}
	musician.Instruments.SetOther(instrument.NewInstrument(other))
	return nil
}

// musicianCommandJoinPromptConfirm checks musician information before creation
func musicianCommandJoinPromptConfirm(musician *musician.Musician) (bool, error) {
	confirmed := true
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("Is '%s' prepared to take the stage?", musician.Name)).
				Value(&confirmed).
				Description(strings.Join(musician.Instruments.InstrumentsF(), "\n")),
		),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	return confirmed, err
}
