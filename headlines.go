package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/mmcdole/gofeed"
	"io"
	"time"
)

type DataSource struct {
	Title string
	Description string
	Link string

	CreatedAt *time.Time
	UpdatedAt *time.Time

}

func insertDataSource(tx *sqlx.Tx, ds DataSource) (int, error) {
	const q = `INSERT into data_sources
               (title, description, link)
               VALUES
               ($1, $2, $3)
               ON CONFLICT (title) 
			   DO UPDATE SET title=EXCLUDED.title
               RETURNING id`

	var dsID int
	err := tx.QueryRow(q, ds.Title, ds.Description, ds.Link).Scan(&dsID)
	if err != nil {
		return -1, err
	}
	return dsID, nil
}


type HeadLine struct {
	Source *DataSource

	Title string
	Description string
	Link string
	LinkHTML string
	Author string

	PublishedAt *time.Time

	CreatedAt *time.Time
	UpdatedAt *time.Time

}

func insertHeadLine(tx *sqlx.Tx, hl HeadLine, dataSourceID int) error {
	const q = `INSERT into headlines 
               (data_source_id, title, description, link, html_link, author, published_at)
               VALUES
               ($1, $2, $3, $4, $5, $6, $7)
               ON CONFLICT DO NOTHING`

	_, err := tx.Exec(q, dataSourceID, hl.Title, hl.Description, hl.Link, hl.LinkHTML, hl.Author, hl.PublishedAt)
	if err != nil {
		return err
	}
	return nil
}

type Parser interface {
	Parse(feed io.Reader) (*gofeed.Feed, error)
	ParseURL(feedURL string) (feed *gofeed.Feed, err error)
	ParseString(feed string) (*gofeed.Feed, error)
}
