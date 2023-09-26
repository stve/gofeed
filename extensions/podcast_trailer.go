package ext

import "strconv"

// PodcastTrailer is used to define the location of an audio or video file to
// be used as a trailer for the entire podcast or a specific season.
//
// More Info: https://podcastindex.org/namespace/1.0#trailer
//
// XML Example:
//
// <podcast:trailer
// pubdate="Thu, 01 Apr 2021 08:00:00 EST"
// url="https://example.org/trailers/teaser"
// length="12345678"
// type="audio/mp3">Coming April 1st, 2021</podcast:trailer>
type PodcastTrailer struct {
	Title   string `json:"title,omitempty"`
	URL     string `json:"url,omitempty"`
	PubDate string `json:"pubDate,omitempty"`
	Length  int    `json:"length,omitempty"`
	Type    string `json:"type,omitempty"`
	Season  int    `json:"season,omitempty"`
}

func parseTrailers(extensions map[string][]Extension) (trailers []*PodcastTrailer) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["trailer"]
	if !ok || len(matches) == 0 {
		return
	}

	trailers = []*PodcastTrailer{}

	for _, cat := range matches {
		t := &PodcastTrailer{}
		if text, ok := cat.Attrs["url"]; ok {
			t.URL = text
		}

		if text, ok := cat.Attrs["pubdate"]; ok {
			t.PubDate = text
		}

		if text, ok := cat.Attrs["length"]; ok {
			if length, err := strconv.Atoi(text); err == nil {
				t.Length = length
			}
		}

		if text, ok := cat.Attrs["type"]; ok {
			t.Type = text
		}

		if text, ok := cat.Attrs["season"]; ok {
			if season, err := strconv.Atoi(text); err == nil {
				t.Season = season
			}
		}

		t.Title = cat.Value

		trailers = append(trailers, t)
	}

	return trailers
}
