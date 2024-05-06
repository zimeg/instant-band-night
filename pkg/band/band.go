package band

import (
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/pkg/bucket"
	"github.com/zimeg/instant-band-night/pkg/instrument"
)

// Band contains the lineup of a performing group
type Band struct {
	Name        string         `json:"name,omitempty"`
	Instruments bucket.Buckets `json:"instruments"`
}

// NewBand forms a new group with a random arrangement of instrumentation from
// the buckets
func NewBand(arrangement Arrangement, buckets bucket.Buckets) (Band, error) {
	band := Band{
		Instruments: bucket.Buckets{},
	}
	for ins, count := range arrangement {
		for range count {
			musicianID, err := band.drawInstrument(buckets, ins)
			if err != nil {
				return Band{}, err
			}
			band.Instruments.GetInstrument(ins).AddMusician(musicianID)
		}
	}
	return band, nil
}

// drawInstrument attempts to draw a unique instrument for the band
func (b *Band) drawInstrument(
	buckets bucket.Buckets,
	ins instrument.Instrument,
) (
	musicianID string,
	err error,
) {
	const RETRIES = 12
	for range RETRIES {
		musicianID, err = buckets.DrawMusician(ins)
		if err != nil {
			return "", err
		}
		if !b.Instruments.ContainsMusician(musicianID) {
			return musicianID, nil
		}
	}
	return "", errors.ErrNotEnoughMusician
}
