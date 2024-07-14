package calendar

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/joelseq/apxlgnds/internal/types"
)

func addMetadataForEvents(events []types.Event) {
	for i, event := range events {
		region := getRegion(event.Title)
		day := getDay(event.Title)

		if region != types.RegionUnknown && day != "" {
			parsedDay, err := strconv.Atoi(day)

			if err == nil {
				isFinals := strings.Contains(event.Title, "Match Point Finals")
				event.Metadata = types.EventMetadata{
					Region:      string(region),
					Day:         parsedDay,
					IsFinals:    isFinals,
					BattlefyURL: getBattlefyURL(region, parsedDay, isFinals),
				}
			}
		}

		events[i] = event
	}
}

func getRegion(title string) types.Region {
	if strings.Contains(title, "APAC N") {
		return types.RegionAPACNorth
	} else if strings.Contains(title, "APAC S") {
		return types.RegionAPACSouth
	} else if strings.Contains(title, "EMEA") {
		return types.RegionEMEA
	} else if strings.Contains(title, "NA") {
		return types.RegionNA
	}

	return types.RegionUnknown
}

func getDay(title string) string {
	dayRegex := regexp.MustCompile(`Day (?P<Day>\d+)`)
	match := dayRegex.FindStringSubmatch(title)

	if len(match) > 0 {
		// The first element in match is the full match, the second is the captured group
		return match[1]
	}

	if strings.Contains(title, "Match Point Finals") {
		// We will encode the finals as day 10 to make it convenient for
		// generating the battlefy url.
		return "10"
	}

	return ""
}

func getBattlefyURL(region types.Region, day int, isFinals bool) string {
	splitURLParam := "pro-league-split-2"
	if isFinals {
		splitURLParam += "-regional-finals"
	}

	return fmt.Sprintf("https://battlefy.com/apex-legends-global-series-year-4/%s/%s/65fc89113fce34803f734707/round/%d/match/%d", splitURLParam, region.URLParam(), day-1, day-1)
}
