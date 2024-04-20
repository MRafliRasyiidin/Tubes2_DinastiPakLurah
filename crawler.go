package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func WebCrawler() {

	c := colly.NewCollector(colly.AllowedDomains(
		"id.wikipedia.org"))

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("https://id.wikipedia.org/wiki/Halaman_Utama")
}
