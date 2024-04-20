package main

import (
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

type Page struct {
	URL   string
	Title string
	Depth int
}

func breadthFirstScrapper(firstWord string, word string) ([]Page, error) {

	url := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(firstWord, " ", "_")
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))

	var (
		isFound bool
		queue   []Page
		visited = make(map[string]bool)
		path    []Page
		mutex   sync.Mutex
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if !strings.HasPrefix(link, "http") {
			link = "https://en.wikipedia.org" + link
		}

		if strings.Contains(link, "%") {
			return
		}

		if !strings.Contains(link, "wikipedia.org") ||
			!strings.Contains(link, "/wiki/") ||
			strings.Contains(link, "Template") ||
			strings.Contains(link, "Special:") ||
			strings.Contains(link, "Wikipedia:") ||
			strings.Contains(link, "Help:") ||
			strings.Contains(link, "Portal:") ||
			strings.Contains(link, "Main_Page") ||
			strings.Contains(link, "Talk:") ||
			strings.Contains(link, "File:") {
			return
		}

		normalizedLink := NormalizeURL(link)
		if visited[normalizedLink] {
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		visited[normalizedLink] = true

		if strings.EqualFold(e.Text, word) && getTitleFromURL(link) == word {
			fmt.Println("Found the specified word:", link)
			fmt.Println("Title: ", e.Text)
			isFound = true
			return
		}

		if strings.Contains(link, "https://en.wikipedia.org") {
			queue = append(queue, Page{URL: link, Title: getTitleFromURL(link), Depth: e.Request.Depth})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	if strings.Contains(url, "https://en.wikipedia.org") {
		queue = append(queue, Page{URL: url, Title: getTitleFromURL(url), Depth: 0})
		visited[NormalizeURL(url)] = true
	}

	for len(queue) > 0 && !isFound {
		currPage := queue[0]
		queue = queue[1:]

		fmt.Println("Visiting: ", currPage.URL)
		fmt.Println("Title: ", currPage.Title)
		fmt.Println("Depth: ", currPage.Depth)

		c.Visit(currPage.URL)
		fmt.Println(queue)
		if isFound && strings.EqualFold(currPage.Title, word) {
			path = append(path, currPage)
			break
		}

	}
	return path, nil
}

func getTitleFromURL(url string) string {
	parts := strings.Split(url, "/")
	title := parts[len(parts)-1]
	title = strings.ReplaceAll(title, "_", " ")
	return title
}

func NormalizeURL(u string) string {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return u
	}
	parsedURL.Fragment = ""
	parsedURL.RawQuery = ""
	return strings.TrimSuffix(parsedURL.String(), "/")
}
