package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

type newsItem struct {
	Title   string
	Summary string
	Tag     string
}

func printNews(n []newsItem) {
	fmt.Println("\n-- NEWS FOR TODAY --")
	for _, item := range n {
		fmt.Printf("Title: %v\nSummary: %v\nTag: %v\n\n", item.Title, item.Summary, item.Tag)
	}
}

func serializeToJSON(n []newsItem) {
	fmt.Println("Serializing news to JSON...")
	for i := range n {
		serialized, _ := json.Marshal(n[i])
		fmt.Println(string(serialized))
	}
}

func main() {
	news := []newsItem{}

	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML(".module--news ul > li", func(e *colly.HTMLElement) {
		item := newsItem{}
		item.Title = e.ChildText(".media__title")
		item.Summary = e.ChildText(".media__summary")
		item.Tag = e.ChildText(".media__tag")
		news = append(news, item)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Looking up the news for today... \nProvider: %v\n", r.URL.String())
	})

	c.Visit("https://www.bbc.com/")
	printNews(news)
	serializeToJSON(news)
}
