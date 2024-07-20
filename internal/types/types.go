package types

import (
	"bytes"
	"encoding/gob"
	"time"
)

type EventGroup struct {
	Upcoming []Event `json:"upcoming,omitempty"`
	Recent   []Event `json:"recent,omitempty"`
}

type CalendarEventsResponse struct {
	ALGS  *EventGroup `json:"algs,omitempty"`
	Other *EventGroup `json:"other,omitempty"`
}

type Event struct {
	StartDate   time.Time     `json:"startDate,omitempty"`
	EndDate     time.Time     `json:"endDate,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Metadata    EventMetadata `json:"metadata,omitempty"`
}

type EventMetadata struct {
	Reddit      *RedditMetadata `json:"reddit,omitempty"`
	Region      string          `json:"region,omitempty"`
	BattlefyURL string          `json:"battlefy_url,omitempty"`
	Day         int             `json:"day,omitempty"`
	IsFinals    bool            `json:"is_finals,omitempty"`
}

type RedditMetadata struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}

func EncodeResponse(res *CalendarEventsResponse) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(*res)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodeResponse(val []byte) (*CalendarEventsResponse, error) {
	buf := bytes.NewBuffer(val)
	dec := gob.NewDecoder(buf)

	var res CalendarEventsResponse

	err := dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type Region string

const (
	RegionAPACNorth Region = "APAC-N"
	RegionAPACSouth Region = "APAC-S"
	RegionEMEA      Region = "EMEA"
	RegionNA        Region = "NA"
	RegionUnknown   Region = "Unknown"
)

func (r Region) URLParam() string {
	switch r {
	case RegionAPACNorth:
		return "asia-pacific-north"
	case RegionAPACSouth:
		return "asia-pacific-south"
	case RegionEMEA:
		return "europe-middle-east-and-africa"
	case RegionNA:
		return "north-america"
	default:
		panic("Unknown region provided")
	}
}
