package main

func Index[T comparable](slice []T, toFind T) int {
	for i, v := range slice {
		if v == toFind {
			return i
		}
	}
	return -1
}

func posAlreadySeen(slice []Node, toFind Pos2D) int {
	for index, v := range slice {
		if v.pos == toFind {
			return index
		}
	}
	return -1
}

func DeepCopyAndAdd[T any](slice []T, elems ...T) []T{
	newSlice := make([]T, len(slice) + len(elems))
	copy(newSlice, slice)
	newSlice = append(newSlice, elems...)
	return newSlice
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
