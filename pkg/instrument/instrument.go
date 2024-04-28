package instrument

import (
	"fmt"

	"github.com/zimeg/instant-band-night/internal/display"
)

// Instrument tends to be something that makes sounds
type Instrument string

// InstrumentF returns a instrument icon with label
func InstrumentF(instrument string) string {
	switch instrument {
	case "guitar":
		return fmt.Sprintf("%s guitar", display.Emoji("guitar"))
	case "bass":
		return fmt.Sprintf("%s bass", display.Emoji("speaker"))
	case "drums":
		return fmt.Sprintf("%s drums", display.Emoji("drums"))
	case "vocals":
		return fmt.Sprintf("%s vocals", display.Emoji("microphone"))
	case "art":
		return fmt.Sprintf("%s art", display.Emoji("art"))
	default:
		return fmt.Sprintf("%s %s", display.Emoji("piano"), instrument)
	}
}
