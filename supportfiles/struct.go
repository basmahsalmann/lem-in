package supportfiles

// Room structure to store room information
type Room struct {
	name  string
	x, y  int      // coordinates
	links []string // connected rooms
}

// Global variables to store start and end room names
var startRoom, endRoom string

// Adjacency list representation of the graph
var rooms map[string]*Room

// this is for the case of paths with length of 2 (start-end); it checks if an ant moved at a step or not
var trackSteps bool

var myPaths [][][]string
