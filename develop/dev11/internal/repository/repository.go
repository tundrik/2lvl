package repository

import (
	"2lvl/develop/dev11/internal/domain"
	"fmt"
	"sync"
	"time"
)

type Repo struct {
	events  map[int][]domain.Event
	sync.RWMutex
}

func New() *Repo {
	return &Repo{
		events:  make(map[int][]domain.Event),
	}
}

// CreateEvent - сохраняет новое событие
func (r *Repo) CreateEvent(e domain.Event) error {
	r.Lock()
	defer r.Unlock()

	if events, ok := r.events[e.UserID]; ok {
		for _, event := range events {
			if event.EventID == e.EventID {
				return fmt.Errorf("event already exist")
			}
		}
	}
	r.events[e.UserID] = append(r.events[e.UserID], e)
	return nil
}

// UpdateEvent - обновляет новое событие
func (r *Repo) UpdateEvent(event domain.Event) error {
	r.Lock()
	defer r.Unlock()

	ind := -1
	events, ok := r.events[event.UserID]
	if !ok {
		return fmt.Errorf("user does not find")
	}

	for i, e := range events {
		if e.EventID == event.EventID {
			ind = i
			break
		}
	}
	if ind == -1 {
		return fmt.Errorf("event does not exist")
	}
	r.events[event.UserID][ind] = event

	return nil
}

// DeleteEvent - удаляет новое событие
func (r *Repo) DeleteEvent(userID, eventID int) (*domain.Event, error) {
	r.Lock()
	defer r.Unlock()

	ind := -1
	events, ok := r.events[userID]
	if !ok {
		return nil, fmt.Errorf("user does not exist")
	}
	for i, event := range events {
		if event.EventID == eventID {
			ind = i
			break
		}
	}
	if ind == -1 {
		return nil, fmt.Errorf("event does not exist")
	}
	eventsLength := len(r.events[userID])
	deletedEvent := r.events[userID][ind]
	r.events[userID][ind] = r.events[userID][eventsLength-1]
	r.events[userID] = r.events[userID][:eventsLength-1]

	return &deletedEvent, nil
}

func (r *Repo) GetEventsForDay(userID int, date time.Time) ([]domain.Event, error) {
	r.RLock()
	defer r.RUnlock()

	var result []domain.Event

	events, ok := r.events[userID]

	if !ok {
		return nil, fmt.Errorf("user does not exist")
	}
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			result = append(result, event)
		}
	}
	return result, nil
}

func (r *Repo) GetEventsForWeek(userID int, date time.Time) ([]domain.Event, error) {
	r.RLock()
	defer r.RUnlock()
	result := make([]domain.Event, 0)
	events, ok := r.events[userID]
	if !ok {
		return nil, fmt.Errorf("user does not exist")
	}
	for _, event := range events {
		eventYear, eventWeek := event.Date.ISOWeek()
		currentYear, currentWeek := date.ISOWeek()
		if eventYear == currentYear && eventWeek == currentWeek {
			result = append(result, event)
		}
	}

	return result, nil
}

func (r *Repo) GetEventsForMonth(userID int, date time.Time) ([]domain.Event, error) {
	r.RLock()
	defer r.RUnlock()
	result := make([]domain.Event, 0)
	events, ok := r.events[userID]
	if !ok {
		return nil, fmt.Errorf("user does not exist")
	}
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			result = append(result, event)
		}
	}
	return result, nil
}