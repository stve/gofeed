package ext

import "strconv"

// PodcastUpdateFrequency allows a podcaster to express their intended release
// schedule as structured data and text.
//
// More Info: https://podcastindex.org/namespace/1.0#update-frequency
//
// XML Example:
//
// <podcast:updateFrequency rrule="FREQ=DAILY">Daily</podcast:updateFrequency>
// <podcast:updateFrequency rrule="FREQ=MONTHLY;INTERVAL=2">Bimonthly</podcast:updateFrequency>
type PodcastUpdateFrequency struct {
	Complete       bool   `json:"complete,omitempty"`
	StartDate      string `json:"dtstart,omitempty"`
	RecurrenceRule string `json:"rrule,omitempty"`
	Label          string `json:"label,omitempty"`
}

func parseUpdateFrequency(extensions map[string][]Extension) (updateFrequency *PodcastUpdateFrequency) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["updateFrequency"]
	if !ok || len(matches) == 0 {
		return
	}

	updateFrequency = &PodcastUpdateFrequency{}

	if text, ok := matches[0].Attrs["rrule"]; ok {
		updateFrequency.RecurrenceRule = text
	}

	if text, ok := matches[0].Attrs["dstart"]; ok {
		updateFrequency.StartDate = text
	}

	if text, ok := matches[0].Attrs["complete"]; ok {
		if val, err := strconv.ParseBool(text); err == nil {
			updateFrequency.Complete = val
		}
	}

	updateFrequency.Label = matches[0].Value

	return updateFrequency
}
