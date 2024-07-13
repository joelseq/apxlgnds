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
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"startDate,omitempty"`
	EndDate     time.Time `json:"endDate,omitempty"`
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
