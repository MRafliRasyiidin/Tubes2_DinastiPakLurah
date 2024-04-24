package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/gocolly/colly/v2"
)

type pairChan struct {
	found bool
	done  bool
}

var (
	controlChan = make(chan pairChan)
)

func crawlerDLS(start string, target string, depth int, path *orderedmap.OrderedMap[string, any]) {
	var mutex sync.Mutex
	// path := orderedmap.NewOrderedMap[string, any]()
	// TODO : Simpan path, cek mencapai target or not
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
		if link == "/wiki/"+target {
			fmt.Println("Found target link at depth", depth+1, ":", link)
			path.Set(e.Request.AbsoluteURL(link), e.Request.URL.String())
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
			mutex.Lock()
			path.Set(e.Request.AbsoluteURL(link), e.Request.URL.String())
			mutex.Unlock()
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

func crawlerIDS(start, target string) {
	// os.RemoveAll("./cache")
	path := orderedmap.NewOrderedMap[string, any]()
	i := 0
incrementLoop:
	for {
		fmt.Println("Depth", i)
		go crawlerDLS(start, target, i, path)
		controlFlow := <-controlChan
		if controlFlow.found && controlFlow.done {
			break incrementLoop
		}
		if controlFlow.done && !controlFlow.found {
			i++
		}
	}

	key := "https://en.wikipedia.org/wiki/" + target
	expPath := []string{}
	for key != "https://en.wikipedia.org/wiki/"+start {
		fmt.Println("Key : ", key)
		expPath = append(expPath, key)
		value, _ := path.Get(key)
		key = value.(string)
	}

	for i, j := 0, len(expPath)-1; i < j; i, j = i+1, j-1 {
		expPath[i], expPath[j] = expPath[j], expPath[i]
	}

	jsonStr, err := json.Marshal(expPath)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
	// for iter := path.Front(); iter != path.Back(); iter = iter.Next() {
	// 	fmt.Println(iter.Key, iter.Value)
	// }
	// key := "https://en.wikipedia.org/wiki/" + target
	// fmt.Println(path.Get(key))
}
