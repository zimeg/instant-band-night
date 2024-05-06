package band

import (
	"github.com/zimeg/instant-band-night/pkg/instrument"
)

// Arrangement counts the number of instruments performing
type Arrangement map[instrument.Instrument]int

// GetInstruments returns the entire instrumentation of an arrangement
func (b *Arrangement) GetInstruments() (instruments []instrument.Instrument) {
	for ins, count := range *b {
		for range count {
			instruments = append(instruments, ins)
		}
	}
	return
}

// ArrangementF organizes a list of instrument counts for the band
func (b *Arrangement) ArrangementF() (lines []string) {
	var instruments instrument.Instruments = b.GetInstruments()
	return instruments.SortInstruments().InstrumentsF()
}
