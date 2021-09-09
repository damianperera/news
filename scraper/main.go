package main

import "fmt"
import "github.com/gocolly/colly/v2"

func main() {
	fmt.Println("Hello, World!")
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		err := e.Request.Visit(e.Attr("href"))
		if err != nil {
			return
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("http://go-colly.org/")
	if err != nil {
		return
	}
}