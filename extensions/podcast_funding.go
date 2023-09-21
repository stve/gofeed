package ext

// A PodcastFunding lists possible donation/funding links for the podcast
//
// More Info: https://podcastindex.org/namespace/1.0#funding
//
// XML Examples:
//
//	<podcast:funding url="https://www.ex.com/donations">Support the show!</podcast:funding>
type PodcastFunding struct {
	Label string `json:"label,omitempty"`
	URL   string `json:"url,omitempty"`
}

func parseFunding(extensions map[string][]Extension) (funding []*PodcastFunding) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["funding"]
	if !ok || len(matches) == 0 {
		return
	}

	funding = []*PodcastFunding{}

	for _, cat := range matches {
		f := &PodcastFunding{}
		if url, ok := cat.Attrs["url"]; ok {
			f.URL = url
		}

		f.Label = cat.Value

		funding = append(funding, f)
	}

	return funding
}
