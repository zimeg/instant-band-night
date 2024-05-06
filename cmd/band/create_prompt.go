package band

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/band"
	"github.com/zimeg/instant-band-night/pkg/instrument"
)

// bandCommandCreatePrompt gathers instrument counts for the band
func bandCommandCreatePrompt(cmd *cobra.Command) (arrangement band.Arrangement, cooldown int, err error) {
	confirm, cooldown := bandCommandCreatePromptFlags(cmd, &arrangement)
	if confirm {
		return arrangement, cooldown, nil
	}
	confirm, err = bandCommandCreatePromptConfirm(arrangement, cooldown)
	if err != nil {
		return band.Arrangement{}, 0, errors.ToIBNError(err)
	}
	if confirm {
		return arrangement, cooldown, nil
	}
	err = bandCommandCreatePromptArrangement(&arrangement, &cooldown)
	if err != nil {
		return band.Arrangement{}, 0, errors.ToIBNError(err)
	}
	return arrangement, cooldown, nil
}

// bandCommandCreatePromptFlags sets instrument arrangement values using command
// flags and returns if this is a confirmed customization
func bandCommandCreatePromptFlags(cmd *cobra.Command, arrangement *band.Arrangement) (confirm bool, cooldown int) {
	*arrangement = band.Arrangement{
		instrument.GUITAR: terminal.FlagToInt(cmd.Flag("guitar")),
		instrument.BASS:   terminal.FlagToInt(cmd.Flag("bass")),
		instrument.DRUMS:  terminal.FlagToInt(cmd.Flag("drums")),
		instrument.VOCALS: terminal.FlagToInt(cmd.Flag("vocals")),
		instrument.ART:    terminal.FlagToInt(cmd.Flag("artist")),
		instrument.OTHER:  terminal.FlagToInt(cmd.Flag("other")),
	}
	confirm = terminal.FlagToBool(cmd.Flag("confirm"))
	cooldown = terminal.FlagToInt(cmd.Flag("cooldown"))
	return
}

// bandCommandCreatePromptConfirm determines if the band arrangement is set
func bandCommandCreatePromptConfirm(arrangement band.Arrangement, cooldown int) (bool, error) {
	confirm := true
	description := arrangement.ArrangementF()
	description = append(description, "")
	description = append(description,
		fmt.Sprintf("%s cooldown: %d", display.Emoji("hourglass"), cooldown))
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Arrange this orchestration").
				Description(strings.Join(description, "\n")).
				Value(&confirm),
		),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	if err != nil {
		return false, err
	}
	return confirm, nil
}

// bandCommandCreatePromptArrangement collects a count for each instrument
func bandCommandCreatePromptArrangement(arrangement *band.Arrangement, cooldown *int) error {
	newInputCount := func(icon string, prompt string, count *int) *huh.Input {
		instrumentCount := strconv.Itoa(*count)
		return huh.NewInput().
			Prompt(fmt.Sprintf("%s %s ", icon, prompt)).
			Validate(func(s string) error {
				if len(s) == 0 {
					return errors.ErrMissingInputInstrumentCount
				}
				if value, err := strconv.Atoi(s); err != nil || value < 0 {
					return errors.ErrMissingInputInstrumentCount
				} else {
					*count = value
				}
				return nil
			}).
			Placeholder(instrumentCount).
			Value(&instrumentCount)
	}
	var (
		guitar int = (*arrangement)[instrument.GUITAR]
		bass   int = (*arrangement)[instrument.BASS]
		drums  int = (*arrangement)[instrument.DRUMS]
		vocals int = (*arrangement)[instrument.VOCALS]
		art    int = (*arrangement)[instrument.ART]
		other  int = (*arrangement)[instrument.OTHER]
	)
	err := huh.NewForm(
		huh.NewGroup(
			newInputCount(display.Emoji("guitar"), "guitar", &guitar).
				Title("Arrange the instrumentation").
				Description(" "),
			newInputCount(display.Emoji("speaker"), "bass", &bass).Inline(true),
			newInputCount(display.Emoji("drums"), "drums", &drums).Inline(true),
			newInputCount(display.Emoji("microphone"), "vocals", &vocals).Inline(true),
			newInputCount(display.Emoji("art"), "art", &art).Inline(true),
			newInputCount(display.Emoji("piano"), "other", &other).Inline(true),
			newInputCount(display.Emoji("hourglass"), "cooldown", cooldown).
				Title("Relax on the cooldown").
				Description(" "),
		),
	).
		WithTheme(huh.ThemeDracula()).
		Run()
	if err != nil {
		return errors.ToIBNError(err)
	}
	(*arrangement)[instrument.GUITAR] = guitar
	(*arrangement)[instrument.BASS] = bass
	(*arrangement)[instrument.DRUMS] = drums
	(*arrangement)[instrument.VOCALS] = vocals
	(*arrangement)[instrument.ART] = art
	(*arrangement)[instrument.OTHER] = other
	return nil
}
