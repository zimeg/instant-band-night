package instrument

import (
	"fmt"
	"strings"

	"github.com/zimeg/instant-band-night/internal/display"
)

// Instrument tends to be something that makes sounds
type Instrument string

const (
	GUITAR Instrument = "guitar"
	BASS   Instrument = "bass"
	DRUMS  Instrument = "drums"
	VOCALS Instrument = "vocals"
	ART    Instrument = "art"
	OTHER  Instrument = "other"
)

// NewInstrument creates the value of a unique instrument
func NewInstrument(instrument string) Instrument {
	instrument = strings.ToLower(instrument)
	instrument = strings.TrimSpace(instrument)
	return Instrument(instrument)
}

// Emoji returns an icon that represents the instrument
func (ins Instrument) Emoji() string {
	switch ins {
	case GUITAR:
		return display.Emoji("guitar")
	case BASS:
		return display.Emoji("speaker")
	case DRUMS:
		return display.Emoji("drums")
	case VOCALS:
		return display.Emoji("microphone")
	case ART:
		return display.Emoji("art")
	default:
		return display.Emoji("piano")
	}
}

// Order returns a unique number for sorting instrument with
func (ins Instrument) Order() int {
	switch ins {
	case GUITAR:
		return 0
	case BASS:
		return 1
	case DRUMS:
		return 2
	case VOCALS:
		return 3
	case ART:
		return 4
	default:
		return 5
	}
}

// InstrumentF returns an instrument icon with an optional label
func (ins Instrument) InstrumentF(label string, a ...any) string {
	if label == "" {
		label = string(ins)
	} else {
		label = fmt.Sprintf(label, a...)
	}
	return fmt.Sprintf("%s %s", ins.Emoji(), label)
}

// IsOther determines if the instrument is part of the other bucket
func (ins Instrument) IsOther() bool {
	return (ins != GUITAR &&
		ins != BASS &&
		ins != DRUMS &&
		ins != VOCALS &&
		ins != ART)
}
