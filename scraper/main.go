package main

import "fmt"
import "github.com/gocolly/colly/v2"

func main() {
	fmt.Println("[..] scraping dw.com")
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("[ok] scraping", r.URL)
	})

	c.Visit("https://www.dw.com/de")
	fmt.Println("[OK] completed scrape of dw.com")
}