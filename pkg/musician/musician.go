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
	Bands       []int                  `json:"-"`
	Instruments instrument.Instruments `json:"instruments,omitempty"`
}

// AddBand adds the ordered band to the musician bands
func (m *Musician) AddBand(band int) {
	m.Bands = append(m.Bands, band)
}

// GetID returns the unique ID of a musician
func (m *Musician) GetID() string {
	return m.id
}

// LastPerformance returns the most recent band if a performance happened
func (m *Musician) LastPerformance() (int, bool) {
	if len(m.Bands) > 0 {
		return m.Bands[len(m.Bands)-1], true
	} else {
		return 0, false
	}
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
