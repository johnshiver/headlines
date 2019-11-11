package main

import (
	"github.com/gocolly/colly"
	"log"
	"strings"
)


func scrapeWebsite(url string, domains []string) {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.AllowedDomains(domains...),
		colly.UserAgent("john-shiver-dev"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// fmt.Println(link)
		if strings.Contains(link, "section") {
			err := e.Request.Visit(link)
			if err != nil {
				return
				// log.Printf("error visiting link %s: %+v", link, err)
			}

		}
	})

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		headline:= e.Attr("h2")
		log.Printf("headline: %s", headline)
	})

	err := c.Visit(url)
	if err != nil {
		log.Printf("error visiting url %s: %+v", url, err)
	}

}
