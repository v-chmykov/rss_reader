package rss_reader

import (
	"github.com/mmcdole/gofeed"
	"reflect"
	"testing"
	"time"
)

var yesterday = time.Now().AddDate(0, 0, -1)
var weekAgo = time.Now().AddDate(0, 0, -7)
var monthAgo = time.Now().AddDate(0, -1, 0)

func TestGetNewChapters(t *testing.T) {
	link3 := "https://example.net/c012/1.html"
	link2 := "https://example.net/c011/1.html"
	link1 := "https://example.net/c010/1.html"

	dummyItems := []*gofeed.Item{
		createDummyItem(link3, yesterday),
		createDummyItem(link2, weekAgo),
		createDummyItem(link1, monthAgo),
	}

	testCases := []struct {
		items    []*gofeed.Item
		expected []string
	}{
		{dummyItems, []string{link3, link2}},
	}

	for i, testCase := range testCases {
		actual := GetNewChapters(testCase.items)
		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf(`GetUnreadChapters[%d]: expected chapter = %v, %v recieved`, i, testCase.expected, actual)
		}
	}
}

func TestIsNew(t *testing.T) {
	testCases := []struct {
		link     string
		pubDate  time.Time
		expected bool
	}{
		{"https://example.net/c081/1.html", yesterday, true},
		{"https://example.net/c082/1.html", weekAgo, true},
		{"https://example.net/c083/1.html", monthAgo, false},
	}

	for i, testCase := range testCases {
		item := createDummyItem(testCase.link, testCase.pubDate)

		actual := IsNew(item)
		if actual != testCase.expected {
			t.Errorf(`%d) GetChapter: expected chapter = %t, %t recieved`, i, testCase.expected, actual)
		}
	}
}

func createDummyItem(link string, pubDate time.Time) *gofeed.Item {
	return &gofeed.Item{
		Link:      link,
		Published: pubDate.Format(TimeLayout),
	}
}
