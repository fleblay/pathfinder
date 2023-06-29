package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type World struct {
	Map [][]byte
}

type Pos2D struct {
	X int
	Y int
}

type Move2D struct {
	X int
	Y int
}

var directions = map[byte]Move2D{
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

func findPlayer(world [][]byte) (pos Pos2D) {
	for lineIndex, line := range world {
		for rowIndex, value := range line {
			if value == 'S' {
				return Pos2D{rowIndex, lineIndex}
			}
		}
	}
	return
}

func isPosOk(world [][]byte, pos Pos2D) bool {
	if pos.X >= len(world[0])-1 ||
		pos.Y >= len(world)-1 ||
		world[pos.Y][pos.X] == '1' {
		return false
	}
	return true
}

func checkMove(world [][]byte, pos Pos2D, move Move2D) bool {
	nextPos := Pos2D{pos.X + move.X, pos.Y + move.Y}
	if isPosOk(world, nextPos) == false {
		return false
	}
	return true
}

func updatePos(world [][]byte, pos Pos2D, move Move2D) {
	nextPos := Pos2D{pos.X + move.X, pos.Y + move.Y}
	world[pos.Y][pos.X] = '_'
	world[nextPos.Y][nextPos.X] = 'S'
}

func getNextMoves(lastGrid [][]byte, moves []byte) (nextMoves [][]byte) {
	player := findPlayer(lastGrid)
	for key, value := range directions {
		if checkMove(lastGrid, player, value) {
			nextMoves = append(nextMoves, append(moves, key))
		}
	}
	return
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
	data, err := readFile(fd)
	if err != nil {
		fmt.Printf("Error opening scanning file %q : %q\n", os.Args[1], err)
		os.Exit(1)
	}
	world := World{data}
	path := []byte{
		'R',
		'R',
		'R',
		'R',
		'U',
		'U',
	}
	printPath(world.Map, path)
	fmt.Println(getNextMoves(world.Map, path))
}
