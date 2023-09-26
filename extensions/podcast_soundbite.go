package ext

// PodcastSoundbite points to one or more soundbites within a podcast episode.
//
// More Info: https://podcastindex.org/namespace/1.0#soundbite
//
// XML Example:
//
// <podcast:soundbite startTime="73.0" duration="60.0" />

type PodcastSoundbite struct {
	Duration  string `json:"duration,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	Title     string `json:"title,omitempty"`
}

func parseSoundbites(extensions map[string][]Extension) (soundbites []*PodcastSoundbite) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["soundbite"]
	if !ok || len(matches) == 0 {
		return
	}

	soundbites = []*PodcastSoundbite{}

	for _, cat := range matches {
		soundbite := &PodcastSoundbite{}

		if text, ok := cat.Attrs["duration"]; ok {
			soundbite.Duration = text
		}

		if text, ok := cat.Attrs["startTime"]; ok {
			soundbite.StartTime = text
		}

		soundbite.Title = cat.Value

		soundbites = append(soundbites, soundbite)
	}

	return soundbites
}
