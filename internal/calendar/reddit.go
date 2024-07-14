package calendar

import (
	"context"
	"os"

	"github.com/vartanbeno/go-reddit/v2/reddit"
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

func GetRedditALGSThreads(ctx context.Context, debug bool) (*RedditResponse, error) {
	credentials := reddit.Credentials{ID: os.Getenv("REDDIT_CLIENT_ID"), Secret: os.Getenv("REDDIT_CLIENT_SECRET"), Username: os.Getenv("REDDIT_USERNAME"), Password: os.Getenv("REDDIT_PASSWORD")}
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, err
	}

	posts, _, err := client.Subreddit.SearchPosts(ctx,
		"flair_name:\"ALGS Y4 S2 :apex:\"",
		"CompetitiveApex",
		&reddit.ListPostSearchOptions{
			ListPostOptions: reddit.ListPostOptions{
				ListOptions: reddit.ListOptions{
					Limit: 50,
				},
				Time: "month",
			},
		},
	)
	if err != nil {
		return nil, err
	}

	data := RedditResponse{
		Kind: "Listing",
		Data: RedditData{
			Children: []Thread{},
		},
	}

	for _, post := range posts {
		data.Data.Children = append(data.Data.Children, Thread{
			Data: ThreadData{
				Title: post.Title,
				URL:   post.URL,
			},
		})
	}

	return &data, nil
}
