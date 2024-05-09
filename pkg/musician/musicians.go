package musician

import (
	"sort"

	"github.com/zimeg/instant-band-night/internal/errors"
)

// Musicians maps IDs to musicians
type Musicians map[string]Musician

// GetMusician finds the musician with a known ID
func (ms *Musicians) GetMusician(musicianID string) (Musician, error) {
	if musician, ok := (*ms)[musicianID]; ok {
		return musician, nil
	}
	return Musician{}, errors.ErrMusicianNotFound
}

// GetMusicians returns all of the musicians at an event
func (ms *Musicians) GetMusicians() (musicians []Musician) {
	for id, musician := range *ms {
		musician.SetID(id)
		musicians = append(musicians, musician)
	}
	sort.Slice(musicians, func(i, j int) bool {
		if musicians[i].Name == musicians[j].Name {
			return musicians[i].GetID() > musicians[j].GetID()
		}
		return musicians[i].Name < musicians[j].Name
	})
	return
}
