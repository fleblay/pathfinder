package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type World struct {
	Map [][]byte
}

// size < 128
type Pos2D struct {
	X uint8
	Y uint8
}

type Move2D struct {
	X int8
	Y int8
}

func readFile(fd *os.File) (world [][]byte, err error) {
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := []byte(strings.Join(strings.Split(scanner.Text(), " "), ""))
		world = append(world, line)
	}
	return world, scanner.Err()
}

func printMap(world [][]byte) (res string) {
	for _, line := range world {
		for pos, value := range line {
			//color
			if value == 'S' {
				res += "\x1b[33m"
			} else if value == 'E' {
				res += "\x1b[32m"
			} else if value == '1' {
				res += "\x1b[31m"
			}
			res += string(value)
			res += "\x1b[0m"
			if pos == len(line)-1 {
				res += "\n"
			} else {
				res += " "
			}
		}
	}
	return
}

func findPlayer(world [][]byte) (pos Pos2D) {
	for lineIndex, line := range world {
		for rowIndex, value := range line {
			if value == 'S' {
				return Pos2D{uint8(rowIndex), uint8(lineIndex)}
			}
		}
	}
	return
}

func isPosOk(world [][]byte, pos Pos2D) bool {
	if pos.X >= uint8(len(world[0]) - 1) ||
		pos.Y >= uint8(len(world) - 1) ||
		world[pos.Y][pos.X] == '1' {
		return false
	}
	return true
}

func checkPos(world [][]byte, pos Pos2D, move Move2D) bool {
	nextPos := Pos2D{uint8(int8(pos.X) + move.X), uint8(int8(pos.Y) + move.Y)}
	if isPosOk(world, nextPos) == false {
		return false
	}
	return true
}

func updatePos(world [][]byte, pos Pos2D, move Move2D) {
	nextPos := Pos2D{uint8(int8(pos.X) + move.X), uint8(int8(pos.Y) + move.Y)}
	world[pos.Y][pos.X] = '0'
	world[nextPos.Y][nextPos.X] = 'S'
}

func printPath(world [][]byte, move []byte) {
	for _, m := range move {
		fmt.Print("\x1b[2J")
		fmt.Print("\x1b[1;1H")
		fmt.Print(printMap(world))
		time.Sleep(300 * time.Millisecond)
		player := findPlayer(world)
		switch m {
		case 'U':
			updatePos(world, player, Move2D{0, -1})
		case 'D':
			updatePos(world, player, Move2D{0, 1})
		case 'L':
			updatePos(world, player, Move2D{-1, 0})
		case 'R':
			updatePos(world, player, Move2D{1, 0})
		}
	}
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[1;1H")
	fmt.Print(printMap(world))
}

func getNextMoves(lastGrid [][]byte, moves []byte) (nextMoves [][]byte) {
	player := findPlayer(lastGrid)
	for _, value := range "UDLR" {
		switch value {
		case 'U':
			if checkPos(lastGrid, player, Move2D{0, -1}) {
				nextMoves = append(nextMoves, append(moves, byte(value)))
			}
		case 'D':
			if checkPos(lastGrid, player, Move2D{0, 1}) {
				nextMoves = append(nextMoves, append(moves, byte(value)))
			}
		case 'L':
			if checkPos(lastGrid, player, Move2D{-1, 0}) {
				nextMoves = append(nextMoves, append(moves, byte(value)))
			}
		case 'R':
			if checkPos(lastGrid, player, Move2D{1, 0}) {
				nextMoves = append(nextMoves, append(moves, byte(value)))
			}
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
