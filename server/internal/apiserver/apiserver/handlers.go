package apiserver

import (
	Constants "Smart-city/internal/apiserver"
	modeltimetable "Smart-city/internal/apiserver/model/timetable"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

// Обработчик для скачивания картинок для новостей (table news)
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
	type answer struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
		Txt   string `json:"text"`
		Time  int    `json:"time"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &modeltimetable.Timetable{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.Err(w, r, http.StatusBadRequest, err)
			s.Logger.Error(err)
			return
		}

		data, err := s.store.Timetable().GetTimetable(req.IdUser)
		if err != nil {
			s.Err(w, r, http.StatusInternalServerError, err)
			s.Logger.Error(err)
		}

		s.respond(w, r, http.StatusOK, data)
	}
}
