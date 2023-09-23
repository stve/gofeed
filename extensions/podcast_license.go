package ext

// PodcastLicense defines a license that is applied to the
// audio/video content of a single episode, or the audio/video of
// the podcast as a whole
//
// More Info: https://podcastindex.org/namespace/1.0#license
//
// XML Examples:
//
// <podcast:license>cc-by-4.0</podcast:license>
// <podcast:license url="https://example.org/mypodcastlicense/full.pdf">my-podcast-license-v1</podcast:license>
type PodcastLicense struct {
	Identifier string `json:"identifier,omitempty"`
	URL        string `json:"url,omitempty"`
}

func parseLicense(extensions map[string][]Extension) *PodcastLicense {
	if extensions == nil {
		return nil
	}

	matches, ok := extensions["license"]
	if !ok || len(matches) == 0 {
		return nil
	}

	license := &PodcastLicense{}

	license.Identifier = matches[0].Value

	if text, ok := matches[0].Attrs["url"]; ok {
		license.URL = text
	}

	return license
}
