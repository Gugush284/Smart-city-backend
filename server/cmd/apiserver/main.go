package main

import (
	"Smart-city/internal/apiserver/apiserver"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var configPath string
var sessionKeyPath string

// Выставляем флаги перед началом программы
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.StringVar(&sessionKeyPath, "sessionKey-Path", "configs/SessionKey.toml", "path to SessionKey")
}

func main() {
	flag.Parse()

	// Получаем файл конфигурации
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Получаем файл с ключом сессии
	_, err = toml.DecodeFile(sessionKeyPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Запускаем apiserver
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
