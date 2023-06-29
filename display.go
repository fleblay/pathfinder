package main

import (
	"fmt"
	"time"
)

func getPrintColor(value byte) (color, updatedVal, reset string) {
	updatedVal = string(value)
	switch value {
	case 'S':
		color = "\x1b[33m"
		reset = "\x1b[0m"
	case 'E':
		color = "\x1b[32m"
		reset = "\x1b[0m"
	case '1':
		color = "\x1b[31m"
		reset = "\x1b[0m"
	case '_':
		color = "\x1b[44;5m"
		reset = "\x1b[0;25m"
		updatedVal = "0"
	default:
		color = ""
		reset = ""
	}
	return
}

func clearScreen() {
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[1;1H")
}

func printMap(world [][]byte) (res string) {
	for _, line := range world {
		for pos, value := range line {
			color, updatedVal,reset := getPrintColor(value)
			res += color + updatedVal + reset
			if pos == len(line)-1 {
				res += "\n"
			} else {
				res += " "
			}
		}
	}
	return
}

func printPath(world [][]byte, move []byte) {
	for _, m := range move {
		clearScreen()
		fmt.Print(printMap(world))
		time.Sleep(300 * time.Millisecond)
		player := findPlayer(world)
		updatePos(world, player, directions[m])
	}
	clearScreen()
	fmt.Print(printMap(world))
}
