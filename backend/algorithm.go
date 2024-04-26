// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// )

// // func index(w http.ResponseWriter, r *http.Request) {
// // fmt.Fprintln(w, "your mum")
// // }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "")

// 		var t, err = template.ParseFiles(("template/main.html"))
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}
// 		t.Execute(w, map[string]string{"Text": "naw"})
// 	})

// 	// http.HandleFunc("/urmumgay", index)

// 	fmt.Println("starting web server at http://localhost:8080/")
// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
// 	http.ListenAndServe(":8080", nil)
// }

package main

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/thalesfsp/go-common-types/safeorderedmap"
)

func dfsPathMaker(node, ender string, adjacencyList *safeorderedmap.SafeOrderedMap[[]string], path []string, paths *[][]string, depth int32) {
	adjacent, _ := adjacencyList.Get(node)
	finish := node == "https://en.wikipedia.org/wiki/"+ender

	path = append(path, node)
	pathCopy := []string{}
	for _, value := range path {
		pathCopy = append(pathCopy, value)
	}
	isExist := false
	for _, value := range *paths {
		if reflect.DeepEqual(value, pathCopy) {
			isExist = true
			break
		}
	}
	if finish && !isExist {
		*paths = append(*paths, pathCopy)
	} else {
		if len(path) < int(depth)+1 {
			for _, neighbor := range adjacent {
				dfsPathMaker(neighbor, ender, adjacencyList, pathCopy, paths, depth)
			}
		}
	}
}

func converter(adjacencyList *safeorderedmap.SafeOrderedMap[[]string], start, end string, depth int32) [][]string {
	depth++
	var path []string
	var paths [][]string
	dfsPathMaker("https://en.wikipedia.org/wiki/"+start, end, adjacencyList, path, &paths, depth)
	return paths
}

func algorithm(startLink string, targetLink string) {
	var wg sync.WaitGroup
	var depth int32 = 0
	start := time.Now()
	path := safeorderedmap.New[[]string]()
	linkStart := strings.ReplaceAll(startLink, " ", "_")
	linkTarget := strings.ReplaceAll(targetLink, " ", "_")
	wg.Add(1)
	go func() {
		defer wg.Done()
		crawlerBFS(linkStart, linkTarget, path, &depth, start)
		// crawlerIDS(linkStart, linkTarget, path, &depth)
	}()
	wg.Wait()
	fmt.Println("DEPTH IS", depth+1)
	paths := converter(path, linkTarget, linkStart, depth)
	for _, res := range paths {
		if res[len(res)-1] == "https://en.wikipedia.org/wiki/"+linkStart {
			fmt.Println(res)
		}
	}
	fmt.Println("Runtime:", time.Since(start))
}
