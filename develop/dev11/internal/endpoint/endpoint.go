package endpoint

import (
	"2lvl/develop/dev11/internal/domain"
	"2lvl/develop/dev11/internal/repository"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

const layout = "2006-01-02"

type Endpoint struct {
	mux   *http.ServeMux
	repo  *repository.Repo
}

func New(repo *repository.Repo) *Endpoint {
	var e Endpoint
	e.mux = http.NewServeMux()
    
	e.mux.HandleFunc("/create_event", middleware(e.createEvent))
	e.mux.HandleFunc("/update_event", middleware(e.updateEvent))
	e.mux.HandleFunc("/delete_event", middleware(e.deleteEvent))

	e.mux.HandleFunc("/events_for_day", middleware(e.eventsForDay))
	e.mux.HandleFunc("/events_for_week", middleware(e.eventsForWeek))
	e.mux.HandleFunc("/events_for_month", middleware(e.eventsForMonth))

	e.repo = repo
	return &e
}

// Run запускает сервер
func (e *Endpoint) Run() error {
	return http.ListenAndServe(":8080", e.mux)
}

// createEvent обрабатывает запрос на создание события
func (e *Endpoint) createEvent(w http.ResponseWriter, r *http.Request) {
	var d domain.Event
	if err := d.Decode(r.Body); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}

	if err := e.repo.CreateEvent(d); err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Event is created", []domain.Event{d}, http.StatusOK)
}

// updateEvent обрабатывает запрос на обновление события
func (e *Endpoint) updateEvent(w http.ResponseWriter, r *http.Request) {
	var d domain.Event
	if err := d.Decode(r.Body); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	if err := e.repo.UpdateEvent(d); err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Event is updated", []domain.Event{d}, http.StatusOK)
}

// deleteEvent обрабатывает запрос на удаление события
func (e *Endpoint) deleteEvent(w http.ResponseWriter, r *http.Request) {
	var d domain.Event
	if err := d.Decode(r.Body); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	deletedEvent, err := e.repo.DeleteEvent(d.UserID, d.EventID)
	if err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Event has been deleted", []domain.Event{*deletedEvent}, http.StatusOK)
}

// eventsForDay обрабатывает запрос на получение событий на конкретный день
func (e *Endpoint) eventsForDay(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	date, err := time.Parse(layout, r.URL.Query().Get("date"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	events, err := e.repo.GetEventsForDay(userID, date)
	if err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Events foud", events, http.StatusOK)
}

// eventsForWeek обрабатывает запрос на получение событий на 7 дней
func (e *Endpoint) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	date, err := time.Parse(layout, r.URL.Query().Get("date"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	events, err := e.repo.GetEventsForWeek(userID, date)
	if err != nil {
		respondWithError(w, err, http.StatusServiceUnavailable)
		return
	}
	respondWithResult(w, "Events foud", events, http.StatusOK)
}

// eventsForMonth обрабатывает запрос на получение событий на месяц
func (e *Endpoint) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	date, err := time.Parse(layout, r.URL.Query().Get("date"))
	if err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	var events []domain.Event
	if events, err = e.repo.GetEventsForMonth(userID, date); err != nil {
		respondWithError(w, err, http.StatusBadRequest)
		return
	}
	respondWithResult(w, "Events foud", events, http.StatusOK)
}

func respondWithError(w http.ResponseWriter, e error, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{e.Error()}

	respond(errorResponse, w, status)

}

func respondWithResult(w http.ResponseWriter, r string, e []domain.Event, status int) {
	resultResponse := struct {
		Result string        `json:"result"`
		Events []domain.Event `json:"events"`
	}{r, e}

	respond(resultResponse, w, status)
}

func respond(response interface{}, w http.ResponseWriter, status int) {
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}