package rss_reader

import (
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"net/http"
	"time"
)

const PeriodDays = 7

type Source struct {
	Link string `json:"link"`
}

func RssReader(w http.ResponseWriter, r *http.Request) {
	var s Source

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		fmt.Fprint(w, "Could not decode the body")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(New(s))
}

func New(s Source) []string {
	feed, _ := gofeed.NewParser().ParseURL(s.Link)
	unread := GetNewChapters(feed.Items)

	return unread
}

func GetNewChapters(items []*gofeed.Item) []string {
	newChapters := []string{}
	for _, item := range items {
		if !IsNew(item) {
			break
		}

		newChapters = append(newChapters, item.Link)
	}

	return newChapters
}

func IsNew(item *gofeed.Item) bool {
	lastCheckDate := time.Now().AddDate(0, 0, -1*PeriodDays-1)

	return item.PublishedParsed.After(lastCheckDate)
}
