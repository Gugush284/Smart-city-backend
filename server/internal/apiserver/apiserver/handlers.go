package apiserver

import (
	Constants "Smart-city/internal/apiserver"
	modelEvents "Smart-city/internal/apiserver/model/event"
	modelScoreboard "Smart-city/internal/apiserver/model/scoreboard"
	modeltimetable "Smart-city/internal/apiserver/model/timetable"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		uuID := uuid.New().String()
		w.Header().Set("X-Request-ID", uuID)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), Constants.CtxKeyId, uuID)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.Logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(Constants.CtxKeyId),
		})
		logger.Infof("started %s, %s", r.Method, r.RequestURI)

		start := time.Now()

		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof("completed with %d: %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start))
	})
}

func (s *server) handleNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.News().GetNews()
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
		}

		s.respond(w, r, http.StatusOK, data)
	}
}

// Обработчик для скачивания картинок
func (s *server) Download() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path

		path := strings.Join([]string{"assets", key}, "/")
		fileBytes, err := ioutil.ReadFile(path)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
	})
}

func (s *server) handleBroadcast() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := s.store.Broadcast().GetBroadcast()
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
		}

		s.respond(w, r, http.StatusOK, data)
	})
}

func (s *server) handleUploadTimetable() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &modeltimetable.Timetable{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Err(w, r, http.StatusBadRequest, err)
			s.Logger.Error(err)
			return
		}

		if err := s.store.Timetable().Create(req); err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		s.respond(w, r, http.StatusOK, req.Id)
	}
}

func (s *server) handleGetTimetable() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.ReplaceAll(r.URL.Path, "/TimeTabel/", "")
		s.Logger.Info(key)
		id, err := strconv.Atoi(key)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		data, err := s.store.Timetable().GetTimetable(id)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		s.respond(w, r, http.StatusOK, data)
	}
}

// Поставили заглушку, так как не успевали
func (s *server) handleGetScoreboard() http.HandlerFunc {
	answer := &modelScoreboard.Scoreboard{
		FirstTeam:    "Спартак",
		SecondTeam:   "Динамо",
		FirstNumber:  "1",
		SecondNumber: "2",
		Type:         "Футбол",
		URL:          "https://www.youtube.com/watch?v=dGaMze_rOWk",
		Term:         "1/2",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, answer)
	}
}

func (s *server) handleGetTeam() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.ReplaceAll(r.URL.Path, "/TeAm/", "")
		s.Logger.Info(key)
		id, err := strconv.Atoi(key)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		data, err := s.store.Teams().GetTeams(id)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
			return
		}

		s.respond(w, r, http.StatusOK, data)
	}
}

// Поставили заглушку, так как не успевали
func (s *server) handleGetRegions() http.HandlerFunc {

	regionList := []struct {
		Place string `json:"place"`
		Name  string `json:"region"`
		Score string `json:"points"`
	}{
		{
			Place: "1",
			Name:  "Мурманская область",
			Score: "15",
		},
		{
			Place: "2",
			Name:  "Иркутская область",
			Score: "12",
		},
		{
			Place: "3",
			Name:  "Московская область",
			Score: "6",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, regionList)
	}
}

// Иммитация push от админки
func (s *server) handlePush() http.HandlerFunc {
	type push struct {
		Title string `json:"title"`
		Txt   string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if Fakepush != 0 {
			Fakepush = 0

			s.respond(w, r, http.StatusCreated, push{
				Title: "Уведомление",
				Txt:   "Мероприятие X переносится",
			})
		} else {
			s.respond(w, r, http.StatusOK, nil)
		}

	}
}

func (s *server) handleMes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Fakepush = 1
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleGetEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, err := s.store.Event().GetEvents()
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
		}

		s.respond(w, r, http.StatusOK, data)
	}
}

func (s *server) handleRegEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req := &modelEvents.EventRegistratePLayers{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Err(w, r, http.StatusBadRequest, err)
			s.Logger.Error(err)
			return
		}

		data, err := s.store.Event().RegEvent(req)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
		}

		s.Logger.Info(data)
		s.respond(w, r, http.StatusOK, data)
	}
}

func (s *server) handleGetEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.ReplaceAll(r.URL.Path, "/event/", "")

		data, err := s.store.Event().GetEvent(key)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
		}

		s.respond(w, r, http.StatusOK, data)
	}
}
