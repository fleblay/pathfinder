package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos2D struct {
	X int
	Y int
}

var directions = map[byte]Pos2D{
	'U': {0, -1},
	'D': {0, 1},
	'L': {-1, 0},
	'R': {1, 0},
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
	if (pos.X > len(world[0])-1) ||
		(pos.Y > len(world)-1) ||
		(pos.X < 0) ||
		(pos.Y < 0) {
		return false
	}
	if world[pos.Y][pos.X] == '1' {
		fmt.Printf("WaLL : [X : %d][Y : %d]\n", pos.X, pos.Y)
		return false
	}
	return true
}

func updatePos(world [][]byte, pos Pos2D, move Pos2D) {
	//fmt.Println("Update pos", move)
	nextPos := Pos2D{pos.X + move.X, pos.Y + move.Y}
	world[pos.Y][pos.X] = '0'
	world[nextPos.Y][nextPos.X] = 'S'
}

func getNextMoves(lastGrid [][]byte, path []byte, player Pos2D, seen []Pos2D) (nextPaths [][]byte, nextPoses []Pos2D, nextSeen []Pos2D) {
	fmt.Println("----------------------------------------")
	defer fmt.Println("----------------------------------------")
	fmt.Println("current Pos :", player, "current Path", string(path))
	fmt.Println(printMap(lastGrid))
	for key, value := range directions {
		nextPos := Pos2D{player.X + value.X, player.Y + value.Y}
		if isPosOk(lastGrid, nextPos) == true &&
			Index(seen, nextPos) == -1 &&
			Index(nextSeen, nextPos) == -1 {
			toAppend := append(path, key)
			fmt.Println("to append path", string(toAppend))
			nextPaths = append(nextPaths, toAppend)
			nextPoses = append(nextPoses, nextPos)
			nextSeen = append(nextSeen, nextPos)
		}
	}
	return
}

func BFS(world [][]byte) (path []byte) {
	seen := make([]Pos2D, 0, 100)
	seen = append(seen, findItem(world, 'S'))

	//pos to check
	posQueue := make([]Pos2D, 0, 100)
	posQueue = append(posQueue, findItem(world, 'S'))

	//path associated
	pathQueue := make([][]byte, 0, 100)
	init := []byte{}
	pathQueue = append(pathQueue, init)

	goalPos := findItem(world, 'E')
	fmt.Printf("Goal : [X : %d][Y : %d]\n", goalPos.X, goalPos.Y)

	for len(posQueue) > 0 {
		//fmt.Println(posQueue, pathQueue)
		currentPos := posQueue[0]
		posQueue = posQueue[1:]

		currentPath := make([]byte, len(pathQueue[0]))
		copy(currentPath, pathQueue[0])
		pathQueue = pathQueue[1:]
		if goalPos == currentPos {
			fmt.Println("found a solution : ", string(currentPath), currentPos)
			return currentPath
		}
		nextPaths, nextPoses, nextSeen := getNextMoves(world, currentPath, currentPos, seen)
		for i := 0; i < len(nextPoses); i++ {
			posQueue = append(posQueue, nextPoses[i])
			pathQueue = append(pathQueue, nextPaths[i])
		}
		for i := 0; i< len(nextSeen); i++ {
			seen = append(seen, nextSeen[i])
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must provide one and only one argument : the map")
		os.Exit(1)
	}
	fd, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file %q : %q\n", os.Args[1], err)
		os.Exit(1)
	}
	world, err := readFile(fd)
	if err != nil {
		fmt.Printf("Error opening scanning file %q : %q\n", os.Args[1], err)
		os.Exit(1)
	}
	path := BFS(world)
	printPath(world, path)
}
