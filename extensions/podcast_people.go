package ext

// PodcastPerson specifies a person of interest to the podcast. It is primarily
// intended to identify people like hosts, co-hosts and guests.
//
// More Info: https://podcastindex.org/namespace/1.0#person
//
// XML Example:
//
// <podcast:person
//
//	href="https://example.com/johnsmith/blog"
//	img="http://example.com/images/johnsmith.jpg"
//
// >John Smith</podcast:person>
// <podcast:person
//
//	role="guest"
//	href="https://www.imdb.com/name/nm0427852888/"
//	img="http://example.com/images/janedoe.jpg"
//
// >Jane Doe</podcast:person>
type PodcastPerson struct {
	Name  string `json:"name,omitempty"`
	Role  string `json:"role,omitempty"`
	Group string `json:"group,omitempty"`
	Image string `json:"img,omitempty"`
	URL   string `json:"href,omitempty"`
}

func parsePeople(extensions map[string][]Extension) (people []*PodcastPerson) {
	if extensions == nil {
		return nil
	}

	matches, ok := extensions["person"]
	if !ok || len(matches) == 0 {
		return nil
	}

	people = []*PodcastPerson{}

	for _, cat := range matches {
		p := &PodcastPerson{}
		if text, ok := cat.Attrs["group"]; ok {
			p.Group = text
		}

		if text, ok := cat.Attrs["role"]; ok {
			p.Role = text
		}

		if text, ok := cat.Attrs["img"]; ok {
			p.Image = text
		}

		if text, ok := cat.Attrs["href"]; ok {
			p.URL = text
		}

		p.Name = cat.Value

		people = append(people, p)
	}

	return people
}
