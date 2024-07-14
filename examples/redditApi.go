package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joelseq/apxlgnds/internal/calendar"
)

func main() {
	ctx := context.Background()
	res, err := calendar.GetRedditALGSThreads(ctx)
	if err != nil {
		log.Fatalf("failed to get reddit threads: %v", err)
	}

	fmt.Printf("Reddit threads: %v\n", res)
}
