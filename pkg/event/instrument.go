package event

import (
	"github.com/zimeg/instant-band-night/internal/display"
	"github.com/zimeg/instant-band-night/pkg/band"
	"github.com/zimeg/instant-band-night/pkg/instrument"
)

// InstrumentF formats the list of musicians on a certain instrument
func (e *Event) InstrumentF(band band.Band, ins instrument.Instrument) []string {
	var musicians []string
	for _, musicianID := range *band.Instruments.GetInstrument(ins) {
		musician := e.GetMusician(musicianID)
		tag := ins.InstrumentF("%s %s", display.Secondary(musicianID), musician.Name)
		musicians = append(musicians, tag)
	}
	return musicians
}
