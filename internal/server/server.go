package server

import (
	"log/slog"

	"github.com/joelseq/apxlgnds/internal/cache"
	"github.com/joelseq/apxlgnds/internal/calendar"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	eventLimit = 8
)

type Server struct {
	port     string
	cache    cache.Cacher
	calendar calendar.Service
	logger   *slog.Logger
}

func NewServer(port string, cache cache.Cacher, calendarService calendar.Service, logger *slog.Logger) *Server {
	return &Server{
		port:     port,
		cache:    cache,
		calendar: calendarService,
		logger:   logger,
	}
}

func (s *Server) Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/events", s.handleEvents)
	e.GET("/", s.handleHealth)
	e.Logger.Fatal(e.Start(":" + s.port))
}
