package event

import (
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/band"
)

// BandMusicianF formats a list of musicians performing in the band
func (e *Event) BandMusicianF(band band.Band) (musicians []string) {
	var instruments [][]string
	for _, ins := range band.Instruments.GetInstruments() {
		instruments = append(instruments, e.InstrumentF(band, ins))
	}
	for _, instrument := range instruments {
		musicians = append(musicians, instrument...)
	}
	return musicians
}

// SaveBand adds a band to the list of performances for the event
func (e *Event) SaveBand(band band.Band) error {
	e.Bands = append(e.Bands, band)
	err := terminal.WriteJSON(e.filepath, e)
	if err != nil {
		return err
	}
	return nil
}
