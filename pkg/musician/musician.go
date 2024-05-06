package musician

import (
	"fmt"
	"strings"

	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/pkg/instrument"
)

// Musician represents the interested instruments of a performer
type Musician struct {
	id          string
	Name        string                 `json:"name,omitempty"`
	Instruments instrument.Instruments `json:"instruments,omitempty"`
}

// GetID returns the unique ID of a musician
func (m *Musician) GetID() string {
	return m.id
}

// SetID sets the unique ID of a musician
func (m *Musician) SetID(id string) {
	m.id = id
}

// MusicianF arranges details about the musician for printing
func (m *Musician) MusicianF() string {
	return fmt.Sprintf("%s %s %s  %s  %s",
		display.Emoji("star"),
		display.Secondary(m.id),
		m.Name,
		display.Faint(strings.Repeat(".", 26-len(m.Name))),
		m.Instruments.InstrumentsInlineF(),
	)
}
