package musician

import (
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/event"
)

// MusicianCommandListNew outputs all musicians and instruments
func MusicianCommandListNew(event *event.Event) *cobra.Command {
	musicianCommandList := &cobra.Command{
		Use:     "list",
		Aliases: []string{"read", "show"},
		Short:   "List instruments of musicians",
		Long: strings.Join([]string{
			"Join different instrument buckets as a participating band member.",
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			return musicianCommandListRunE(event)
		},
	}
	return musicianCommandList
}

// musicianCommandListRunE prompts and saves information about the new musician
func musicianCommandListRunE(event *event.Event) error {
	var musicians []Musician
	for id, musician := range event.Musicians {
		musician.SetID(id)
		musicians = append(musicians, musician)
	}
	sort.Slice(musicians, func(i, j int) bool {
		if musicians[i].Name == musicians[j].Name {
			return musicians[i].GetID() > musicians[j].GetID()
		}
		return musicians[i].Name < musicians[j].Name
	})
	for _, musician := range musicians {
		terminal.PrintInfo(musician.MusicianF())
	}
	return nil
}
