package main

import (
	"fmt"
	"strings"
)

func main() {
	p := `0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5`

	antFarm := AntFarm{
		Start: "1",
		End:   "0",
		Ants:  3,
		Rooms: make(map[string]*Room),
	}

	// Parse input paths
	for _, t := range strings.Split(p, "\n") {
		r := strings.Split(t, "-")

		for _, room := range r {
			if _, exists := antFarm.Rooms[room]; !exists {
				antFarm.Rooms[room] = &Room{Name: room}
			}
		}

		antFarm.Rooms[r[0]].Links = appendIfNotExists(antFarm.Rooms[r[0]].Links, r[1])
		antFarm.Rooms[r[1]].Links = appendIfNotExists(antFarm.Rooms[r[1]].Links, r[0])
	}

	antFarm.findPaths(antFarm.Start, []string{antFarm.Start})

	fmt.Println("All Paths:")
	for i, path := range antFarm.Paths {
		fmt.Println(i, ":", path)
	}
	fmt.Println("done")

	antFarm.removeInvalidPaths()
	fmt.Println("Valid Paths:", antFarm.ValidPaths)

	antFarm.findOptimalPath()
	fmt.Println("Optimal Paths:", antFarm.ValidPaths)
}

// Append item to slice only if it doesn't exist
func appendIfNotExists(slice []string, item string) []string {
	for _, v := range slice {
		if v == item {
			return slice
		}
	}
	return append(slice, item)
}

// Recursively find all paths
func (antFarm *AntFarm) findPaths(current string, path []string) {
	nPath := append([]string{}, path...)
	for _, room := range antFarm.Rooms[current].Links {
		if contains(path, room) {
			continue
		}

		newPath := append(nPath, room)
		if room == antFarm.End {
			antFarm.Paths = append(antFarm.Paths, newPath)
		} else {
			antFarm.findPaths(room, newPath)
		}
	}
}

// Remove conflicting paths
func (antFarm *AntFarm) removeInvalidPaths() {
	if len(antFarm.Paths) == 0 {
		fmt.Println("done removing")
		return
	}
	
	for len(antFarm.Paths) > 0 {
		// fmt.Println("length of paths", len(antFarm.Paths))
		i := antFarm.findShortestPath()
		shortest := append([]string{},antFarm.Paths[i]...)
		antFarm.ValidPaths = append(antFarm.ValidPaths, shortest)
		// fmt.Println(shortest)

		toRemove := make(map[int]bool)
		toRemove[i] = true

		for k, path := range antFarm.Paths {
			if k == i {
				continue
			}
			for _, room := range shortest[1:] {
				if contains(path, room) && room != antFarm.End{
					fmt.Println("path",k,path,"has",room)
					toRemove[k] = true
					break
				}
			}
		}
		fmt.Println("To be removed", toRemove, "length of paths", len(antFarm.Paths))

		// Rebuild Paths excluding removed ones
		var newPaths [][]string
		for k, path := range antFarm.Paths {
			if !toRemove[k] {
				newPaths = append(newPaths, path)
			}
		}
		toRemove = make(map[int]bool)		
		antFarm.Paths = newPaths
		fmt.Println("valid paths", antFarm.ValidPaths)
		fmt.Println("To be removed", toRemove, "length of paths", len(antFarm.Paths), "\nPaths", antFarm.Paths)
	}
}

// Find the most balanced optimal paths
func (antFarm *AntFarm) findOptimalPath() {
	pathLengths := make([]int, len(antFarm.ValidPaths))
	antsLeft := antFarm.Ants

	for i, path := range antFarm.ValidPaths {
		pathLengths[i] = len(path)
	}

	for antsLeft > 0 {
		index := findSmallestIndex(pathLengths)
		pathLengths[index]++
		antsLeft--
	}

	var newValidPaths [][]string
	for i := range antFarm.ValidPaths {
		if pathLengths[i] <= len(antFarm.ValidPaths[i]) {
			newValidPaths = append(newValidPaths, antFarm.ValidPaths[i])
		}
	}
	antFarm.ValidPaths = newValidPaths
}

// Find index of the shortest path
func (antFarm *AntFarm) findShortestPath() int {
	minIdx, minLen := 0, len(antFarm.Paths[0])
	for i, path := range antFarm.Paths {
		if len(path) < minLen {
			minLen = len(path)
			minIdx = i
		}
	}
	return minIdx
}

// Find index of the smallest element in a slice
func findSmallestIndex(slice []int) int {
	minIdx, minVal := 0, slice[0]
	for i, val := range slice {
		if val < minVal {
			minIdx = i
			minVal = val
		}
	}
	return minIdx
}

// Check if a slice contains an item
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Structs for Room and AntFarm
type Room struct {
	Name  string
	Links []string
}

type AntFarm struct {
	Ants       int
	Rooms      map[string]*Room
	Start      string
	End        string
	Paths      [][]string
	ValidPaths [][]string
}
