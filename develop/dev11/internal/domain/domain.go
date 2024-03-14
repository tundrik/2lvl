package domain

import (
	"encoding/json"
	"io"
	"time"
)


type Event struct {
	EventID     int       `json:"event_id"`
	UserID      int       `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" time_format:"2006-01-02" time_utc:"true"`
}

func (e *Event) Decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		return err
	}
	return nil
}