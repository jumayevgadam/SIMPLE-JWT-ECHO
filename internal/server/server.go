package server

import (
	"SIMPLE-JWT-ECHO/config"
	"SIMPLE-JWT-ECHO/internal/database"

	"github.com/labstack/echo/v4"
)

// Server is
type Server struct {
	echo      *echo.Echo
	cfg       *config.Config
	dataStore database.DataStore
}

// NewServer is
func NewServer(
	cfg *config.Config,
	dataStore database.DataStore,
) *Server {
	server := &Server{
		echo:      echo.New(),
		cfg:       cfg,
		dataStore: dataStore,
	}

	return server
}

// Run is
func (s *Server) Run() error {
	s.MapHandlers(s.echo)
	return s.echo.Start(":" + s.cfg.Server.PORT)
}
