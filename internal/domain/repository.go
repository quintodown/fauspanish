package domain

type TweetRepository interface {
	GetFrom(listID, sinceID int64) ([]Tweet, error)
}
