package ext

import "strings"

// PodcastFeedExtension is a set of extension
// fields for RSS feeds.
type PodcastFeedExtension struct {
	Blocks          []*PodcastBlock         `json:"blocks,omitempty"`
	Funding         []*PodcastFunding       `json:"funding,omitempty"`
	GUID            string                  `json:"guid,omitempty"`
	Images          *PodcastImages          `json:"images,omitempty"`
	License         *PodcastLicense         `json:"license,omitempty"`
	LiveItems       []*PodcastLiveItem      `json:"liveItems,omitempty"`
	Location        *PodcastLocation        `json:"location,omitempty"`
	Locked          *PodcastLocked          `json:"locked,omitempty"`
	Medium          *PodcastMedium          `json:"medium,omitempty"`
	People          []*PodcastPerson        `json:"people,omitempty"`
	Podping         *PodcastPodping         `json:"podping,omitempty"`
	Podroll         *PodcastPodroll         `json:"podroll,omitempty"`
	Trailers        []*PodcastTrailer       `json:"trailers,omitempty"`
	Txt             []*PodcastTxt           `json:"txt,omitempty"`
	UpdateFrequency *PodcastUpdateFrequency `json:"updateFrequency,omitempty"`
	Value           *PodcastValue           `json:"value,omitempty"`
}

// NewPodcastFeedExtension creates a PodcastFeedExtension given an
// extension map for the "podcast" key.
func NewPodcastFeedExtension(extensions map[string][]Extension) *PodcastFeedExtension {
	feed := &PodcastFeedExtension{}
	feed.Blocks = parseBlocks(extensions)
	feed.Funding = parseFunding(extensions)
	feed.GUID = parseTextExtension("guid", extensions)
	feed.Images = parseImages(extensions)
	feed.License = parseLicense(extensions)
	feed.LiveItems = parseLiveItems(extensions)
	feed.Location = parseLocation(extensions)
	feed.Locked = parseLocked(extensions)
	feed.Medium = parseMedium(extensions)
	feed.People = parsePeople(extensions)
	feed.Podping = parsePodping(extensions)
	feed.Podroll = parsePodroll(extensions)
	feed.Trailers = parseTrailers(extensions)
	feed.Txt = parseTxt(extensions)
	feed.UpdateFrequency = parseUpdateFrequency(extensions)
	feed.Value = parseValue(extensions)
	return feed
}

// PodcastItemExtension is a set of extension
// fields for RSS items.
type PodcastItemExtension struct {
	AlternateEnclosures []*PodcastAlternateEnclosure `json:"alternateEnclosures,omitempty"`
	Chapters            *PodcastChapter              `json:"chapters,omitempty"`
	Episode             *PodcastEpisode              `json:"episode,omitempty"`
	Images              *PodcastImages               `json:"images,omitempty"`
	License             *PodcastLicense              `json:"license,omitempty"`
	People              []*PodcastPerson             `json:"people,omitempty"`
	Season              *PodcastSeason               `json:"season,omitempty"`
	SocialInteractions  []*PodcastSocialInteract     `json:"socialInteract,omitempty"`
	Soundbites          []*PodcastSoundbite          `json:"soundbites,omitempty"`
	Transcripts         []*PodcastTranscript         `json:"transcripts,omitempty"`
	Txt                 []*PodcastTxt                `json:"txt,omitempty"`
	Value               *PodcastValue                `json:"value,omitempty"`
}

// NewPodcastItemExtension creates a PodcastItemExtension given an
// extension map for the "podcast" key.
func NewPodcastItemExtension(extensions map[string][]Extension) *PodcastItemExtension {
	entry := &PodcastItemExtension{}
	entry.AlternateEnclosures = parseAlternateEnclosures(extensions)
	entry.Chapters = parseChapters(extensions)
	entry.Episode = parseEpisode(extensions)
	entry.Images = parseImages(extensions)
	entry.License = parseLicense(extensions)
	entry.People = parsePeople(extensions)
	entry.Season = parseSeason(extensions)
	entry.SocialInteractions = parseSocialInteractions(extensions)
	entry.Soundbites = parseSoundbites(extensions)
	entry.Transcripts = parseTranscripts(extensions)
	entry.Txt = parseTxt(extensions)
	entry.Value = parseValue(extensions)
	return entry
}

func parseBoolean(val string) bool {
	lower := strings.ToLower(val)

	if lower == "yes" {
		return true
	}

	return false
}
