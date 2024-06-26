package pkg

import (
	"context"
	"fmt"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"regexp"
	"time"
)

func AddEvent(title, description, colorId, start, end string) {
	client, err := GetClient()
	if err != nil {
		return
	}

	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create Calendar service: %v", err)
	}

	var startTime, endTime time.Time
	var err1, err2 error

	if start == "" && end != "" {
		fmt.Println("Please enter start time")
		return
	} else if start != "" {
		pattern := `^([01]\d|2[0-3]):([0-5]\d):([0-5]\d)$`
		re := regexp.MustCompile(pattern)

		if re.MatchString(start) {
			currDate := time.Now().Format("2006-01-02")
			start = currDate + " " + start
		}
		startTime, err1 = time.Parse("2006-01-02 15:04:05", start)
		if err1 != nil {
			fmt.Println("Please enter correct time")
			return
		}
		if end == "" {
			endTime = startTime.Add(1 * time.Hour)
		}
	} else {
		fmt.Println(start)
		startTime, err1 = time.Parse("2006-01-02 15:04:05", start)
		endTime, err2 = time.Parse("2006-01-02 15:04:05", end)
	}

	if err1 != nil && err2 != nil {
		fmt.Println("Please enter correct time format")
		return
	}

	event := &calendar.Event{
		Summary:     title,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), 0, time.Now().Location()).Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: time.Date(endTime.Year(), endTime.Month(), endTime.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, time.Now().Location()).Format(time.RFC3339),
		},
		ColorId: colorId,
	}

	event, err = srv.Events.Insert("primary", event).Do()
	if err != nil {
		log.Fatalf("Unable to create event: %v", err)
	}

	fmt.Printf("Event created: %v\n", event.Summary)
}
