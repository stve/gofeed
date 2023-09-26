package ext

import "strconv"

// PodcastValue designates the cryptocurrency or payment layer that will be
// used, the transport method for transacting the payments, and a suggested
// amount denominated in the given cryptocurrency.
//
// More Info: https://podcastindex.org/namespace/1.0#value
//
// XML Example:
//
// <podcast:value type="lightning" method="keysend" suggested="0.00000005000"></podcast:value>
type PodcastValue struct {
	Type            string                   `json:"type,omitempty"`
	Method          string                   `json:"method,omitempty"`
	Suggested       string                   `json:"suggested,omitempty"`
	ValueRecipients []*PodcastValueRecipient `json:"recipients,omitempty"`
	ValueTimeSplit  []*PodcastValueTimeSplit `json:"timeSplit,omitempty"`
}

// PodcastValueRecipient designates various destinations for payments to be
// sent to during consumption of the enclosed media. Each recipient is
// considered to receive a "split" of the total payment according to the number
// of shares given in the split attribute.
//
// More Info: https://podcastindex.org/namespace/1.0#value-recipient
//
// XML Example:
//
// <podcast:valueRecipient name="Alice (Podcaster)" type="node" address="02d5c1bf8b940dc9cadca86d1b0a3c37fbe39cee4c7e839e33bef9174531d27f52" split="40" />
type PodcastValueRecipient struct {
	Name        string `json:"name,omitempty"`
	CustomKey   string `json:"customKey,omitempty"`
	CustomValue string `json:"customValue,omitempty"`
	Type        string `json:"type,omitempty"`
	Address     string `json:"address,omitempty"`
	Split       string `json:"split,omitempty"`
	Fee         bool   `json:"fee,omitempty"`
}

// PodcastValueTimeSplit allows different value splits for a certain period
// of time.
//
// More Info: https://podcastindex.org/namespace/1.0#value-time-split
//
// XML Example:
//
// <podcast:valueTimeSplit startTime="60" duration="237" remotePercentage="95">
//
//	<podcast:remoteItem itemGuid="https://podcastindex.org/podcast/4148683#1" feedGuid="a94f5cc9-8c58-55fc-91fe-a324087a655b" medium="music" />
//
// </podcast:valueTimeSplit>
type PodcastValueTimeSplit struct {
	StartTime        int                      `json:"startTime,omitempty"`
	Duration         int                      `json:"duration,omitempty"`
	RemoteStartTime  int                      `json:"remoteStartTime,omitempty"`
	RemotePercentage int                      `json:"remotePercentage,omitempty"`
	RemoteItem       *PodcastRemoteItem       `json:"remoteItem,omitempty"`
	ValueRecipients  []*PodcastValueRecipient `json:"recipients,omitempty"`
}

func parseValue(extensions map[string][]Extension) (value *PodcastValue) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["value"]
	if !ok || len(matches) == 0 {
		return
	}

	value = &PodcastValue{}

	if text, ok := matches[0].Attrs["method"]; ok {
		value.Method = text
	}

	if text, ok := matches[0].Attrs["suggested"]; ok {
		value.Suggested = text
	}

	if text, ok := matches[0].Attrs["type"]; ok {
		value.Type = text
	}

	value.ValueRecipients = parseValueRecipients(matches[0].Children)

	value.ValueTimeSplit = parseValueTimeSplit(matches[0].Children)

	return value
}

func parseValueRecipients(extensions map[string][]Extension) (recipients []*PodcastValueRecipient) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["valueRecipient"]
	if !ok || len(matches) == 0 {
		return
	}

	recipients = []*PodcastValueRecipient{}

	for _, cat := range matches {
		r := &PodcastValueRecipient{}
		if text, ok := cat.Attrs["name"]; ok {
			r.Name = text
		}

		if text, ok := cat.Attrs["customKey"]; ok {
			r.CustomKey = text
		}

		if text, ok := cat.Attrs["customValue"]; ok {
			r.CustomValue = text
		}

		if text, ok := cat.Attrs["type"]; ok {
			r.Type = text
		}

		if text, ok := cat.Attrs["address"]; ok {
			r.Address = text
		}

		if text, ok := cat.Attrs["split"]; ok {
			r.Split = text
		}

		if text, ok := cat.Attrs["fee"]; ok {
			if val, err := strconv.ParseBool(text); err == nil {
				r.Fee = val
			}
		}

		recipients = append(recipients, r)
	}

	return recipients
}

func parseValueTimeSplit(extensions map[string][]Extension) (valueTimeSplits []*PodcastValueTimeSplit) {
	if extensions == nil {
		return
	}

	matches, ok := extensions["valueTimeSplit"]
	if !ok || len(matches) == 0 {
		return
	}

	valueTimeSplits = []*PodcastValueTimeSplit{}

	for _, cat := range matches {
		vts := &PodcastValueTimeSplit{}
		if text, ok := cat.Attrs["duration"]; ok {
			if val, err := strconv.Atoi(text); err == nil {
				vts.Duration = val
			}
		}

		if text, ok := cat.Attrs["startTime"]; ok {
			if val, err := strconv.Atoi(text); err == nil {
				vts.StartTime = val
			}
		}

		if text, ok := cat.Attrs["remoteStartTime"]; ok {
			if val, err := strconv.Atoi(text); err == nil {
				vts.RemoteStartTime = val
			}
		}

		// If not defined, defaults to 100.
		vts.RemotePercentage = 100
		if text, ok := cat.Attrs["remotePercentage"]; ok {
			if val, err := strconv.Atoi(text); err == nil {
				vts.RemotePercentage = val
			}
		}

		remoteItems := parseRemoteItems(matches[0].Children)
		if len(remoteItems) > 0 {
			vts.RemoteItem = remoteItems[0]
		} else {
			vts.ValueRecipients = parseValueRecipients(matches[0].Children)
		}

		valueTimeSplits = append(valueTimeSplits, vts)
	}

	return valueTimeSplits
}
