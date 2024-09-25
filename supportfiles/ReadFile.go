package supportfiles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numberOfAnts int

func ReadFile(filename string) map[string]*Room {
	// Adjacency list representation of the graph
	var rooms map[string]*Room

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	rooms = make(map[string]*Room)

	scanner := bufio.NewScanner(file)

	// To ensure taking the line after '##start' or '##end' and prevent unncessary updates
	var processingStart, processingEnd bool

	// To check for start and end rooms
	var startRoomSet, endRoomSet bool

	seenCoordinates := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		if numberOfAnts == 0 {
			numberOfAnts, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println("ERROR: invalid data format, invalid number of ants")
				os.Exit(0)
			}

			if numberOfAnts > 100000 {
				fmt.Println("ERROR: invalid data format, ant limit is 100000")
				os.Exit(0)
			}
			continue
		}

		if strings.Contains(line, "##start") {
			if startRoomSet {
				fmt.Println("ERROR: Multiple start rooms found")
				os.Exit(1)
			}
			processingStart = true
			continue
		} else if strings.Contains(line, "##end") {
			if endRoomSet {
				fmt.Println("ERROR: Multiple end rooms found")
				os.Exit(1)
			}
			processingEnd = true
			continue
		} else if processingStart {
			roomInfo := strings.Fields(line)
			if len(roomInfo) == 3 {
				startRoom = roomInfo[0]
				coordX, errx := strconv.Atoi(roomInfo[1])
				if errx != nil {
					fmt.Println("Error: ", errx)
					os.Exit(0)
				}
				coordY, erry := strconv.Atoi(roomInfo[2]) //!!add error checking for when the next line is not lenght of 3
				if erry != nil {
					fmt.Println("Error: ", erry)
					os.Exit(0)
				}
				coordKey := fmt.Sprintf("%d,%d", coordX, coordY)
				if _, exists := seenCoordinates[coordKey]; exists {
					fmt.Println("ERROR: Duplicate coordinates found for room:", startRoom)
					os.Exit(1)
				}
				if _, exists := rooms[startRoom]; exists {
					fmt.Println("ERROR: Duplicate room name found:", startRoom)
					os.Exit(1)
				}
				rooms[startRoom] = &Room{name: startRoom, x: coordX, y: coordY, links: make([]string, 0)}
				processingStart = false
				startRoomSet = true
				seenCoordinates[coordKey] = true
				continue
			} else {
				fmt.Println("ERROR: Invalid format for start room details")
				os.Exit(1)
			}
		} else if processingEnd {
			roomInfo := strings.Fields(line)
			if len(roomInfo) == 3 {
				endRoom = roomInfo[0]
				coordX, errx := strconv.Atoi(roomInfo[1])
				if errx != nil {
					fmt.Println("Error: ", errx)
					os.Exit(0)
				}
				coordY, erry := strconv.Atoi(roomInfo[2])
				if erry != nil {
					fmt.Println("Error: ", erry)
					os.Exit(0)
				}
				coordKey := fmt.Sprintf("%d,%d", coordX, coordY)
				if _, exists := seenCoordinates[coordKey]; exists {
					fmt.Println("ERROR: Duplicate coordinates found for room:", endRoom)
					os.Exit(1)
				}
				if _, exists := rooms[endRoom]; exists {
					fmt.Println("ERROR: Duplicate room name found:", endRoom)
					os.Exit(1)
				}
				rooms[endRoom] = &Room{name: endRoom, x: coordX, y: coordY, links: make([]string, 0)}
				processingEnd = false
				endRoomSet = true
				seenCoordinates[coordKey] = true
				continue
			} else {
				fmt.Println("ERROR: Invalid format for end room details")
				os.Exit(1)
			}
		} else if strings.Contains(line, "-") {
			//makes sure it's not a room with a negative coordinate
			if isRoom(line) {
				continue
			}
			linkParts := strings.Split(line, "-")
			if len(linkParts) == 2 {
				// check if room exist
				if rooms[linkParts[0]] == nil || rooms[linkParts[1]] == nil {
					fmt.Println("ERROR: invalid data format, invalid room in link")
					os.Exit(0)
				}
				//makes sure the link is not a duplicate
				if duplicateLink(rooms, linkParts, linkParts) {
					fmt.Println("ERROR: invalid data format, duplicate link")
					os.Exit(0)
				}
				addLink(rooms, linkParts[0], linkParts[1])
			} else {
				fmt.Println("ERROR: invalid data format, invalid link")
				os.Exit(0)
			}
		} else {
			roomInfo := strings.Fields(line)
			if len(roomInfo) == 3 {
				roomName := roomInfo[0]
				coordX, errx := strconv.Atoi(roomInfo[1])
				if errx != nil {
					fmt.Println("Error: ", errx)
					os.Exit(0)
				}
				coordY, erry := strconv.Atoi(roomInfo[2])
				if erry != nil {
					fmt.Println("Error: ", erry)
					os.Exit(0)
				}
				coordKey := fmt.Sprintf("%d,%d", coordX, coordY)
				if _, exists := seenCoordinates[coordKey]; exists {
					fmt.Println("ERROR: Duplicate coordinates found for room:", roomName)
					os.Exit(1)
				}
				seenCoordinates[coordKey] = true

				if _, exists := rooms[roomName]; exists {
					fmt.Println("ERROR: Duplicate room name found:", roomName)
					os.Exit(1)
				}
				rooms[roomName] = &Room{name: roomName, x: coordX, y: coordY, links: make([]string, 0)}
			}
		}
	}

	if !startRoomSet {
		fmt.Println("ERROR: invalid data format, no start room found")
		os.Exit(0)
	}

	if !endRoomSet {
		fmt.Println("ERROR: invalid data format, no end room found")
		os.Exit(0)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return rooms
}

// Prints file contents
func ReadFileAndPrint(filename string) { // should be called after error checking
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Check for errors during the scanning process
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println()
}

func GetNumberofAnts() int {
	return numberOfAnts
}
