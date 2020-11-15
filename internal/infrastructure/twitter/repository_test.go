package twitter_test

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/jarcoal/httpmock"
	"github.com/quintodown/fauspanish/internal/domain"
	twitter2 "github.com/quintodown/fauspanish/internal/infrastructure/twitter"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestRepository_GetFrom(t *testing.T) {
	httpClient := &http.Client{}
	httpmock.ActivateNonDefault(httpClient)
	defer httpmock.DeactivateAndReset()

	repo := twitter2.NewRepository(twitter.NewClient(httpClient))

	t.Run("it fails getting statuses from API", func(t *testing.T) {
		httpmock.RegisterResponder(
			"GET",
			"https://api.twitter.com/1.1/lists/statuses.json?list_id=2",
			httpmock.NewStringResponder(500, "{}"),
		)

		tweets, err := repo.GetFrom(2, 0)

		assert.EqualError(t, err, "response error, status: 500")
		assert.Nil(t, tweets)
	})

	t.Run("it is possible to get statuses from list", func(t *testing.T) {
		statuses, err := ioutil.ReadFile("testdata/statuses.json")
		if err != nil {
			t.Fatal(err)
		}

		httpmock.RegisterResponder(
			"GET",
			"https://api.twitter.com/1.1/lists/statuses.json?list_id=1",
			httpmock.NewStringResponder(200, string(statuses)),
		)

		tweets, err := repo.GetFrom(1, 0)

		assert.NoError(t, err)
		assert.Equal(t, []domain.Tweet{
			domain.NewTweetBuilder().
				Id(245160944223793152).
				Text("Create your own TFC ESQ by Movado Watch: http://t.co/W2tON3OK in support of @TeamUpFdn #TorontoFC #MLS").
				Media("").
				URL("http://bit.ly/MuCCDo").
				PublishedAt(time.Date(2012, 9, 10, 14, 04, 58, 0, time.UTC)).
				Build(),
		}, tweets)
	})
}
