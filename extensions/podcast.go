package ext

// PodcastFeedExtension is a set of extension
// fields for RSS feeds.
type PodcastFeedExtension struct {
	Blocks  []*PodcastBlock   `json:"blocks,omitempty"`
	Funding []*PodcastFunding `json:"funding,omitempty"`
	GUID    string            `json:"guid,omitempty"`
}

// PodcastItemExtension is a set of extension
// fields for RSS items.
type PodcastItemExtension struct{}

// NewPodcastFeedExtension creates a PodcastFeedExtension given an
// extension map for the "podcast" key.
func NewPodcastFeedExtension(extensions map[string][]Extension) *PodcastFeedExtension {
	feed := &PodcastFeedExtension{}
	feed.Blocks = parseBlocks(extensions)
	feed.Funding = parseFunding(extensions)
	feed.GUID = parseTextExtension("guid", extensions)
	return feed
}

// NewPodcastItemExtension creates a PodcastItemExtension given an
// extension map for the "podcast" key.
func NewPodcastItemExtension(extensions map[string][]Extension) *PodcastItemExtension {
	entry := &PodcastItemExtension{}
	return entry
}
