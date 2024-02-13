package ext

// PodcastEpisode exists largely for compatibility with the season tag. But,
// it also allows for a similar idea to what "name" functions as in that
// element.
//
// More Info: https://podcastindex.org/namespace/1.0#episode
//
// XML Example:
//
// <podcast:episode display="Ch.3">204</podcast:episode>
type PodcastEpisode struct {
	Display string `json:"display,omitempty"`
	Number  string `json:"number,omitempty"`
}

func parseEpisode(extensions map[string][]Extension) (episode *PodcastEpisode) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["episode"]
	if !ok || len(matches) == 0 {
		return
	}

	episode = &PodcastEpisode{}
	episode.Display = parseTextAttrExtension("display", &matches[0])
	episode.Number = matches[0].Value

	return episode
}
