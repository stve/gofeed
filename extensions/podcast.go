package ext

// PodcastFeedExtension is a set of extension
// fields for RSS feeds.
type PodcastFeedExtension struct{}

// PodcastItemExtension is a set of extension
// fields for RSS items.
type PodcastItemExtension struct{}

// NewPodcastFeedExtension creates a PodcastFeedExtension given an
// extension map for the "podcast" key.
func NewPodcastFeedExtension(extensions map[string][]Extension) *PodcastFeedExtension {
	feed := &PodcastFeedExtension{}
	return feed
}

// NewPodcastItemExtension creates a PodcastItemExtension given an
// extension map for the "podcast" key.
func NewPodcastItemExtension(extensions map[string][]Extension) *PodcastItemExtension {
	entry := &PodcastItemExtension{}
	return entry
}
