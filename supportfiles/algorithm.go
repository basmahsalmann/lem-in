package supportfiles

import (
	"fmt"
	"os"
)

func Algorithm(fileName string, numberOfAnts int) {
	var results [][]string

	rooms := ReadFile(fileName)

	allPaths := bfsPaths(rooms, startRoom, endRoom)

	if len(allPaths) == 0 {
		fmt.Println("ERROR: invalid data format, no paths found")
		os.Exit(0)
	} else if len(allPaths) > 1 {
		generateCombinations(allPaths)
		results = findMaxLength2DArray(myPaths)

	} else {
		// if length is 1, then take that path
		results = allPaths
	}

	ReadFileAndPrint(fileName)
	paths := assignPaths(results, numberOfAnts)
	MoveAnts(paths, numberOfAnts)
}

func bfsPaths(graph map[string]*Room, start string, end string) [][]string {
	var paths [][]string
	queue := [][]string{{start}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastNode := path[len(path)-1]
		if lastNode == end {
			paths = append(paths, path)
			continue
		}
		for _, neighbor := range graph[lastNode].links {
			if !contains(path, neighbor) {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return paths
}

func generateCombinations(paths [][]string) {
	// tracks the return of paths with similar rooms to eliminate repetitive checking
	var trackPaths [][][]string

	for _, path := range paths {
		for i, str := range path {
			// ignore start and end rooms
			if i == 0 || i == len(path)-1 {
				continue
			}

			// returns the paths that shares a similar room
			shared, samePaths := sharesRoom(paths, str)

			if shared && !containsPath(trackPaths, samePaths) && len(samePaths) != 1 {
				trackPaths = append(trackPaths, samePaths)
				for _, each := range samePaths {
					FilterPaths(paths, samePaths, each)
				}

			} else if (len(samePaths) == 1 || len(samePaths) == 2) && !shared {
				myPaths = append(myPaths, paths)
			}

		}
	}
}

func findMaxLength2DArray(arr [][][]string) [][]string {
	var maxLength int
	var max2DArray [][]string

	// Loop through the 3D array
	for _, twoDArr := range arr {
		// Check if the current 2D array has a greater length than the current maximum
		if len(twoDArr) > maxLength {
			maxLength = len(twoDArr)
			max2DArray = twoDArr
		}
	}
	return max2DArray
}

func assignPaths(allPaths [][]string, numberOfAnts int) [][]string {
	queue := make(map[int]int)              // each path has a queue where we increment whenever we assign a path to an ant
	cost := make(map[int]int)               // number of ants in queue + number of rooms in the path
	paths := make([][]string, numberOfAnts) // each ant gets a path

	for i := 0; i < numberOfAnts; i++ {
		for j := range allPaths {
			numOfRooms := len(allPaths[j])
			countAnts := queue[j]

			cost[j] = numOfRooms + countAnts
		}
		antPath := bestCost(cost)
		if antPath < 0 || antPath >= len(allPaths) { // ensure that the index is within bounds
			panic(fmt.Sprintf("bestCost returned an invalid index: %d", antPath))
		}
		paths[i] = allPaths[antPath]
		queue[antPath]++
	}
	return paths
}

func containsPath(arrs [][][]string, arr [][]string) bool {
	for _, paths := range arrs {
		if len(paths) == len(arr) {
			allEqual := true
			for i, path := range paths {
				if len(path) != len(arr[i]) {
					allEqual = false
					break
				}
				for j := range path {
					if path[j] != arr[i][j] {
						allEqual = false
						break
					}
				}
				if !allEqual {
					break
				}
			}
			if allEqual {
				return true
			}
		}
	}
	return false
}

// Function to remove specific paths and call the recursive filtering
func FilterPaths(allPaths, paths [][]string, onePath []string) [][]string {
	// Create a copy of allPaths
	var filteredPaths [][]string
	for _, p := range allPaths {
		filteredPaths = append(filteredPaths, append([]string(nil), p...))
	}

	// Remove paths that are in `paths` but keep `onePath`
	var tempPaths [][]string
	for _, fp := range filteredPaths {
		keep := true
		for _, p := range paths {
			if fmt.Sprintf("%v", fp) == fmt.Sprintf("%v", p) && fmt.Sprintf("%v", fp) != fmt.Sprintf("%v", onePath) {
				keep = false
				break
			}
		}
		if keep {
			tempPaths = append(tempPaths, fp)
		}
	}

	filteredPaths = tempPaths

	if shared := hasSharedString(filteredPaths); !shared && len(filteredPaths) != 0 {
		myPaths = append(myPaths, filteredPaths)
		return filteredPaths
	} else {
		// Call the recursive function to filter out paths that branch out
		return filterBranchingPaths(filteredPaths)
	}
}

func filterBranchingPaths(filteredPaths [][]string) [][]string {
	for _, path := range filteredPaths {
		// Check each room in the path (excluding first and last)
		for i := 1; i < len(path)-1; i++ {
			room := path[i]
			if shared, arrOfPaths := sharesRoom(filteredPaths, room); shared { //
				// Call FilterPaths recursively for paths sharing the same room until there are no clashing rooms
				for _, p := range arrOfPaths {
					filteredPaths = FilterPaths(filteredPaths, arrOfPaths, p)
				}
			}
		}
	}
	return filteredPaths
}

func bestCost(cost map[int]int) int {
	if len(cost) == 0 {
		return -1 // handle empty cost map case
	}

	min := cost[0]
	minIndex := 0
	for k, v := range cost {
		if v < min {
			min = v
			minIndex = k
		}
	}

	return minIndex
}

func BfsPaths(graph map[string]*Room, start string, end string) [][]string {
	var paths [][]string
	queue := [][]string{{start}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastNode := path[len(path)-1]
		if lastNode == end {
			paths = append(paths, path)
			continue
		}
		for _, neighbor := range graph[lastNode].links {
			if !contains(path, neighbor) {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return paths
}

func AssignPaths(allPaths [][]string, numberOfAnts int) [][]string {
	queue := make(map[int]int)              // each path has a queue where we increment whenever we assign a path to an ant
	cost := make(map[int]int)               // number of ants in queue + number of rooms in the path
	paths := make([][]string, numberOfAnts) // each ant gets a path

	for i := 0; i < numberOfAnts; i++ {
		for j := range allPaths {
			numOfRooms := len(allPaths[j])
			countAnts := queue[j]

			cost[j] = numOfRooms + countAnts
		}
		antPath := bestCost(cost)
		if antPath < 0 || antPath >= len(allPaths) { // ensure that the index is within bounds
			panic(fmt.Sprintf("bestCost returned an invalid index: %d", antPath))
		}
		paths[i] = allPaths[antPath]
		queue[antPath]++
	}
	return paths
}

func GenerateCombinations(paths [][]string) {
	// tracks the return of paths with similar rooms to eliminate repetitive checking
	var trackPaths [][][]string

	for _, path := range paths {
		for i, str := range path {
			// ignore start and end rooms
			if i == 0 || i == len(path)-1 {
				continue
			}

			// returns the paths that shares a similar room
			shared, samePaths := sharesRoom(paths, str)

			if shared && !containsPath(trackPaths, samePaths) && len(samePaths) != 1 {
				trackPaths = append(trackPaths, samePaths)
				for _, each := range samePaths {
					FilterPaths(paths, samePaths, each)
				}

			} else if (len(samePaths) == 1 || len(samePaths) == 2) && !shared {
				myPaths = append(myPaths, paths)
			}

		}
	}
}
