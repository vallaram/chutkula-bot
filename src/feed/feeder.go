package feed

import "time"

type Feeder interface {
	ParseContent(content string, title string) (parsedItem string)

	FetchFeed() (items []string, err error)

	IsSyncedTime(updatedTime *time.Time) (bool, error)

	EmojiInjector() (emojis []string)
}