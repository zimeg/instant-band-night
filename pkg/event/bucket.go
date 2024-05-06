package event

// FilterCooldown removes musicians that performed since the last band
func (e *Event) FilterCooldown(since int) {
	for _, bucket := range *e.Buckets {
		for _, musicianID := range bucket.GetMusicians() {
			musician := e.GetMusician(musicianID)
			performance, ok := musician.LastPerformance()
			if !ok || performance < len(e.Bands)-since {
				continue
			} else {
				bucket.RemoveMusician(musicianID)
			}
		}
	}
}
