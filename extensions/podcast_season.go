package ext

import "strconv"

// PodcastSeason allows for identifying which episodes in a podcast are part
// of a particular "season"
//
// More Info: https://podcastindex.org/namespace/1.0#season
//
// XML Example:
//
// <podcast:season name="Egyptology: The 19th Century">1</podcast:season>
type PodcastSeason struct {
	Name   string `json:"name,omitempty"`
	Number int    `json:"number,omitempty"`
}

func parseSeason(extensions map[string][]Extension) (season *PodcastSeason) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["season"]
	if !ok || len(matches) == 0 {
		return
	}

	season = &PodcastSeason{}

	if text, ok := matches[0].Attrs["name"]; ok {
		season.Name = text
	}

	if val, err := strconv.Atoi(matches[0].Value); err == nil {
		season.Number = val
	}

	return season
}
