package event

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/musician"
)

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
