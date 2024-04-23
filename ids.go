package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/gocolly/colly/v2"
)

func crawlerDLS(start string, target string, depth int, startChan, doneChan, found chan bool, path *orderedmap.OrderedMap[string, any]) {
	var mutex sync.Mutex
	var targetFound int32 = 0
	// path := orderedmap.NewOrderedMap[string, any]()
	// TODO : Simpan path, cek mencapai target or not
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.MaxDepth(depth),
		colly.Async(true),
		colly.AllowURLRevisit(),
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
		if link == "/wiki/"+target {
			mutex.Lock()
			path.Set(e.Request.AbsoluteURL(link), e.Request.URL.String())
			atomic.StoreInt32(&targetFound, 1)
			mutex.Unlock()
			return
		}
		mutex.Lock()
		_, exists2 := path.Get(e.Request.AbsoluteURL(link))
		mutex.Unlock()
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
			if !exists2 {
				mutex.Lock()
				path.Set(e.Request.AbsoluteURL(link), e.Request.URL.String())
				mutex.Unlock()
			}
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
	if atomic.LoadInt32(&targetFound) == 1 {
		found <- true
	} else {
		found <- false
	}
	// found <- false
}

func crawlerIDS(start, target string) {
	// os.RemoveAll("./cache")
	startChan := make(chan bool)
	doneChan := make(chan bool)
	foundChan := make(chan bool, 1)
	path := orderedmap.NewOrderedMap[string, any]()

incrementLoop:
	for i := 1; ; i++ {
		fmt.Println("DEPTH KE", i)
		go crawlerDLS(start, target, i, startChan, doneChan, foundChan, path)
		startChan <- true
		<-doneChan
		var isFound bool
		select {
		case isFound = <-foundChan:
		default:
			isFound = false
		}
		if isFound {
			break incrementLoop
		}
	}
	key := "https://en.wikipedia.org/wiki/" + target
	expPath := []string{}
	for key != "https://en.wikipedia.org/wiki/"+start {
		expPath = append(expPath, key)
		value, _ := path.Get(key)
		key = value.(string)
	}

	for i, j := 0, len(expPath)-1; i < j; i, j = i+1, j-1 {
		expPath[i], expPath[j] = expPath[j], expPath[i]
	}
	expPath = expPath[1:]

	jsonStr, err := json.Marshal(expPath)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
}
