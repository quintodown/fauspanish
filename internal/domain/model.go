package domain

import "time"

type Tweet struct {
	id         string
	text       string
	media      string
	url        string
	scheduleAt time.Time
}

func (t Tweet) Id() string {
	return t.id
}

func (t Tweet) Text() string {
	return t.text
}

func (t Tweet) Media() string {
	return t.media
}

func (t Tweet) Url() string {
	return t.url
}

func (t Tweet) ScheduleAt() time.Time {
	return t.scheduleAt
}

type TweetBuilder struct {
	tweet *Tweet
}

func NewTweetBuilder() *TweetBuilder {
	return &TweetBuilder{tweet: &Tweet{}}
}

func (b *TweetBuilder) id(id string) *TweetBuilder {
	b.tweet.id = id

	return b
}

func (b *TweetBuilder) text(text string) *TweetBuilder {
	b.tweet.text = text

	return b
}

func (b *TweetBuilder) media(media string) *TweetBuilder {
	b.tweet.media = media

	return b
}

func (b *TweetBuilder) url(url string) *TweetBuilder {
	b.tweet.url = url

	return b
}

func (b *TweetBuilder) scheduleAt(scheduleAt time.Time) *TweetBuilder {
	b.tweet.scheduleAt = scheduleAt

	return b
}

func (b *TweetBuilder) Build() Tweet {
	return *b.tweet
}

