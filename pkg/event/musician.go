package event

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"

	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/musician"
)

// GetMusician finds the musician with a known ID
func (e *Event) GetMusician(musicianID string) musician.Musician {
	return e.Musicians[musicianID]
}

// GetMusicians returns all of the musicians at an event
func (e *Event) GetMusicians() (musicians []musician.Musician) {
	for _, musician := range e.Musicians {
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

// SaveMusician adds a musician with a unique ID to the current event
func (e *Event) SaveMusician(m musician.Musician) error {
	if m.GetID() == "" {
		var id string
		for id = fmt.Sprintf("m%s", randomID(6)); e.Musicians[id].Name != ""; {
		}
		m.SetID(id)
	}
	e.Musicians[m.GetID()] = m
	err := terminal.WriteJSON(e.filepath, e)
	if err != nil {
		return err
	}
	return nil
}

// randomID creates a random identifier for unique purposes
func randomID(length int) string {
	const charset = "abcdef0123456789"
	var result string
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		result += string(charset[index.Int64()])
	}
	return result
}
