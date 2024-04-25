package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/thalesfsp/go-common-types/safeorderedmap"
)

type pairChan struct {
	found bool
	done  bool
}

var (
	controlChan = make(chan pairChan)
)

func crawlerDLS(start string, target string, depth int, path *safeorderedmap.SafeOrderedMap[[]string]) {
	var inserter pairChan
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.MaxDepth(depth+1),
		colly.AllowURLRevisit(),
		colly.Async(true),
		// DELETE CACHE SETIAP KALI MAU SEARCH BARU,
		// PENCEGAHAN RACE CONDITION & RAM MELEDAK
		// colly.CacheDir("./cache"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100, RandomDelay: 25 * time.Millisecond})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// By Default, this is already Depth-Limited Search LMAOOOOO
		// Sprinkle some async + increment the depth :D
		link := e.Attr("href")
		pathInserter, _ := path.Get(e.Request.AbsoluteURL(link))

		if link == "/wiki/"+target {
			fmt.Println("Found target link at depth", depth+1, ":", link)
			path.Add(e.Request.AbsoluteURL(link), append(pathInserter, e.Request.URL.String()))
			inserter.done = true
			inserter.found = true
			controlChan <- inserter
			return
		}
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
			path.Add(e.Request.AbsoluteURL(link), append(pathInserter, e.Request.URL.String()))
			e.Request.Visit(link)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://en.wikipedia.org/wiki/" + start)
	c.Wait()
	inserter.found = false
	inserter.done = true
	controlChan <- inserter
}

func crawlerIDS(start, target string, path *safeorderedmap.SafeOrderedMap[[]string], depth *int32) {
	// os.RemoveAll("./cache")
	i := 0
incrementLoop:
	for {
		fmt.Println("Searching at depth:", i)
		go crawlerDLS(start, target, i, path)
		controlFlow := <-controlChan
		if controlFlow.found && controlFlow.done {
			time.Sleep(5 * time.Second)
			break incrementLoop
		}
		if controlFlow.done && !controlFlow.found {
			i++
			atomic.StoreInt32(depth, int32(i))
		}
	}
}
