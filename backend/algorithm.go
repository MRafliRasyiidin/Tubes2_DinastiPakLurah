package main

import (
	//"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

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

func extractTitle(url string) string {
	parts := strings.Split(url, "/")
	title := parts[len(parts)-1]
	title = strings.ReplaceAll(title, "_", " ")
	return title
}

func dfsPathMaker(node, ender string, adjacencyList *safeorderedmap.SafeOrderedMap[[]string], path []string, paths *[][]string, depth int32) {
	adjacent, _ := adjacencyList.Get(node)
	finish := node == "https://en.wikipedia.org/wiki/"+ender

	path = append(path, node)
	pathCopy := []string{}
	pathCopy = append(pathCopy, path...)
	isExist := false
	for _, value := range *paths {
		if reflect.DeepEqual(value, pathCopy) {
			isExist = true
			break
		}
	}
	if finish && !isExist {
		for i, j := 0, len(pathCopy)-1; i < j; i, j = i+1, j-1 {
			fmt.Println(extractTitle(pathCopy[i]))
			pathCopy[i], pathCopy[j] = extractTitle(pathCopy[j]), extractTitle(pathCopy[i])
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

func caller(linkStart, linkTarget string, isBFS, searchAll bool, depth, visitCount *int32, timer *time.Duration) [][]string {
	start := time.Now()
	path := safeorderedmap.New[[]string]()
	if isBFS {
		crawlerBFS(linkStart, linkTarget, path, depth, visitCount, searchAll, start)
	} else {
		crawlerIDS(linkStart, linkTarget, path, depth, visitCount, searchAll, start)
	}
	paths := converter(path, linkTarget, linkStart, *depth)

	*timer = time.Since(start)
	return paths
}

func runAlgorithm(start string, target string, bfs bool, all bool) ([][]string, time.Duration, int32) {
	var visitCount int32
	var depth int32
	var timer time.Duration
	result := caller(start, target, bfs, all, &depth, &visitCount, &timer)
	fmt.Println(result)
	fmt.Println(timer)
	return result, timer, visitCount
}
