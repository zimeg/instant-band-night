package musician

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/errors"
)

// musicianCommandJoinPrompt gathers information about the performing musician
// from flags and interactive input
func musicianCommandJoinPrompt(cmd *cobra.Command) (musician Musician, err error) {
	var otherInstrument string
	if cmd.Flag("guitar").Value.String() == "true" {
		musician.Instruments = append(musician.Instruments, "guitar")
	}
	if cmd.Flag("bass").Value.String() == "true" {
		musician.Instruments = append(musician.Instruments, "bass")
	}
	if cmd.Flag("drums").Value.String() == "true" {
		musician.Instruments = append(musician.Instruments, "drums")
	}
	if cmd.Flag("vocals").Value.String() == "true" {
		musician.Instruments = append(musician.Instruments, "vocals")
	}
	if cmd.Flag("other").Changed {
		musician.Instruments = append(musician.Instruments, "other")
		otherInstrument = cmd.Flag("other").Value.String()
	}
	if cmd.Flag("artist").Value.String() == "true" {
		musician.Instruments = append(musician.Instruments, "art")
	}
	if cmd.Flag("name").Changed {
		if len(strings.TrimSpace(cmd.Flag("name").Value.String())) > 0 {
			musician.Name = strings.TrimSpace(cmd.Flag("name").Value.String())
		} else {
			return Musician{}, errors.ErrMissingInputMusicianName
		}
	}
	joinFlagsUsed := musician.Name != "" && len(musician.Instruments) > 0
	if !joinFlagsUsed {
		err = huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Sign a stage name").
					Value(&musician.Name).
					CharLimit(60).
					Validate(func(s string) error {
						if len(strings.TrimSpace(s)) == 0 {
							return errors.ErrMissingInputMusicianName
						}
						return nil
					}).
					Placeholder("Rockstar"),
				huh.NewMultiSelect[string]().
					Title("Pick some instruments").
					Options(
						huh.NewOption("Guitar", "guitar"),
						huh.NewOption("Bass", "bass"),
						huh.NewOption("Drums", "drums"),
						huh.NewOption("Vocals", "vocals"),
						huh.NewOption("Art", "art"),
						huh.NewOption("Other", "other"),
					).
					Validate(func(s []string) error {
						if len(s) == 0 {
							return errors.ErrMissingInputMusicianInstrument
						}
						return nil
					}).
					Filterable(false).
					Value(&musician.Instruments),
			),
		).
			WithTheme(huh.ThemeDracula()).
			Run()
		if err != nil {
			return Musician{}, errors.ToIBNError(err)
		}
	}
	if slices.Contains(musician.Instruments, "other") && otherInstrument == "" {
		err = huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("What is the other instrument?").
					Value(&otherInstrument).
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
			return Musician{}, errors.ToIBNError(err)
		}
	}
	musician.Name = strings.TrimSpace(musician.Name)
	if len(strings.TrimSpace(otherInstrument)) > 0 {
		otherInstrument = strings.ToLower(otherInstrument)
		otherInstrument = strings.TrimSpace(otherInstrument)
		otherInstrument = fmt.Sprintf("other: %s", otherInstrument)
		musician.Instruments[len(musician.Instruments)-1] = otherInstrument
	}
	if !joinFlagsUsed {
		confirmed := true
		err = huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title(fmt.Sprintf("Is '%s' prepared to take the stage?", musician.Name)).
					Value(&confirmed).
					Description(strings.Join(musician.InstrumentsF(), "\n")),
			),
		).
			WithTheme(huh.ThemeDracula()).
			Run()
		if err != nil {
			return Musician{}, errors.ToIBNError(err)
		}
		if !confirmed {
			return Musician{}, errors.ErrPromptInterrupt.
				WithMessage("No musician created")
		}
	}
	return
}
