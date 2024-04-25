package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/gocolly/colly/v2"
	"github.com/thalesfsp/go-common-types/safeorderedmap"
)

func linkNotInside(linkEntry, linkTarget string, pathMap *safeorderedmap.SafeOrderedMap[[]string]) bool {
	slicePath, _ := pathMap.Get(linkEntry)
	for _, link := range slicePath {
		if link == linkTarget {
			return false
		}
	}
	return true
}

func crawlerBFS(start string, target string, path *safeorderedmap.SafeOrderedMap[[]string], depth *int32, timer time.Time) {
	var mutex sync.Mutex
	queue := orderedmap.NewOrderedMap[string, any]()
	visits := safeorderedmap.New[bool]()
	// path := safeorderedmap.New[[]string]()
	queueChild := orderedmap.NewOrderedMap[string, any]()
	var found int32

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
		// RAMKU MELEDAK KALAU BFS PAKE CACHE
		// colly.CacheDir("./cache"),
	)

	// Wtf is even Parallelism: 1000?? Me brainrot big number equals good
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100, RandomDelay: 25 * time.Millisecond})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		// Pas error queuenya juga dihapus biar ngga ngeblock
		log.Println("Something went wrong:", err)
		mutex.Lock()
		fmt.Println("SISA", queue.Len())
		queue.Delete(r.Request.URL.String())
		if queue.Len() == 0 && atomic.LoadInt32(&found) != 1 {
			queue, queueChild = queueChild, queue
			queueChild = orderedmap.NewOrderedMap[string, any]()
			atomic.StoreInt32(depth, atomic.LoadInt32(depth)+1)
			fmt.Println("Searching at depth:", atomic.LoadInt32(depth))
		}
		mutex.Unlock()
	})

	c.OnResponse(func(r *colly.Response) {
		// I need to track how many links are fully visited
		// fmt.Println("Visited", r.Request.URL.String())
	})

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		// time.Sleep(1 * time.Millisecond)
		link := h.Attr("href")
		pathInserter, _ := path.Get(h.Request.AbsoluteURL(link))
		if link == "/wiki/"+target {
			// path.Set(h.Request.AbsoluteURL(link), h.Request.URL.String())
			fmt.Println("Found target link at depth", atomic.LoadInt32(depth)+1, ":", link)
			if linkNotInside(h.Request.AbsoluteURL(link), h.Request.URL.String(), path) {
				path.Add(h.Request.AbsoluteURL(link), append(pathInserter, h.Request.URL.String()))
			}
			atomic.StoreInt32(&found, atomic.LoadInt32(&found)+1)
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
			link != "/wiki/Main_Page" &&
			h.Request.AbsoluteURL(link) != h.Request.URL.String() {
			mutex.Lock()
			visited, _ := visits.Get(h.Request.AbsoluteURL(link))
			if !visited {
				queueChild.Set(h.Request.AbsoluteURL(link), true)
			}
			mutex.Unlock()
			if linkNotInside(h.Request.AbsoluteURL(link), h.Request.URL.String(), path) {
				path.Add(h.Request.AbsoluteURL(link), append(pathInserter, h.Request.URL.String()))
			}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println("Finished", r.Request.URL)
		mutex.Lock()
		fmt.Println("SISA", queue.Len())
		queue.Delete(r.Request.URL.String())
		if queue.Len() == 0 && atomic.LoadInt32(&found) != 1 {
			queue, queueChild = queueChild, queue
			queueChild = orderedmap.NewOrderedMap[string, any]()
			atomic.StoreInt32(depth, atomic.LoadInt32(depth)+1)
			fmt.Println("Searching at depth:", atomic.LoadInt32(depth))
		}
		mutex.Unlock()
	})

	queue.Set("https://en.wikipedia.org/wiki/"+start, true)

queueIteration:
	for {
		mutex.Lock()
		for el := queue.Front(); el != nil; el = el.Next() {
			visits.Add(el.Key, true)
			c.Visit(el.Key)
		}
		mutex.Unlock()
		if queue.Len() == 0 || atomic.LoadInt32(&found) >= 1 && time.Since(timer) >= 58*time.Second {
			c.AllowedDomains = []string{""}
			break queueIteration
		}
	}
}
