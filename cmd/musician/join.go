package musician

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/internal/terminal"
)

// musicianCommandJoinFlags contains the flags for this command
type musicianCommandJoinFlagSet struct {
	artFlag    bool
	bassFlag   bool
	guitarFlag bool
	drumsFlag  bool
	nameFlag   string
	otherFlag  string
	vocalsFlag bool
}

// musicianCommandJoinFlags implements the flags for this command
var musicianCommandJoinFlags musicianCommandJoinFlagSet

// MusicianCommandJoinNew adds musicians to certain instrument buckets
func MusicianCommandJoinNew() *cobra.Command {
	musicianCommandJoin := &cobra.Command{
		Use:     "join",
		Aliases: []string{"add", "create"},
		Short:   "Join buckets of instruments",
		Long: strings.Join([]string{
			"Join different instrument buckets as a participating band member.",
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			return musicianCommandJoinRunE(cmd)
		},
	}
	musicianCommandJoin.Flags().BoolVar(&musicianCommandJoinFlags.artFlag, "artist", false, "the one drawing band posters")
	musicianCommandJoin.Flags().BoolVar(&musicianCommandJoinFlags.bassFlag, "bass", false, "the one making lower sounds")
	musicianCommandJoin.Flags().BoolVar(&musicianCommandJoinFlags.drumsFlag, "drums", false, "the one often with the beat")
	musicianCommandJoin.Flags().BoolVar(&musicianCommandJoinFlags.guitarFlag, "guitar", false, "the one playing the six strings")
	musicianCommandJoin.Flags().StringVar(&musicianCommandJoinFlags.nameFlag, "name", "", "the title of the one performing")
	musicianCommandJoin.Flags().StringVar(&musicianCommandJoinFlags.otherFlag, "other", "", "the one with another instrument")
	musicianCommandJoin.Flags().BoolVar(&musicianCommandJoinFlags.vocalsFlag, "vocals", false, "the one making person noises")
	return musicianCommandJoin
}

// musicianCommandJoinRunE prompts and saves information about the new musician
func musicianCommandJoinRunE(cmd *cobra.Command) error {
	musician, err := musicianCommandJoinPrompt(cmd)
	if err != nil {
		return err
	}
	terminal.PrintInfo(display.Section(display.SectionF{
		Icon:   "star",
		Header: fmt.Sprintf("Welcome to instant band night '%s'!", musician.Name),
		Body:   musician.InstrumentsF(),
	}))
	return nil
}
