package apiserver

import (
	"net/http"

	"Smart-city/internal/apiserver/store"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

type server struct {
	router       *mux.Router
	Logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

// Создаем сервер с определенной конфигурацией
func NewServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		Logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Configuration of router ...
func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/news", s.handleNews()).Methods("GET")
	s.router.HandleFunc("/news/{key}", s.Download()).Methods("GET")
	s.router.HandleFunc("/broadcast", s.handleBroadcast()).Methods("GET")
	s.router.HandleFunc("/broadcast/{key}", s.Download()).Methods("GET")

	s.router.HandleFunc("/upload/timetable", s.handleUploadTimetable()).Methods("POST")
	s.router.HandleFunc("/timetabel", s.handleGetTimetable()).Methods("POST")
}

// Configuration of logger ...
func (s *server) configureLogger(config *Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	s.Logger.SetLevel(level)

	return nil
}