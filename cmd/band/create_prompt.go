package band

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/band"
	"github.com/zimeg/instant-band-night/pkg/instrument"
)

// bandCommandCreatePrompt gathers instrument counts for the band
func bandCommandCreatePrompt(cmd *cobra.Command) (arrangement band.Arrangement, err error) {
	confirm := bandCommandCreatePromptFlags(cmd, &arrangement)
	if confirm {
		return arrangement, nil
	}
	confirm, err = bandCommandCreatePromptConfirm(arrangement)
	if err != nil {
		return band.Arrangement{}, errors.ToIBNError(err)
	}
	if confirm {
		return arrangement, nil
	}
	err = bandCommandCreatePromptArrangement(&arrangement)
	if err != nil {
		return band.Arrangement{}, errors.ToIBNError(err)
	}
	return arrangement, nil
}

// bandCommandCreatePromptFlags sets instrument arrangement values using command
// flags and returns if this is a confirmed customization
func bandCommandCreatePromptFlags(cmd *cobra.Command, arrangement *band.Arrangement) (confirm bool) {
	*arrangement = band.Arrangement{
		instrument.GUITAR: terminal.FlagToInt(cmd.Flag("guitar")),
		instrument.BASS:   terminal.FlagToInt(cmd.Flag("bass")),
		instrument.DRUMS:  terminal.FlagToInt(cmd.Flag("drums")),
		instrument.VOCALS: terminal.FlagToInt(cmd.Flag("vocals")),
		instrument.ART:    terminal.FlagToInt(cmd.Flag("artist")),
		instrument.OTHER:  terminal.FlagToInt(cmd.Flag("other")),
	}
	return terminal.FlagToBool(cmd.Flag("confirm"))
}

// bandCommandCreatePromptConfirm determines if the band arrangement is set
func bandCommandCreatePromptConfirm(arrangement band.Arrangement) (bool, error) {
	confirm := true
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Arrange this orchestration").
				Description(strings.Join(arrangement.ArrangementF(), "\n")).
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
func bandCommandCreatePromptArrangement(arrangement *band.Arrangement) error {
	newInputCount := func(prompt string, count *int) *huh.Input {
		instrumentCount := strconv.Itoa(*count)
		return huh.NewInput().
			Prompt(prompt).
			Validate(func(s string) error {
				if len(s) == 0 {
					return errors.ErrMissingInputInstrumentCount
				}
				if value, err := strconv.Atoi(s); err != nil {
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
			newInputCount("🎸 guitar ", &guitar).
				Title("Arrange the instrumentation").
				Description(" "),
			newInputCount("🔉 bass ", &bass).Inline(true),
			newInputCount("🥁 drums ", &drums).Inline(true),
			newInputCount("🎤 vocals ", &vocals).Inline(true),
			newInputCount("🎨 art ", &art).Inline(true),
			newInputCount("🎹 other ", &other).Inline(true),
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
