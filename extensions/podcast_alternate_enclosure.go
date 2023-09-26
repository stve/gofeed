package ext

import "strconv"

// PodcastAlternateEnclosure is meant to provide different versions of, or
// companion media to the main <enclosure> file.
//
// More Info: https://podcastindex.org/namespace/1.0#alternate-enclosure
//
// XML Example:
//
// <podcast:alternateEnclosure type="audio/opus" length="32400000" bitrate="96000" title="High quality">
//
//	<podcast:source uri="https://example.com/file-high.opus" />
//	<podcast:source uri="ipfs://someRandomHighBitrateOpusFile" />
//
// </podcast:alternateEnclosure>
type PodcastAlternateEnclosure struct {
	Type      string            `json:"type,omitempty"`
	Length    int               `json:"length,omitempty"`
	BitRate   float32           `json:"bitRate,omitempty"`
	Lang      string            `json:"lang,omitempty"`
	Title     string            `json:"title,omitempty"`
	Rel       string            `json:"rel,omitempty"`
	Codecs    string            `json:"codecs,omitempty"`
	Default   bool              `json:"default,omitempty"`
	Sources   []*PodcastSource  `json:"sources,omitempty"`
	Integrity *PodcastIntegrity `json:"integrity,omitempty"`
}

// PodcastSource element defines a uri location for a
// <podcast:alternateEnclosure> media file is meant to provide different versions of, or
//
// More Info: https://podcastindex.org/namespace/1.0#source
//
// XML Example:
//
// <podcast:alternateEnclosure type="video/mp4" length="7924786" bitrate="511276.52" height="720">
//
//	<podcast:source uri="https://example.com/file-720.mp4" />
//	<podcast:source uri="ipfs://QmX33FYehk6ckGQ6g1D9D3FqZPix5JpKstKQKbaS8quUFb" />
//
// </podcast:alternateEnclosure>
type PodcastSource struct {
	URI         string `json:"uri,omitempty"`
	ContentType string `json:"contentType,omitempty"`
}

// PodcastIntegrity element defines a uri location for a
// <podcast:alternateEnclosure> media file is meant to provide different versions of, or
//
// More Info: https://podcastindex.org/namespace/1.0#integrity
//
// XML Example:
//
// <podcast:alternateEnclosure type="video/mp4" length="7924786" bitrate="511276.52" height="720">
//
//	  <podcast:source uri="https://example.com/file-720.mp4" />
//		<podcast:integrity type="sri" value="sha384-ExVqijgYHm15PqQqdXfW95x+Rs6C+d6E/ICxyQOeFevnxNLR/wtJNrNYTjIysUBo" />
//
// </podcast:alternateEnclosure>
type PodcastIntegrity struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func parseAlternateEnclosures(extensions map[string][]Extension) (enclosures []*PodcastAlternateEnclosure) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["alternateEnclosure"]
	if !ok || len(matches) == 0 {
		return
	}

	enclosures = []*PodcastAlternateEnclosure{}

	for _, cat := range matches {
		enclosure := &PodcastAlternateEnclosure{}
		if text, ok := cat.Attrs["type"]; ok {
			enclosure.Type = text
		}

		if text, ok := cat.Attrs["length"]; ok {
			if val, err := strconv.Atoi(text); err == nil {
				enclosure.Length = val
			}
		}

		if text, ok := cat.Attrs["bitRate"]; ok {
			if val, err := strconv.ParseFloat(text, 32); err == nil {
				enclosure.BitRate = float32(val)
			}
		}

		if text, ok := cat.Attrs["lang"]; ok {
			enclosure.Lang = text
		}

		if text, ok := cat.Attrs["title"]; ok {
			enclosure.Title = text
		}

		if text, ok := cat.Attrs["rel"]; ok {
			enclosure.Rel = text
		}

		if text, ok := cat.Attrs["codecs"]; ok {
			enclosure.Codecs = text
		}

		if text, ok := cat.Attrs["default"]; ok {
			if val, err := strconv.ParseBool(text); err == nil {
				enclosure.Default = val
			}
		}

		enclosure.Sources = parseSources(cat.Children)
		enclosure.Integrity = parseIntegrity(cat.Children)

		enclosures = append(enclosures, enclosure)
	}

	return enclosures
}

func parseSources(extensions map[string][]Extension) (sources []*PodcastSource) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["source"]
	if !ok || len(matches) == 0 {
		return
	}

	sources = []*PodcastSource{}

	for _, cat := range matches {
		source := &PodcastSource{}

		if text, ok := cat.Attrs["contentType"]; ok {
			source.ContentType = text
		}

		if text, ok := cat.Attrs["uri"]; ok {
			source.URI = text
		}

		sources = append(sources, source)
	}

	return sources
}

func parseIntegrity(extensions map[string][]Extension) (integrity *PodcastIntegrity) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["integrity"]
	if !ok || len(matches) == 0 {
		return
	}

	integrity = &PodcastIntegrity{}

	if text, ok := matches[0].Attrs["type"]; ok {
		integrity.Type = text
	}

	if text, ok := matches[0].Attrs["value"]; ok {
		integrity.Value = text
	}

	return integrity
}
