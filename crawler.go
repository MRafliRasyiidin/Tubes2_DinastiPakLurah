package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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
		for i, j := 0, len(pathCopy)-1; i < j; i, j = i+1, j-1 {
			pathCopy[i], pathCopy[j] = pathCopy[j], pathCopy[i]
		}
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

func caller(linkStart, linkTarget string, isBFS, searchAll bool, depth, visitCount *int32, timer *time.Duration) []byte {
	start := time.Now()
	path := safeorderedmap.New[[]string]()
	if isBFS {
		crawlerBFS(linkStart, linkTarget, path, depth, visitCount, searchAll, start)
	} else {
		crawlerIDS(linkStart, linkTarget, path, depth, visitCount, searchAll, start)
	}
	paths := converter(path, linkTarget, linkStart, *depth)

	*timer = time.Since(start)
	jsonString, err := json.Marshal(paths)
	if err != nil {
		fmt.Println(err)
	}
	return jsonString
}

func main() {
	var visitCount int32
	var depth int32
	var timer time.Duration
	result := caller("Medan_Prijaji", "Adolf_Hitler", true, false, &depth, &visitCount, &timer)
	fmt.Println(string(result))
	fmt.Println(timer)
}
