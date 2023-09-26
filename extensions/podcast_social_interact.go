package ext

import "strconv"

// PodcastSocialInteract allows a podcaster to attach the url of a "root post"
// of a comment thread to an episode.
//
// More Info: https://podcastindex.org/namespace/1.0#social-interact
//
// XML Example:
//
// <podcast:socialInteract
//
//	uri="https://podcastindex.social/web/@dave/108013847520053258"
//	protocol="activitypub"
//	accountId="@dave"
//
// />
type PodcastSocialInteract struct {
	URI        string `json:"uri,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
	AccountID  string `json:"accountId,omitempty"`
	AccountURL string `json:"accountUrl,omitempty"`
	Priority   int    `json:"priority,omitempty"`
}

func parseSocialInteractions(extensions map[string][]Extension) (socialInteracts []*PodcastSocialInteract) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["socialInteract"]
	if !ok || len(matches) == 0 {
		return
	}

	socialInteracts = []*PodcastSocialInteract{}

	for _, cat := range matches {
		si := &PodcastSocialInteract{}

		if text, ok := cat.Attrs["priority"]; ok {
			if val, err := strconv.Atoi(text); err == nil {
				si.Priority = val
			}
		}

		if text, ok := cat.Attrs["uri"]; ok {
			si.URI = text
		}

		if text, ok := cat.Attrs["protocol"]; ok {
			si.Protocol = text
		}

		if text, ok := cat.Attrs["accountId"]; ok {
			si.AccountID = text
		}

		if text, ok := cat.Attrs["accountUrl"]; ok {
			si.AccountURL = text
		}

		socialInteracts = append(socialInteracts, si)
	}

	return socialInteracts
}
