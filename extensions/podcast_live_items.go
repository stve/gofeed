package ext

// PodcastLiveItem  tag is used for a feed to deliver a live audio or video stream to podcast apps.
//
// More Info: https://podcastindex.org/namespace/1.0#live-item
//
// XML Examples:
//
// <podcast:liveItem status="live" start="2021-09-26T07:30:00.000-0600" end="2021-09-26T09:30:00.000-0600">
//
//	<title>Podcasting 2.0 Live Show</title>
//
// </podcast:liveItem>
type PodcastLiveItem struct {
	Status       string                `json:"status,omitempty"`
	Start        string                `json:"start,omitempty"`
	End          string                `json:"end,omitempty"`
	ContentLinks []*PodcastContentLink `json:"contentLinks,omitempty"`
}

// PodcastContentlink is used to indicate that the content begin delivered
// by the parent element can be found at an external location instead of, or
// in addition to, being delivered directly to the tag itself within an app.
//
// More Info: https://podcastindex.org/namespace/1.0#block
//
// XML Examples:
//
//	<podcast:contentLink href="https://youtube.com/blahblah/livestream">Live on YouTube!</podcast:contentLink>
type PodcastContentLink struct {
	Label string `json:"label,omitempty"`
	URL   string `json:"url,omitempty"`
}

func parseLiveItems(extensions map[string][]Extension) (liveItems []*PodcastLiveItem) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["liveItem"]
	if !ok || len(matches) == 0 {
		return
	}

	liveItems = []*PodcastLiveItem{}

	for _, cat := range matches {
		li := &PodcastLiveItem{}
		if text, ok := cat.Attrs["status"]; ok {
			li.Status = text
		}

		if text, ok := cat.Attrs["start"]; ok {
			li.Start = text
		}

		if text, ok := cat.Attrs["end"]; ok {
			li.End = text
		}

		li.ContentLinks = parseContentLinks(cat.Children)

		liveItems = append(liveItems, li)
	}

	return liveItems
}

func parseContentLinks(extensions map[string][]Extension) (contentLinks []*PodcastContentLink) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["contentLink"]
	if !ok || len(matches) == 0 {
		return
	}

	contentLinks = []*PodcastContentLink{}

	for _, cat := range matches {
		cl := &PodcastContentLink{}

		cl.Label = cat.Value

		if text, ok := cat.Attrs["href"]; ok {
			cl.URL = text
		}

		contentLinks = append(contentLinks, cl)
	}

	return contentLinks
}
