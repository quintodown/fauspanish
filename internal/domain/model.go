package domain

import "time"

type Tweet struct {
	id         int64
	text       string
	media      string
	url        string
	publishedAt time.Time
	scheduleAt time.Time
}

func (t Tweet) Id() int64 {
	return t.id
}

func (t Tweet) Text() string {
	return t.text
}

func (t Tweet) Media() string {
	return t.media
}

func (t Tweet) URL() string {
	return t.url
}

func (t Tweet) PublishedAt() time.Time {
	return t.publishedAt
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

func (b *TweetBuilder) Id(id int64) *TweetBuilder {
	b.tweet.id = id

	return b
}

func (b *TweetBuilder) Text(text string) *TweetBuilder {
	b.tweet.text = text

	return b
}

func (b *TweetBuilder) Media(media string) *TweetBuilder {
	b.tweet.media = media

	return b
}

func (b *TweetBuilder) URL(url string) *TweetBuilder {
	b.tweet.url = url

	return b
}

func (b *TweetBuilder) PublishedAt(publishedAt time.Time) *TweetBuilder {
	b.tweet.publishedAt = publishedAt

	return b
}

func (b *TweetBuilder) ScheduleAt(scheduleAt time.Time) *TweetBuilder {
	b.tweet.scheduleAt = scheduleAt

	return b
}

func (b *TweetBuilder) Build() Tweet {
	return *b.tweet
}

