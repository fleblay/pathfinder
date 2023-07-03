package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"flag"
)

type Pos2D struct {
	X int
	Y int
}

type Node struct {
	pos   Pos2D
	score int
}

var directions = map[byte]Pos2D{
	'U': {0, -1},
	'D': {0, 1},
	'L': {-1, 0},
	'R': {1, 0},
}

var heuristic = map[string]func(pos, startPos, goalPos Pos2D)int{
	"bfs": FIFO(),
	"dij": startToCurrent,
	"greed": currentToGoal,
	"astar": ASTAR,
}

func readFile(fd *os.File) (world [][]byte, err error) {
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := []byte(strings.Join(strings.Split(scanner.Text(), " "), ""))
		world = append(world, line)
	}
	return world, scanner.Err()
}

func findItem(world [][]byte, toFind byte) (pos Pos2D) {
	pos = Pos2D{-1, -1}
	for lineIndex, line := range world {
		for rowIndex, value := range line {
			if value == toFind {
				return Pos2D{rowIndex, lineIndex}
			}
		}
	}
	return
}

func isPosOk(world [][]byte, pos Pos2D) bool {
	if pos.X > len(world[0])-1 ||
		pos.Y > len(world)-1 ||
		pos.X < 0 ||
		pos.Y < 0 ||
		world[pos.Y][pos.X] == '1' {
		return false
	}
	return true
}

func getNextNodeIndex(queue []Node) int {
	retScore := queue[0].score
	ret := 0
	for index, value := range queue {
		if value.score < retScore {
			retScore = value.score
			ret = index
		}
	}
	return ret
}

func getNextMoves(world [][]byte, startPos, goalPos Pos2D, scoreFx func(pos, startPos, goalPos Pos2D) int, path []byte, currentNode Node, seenNodes []Node) (nextPaths [][]byte, nextNodes []Node, nextSeen []Node) {
	for key, value := range directions {
		nextPos := Pos2D{currentNode.pos.X + value.X, currentNode.pos.Y + value.Y}
		score := scoreFx(nextPos, startPos, goalPos)
		nextNode := Node{nextPos, score}
		if isPosOk(world, nextPos) == true &&
			(posAlreadySeen(seenNodes, nextPos) == -1 ||
			score < seenNodes[posAlreadySeen(seenNodes, nextPos)].score) {
			nextPaths = append(nextPaths, DeepCopyAndAdd(path, key))
			nextNodes = append(nextNodes, nextNode)
			nextSeen = append(nextSeen, nextNode)
		}
	}
	return
}

func algo(world [][]byte, scoreFx func(pos, startPos, goalPos Pos2D) int) (currentPath []byte, seenPos []Node, tries int, maxSizeQueue int) {
	goalPos := findItem(world, 'E')
	startPos := findItem(world, 'S')
	seenPos = []Node{{findItem(world, 'S'), 0}}
	posQueue := DeepCopyAndAdd(seenPos)
	pathQueue := [][]byte{{}}

	for ; len(posQueue) > 0; tries++ {
		maxSizeQueue = Max(maxSizeQueue, len(posQueue))

		nextIndex := getNextNodeIndex(posQueue)
		currentPos := posQueue[nextIndex]
		posQueue = append(posQueue[:nextIndex], posQueue[nextIndex+1:]...)

		currentPath = pathQueue[nextIndex]
		pathQueue = append(pathQueue[:nextIndex], pathQueue[nextIndex+1:]...)

		if goalPos == currentPos.pos {
			return
		}
		nextPaths, nextPoses, nextSeen := getNextMoves(world, startPos, goalPos, scoreFx, currentPath, currentPos, seenPos)
		posQueue = append(posQueue, nextPoses...)
		pathQueue = append(pathQueue, nextPaths...)
		seenPos = append(seenPos, nextSeen...)
	}
	return nil, seenPos, tries, maxSizeQueue
}

func main() {
	var selectedAlgo, inputFile string
	flag.StringVar(&selectedAlgo, "a", "astar", "usage : [bfs|dij|greed|astar]")
	flag.StringVar(&inputFile, "f", "map.txt", "usage : file")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("You must provide at least one argument : the map")
		os.Exit(1)
	}
	fd, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file %q : %q\n", os.Args[1], err)
		os.Exit(1)
	}
	world, err := readFile(fd)
	if err != nil {
		fmt.Printf("Error opening scanning file %q : %q\n", os.Args[1], err)
		os.Exit(1)
	}
	path, seenPos, tries, sizeMax := algo(world, heuristic[selectedAlgo])
	if path != nil {
		printPath(world, path)
		printMapAndTries(DeepCopyAndAdd(world), seenPos)
		fmt.Printf("Solution found in %d tries. Max size is %d \n", tries, sizeMax)
	} else {
		fmt.Println("No solution !")
	}
}
