package instrument

import (
	"slices"
	"strings"

	"github.com/zimeg/instant-band-night/internal/display"
)

// Instruments contains a collection of instruments for a musician
type Instruments []Instrument

// AddInstrument adds an instrument to the collection of instruments
func (is *Instruments) AddInstrument(ins Instrument) {
	if ins.IsOther() {
		is.SetOther(ins)
	} else if ok, _ := is.HasInstrument(ins); !ok {
		*is = append(*is, ins)
	}
}

// GetInstruments returns a slice of instrument to make instruments
func (is *Instruments) GetInstruments() (instruments []Instrument) {
	return *is
}

// HasOther denotes if some kind of instrument exists
func (is *Instruments) HasInstrument(find Instrument) (bool, int) {
	for ii, ins := range is.GetInstruments() {
		if find == ins || (find.IsOther() && ins.IsOther()) {
			return true, ii
		}
	}
	return false, 0
}

// SetInstruments overwrites the existing instrument set with those provided
func (is *Instruments) SetInstruments(instruments []Instrument) {
	*is = instruments
}

// SetOther sets an instrument value for the other instrument
func (is *Instruments) SetOther(other Instrument) {
	for ii, ins := range *is {
		if ins.IsOther() {
			(*is)[ii] = other
			return
		}
	}
	*is = append(*is, other)
}

// SortInstruments orders the instruments a standard order
func (is *Instruments) SortInstruments() *Instruments {
	instruments := is.GetInstruments()
	slices.SortFunc(instruments, func(a Instrument, b Instrument) int {
		return a.Order() - b.Order()
	})
	is.SetInstruments(instruments)
	return is
}

// InstrumentsF formats instruments on individual lines
func (is *Instruments) InstrumentsF() (lines []string) {
	for _, instrument := range is.SortInstruments().GetInstruments() {
		lines = append(lines, instrument.InstrumentF(""))
	}
	return lines
}

// InstrumentsInlineF formats instrument icons with space for missing instrument
func (is *Instruments) InstrumentsInlineF() string {
	emoji := func(instrument Instrument) string {
		if instrument.IsOther() {
			if ok, ii := is.HasInstrument(OTHER); ok {
				other := is.GetInstruments()[ii]
				return instrument.InstrumentF(display.Secondary(string(other)))
			} else {
				return "  "
			}
		}
		if ok, _ := is.HasInstrument(instrument); ok {
			return instrument.Emoji()
		} else {
			return "  "
		}
	}
	instruments := []string{
		emoji(GUITAR),
		emoji(BASS),
		emoji(DRUMS),
		emoji(VOCALS),
		emoji(ART),
		emoji(OTHER),
	}
	return strings.Join(instruments, " ")
}
