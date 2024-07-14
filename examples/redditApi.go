package main

import (
	"context"
	"log"

	"github.com/joelseq/apxlgnds/internal/calendar"
)

func main() {
	ctx := context.Background()
	_, err := calendar.GetRedditALGSThreads(ctx, false)
	if err != nil {
		log.Fatalf("failed to get reddit threads: %v", err)
	}

	// fmt.Printf("Reddit threads: %v\n", res)
}
