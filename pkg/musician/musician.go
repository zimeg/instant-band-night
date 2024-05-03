package musician

import (
	"fmt"
	"slices"
	"strings"

	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/pkg/instrument"
)

// Musician represents the interested instruments of a performer
type Musician struct {
	id          string
	Name        string   `json:"name,omitempty"`
	Instruments []string `json:"instruments,omitempty"`
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
		display.Secondary(fmt.Sprintf("%s:", m.id)),
		m.Name,
		display.Faint(strings.Repeat(".", 26-len(m.Name))),
		m.instrumentIcons(),
	)
}

// InstrumentsF arranges the instruments used by the musician
func (m *Musician) InstrumentsF() []string {
	var instruments []string
	for _, ins := range m.Instruments {
		instruments = append(instruments, instrument.InstrumentF(ins))
	}
	return instruments
}

// instrumentIcons returns an ordered list of instruments for the musician
func (m *Musician) instrumentIcons() string {
	var icons []string
	if slices.Contains(m.Instruments, "guitar") {
		icons = append(icons, display.Emoji("guitar"))
	} else {
		icons = append(icons, "  ")
	}
	if slices.Contains(m.Instruments, "bass") {
		icons = append(icons, display.Emoji("speaker"))
	} else {
		icons = append(icons, "  ")
	}
	if slices.Contains(m.Instruments, "drums") {
		icons = append(icons, display.Emoji("drums"))
	} else {
		icons = append(icons, "  ")
	}
	if slices.Contains(m.Instruments, "vocals") {
		icons = append(icons, display.Emoji("microphone"))
	} else {
		icons = append(icons, "  ")
	}
	if slices.Contains(m.Instruments, "art") {
		icons = append(icons, display.Emoji("art"))
	} else {
		icons = append(icons, "  ")
	}
	if slices.ContainsFunc(m.Instruments, func(instrument string) bool {
		return strings.HasPrefix(instrument, "other")
	}) {
		icons = append(icons, display.Emoji("piano"))
	} else {
		icons = append(icons, "  ")
	}
	return strings.Join(icons, " ")
}
