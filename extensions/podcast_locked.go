package ext

// PodcastLocked tells podcast hosting platforms whether they are allowed to
// import this feed. A value of yes means that any attempt to import this feed
// into a new platform should be rejected.
//
// More Info: https://podcastindex.org/namespace/1.0#locked
//
// XML Example:
//
// <podcast:locked owner="email@example.com">no</podcast:locked>
type PodcastLocked struct {
	Locked bool   `json:"locked"`
	Owner  string `json:"owner,omitempty"`
}

func parseLocked(extensions map[string][]Extension) *PodcastLocked {
	if extensions == nil {
		return nil
	}

	matches, ok := extensions["locked"]
	if !ok || len(matches) == 0 {
		return nil
	}

	locked := &PodcastLocked{}

	locked.Locked = parseBoolean(matches[0].Value)
	locked.Owner = parseTextAttrExtension("owner", &matches[0])

	return locked
}
