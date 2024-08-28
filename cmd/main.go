package main

import (
	"SIMPLE-JWT-ECHO/config"
	"SIMPLE-JWT-ECHO/internal/connection"
	"SIMPLE-JWT-ECHO/internal/database/postgres"
	"SIMPLE-JWT-ECHO/internal/server"
	"context"
	"log"

	"github.com/sirupsen/logrus"
)

func main() {
	// LoadConfig here
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Errorf("main.LoadConfig: %v", err.Error())
	}

	DB, err := connection.GetDBClient(context.Background())
	if err != nil {
		log.Printf("failed to connect to DB")
	}

	defer func() {
		err = DB.Close()
		if err != nil {
			logrus.Errorf("main.CloseDB: %v", err.Error())
		}
	}()

	// dataStore is
	dataStore := postgres.NewDataStore(DB)

	// Server is
	s := server.NewServer(cfg, dataStore)
	if err := s.Run(); err != nil {
		logrus.Errorf("main.Run: %v", err.Error())
	}
}
