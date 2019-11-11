package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
)

var rssFeeds= [4]string{
     "https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	 "https://feeds.a.dj.com/rss/RSSWorldNews.xml",
	"http://www.reddit.com/r/politics/.rss",
	"http://feeds.washingtonpost.com/rss/opinions",
}

type env struct {
	RssParser Parser
}

func main() {
	e := env{
		RssParser: gofeed.NewParser(),
	}
	e.scapeRssFeeds()
}

func (e *env) scapeRssFeeds() {
	for _, feed := range rssFeeds {
		result,err := e.RssParser.ParseURL(feed)
		if err != nil {
			log.Printf("error getting feed: %s",feed)
			continue
		}
		fmt.Println(result.Title)
	}
}



