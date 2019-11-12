package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mmcdole/gofeed"
	"log"
)

var rssFeeds = map[string]string{
	"NYTimes":      "https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	"WSJ: World":   "https://feeds.a.dj.com/rss/RSSWorldNews.xml",
	"WSJ: US Biz": "https://feeds.a.dj.com/rss/WSJcomUSBusiness.xml",
	"WSJ: Opinion": "https://feeds.a.dj.com/rss/RSSOpinion.xml",
	"WP: Opinion":  "http://feeds.washingtonpost.com/rss/opinions",
	"WP: Politics": "http://feeds.washingtonpost.com/rss/politics",
	"Reddit: Politics": "http://www.reddit.com/r/politics/.rss",
	"Google News: Nation": "https://news.google.com/rss/topics/CAAqIggKIhxDQkFTRHdvSkwyMHZNRGxqTjNjd0VnSmxiaWdBUAE",
	"BuzzFeed: Politics": "https://www.buzzfeed.com/politics.xml",
}

type env struct {
	rss Parser
	db  sqlx.DB
}

func main() {

	db, err := sqlx.Connect("postgres", "user=secret password=verysecret dbname=headlines sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	e := env{
		rss: gofeed.NewParser(),
		db:  *db,
	}

	e.scrapeRssFeeds()
}

func (e *env) scrapeRssFeeds() {
	for feed_name, feed_source := range rssFeeds {

		result, err := e.rss.ParseURL(feed_source)
		if err != nil {
			log.Printf("error getting feed: %s: %+v", feed_name, err)
			continue
		}
		log.Printf("found %d results for %s", len(result.Items), feed_name)
		ds := DataSource{
			Title:       feed_name,
			Description: result.Description,
			Link:        result.Link,
		}

		tx := e.db.MustBegin()
		dsID, err := insertDataSource(tx, ds)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range result.Items {
			var author string
			if item.Author != nil {
				author = item.Author.Name
			} else {
				author = ""
			}
			newHeadline := HeadLine{
				Title:       item.Title,
				Description: item.Description,
				Link:        item.Link,
				LinkHTML:    item.GUID,
				Author:      author,
				PublishedAt: item.PublishedParsed,
			}
			err := insertHeadLine(tx, newHeadline, dsID)
			if err != nil {
				log.Println(err)
			}
		}

		err = tx.Commit()
		if err != nil {
			log.Println("error committing to db %+v", err)
		}

	}
}
