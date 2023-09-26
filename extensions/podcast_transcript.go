package ext

// PodcastTranscript is an element used to link to a transcript or closed
// captions file.
//
// More Info: https://podcastindex.org/namespace/1.0#transcript
//
// XML Example:
//
// <podcast:transcript url="https://example.com/episode1/transcript.vtt" type="text/vtt" />
type PodcastTranscript struct {
	URL      string `json:"url,omitempty"`
	Type     string `json:"type,omitempty"`
	Language string `json:"language,omitempty"`
	Rel      string `json:"rel,omitempty"`
}

func parseTranscripts(extensions map[string][]Extension) (transcripts []*PodcastTranscript) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["transcript"]
	if !ok || len(matches) == 0 {
		return
	}

	transcripts = []*PodcastTranscript{}

	for _, cat := range matches {
		t := &PodcastTranscript{}
		if text, ok := cat.Attrs["url"]; ok {
			t.URL = text
		}

		if text, ok := cat.Attrs["rel"]; ok {
			t.Rel = text
		}

		if text, ok := cat.Attrs["language"]; ok {
			t.Language = text
		}

		if text, ok := cat.Attrs["type"]; ok {
			t.Type = text
		}

		transcripts = append(transcripts, t)
	}

	return transcripts
}
