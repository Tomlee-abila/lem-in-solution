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

	// for _, room := range antFarm.Rooms{
	// 	fmt.Println(room.Name, room.Links)
	// }

	antFarm.pathsCreation(antFarm.Start, []string{"1"})

	fmt.Println("paths\n", antFarm.Paths)
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
				}else{
					antFarm.Paths = append(antFarm.Paths, path)
					fmt.Println("complete", path)
				}
			}
		} else {
			newPath, cnd = appendIfNotExists(newPath, room)
			if cnd {
				
				if antFarm.End != room {
					antFarm.pathsCreation(room, newPath)
				}else{
					antFarm.Paths = append(antFarm.Paths, newPath)
					fmt.Println("complete", newPath)
				}
			}
		}

		// if room != antFarm.End {

		// }

	}
}

// func (antFarm *AntFarm) pathIndex(path []string)int{
// 	var index int
// 	for i, p := range antFarm.Paths{
// 		if path == p{
// 			index = i
// 			break
// 		}
// 	}
// 	return index
// }

type Room struct {
	Name  string
	Links []string
}

type AntFarm struct {
	Ants    int
	Rooms   map[string]*Room
	Start   string
	End     string
	Tunnels []string
	Paths   [][]string
}
