package event

import (
	"github.com/zimeg/instant-band-night/internal/errors"
	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/band"
	ibninstrument "github.com/zimeg/instant-band-night/pkg/instrument"
	ibnmusician "github.com/zimeg/instant-band-night/pkg/musician"
)

// BandMusician returns the performer and instrument in a band
func (e *Event) BandMusician(
	band band.Band,
	musicianID string,
) (
	musician ibnmusician.Musician,
	instrument ibninstrument.Instrument,
	err error,
) {
	for _, ins := range band.Instruments.GetInstruments() {
		for _, mID := range *band.Instruments.GetInstrument(ins) {
			if mID == musicianID {
				musician, err := e.Musicians.GetMusician(mID)
				if err != nil {
					return ibnmusician.Musician{}, ibninstrument.NewInstrument(""), err
				}
				return musician, ins, nil
			}
		}
	}
	return ibnmusician.Musician{}, ibninstrument.NewInstrument(""), errors.ErrMusicianNotFound
}

// BandMusiciansF formats a list of musicians performing in the band
func (e *Event) BandMusiciansF(band band.Band) (musicianIDs, musicians []string) {
	for _, ins := range band.Instruments.GetInstruments() {
		instrumentMusicianIDs, instrumentMusicians := e.InstrumentF(band, ins)
		for ii, musicianID := range instrumentMusicianIDs {
			musicianIDs = append(musicianIDs, musicianID)
			musicians = append(musicians, instrumentMusicians[ii])
		}
	}
	return
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
