package main

import (
	"github.com/mmcdole/gofeed"
	"io"
	"time"
)

type DataSource struct {
	Title string
	Description string
	Link string

	CreatedAt time.Time
	UpdatedAt time.Time

}


type Headline struct {
	Source *DataSource

	Title string
	Description string
	Link string
	LinkHTML string
	Authors string

	PublishedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time

}

type Parser interface {
	Parse(feed io.Reader) (*gofeed.Feed, error)
	ParseURL(feedURL string) (feed *gofeed.Feed, err error)
	ParseString(feed string) (*gofeed.Feed, error)
}
