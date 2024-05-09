package band

import (
	"github.com/zimeg/instant-band-night/internal/errors"
)

// Bands contains the collection of collaborating bands
type Bands []Band

// GetBand returns the band matching the ID
func (b Bands) GetBand(bandID int) (*Band, error) {
	if 0 <= bandID && bandID < len(b) {
		return &b[bandID], nil
	}
	return &Band{}, errors.ErrBandNotFound
}
