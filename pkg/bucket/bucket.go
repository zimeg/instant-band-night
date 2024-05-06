package bucket

import (
	"math/rand/v2"
	"slices"

	"github.com/zimeg/instant-band-night/internal/errors"
)

// Bucket is a slice of musician for a certain instrument
type Bucket []string

// NewBucket returns a bucket with no entries
func NewBucket() *Bucket {
	var bucket Bucket = []string{}
	return &bucket
}

// AddMusician adds the musicianID to a bucket
func (b *Bucket) AddMusician(musicianID string) {
	*b = append(*b, musicianID)
}

// ContainsMusician returns if the musicianID is in the bucket
func (b Bucket) ContainsMusician(musicianID string) bool {
	return slices.Contains(b, musicianID)
}

// DrawMusician returns a single musician from the possible musicians
func (b Bucket) DrawMusician() (musicianID string, err error) {
	if len(b) <= 0 {
		return "", errors.ErrNotEnoughMusician
	}
	return b[rand.IntN(len(b))], nil
}

// GetMusicians returns the musicianIDs in a bucket
func (b Bucket) GetMusicians() []string {
	return b
}

// RemoveMusician replaces the bucket with a musician with a bucket without
func (b *Bucket) RemoveMusician(musicianID string) {
	var musicians Bucket
	for _, id := range *b {
		if id != musicianID {
			musicians = append(musicians, id)
		}
	}
	*b = musicians
}
