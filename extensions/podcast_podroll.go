package ext

// PodcastPodroll allows for a podcaster to include references to one or more
// podcasts in it's <channel> as a way of "recommending" other podcasts to
// their listener.
//
// More Info: https://podcastindex.org/namespace/1.0#podroll
//
// XML Example:
//
// <podcast:podroll>
//
//	<podcast:remoteItem feedGuid="29cdca4a-32d8-56ba-b48b-09a011c5daa9" />
//	<podcast:remoteItem feedGuid="396d9ae0-da7e-5557-b894-b606231fa3ea" />
//	<podcast:remoteItem feedGuid="917393e3-1b1e-5cef-ace4-edaa54e1f810" />
//
// </podcast:podroll>
type PodcastPodroll struct {
	RemoteItems []*PodcastRemoteItem `json:"remoteItems,omitempty"`
}

func parsePodroll(extensions map[string][]Extension) (podroll *PodcastPodroll) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["podroll"]
	if !ok || len(matches) == 0 {
		return
	}

	podroll = &PodcastPodroll{}
	podroll.RemoteItems = parseRemoteItems(matches[0].Children)

	return podroll
}
