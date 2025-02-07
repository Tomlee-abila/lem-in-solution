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

	var antFarm AntFarm
	antFarm.Start = "1"
	antFarm.End = "0"
	antFarm.Ants = 3
	antFarm.Rooms = make(map[string]*Room)

	arrP := strings.Split(p, "\n")

	fmt.Println(arrP)

	for _, t := range arrP {
		r := strings.Split(t, "-")

		// Ensure rooms exist before modifying them
		if _, exists := antFarm.Rooms[r[0]]; !exists {
			antFarm.Rooms[r[0]] = &Room{Name: r[0], Links: []string{}}
		}

		if _, exists := antFarm.Rooms[r[1]]; !exists {
			antFarm.Rooms[r[1]] = &Room{Name: r[1], Links: []string{}}
		}

		antFarm.Rooms[r[0]].Links, _ = appendIfNotExists(antFarm.Rooms[r[0]].Links, r[1])
		antFarm.Rooms[r[1]].Links, _ = appendIfNotExists(antFarm.Rooms[r[1]].Links, r[0])

		antFarm.Rooms[r[0]].Name = r[0]
		antFarm.Rooms[r[1]].Name = r[1]
	}

	antFarm.pathsCreation(antFarm.Start, []string{"1"})
	for i, path := range antFarm.Paths {
		fmt.Println(i, ":", path)
	}
	// fmt.Println("paths\n", antFarm.Paths)

	antFarm.removeInvalidPaths()
	fmt.Println("Valid paths\n", antFarm.ValidPaths)
	antFarm.optimalPath()
	fmt.Println("Optimal paths\n", antFarm.ValidPaths)
}

func appendIfNotExists(slice []string, item string) ([]string, bool) {
	if !contains(slice, item) {
		slice = append(slice, item)
		return slice, true
	}
	return slice, false
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (antFarm *AntFarm) pathsCreation(start string, path []string) {
	nPath := append([]string{}, path[0:]...)

	for i, room := range antFarm.Rooms[start].Links {
		newPath := append([]string{}, nPath...)
		cnd := false
		if i == 0 {
			path, cnd = appendIfNotExists(path, room)
			if cnd {
				if antFarm.End != room {
					antFarm.pathsCreation(room, path)
				} else {
					antFarm.Paths = append(antFarm.Paths, path)
				}
			}
		} else {
			newPath, cnd = appendIfNotExists(newPath, room)
			if cnd {
				if antFarm.End != room {
					antFarm.pathsCreation(room, newPath)
				} else {
					antFarm.Paths = append(antFarm.Paths, newPath)
					fmt.Println("complete 2", newPath)
				}
			}
		}
	}
}

func (antFarm *AntFarm) removeInvalidPaths() {
	// if len(antFarm.Paths) <= len(antFarm.Rooms[antFarm.Start].Links) && len(antFarm.Paths) <= antFarm.Ants {
	// 	return // No need to check if there's only one or no paths
	// }
	if len(antFarm.Paths) < 1{
		return
	}

	// Assuming this returns an index of the shortest path
	i := antFarm.findShortestPath()
	shortest := append([]string{}, antFarm.Paths[i]...)
	antFarm.ValidPaths = append(antFarm.ValidPaths, shortest)
	fmt.Println("shortest parth", shortest, "\npaths", antFarm.Paths, "length", len(antFarm.Paths), "\nvalid paths", antFarm.ValidPaths)

	for j := 1; j < len(shortest); j++ { // Start from 1 to avoid start room
		toRemove := []int{} // Track indices to remove

		farm := append([][]string{}, antFarm.Paths...)
		room := antFarm.Paths[i][j]

		if room == antFarm.End {
			break // Skip if it's the end room
		}

		for k, path := range farm {
			if k == i {
				// toRemove = append(toRemove, k)
				continue // Skip comparing to itself
			}

			if contains(path, room) {
				toRemove = append(toRemove, k)
			}
			// if j < len(path) && path[j] == room { // Check if another path shares the same room
			//     toRemove = append(toRemove, k)
			// }
		}
		// Remove paths from the list (in reverse order to avoid shifting issues)

		for x := len(toRemove) - 1; x >= 0; x-- {
			farm = append(farm[:toRemove[x]], farm[toRemove[x]+1:]...)
		}
		antFarm.Paths = farm
		i = antFarm.findShortestPath()
	}

	antFarm.Paths = append(antFarm.Paths[:i], antFarm.Paths[i+1:]...)
	antFarm.removeInvalidPaths()
}

func (antFarm AntFarm) optimalPath() {
	paths := []int{}
	count := antFarm.Ants

	for _, path := range antFarm.ValidPaths {
		paths = append(paths, len(path))
	}

	for count > 0 {
		index := smallest(paths)
		paths[index] += 1
		count--
	}

	for i := len(paths) - 1; i >= 0; i-- {
		if paths[i] == len(antFarm.ValidPaths[i]) {
			antFarm.ValidPaths = append(antFarm.ValidPaths[:i], antFarm.Paths[i+1:]...)
		}
	}
}

func smallest(slice []int) int {
	min := -1
	var index int
	for i := range slice {
		if min == -1 {
			min = slice[i]
			index = i
		}

		if slice[i] < min {
			min = slice[i]
			index = i
		}
	}
	return index
}

func (antFarm AntFarm) findShortestPath() int {
	min := 0
	var result int
	for i, path := range antFarm.Paths {
		if min == 0 {
			min = len(path)
			result = i
		}
		if len(path) < min {
			min = len(path)
			result = i
		}
	}
	return result
}

type Room struct {
	Name  string
	Links []string
}

type AntFarm struct {
	Ants       int
	Rooms      map[string]*Room
	Start      string
	End        string
	Tunnels    []string
	Paths      [][]string
	ValidPaths [][]string
}
