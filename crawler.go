package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func crawler() {
	found := false
	queue := make(map[string]bool)
	c := colly.NewCollector(colly.AllowedDomains(
		"en.wikipedia.org"))

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		if strings.HasPrefix(link, "/wiki/") &&
			!strings.Contains(link, "File:") &&
			!strings.Contains(link, "Help:") &&
			!strings.Contains(link, "Category:") &&
			!strings.Contains(link, "Wikipedia:") &&
			!strings.Contains(link, "Talk:") &&
			!strings.Contains(link, "Special:") &&
			!strings.Contains(link, "Portal:") &&
			!strings.Contains(link, "Template:") &&
			("/wiki/Main_Page" != link) &&
			!queue["https://en.wikipedia.org"+link] {
			queue["https://en.wikipedia.org"+link] = true
		}
		if link == "/wiki/Novel_Baswaaaedan" {
			fmt.Println("KETEMU")
			found = true
		}
	})

	queue["https://en.wikipedia.org/wiki/Gibran_Rakabuming"] = true
	c.Visit("https://en.wikipedia.org/wiki/Gibran_Rakabuming")

}

func main() {
	crawler()
}
