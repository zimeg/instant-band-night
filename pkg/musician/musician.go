package musician

import (
	"github.com/zimeg/instant-band-night/pkg/instrument"
)

// Musician represents the interested instruments of a performer
type Musician struct {
	Name        string   `json:"name,omitempty"`
	Instruments []string `json:"instruments,omitempty"`
}

// InstrumentsF arranges the instruments used by the musician
func (m *Musician) InstrumentsF() []string {
	var instruments []string
	for _, ins := range m.Instruments {
		instruments = append(instruments, instrument.InstrumentF(ins))
	}
	return instruments
}
