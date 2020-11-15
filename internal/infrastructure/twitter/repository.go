package twitter

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/quintodown/fauspanish/internal/domain"
	"net/http"
)

type Repository struct {
	c *twitter.Client
}

func NewRepository(c *twitter.Client) *Repository {
	return &Repository{c: c}
}

func (r *Repository) GetFrom(listID, sinceID int64) ([]domain.Tweet, error) {
	statuses, err := r.executeRequest(listID, sinceID)
	if err != nil {
		return nil, err
	}

	tweets := make([]domain.Tweet, 0, len(statuses))
	for i := range statuses {
		atTime, err := statuses[i].CreatedAtTime()
		if err != nil {
			return nil, err
		}

		var url string
		if len(statuses[i].Entities.Urls) > 0 {
			url = statuses[i].Entities.Urls[0].ExpandedURL
		}

		var media string
		if len(statuses[i].Entities.Media) > 0 {
			media = statuses[i].Entities.Media[0].MediaURLHttps
		}

		tweets = append(
			tweets,
			domain.NewTweetBuilder().
				Id(statuses[i].ID).
				Media(media).
				PublishedAt(atTime.UTC()).
				Text(statuses[i].Text).
				URL(url).
				Build(),
		)
	}

	return tweets, nil
}

func (r *Repository) executeRequest(listID int64, sinceID int64) ([]twitter.Tweet, error) {
	statuses, httpResp, err := r.c.Lists.Statuses(&twitter.ListsStatusesParams{ListID: listID, SinceID: sinceID})
	if err != nil  {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error, status: %v", httpResp.StatusCode)
	}

	return statuses, nil
}
