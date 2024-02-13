package ext

// PodcastBlock allows a podcaster to express which platforms
// are allowed to publicly display this feed and its contents.
//
// More Info: https://podcastindex.org/namespace/1.0#block
//
// XML Examples:
//
//	<!-- This means "block everything" -->
//	<podcast:block>yes</podcast:block>
//
//	<!-- This means "block only google and amazon" -->
//	<podcast:block id="google">yes</podcast:block>
//	<podcast:block id="amazon">yes</podcast:block>
type PodcastBlock struct {
	PlatformID string `json:"ID,omitempty"`
	Blocked    bool   `json:"blocked"`
}

func parseBlocks(extensions map[string][]Extension) (blocks []*PodcastBlock) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["block"]
	if !ok || len(matches) == 0 {
		return
	}

	blocks = []*PodcastBlock{}

	for _, cat := range matches {
		f := &PodcastBlock{}
		f.PlatformID = parseTextAttrExtension("id", &cat)
		f.Blocked = parseBoolean(cat.Value)

		blocks = append(blocks, f)
	}

	return blocks
}
