package main

import (
	"fmt"
	"gostore/cmd/app"
	"gostore/internal/pkg/config"
	"gostore/internal/pkg/db"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var cfg config.Config
var DBConn *gorm.DB

func init() {
	// read config
	loadConfig, err := config.LoadConfig("./config/")
	if err != nil {
		log.Error(err.Error())
	}

	cfg = loadConfig

	// connect database
	db, err := db.ConnectDB(cfg.PostgresUri)
	if err != nil {
		log.Error(err.Error())
	}

	DBConn = db

	// setup logrus
	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)                 // apply log level
	log.SetFormatter(&log.JSONFormatter{}) // define format using json
}

func main() {
	server, err := app.NewServer(cfg, DBConn)
	if err != nil {
		log.Fatal("cannot init server")
	}

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	if err := server.Start(appPort); err != nil {
		log.Fatal(fmt.Errorf("cannot start app: %w", err))
	}
}
