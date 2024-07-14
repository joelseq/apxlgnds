package calendar

import (
	"context"
	"encoding/json"
	"net/http"
)

const redditAPIURL = "https://www.reddit.com/r/CompetitiveApex/search.json?q=flair_name%3A%22ALGS%20Y4%20S2%20%3Aapex%3A%22"

type RedditResponse struct {
	Kind string     `json:"kind,omitempty"`
	Data RedditData `json:"data,omitempty"`
}

type RedditData struct {
	Children []Thread `json:"children,omitempty"`
}

type Thread struct {
	Data ThreadData `json:"data,omitempty"`
}

type ThreadData struct {
	Title string
	URL   string
}

func GetRedditALGSThreads(ctx context.Context) (*RedditResponse, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", redditAPIURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	data := RedditResponse{}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
