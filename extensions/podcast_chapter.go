package ext

// PodcastChapter links to an external file (see example file) containing
// chapter data for the episode is meant to provide different versions of, or
//
// More Info: https://podcastindex.org/namespace/1.0#chapters
//
// XML Example:
//
// <podcast:chapters url="https://example.com/episode1/chapters.json" type="application/json+chapters" />
type PodcastChapter struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

func parseChapters(extensions map[string][]Extension) (chapters *PodcastChapter) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["chapters"]
	if !ok || len(matches) == 0 {
		return
	}

	chapters = &PodcastChapter{}
	chapters.Type = parseTextAttrExtension("type", &matches[0])
	chapters.URL = parseTextAttrExtension("url", &matches[0])

	return chapters
}
