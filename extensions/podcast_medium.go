package ext

import "strings"

// PodcastMedium tells an application what the content contained within
// the feed IS, as opposed to what the content is ABOUT in the case of a
// category.
//
// More Info: https://podcastindex.org/namespace/1.0#medium
//
// XML Example:
//
// <podcast:medium>podcast</podcast:medium>
// <podcast:medium>musicL</podcast:medium>
type PodcastMedium struct {
	Type        string               `json:"type,omitempty"`
	RemoteItems []*PodcastRemoteItem `json:"remoteItems,omitempty"`
}

func parseMedium(extensions map[string][]Extension) *PodcastMedium {
	if extensions == nil {
		return nil
	}

	matches, ok := extensions["medium"]
	if !ok || len(matches) == 0 {
		return nil
	}

	medium := &PodcastMedium{}
	medium.Type = matches[0].Value

	// Check if list type
	if strings.HasSuffix(medium.Type, "L") == false {
		return medium
	}

	medium.RemoteItems = parseRemoteItems(extensions)

	return medium
}
