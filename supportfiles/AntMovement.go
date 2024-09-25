package supportfiles

import (
	"fmt"
	"strings"
)

// Refactored moveAnts function
func MoveAnts(paths [][]string, numAnts int) {
	steps, positions, occupied, finished := initialize(numAnts)
	remaining := numAnts

	for remaining > 0 {
		stepOutput, moveOccurred := processAnts(numAnts, paths, positions, occupied, finished, &remaining)

		if len(stepOutput) > 0 {
			steps = append(steps, strings.Join(stepOutput, " "))
			trackSteps = false
		}

		if !moveOccurred {
			break
		}
	}

	outputSteps(steps)
	checkUnfinishedAnts(finished, endRoom)
}

func processAnts(numAnts int, paths [][]string, positions []int, occupied map[string]int, finished []bool, remaining *int) ([]string, bool) {
	moveOccurred := false
	stepOutput := []string{}
	trackSteps = false

	// Attempt to move each ant one at a time
	for i := 0; i < numAnts; i++ {
		antMoved, step := moveAnt(i, paths, positions, occupied, finished, remaining)
		if antMoved {
			moveOccurred = true
			stepOutput = append(stepOutput, step)
		}
	}
	return stepOutput, moveOccurred
}

func initialize(numAnts int) ([]string, []int, map[string]int, []bool) {
	steps := make([]string, 0)
	positions := make([]int, numAnts)
	occupied := make(map[string]int)
	finished := make([]bool, numAnts)
	return steps, positions, occupied, finished
}

func outputSteps(steps []string) {
	for _, step := range steps {
		fmt.Println(step)
	}
}

func checkUnfinishedAnts(finished []bool, endRoom string) {
	for i, m := range finished {
		if !m {
			fmt.Printf("Ant %d did not reach end room (%s)\n", i+1, endRoom)
		}
	}
}

func moveAnt(antIndex int, paths [][]string, positions []int, occupied map[string]int, finished []bool, remaining *int) (bool, string) {
	moveOccurred := false
	stepOutput := ""

	// special case for paths with two rooms (start and end)
	if len(paths[antIndex]) == 2 && !trackSteps && !finished[antIndex] && positions[antIndex] < len(paths[antIndex])-1 {
		nextRoom := paths[antIndex][positions[antIndex]+1]
		if positions[antIndex] == 0 {
			positions[antIndex]++
			stepOutput = fmt.Sprintf("L%d-%s", antIndex+1, paths[antIndex][1])
			trackSteps = true
			moveOccurred = true

			if nextRoom == endRoom {
				finished[antIndex] = true
				(*remaining)--
			}
		}
		return moveOccurred, stepOutput
	} else if len(paths[antIndex]) == 2 && trackSteps {
		return moveOccurred, stepOutput
	}

	if !finished[antIndex] && positions[antIndex] < len(paths[antIndex])-1 {
		currentRoom := paths[antIndex][positions[antIndex]]
		nextRoom := paths[antIndex][positions[antIndex]+1]

		// Check if the next room is unoccupied or is the end room
		if occupied[nextRoom] == 0 || nextRoom == endRoom {

			// If moving from the start room, don't decrement the occupied count
			if currentRoom != startRoom {
				occupied[currentRoom]--
			}

			positions[antIndex]++
			occupied[nextRoom]++
			stepOutput = fmt.Sprintf("L%d-%s", antIndex+1, nextRoom)
			moveOccurred = true

			if nextRoom == endRoom {
				finished[antIndex] = true
				(*remaining)--
			}
		}
	}

	return moveOccurred, stepOutput
}
