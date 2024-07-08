package types

import (
	"bytes"
	"encoding/gob"
	"time"
)

type CalendarEventsResponse struct {
	Events []Event `json:"events"`
}

type Event struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date,omitempty"`
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
