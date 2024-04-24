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
	"sync"
	"time"

	"github.com/thalesfsp/go-common-types/safeorderedmap"
)

func dfs(node, ender string, visited map[string]bool, adjacencyList *safeorderedmap.SafeOrderedMap[[]string], path []string, paths *[][]string, depth int32) {
	adjacent, _ := adjacencyList.Get(node)
	// fmt.Println(path)
	finish := node == "https://en.wikipedia.org/wiki/"+ender
	path = append(path, node)
	if finish {
		*paths = append(*paths, path)
	} else {
		if len(path) < int(depth)+1 {
			for _, neighbor := range adjacent {
				if !visited[neighbor] {
					dfs(neighbor, ender, visited, adjacencyList, path, paths, depth)
				}
			}
		}
	}
	path = path[:len(path)-1]
	visited[node] = false
}

func converter(adjacencyList *safeorderedmap.SafeOrderedMap[[]string], start, end string, depth int32) [][]string {
	visited := make(map[string]bool)
	var path []string
	var paths [][]string

	dfs("https://en.wikipedia.org/wiki/"+start, end, visited, adjacencyList, path, &paths, depth)

	return paths
}

func main() {
	var wg sync.WaitGroup
	var depth int32 = 0
	start := time.Now()
	path := safeorderedmap.New[[]string]()
	linkStart := "Medan_Prijaji"
	linkTarget := "Adolf_Hitler"
	wg.Add(1)
	go func() {
		defer wg.Done()
		crawlerBFS(linkStart, linkTarget, path, &depth)
	}()
	wg.Wait()
	// path.Each(func(key string, value []string) {
	// 	fmt.Println(key)
	// 	fmt.Println(value)
	// })
	fmt.Println("DEPTH IS", depth)
	paths := converter(path, linkTarget, linkStart, depth)
	// // fmt.Println(paths)
	min := 1000
	for _, res := range paths {
		fmt.Println("AAA")
		if min > len(res) && res[len(res)-1] == "https://en.wikipedia.org/wiki/"+linkStart {
			fmt.Println("ABA")
			min = len(res)
		}
	}
	for _, res := range paths {

		if len(res) > 0 {

			if res[len(res)-1] == "https://en.wikipedia.org/wiki/"+linkStart && len(res) <= min {

				fmt.Println(res)
			}
		}
	}
	fmt.Println("IT IS DONE", time.Since(start), paths[0])
	// path.Each(func(key string, value []string) {
	// 	fmt.Println("Key", key)
	// 	fmt.Println(value)
	// })
	// path.Each(func(key string, value []string) {
	// 	fmt.Println("Key", key)
	// 	fmt.Println("Value,", value)
	// })

}

// path.Each(func(key string, value []string) {
// 	fmt.Println(value)
// })
// }

// key := "https://en.wikipedia.org/wiki/" + target
// expPath := []string{}
// for key != "https://en.wikipedia.org/wiki/"+start {
// 	fmt.Println("Key : ", key)
// 	expPath = append(expPath, key)
// 	value, _ := path.Get(key)
// 	key = value.(string)
// }

// for i, j := 0, len(expPath)-1; i < j; i, j = i+1, j-1 {
// 	expPath[i], expPath[j] = expPath[j], expPath[i]
// }

// jsonStr, err := json.Marshal(expPath)
// if err != nil {
// 	fmt.Printf("Error: %s", err.Error())
// } else {
// 	fmt.Println(string(jsonStr))
// }
