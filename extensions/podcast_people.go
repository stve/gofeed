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
		p.Group = parseTextAttrExtension("group", &cat)
		p.Role = parseTextAttrExtension("role", &cat)
		p.Image = parseTextAttrExtension("img", &cat)
		p.URL = parseTextAttrExtension("href", &cat)
		p.Name = cat.Value

		people = append(people, p)
	}

	return people
}
