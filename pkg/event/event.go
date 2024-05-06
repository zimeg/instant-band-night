package event

import (
	"fmt"
	"path"

	"github.com/zimeg/instant-band-night/internal/terminal"
	"github.com/zimeg/instant-band-night/pkg/band"
	"github.com/zimeg/instant-band-night/pkg/bucket"
	"github.com/zimeg/instant-band-night/pkg/musician"
)

// Event represents the attendance and outcomes of the night
type Event struct {
	Date      string             `json:"date,omitempty"`
	Bands     band.Bands         `json:"bands,omitempty"`
	Buckets   *bucket.Buckets    `json:"-"`
	Musicians musician.Musicians `json:"musicians,omitempty"`
	filepath  string
}

// LoadEvent initializes information about the stored event for a given date
func LoadEvent(config string, date string) (event Event, err error) {
	filepath := path.Join(config, fmt.Sprintf("events/%s.json", date))
	err = terminal.ReadJSON(filepath, &event)
	if err != nil {
		return Event{}, err
	}
	event.filepath = filepath
	event.Date = date
	if event.Musicians == nil {
		event.Musicians = make(map[string]musician.Musician)
	}
	for id, musician := range event.Musicians {
		musician.SetID(id)
		event.Musicians[id] = musician
	}
	for order, band := range event.Bands {
		for _, ins := range band.Instruments.GetInstruments() {
			for _, musicianID := range *band.Instruments.GetInstrument(ins) {
				musician := event.Musicians[musicianID]
				musician.AddBand(order)
				event.Musicians[musicianID] = musician
			}
		}
	}
	event.Buckets = bucket.NewBuckets(event.Musicians)
	return event, nil
}
