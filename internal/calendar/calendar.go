package calendar

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joelseq/apxlgnds/internal/types"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type Service interface {
	FetchEvents(ctx context.Context) (*types.CalendarEventsResponse, error)
}

type service struct {
	apiKey string
}

// CalendarID for the Comp Apex calendar
const calendarID = "9ad286735043bbfc1494408580cbe6246b9d92988537e4549053f9e6866d63b3@group.calendar.google.com"

func NewService(apiKey string) Service {
	return &service{
		apiKey: apiKey,
	}
}

func (s *service) FetchEvents(ctx context.Context) (*types.CalendarEventsResponse, error) {
	srv, err := calendar.NewService(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		log.Fatal("unable to retrieve Calendar client: ", err)
	}

	// Get a date string for 2 weeks ago
	timeMinDate := time.Now().AddDate(0, 0, -14).Format(time.RFC3339)
	timeMaxDate := time.Now().AddDate(0, 0, 7).Format(time.RFC3339)

	events, err := srv.Events.List(calendarID).ShowDeleted(false).TimeMin(timeMinDate).TimeMax(timeMaxDate).OrderBy("startTime").SingleEvents(true).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}

	res := &types.CalendarEventsResponse{}

	fmt.Printf("Fetched events count: %d\n", len(events.Items))
	for _, item := range events.Items {
		// Using the start as a filter for deleted events
		if item.Start != nil {
			event := types.Event{}
			dateStr := item.Start.DateTime
			if dateStr == "" {
				dateStr = item.Start.Date
			}
			date := getDateFromString(dateStr)
			fmt.Printf("%v (%v)\n", item.Summary, date)
			event.Title = item.Summary
			event.Description = item.Description
			event.Date = date
			res.Events = append(res.Events, event)
		}
	}

	// err = s.cache.SetResult(ctx, res)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to set events: %w", err)
	// }

	return res, nil
}

func getDateFromString(dateStr string) time.Time {
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		log.Fatalf("failed to parse date string: %v", err)
	}
	return date
}
