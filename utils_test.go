package main

import "testing"

type indexTest struct {
	arg1     []int
	arg2     int
	expected int
}

var indexTests = []indexTest{
	{[]int{}, 1, -1},
	{[]int{1, 2, 3, 4}, 1, 0},
	{[]int{1, 2, 3, 4}, 8, -1},
}

func TestIndex(t *testing.T) {
	for index, test := range indexTests {
		if output := Index(test.arg1, test.arg2); output != test.expected {
			t.Fatalf("Test %d : Output %d different from expected %d", index, output, test.expected)
		}
	}
}

type posAlreadySeenTest struct {
	arg1     []Node
	arg2     Pos2D
	expected int
}

var posAlreadySeenTests = []posAlreadySeenTest{
	{[]Node{{Pos2D{X: 3, Y: 7}, 12}}, Pos2D{3, 7}, 0},
	{[]Node{{Pos2D{X: 3, Y: 7}, 12}}, Pos2D{2, 7}, -1},
}

func TestPosAlreadySeen(t *testing.T) {
	for index, test := range posAlreadySeenTests {
		if output := posAlreadySeen(test.arg1, test.arg2); output != test.expected {
			t.Fatalf("Test %d : Output %d different from expected %d", index, output, test.expected)
		}
	}
}

func TestDeepCopyAndAdd(t *testing.T) {
	initialSlice := []int{1, 2, 3, 4, 5}
	copiedSlice := DeepCopyAndAdd(initialSlice, 6, 7, 8)
	copiedSlice[3] = -4
	if initialSlice[3] != 4 {
		t.Fatalf("Initial Slice is modified")
	}
}

func TestMax(t *testing.T) {
	if Max(12, 4) != 12 ||
		Max(3, 8) != 8 {
			t.Fatalf("Max failure")
	}
}
