package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

type newsItem struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Tag     string `json:"tag"`
}

func writeFile(file []byte) {
	if err := ioutil.WriteFile("output.json", file, 0644); err != nil {
		log.Fatalf("Unable to write file! %v", err)
	}
}

func printNews(n []newsItem) {
	fmt.Println("\n-- NEWS FOR TODAY --")
	for _, item := range n {
		fmt.Printf("Title: %v\nSummary: %v\nTag: %v\n\n", item.Title, item.Summary, item.Tag)
	}
}

func serializeToJSON(n []newsItem) {
	fmt.Println("Serializing news to JSON...")
	serialized, _ := json.Marshal(n)
	writeFile(serialized)
	fmt.Println("Successfully serialized the news to 'output.json'")
}

func main() {
	// Store the scraped news items.
	news := []newsItem{}

	c := colly.NewCollector()

	c.OnHTML(".module--news ul > li", func(e *colly.HTMLElement) {
		item := newsItem{}
		item.Title = e.ChildText(".media__title")
		item.Summary = e.ChildText(".media__summary")
		item.Tag = e.ChildText(".media__tag")
		news = append(news, item)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Looking up the news for today... \nProvider: %v\n", r.URL.String())
	})

	c.Visit("https://www.bbc.com/")

	printNews(news)
	serializeToJSON(news)
}
