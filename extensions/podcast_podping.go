package ext

import "strconv"

// PodcastPodping allows feed owners to signal to aggregators that the feed
// sends out Podping notifications when changes are made to it
//
// More Info: https://podcastindex.org/namespace/1.0#podping
//
// XML Example:
//
// <podcast:podping usesPodping="true"/>
type PodcastPodping struct {
	UsesPodping bool `json:"podping,omitempty"`
}

func parsePodping(extensions map[string][]Extension) (podping *PodcastPodping) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["podping"]
	if !ok || len(matches) == 0 {
		return
	}

	podping = &PodcastPodping{}

	if text, ok := matches[0].Attrs["usesPodping"]; ok {
		if val, err := strconv.ParseBool(text); err == nil {
			podping.UsesPodping = val
		}
	}

	return podping
}
