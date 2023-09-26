package ext

type PodcastRemoteItem struct {
	FeedGUID string `json:"feedGuid,omitempty"`
	FeedURL  string `json:"feedUrl,omitempty"`
	ItemGUID string `json:"itemGuid,omitempty"`
	Medium   string `json:"medium,omitempty"`
}

func parseRemoteItems(extensions map[string][]Extension) (remoteItems []*PodcastRemoteItem) {
	if extensions == nil {
		return nil
	}

	matches, ok := extensions["remoteItem"]
	if !ok || len(matches) == 0 {
		return nil
	}

	remoteItems = []*PodcastRemoteItem{}

	for _, cat := range matches {
		ri := &PodcastRemoteItem{}
		if text, ok := cat.Attrs["feedGuid"]; ok {
			ri.FeedGUID = text
		}

		if text, ok := cat.Attrs["feedUrl"]; ok {
			ri.FeedURL = text
		}

		if text, ok := cat.Attrs["itemGuid"]; ok {
			ri.ItemGUID = text
		}

		if text, ok := cat.Attrs["medium"]; ok {
			ri.Medium = text
		}

		remoteItems = append(remoteItems, ri)
	}

	return remoteItems
}
