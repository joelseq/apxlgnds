package calendar

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/joelseq/apxlgnds/internal/types"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type Service interface {
	FetchEvents(ctx context.Context, eventLimit int) (*types.CalendarEventsResponse, error)
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

func (s *service) FetchEvents(ctx context.Context, eventLimit int) (*types.CalendarEventsResponse, error) {
	srv, err := calendar.NewService(ctx, option.WithAPIKey(s.apiKey))
	if err != nil {
		log.Fatal("unable to retrieve Calendar client: ", err)
	}

	timeMinDate := time.Now().AddDate(0, 0, -60).Format(time.RFC3339)
	timeMaxDate := time.Now().AddDate(0, 0, 60).Format(time.RFC3339)

	events, err := srv.Events.List(calendarID).ShowDeleted(false).TimeMin(timeMinDate).TimeMax(timeMaxDate).OrderBy("startTime").SingleEvents(true).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}

	fmt.Printf("Fetched events count: %d\n", len(events.Items))
	redditResponse, err := GetRedditALGSThreads(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}

	return generateResponse(events, eventLimit, redditResponse)
}

func generateResponse(events *calendar.Events, eventLimit int, redditResponse *RedditResponse) (*types.CalendarEventsResponse, error) {
	res := types.CalendarEventsResponse{}

	parsedEvents := parseEvents(events)
	algsEvents := filterEvent(parsedEvents, func(event types.Event) bool {
		return strings.Contains(event.Title, "ALGS")
	})
	otherEvents := filterEvent(parsedEvents, func(event types.Event) bool {
		return !strings.Contains(event.Title, "ALGS")
	})

	res.ALGS = groupEvents(algsEvents, eventLimit)
	res.Other = groupEvents(otherEvents, eventLimit)
	addALGSMetadata(res.ALGS, redditResponse)

	return &res, nil
}

func parseEvents(events *calendar.Events) []types.Event {
	var res []types.Event
	for _, item := range events.Items {
		// Using the start as a filter for deleted events
		if item.Start != nil {
			event := types.Event{}
			startDateStr := item.Start.DateTime
			if startDateStr == "" {
				startDateStr = item.Start.Date
			}
			startDate := getDateFromString(startDateStr)

			endDateStr := item.End.DateTime
			if endDateStr == "" {
				endDateStr = item.Start.Date
			}
			endDate := getDateFromString(endDateStr)

			fmt.Printf("%v (%v)\n", item.Summary, startDateStr)

			event.Title = item.Summary
			event.Description = item.Description
			event.StartDate = startDate
			event.EndDate = endDate

			res = append(res, event)
		}
	}

	return res
}

func groupEvents(events []types.Event, limit int) *types.EventGroup {
	var upcoming []types.Event
	var recent []types.Event
	for _, event := range events {
		if event.StartDate.After(time.Now()) || event.EndDate.After(time.Now()) {
			upcoming = append(upcoming, event)
		} else {
			recent = append(recent, event)
		}
	}

	// Sort the upcoming events in ascending order
	sort.Slice(upcoming, func(i, j int) bool {
		return upcoming[i].StartDate.Before(upcoming[j].StartDate)
	})
	// Sort the recent events in descending order
	sort.Slice(recent, func(i, j int) bool {
		return recent[j].StartDate.Before(recent[i].StartDate)
	})

	return &types.EventGroup{
		Upcoming: lenOrLimit(upcoming, limit),
		Recent:   lenOrLimit(recent, limit),
	}
}

func addALGSMetadata(events *types.EventGroup, redditResponse *RedditResponse) {
	addMetadataForEvents(events.Upcoming, redditResponse)
	addMetadataForEvents(events.Recent, redditResponse)
}

func lenOrLimit(slice []types.Event, limit int) []types.Event {
	if len(slice) < int(limit) {
		return slice
	}
	return slice[:limit]
}

func filterEvent(events []types.Event, filter func(types.Event) bool) []types.Event {
	var res []types.Event
	for _, event := range events {
		if filter(event) {
			res = append(res, event)
		}
	}
	return res
}

func getDateFromString(dateStr string) time.Time {
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		log.Fatalf("failed to parse date string: %v", err)
	}
	return date
}
