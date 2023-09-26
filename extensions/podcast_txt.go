package ext

// PodcastTxt holds free-form text and is modeled after the DNS "TXT" record.
//
// More Info: https://podcastindex.org/namespace/1.0#txt
//
// XML Example:
//
// <podcast:txt>naj3eEZaWVVY9a38uhX8FekACyhtqP4JN</podcast:txt>
// <podcast:txt purpose="verify">S6lpp-7ZCn8-dZfGc-OoyaG</podcast:txt>
// <podcast:txt purpose="release">2022-10-26T04:45:30.742Z</podcast:txt>
type PodcastTxt struct {
	Purpose string `json:"purpose,omitempty"`
	Value   string `json:"value,omitempty"`
}

func parseTxt(extensions map[string][]Extension) (txts []*PodcastTxt) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["txt"]
	if !ok || len(matches) == 0 {
		return
	}

	txts = []*PodcastTxt{}

	for _, cat := range matches {
		t := &PodcastTxt{}
		if text, ok := cat.Attrs["purpose"]; ok {
			t.Purpose = text
		}

		t.Value = cat.Value

		txts = append(txts, t)
	}

	return txts
}
