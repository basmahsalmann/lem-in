package supportfiles

import (
	"strings"
	"strconv"
)

func addLink(rooms map[string]*Room, room1, room2 string) {
	if roomPtr1, ok := rooms[room1]; ok {
		roomPtr1.links = append(roomPtr1.links, room2)
	}
	if roomPtr2, ok := rooms[room2]; ok {
		roomPtr2.links = append(roomPtr2.links, room1) // Adding bidirectional link
	}
}

func sharesRoom(arr [][]string, str string) (bool, [][]string) {
	found := false
	arrOfPaths := [][]string{}
	for i := range arr {
		for _, s := range arr[i] {
			if s == str {
				found = true
				arrOfPaths = append(arrOfPaths, arr[i])
			}
		}
	}

	// If the path returns one, it is the path we are ranging over in generateCombination func, thus not counted
	if found && len(arrOfPaths) != 1 {
		return true, arrOfPaths
	} else {
		return false, arrOfPaths
	}
}

// Checks if a set of paths share a room and returns a boolean
func hasSharedString(paths [][]string) bool {
	// Use a map to track the occurrence of strings
	seen := make(map[string]struct{})

	// Iterate over each path in the 2D array
	for _, path := range paths {
		// Skip paths that are too short to have middle elements
		if len(path) <= 2 {
			continue
		}

		// Check the middle elements, excluding the first and last
		for i := 1; i < len(path)-1; i++ {
			// If the string has been seen before, return true
			if _, exists := seen[path[i]]; exists {
				return true
			}
			// Mark the string as seen
			seen[path[i]] = struct{}{}
		}
	}
	// If no shared string is found, return false
	return false
}


func isRoom(str string) bool {
	room := strings.Split(str, " ")
	
	if len(room) == 3 {
		if _, err := strconv.Atoi(room[1]); err == nil {
			return false
		}
		if _, err := strconv.Atoi(room[2]); err == nil {
			return false
		}
		return true
	}

	return false
}

//checks if link has already been made
func duplicateLink(rooms map[string]*Room, link1 []string, link2 []string) bool {
	for _, room := range rooms {
		if room.name == link1[0] {
			for _, link := range room.links {
				if link == link1[1] {
					return true
				}
			}
		}
		if room.name == link1[1] {
			for _, link := range room.links {
				if link == link1[0] {
					return true
				}
			}
		}
		if room.name == link2[0] {
			for _, link := range room.links {
				if link == link2[1] {
					return true
				}
			}
		}
		if room.name == link2[1] {
			for _, link := range room.links {
				if link == link2[0] {
					return true
				}
			}
		}
	}
	return false
}