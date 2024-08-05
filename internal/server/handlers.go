package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/joelseq/apxlgnds/internal/cache"
	"github.com/joelseq/apxlgnds/internal/reddit"
	"github.com/joelseq/apxlgnds/internal/types"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleEvents(c echo.Context) error {
	ctx := c.Request().Context()
	events, err := s.GetEvents(ctx, eventLimit)
	if err != nil {
		return err
	}

	return c.JSON(200, events)
}

func (s *Server) handleHealth(c echo.Context) error {
	return c.String(200, "ok")
}

func (s *Server) GetEvents(ctx context.Context, eventLimit int) (*types.CalendarEventsResponse, error) {
	cachedRes, err := s.cache.GetResult(ctx)
	if err != nil && err != cache.ErrCacheEmpty {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}
	if cachedRes != nil {
		s.logger.Info("returning cached events")
		return cachedRes, nil
	}

	events, err := s.calendar.FetchEvents(ctx, eventLimit)
	if err != nil {
		return nil, err
	}

	err = s.cache.SetResult(ctx, events)
	if err != nil {
		s.logger.Error("failed to cache events", slog.String("error", err.Error()))
	}

	return events, nil
}

func (s *Server) handleReddit(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := reddit.GetRedditALGSThreads(ctx, true)
	if err != nil {
		fmt.Printf("failed to get reddit threads: %v", err)
		return err
	}

	return c.JSON(200, res)
}
