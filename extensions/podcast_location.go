package ext

// PodcastLocation  is intended to describe the location of editorial focus
// for a podcast's content
//
// More Info: https://podcastindex.org/namespace/1.0#location
//
// XML Example:
//
// <podcast:location geo="geo:30.2672,97.7431" osm="R113314">Austin, TX</podcast:location>
type PodcastLocation struct {
	Name          string `json:"name,omitempty"`
	Geo           string `json:"geo,omitempty"`
	OpenStreetMap string `json:"osm,omitempty"`
}

func parseLocation(extensions map[string][]Extension) *PodcastLocation {
	if extensions == nil {
		return nil
	}

	matches, ok := extensions["location"]
	if !ok || len(matches) == 0 {
		return nil
	}

	location := &PodcastLocation{}
	location.Name = matches[0].Value
	location.Geo = parseTextAttrExtension("geo", &matches[0])
	location.OpenStreetMap = parseTextAttrExtension("osm", &matches[0])

	return location
}
