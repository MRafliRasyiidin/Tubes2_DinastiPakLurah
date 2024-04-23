package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func crawlerDLS(start string, target string, depth int, startChan, doneChan chan bool) {
	// TODO : Simpan path, cek mencapai target or not
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.MaxDepth(depth),
		colly.Async(true),
		colly.AllowURLRevisit(),
		// DELETE CACHE SETIAP KALI MAU SEARCH BARU,
		// PENCEGAHAN RACE CONDITION & RAM MELEDAK
		colly.CacheDir("./cache"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 20, RandomDelay: 25 * time.Millisecond})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		link := r.Request.URL.String()
		if link == "/wiki/"+target {
			return
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// By Default, this is already Depth-Limited Search LMAOOOOO
		// Sprinkle some async + increment the depth :D
		link := e.Attr("href")
		if strings.HasPrefix(link, "/wiki/") &&
			!strings.Contains(link, "File:") &&
			!strings.Contains(link, "Help:") &&
			!strings.Contains(link, "Category:") &&
			!strings.Contains(link, "Wikipedia:") &&
			!strings.Contains(link, "Talk:") &&
			!strings.Contains(link, "Special:") &&
			!strings.Contains(link, "Portal:") &&
			!strings.Contains(link, "Template:") &&
			!strings.Contains(link, "MediaWiki:") &&
			!strings.Contains(link, "User:") &&
			!strings.Contains(link, "_talk:") &&
			e.Request.AbsoluteURL(link) != e.Request.URL.String() &&
			link != "/wiki/Main_Page" {
			e.Request.Visit(link)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://en.wikipedia.org/wiki/" + start)
	c.Wait()
	<-startChan
	doneChan <- true
}

func crawlerIDS(start, target string) {
	os.RemoveAll("./cache")
	var notFound = true
	startChan := make(chan bool)
	doneChan := make(chan bool)
	for i := 1; notFound; i++ {
		fmt.Println("DEPTH KE", i)
		go crawlerDLS(start, target, i, startChan, doneChan)
		startChan <- true
		<-doneChan
	}
}
