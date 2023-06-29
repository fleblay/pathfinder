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
	if pos.X > len(world[0])-1 ||
		pos.Y > len(world)-1 ||
		pos.X < 0 ||
		pos.Y < 0 ||
		world[pos.Y][pos.X] == '1' {
		return false
	}
	return true
}

func getNextMoves(world [][]byte, path []byte, player Pos2D, seen []Pos2D) (nextPaths [][]byte, nextPoses []Pos2D, nextSeen []Pos2D) {
	for key, value := range directions {
		nextPos := Pos2D{player.X + value.X, player.Y + value.Y}
		if isPosOk(world, nextPos) == true &&
			Index(seen, nextPos) == -1 &&
			Index(nextSeen, nextPos) == -1 {
			nextPaths = append(nextPaths, DeepCopyAndAdd(path, key))
			nextPoses = append(nextPoses, nextPos)
			nextSeen = append(nextSeen, nextPos)
		}
	}
	return
}

func BFS(world [][]byte) (path []byte) {
	goalPos := findItem(world, 'E')
	seen := []Pos2D{findItem(world, 'S')}
	posQueue := []Pos2D{findItem(world, 'S')}
	pathQueue := [][]byte{{}}

	for len(posQueue) > 0 {
		currentPos := posQueue[0]
		posQueue = posQueue[1:]

		currentPath := pathQueue[0]
		pathQueue = pathQueue[1:]

		if goalPos == currentPos {
			return currentPath
		}
		nextPaths, nextPoses, nextSeen := getNextMoves(world, currentPath, currentPos, seen)
		posQueue = append(posQueue, nextPoses...)
		pathQueue = append(pathQueue, nextPaths...)
		seen = append(seen, nextSeen...)
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
	if path != nil {
		printPath(world, path)
	} else {
		fmt.Println("No solution !")
	}
}
