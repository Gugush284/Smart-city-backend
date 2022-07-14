package apiserver

import (
	"Smart-city/internal/apiserver/store/sqlstore"
	"net/http"

	"github.com/gorilla/sessions"
)

// Функция запускает apiserver
func Start(config *Config) error {
	// Создаем хранилище
	store := sqlstore.New(config.DatabaseURL)

	// Создаем хранилище куков
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))

	// Создаем экземпляр сервера
	srv := NewServer(store, sessionStore)

	// Конфиг для логера
	if err := srv.configureLogger(config); err != nil {
		return err
	}

	// Создаем таблицы в нашем хранилище
	if err := store.CreateTables(); err != nil {
		return err
	}

	srv.Logger.Info("starting api server")
	srv.Logger.Debug(config.SessionKey)
	defer store.Db.Close()

	// Прослушиваем порт
	return http.ListenAndServe(config.BindAddr, srv)
}
