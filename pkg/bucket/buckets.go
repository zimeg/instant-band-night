package bucket

import (
	"github.com/zimeg/instant-band-night/pkg/instrument"
	"github.com/zimeg/instant-band-night/pkg/musician"
)

// Buckets contains buckets of musicians on each instrument
type Buckets map[instrument.Instrument]*Bucket

// NewBuckets organizes musicians into interested instrument buckets
func NewBuckets(musicians musician.Musicians) *Buckets {
	var buckets Buckets = map[instrument.Instrument]*Bucket{}
	for _, musician := range musicians {
		for _, ins := range musician.Instruments.GetInstruments() {
			buckets.GetInstrument(ins).AddMusician(musician.GetID())
		}
	}
	return &buckets
}

// ContainsMusician returns if the musicianID is found in the buckets
func (b *Buckets) ContainsMusician(musicianID string) bool {
	for _, instrument := range b.GetInstruments() {
		if b.GetInstrument(instrument).ContainsMusician(musicianID) {
			return true
		}
	}
	return false
}

// DrawMusician returns a unique and random musician for the band
func (b *Buckets) DrawMusician(ins instrument.Instrument) (musicianID string, err error) {
	musicianID, err = b.GetInstrumentGroup(ins).DrawMusician()
	if err != nil {
		return "", err
	}
	return musicianID, nil
}

// GetInstrument returns the bucket for a certain instrument
func (b Buckets) GetInstrument(ins instrument.Instrument) *Bucket {
	if b == nil {
		b = *NewBuckets(make(musician.Musicians))
	}
	if b[ins] == nil {
		bucket := NewBucket()
		b[ins] = bucket
	}
	return b[ins]
}

// GetInstrumentGroup returns the bucket of a certain instrument group
//
// Groupings are based on definitions in pkg/instrument/instrument.go
func (b Buckets) GetInstrumentGroup(group instrument.Instrument) Bucket {
	if !group.IsOther() {
		return *b.GetInstrument(group)
	}
	var musicians Bucket = []string{}
	for _, ins := range b.GetInstruments() {
		if ins.IsOther() {
			musicians = append(musicians, *b.GetInstrument(ins)...)
		}
	}
	return musicians
}

// GetInstruments returns the instruments with musicians in buckets
func (b *Buckets) GetInstruments() []instrument.Instrument {
	var instruments instrument.Instruments
	for ins := range *b {
		instruments.AddInstrument(ins)
	}
	instruments.SortInstruments()
	return instruments.GetInstruments()
}
