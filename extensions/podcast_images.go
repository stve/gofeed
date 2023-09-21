package ext

import "strings"

// PodcastImages allows for specifying many different image sizes.
// The syntax is borrowed from the HTML5 srcset syntax. It allows for
// describing multiple image sources with width and pixel hints directly
// in the attribute.
//
// More Info: https://podcastindex.org/namespace/1.0#images
//
// XML Examples:
//
//		<podcast:images
//		  srcset="https://example.com/images/ep1/pci_avatar-massive.jpg 1500w,
//	 	  https://example.com/images/ep1/pci_avatar-middle.jpg 600w,
//			https://example.com/images/ep1/pci_avatar-small.jpg 300w,
//			https://example.com/images/ep1/pci_avatar-tiny.jpg 150w" />
type PodcastImages struct {
	SrcSet map[string]string `json:"srcset,omitempty"`
}

func parseImages(extensions map[string][]Extension) *PodcastImages {
	if extensions == nil {
		return nil
	}

	matches, ok := extensions["images"]
	if !ok || len(matches) == 0 {
		return nil
	}

	images := &PodcastImages{}

	if srcset, ok := matches[0].Attrs["srcset"]; ok {
		sourceStrings := strings.Split(srcset, ",")

		sources := make(map[string]string)

		for _, source := range sourceStrings {
			// remove outer and duplicate whitespace in string
			source = strings.Join(strings.Fields(strings.TrimSpace(source)), " ")

			sourceKeyVals := strings.Split(source, " ")

			url := sourceKeyVals[0]
			size := sourceKeyVals[1]

			sources[size] = url
		}

		images.SrcSet = sources
	}

	return images
}
