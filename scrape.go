package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type newsItem struct {
	title   string
	summary string
	tag     string
}

func main() {
	news := []newsItem{}

	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML(".module--news ul > li", func(e *colly.HTMLElement) {
		item := newsItem{}
		item.title = e.ChildText(".media__title")
		item.summary = e.ChildText(".media__summary")
		item.tag = e.ChildText(".media__tag")
		news = append(news, item)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.bbc.com/")
	fmt.Println("News: ", news[0])
}
