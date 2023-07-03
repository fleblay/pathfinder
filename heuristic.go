package main

import "math"

//HEURISTIC : Estimated cost form current to Goal
func currentToGoal(pos, startPos, goalPos Pos2D) int {
	return int(math.Abs(float64(pos.X-goalPos.X)) + math.Abs(float64(pos.Y-goalPos.Y)))
}

func startToCurrent(pos, startPos, goalPos Pos2D) int {
	return int(math.Abs(float64(pos.X-startPos.X)) + math.Abs(float64(pos.Y-startPos.Y)))
}

func FIFO() func(pos, startPos, goalPos Pos2D) int {
	count := 0
	return func(pos, startPos, goalPos Pos2D) int {
		count++
		return count
	}
}

func ASTAR(pos, startPos, goalPos Pos2D) int {
	return startToCurrent(pos, startPos, goalPos) + currentToGoal(pos, startPos, goalPos)
}
