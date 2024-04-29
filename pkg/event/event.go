package event

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"path"

	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/musician"
)

// Event represents the attendance and outcomes of the night
type Event struct {
	Date      string                       `json:"date,omitempty"`
	Musicians map[string]musician.Musician `json:"musicians,omitempty"`
	filepath  string
}

// LoadEvent parses the stored event for a given date
func LoadEvent(config string, date string) (event Event, err error) {
	datepath := fmt.Sprintf("events/%s.json", date)
	filepath := path.Join(config, datepath)
	content, err := terminal.ReadFile(filepath)
	if err != nil {
		if err != errors.ErrMissingFile {
			return Event{}, errors.ToIBNError(err).
				WithMessage("Failed to load event configurations")
		} else {
			terminal.PrintInfo(display.Section(display.SectionF{
				Icon:   "star",
				Header: "Creating a new event for a new night",
			}))
		}
	} else if err = json.Unmarshal(content, &event); err != nil {
		return Event{}, errors.ToIBNError(err).
			WithMessage("Failed to parse event configurations")
	}
	event.filepath = filepath
	if event.Musicians == nil {
		event.Musicians = make(map[string]musician.Musician)
	}
	return event, nil
}

// SaveMusician adds a musician with a unique ID to the current event
func (e *Event) SaveMusician(m musician.Musician) (musician.Musician, error) {
	var id string
	for id = fmt.Sprintf("m%s", randomID(6)); e.Musicians[id].Name != ""; {
	}
	e.Musicians[id] = m
	err := terminal.WriteJSON(e.filepath, e)
	if err != nil {
		return musician.Musician{}, err
	}
	return m, nil
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
