package main

import "testing"

func TestSolveCheckmate_SimpleRun(t *testing.T) {
	board := Board{Melds: []Meld{
		{Tiles: []Tile{R1, R2, R3}},
	}}
	hand := Hand{Tiles: []Tile{B5, B6, B7}}

	hasCheckmate, solution := SolveCheckmate(board, hand)
	if !hasCheckmate {
		t.Error("Expected checkmate, but got none")
	}
	if len(solution) != 2 {
		t.Errorf("Expected 2 melds in solution, got %d", len(solution))
	}
}

func TestSolveCheckmate_WithJoker(t *testing.T) {
	board := Board{Melds: []Meld{
		{Tiles: []Tile{R1, R2, R3}},
		{Tiles: []Tile{R7, B7, Y7}},
	}}
	hand := Hand{Tiles: []Tile{B5, B6, B7, JK}}

	hasCheckmate, solution := SolveCheckmate(board, hand)
	if !hasCheckmate {
		t.Error("Expected checkmate, but got none")
	}
	if len(solution) != 3 {
		t.Errorf("Expected 3 melds in solution, got %d", len(solution))
	}
}

func TestSolveCheckmate_NoSolution(t *testing.T) {
	board := Board{Melds: []Meld{
		{Tiles: []Tile{R1, R2, R3}},
	}}
	hand := Hand{Tiles: []Tile{B5, B6, Y7}}

	hasCheckmate, _ := SolveCheckmate(board, hand)
	if hasCheckmate {
		t.Error("Expected no checkmate, but got one")
	}
}

func TestSolveCheckmate_EmptyHand(t *testing.T) {
	board := Board{Melds: []Meld{
		{Tiles: []Tile{R1, R2, R3}},
	}}
	hand := Hand{Tiles: []Tile{}}

	hasCheckmate, _ := SolveCheckmate(board, hand)
	if !hasCheckmate {
		t.Error("Expected checkmate with empty hand, but got none")
	}
}

func TestSolveCheckmate_GroupOnly(t *testing.T) {
	board := Board{Melds: []Meld{}}
	hand := Hand{Tiles: []Tile{R7, B7, Y7}}

	hasCheckmate, solution := SolveCheckmate(board, hand)
	if !hasCheckmate {
		t.Error("Expected checkmate, but got none")
	}
	if len(solution) != 1 {
		t.Errorf("Expected 1 meld in solution, got %d", len(solution))
	}
}

func TestSolveCheckmate_FourColorGroup(t *testing.T) {
	board := Board{Melds: []Meld{}}
	hand := Hand{Tiles: []Tile{R7, B7, Y7, K7}}

	hasCheckmate, solution := SolveCheckmate(board, hand)
	if !hasCheckmate {
		t.Error("Expected checkmate, but got none")
	}
	if len(solution) != 1 {
		t.Errorf("Expected 1 meld in solution, got %d", len(solution))
	}
}

func TestSolveCheckmate_JokerFillsGap(t *testing.T) {
	board := Board{Melds: []Meld{}}
	hand := Hand{Tiles: []Tile{R1, R3, JK}}

	hasCheckmate, solution := SolveCheckmate(board, hand)
	if !hasCheckmate {
		t.Error("Expected checkmate with joker filling gap, but got none")
	}
	if len(solution) != 1 {
		t.Errorf("Expected 1 meld in solution, got %d", len(solution))
	}
}

func TestSolveCheckmate_ComplexRearrangement(t *testing.T) {
	board := Board{Melds: []Meld{
		{Tiles: []Tile{R1, R2, R3}},
		{Tiles: []Tile{B1, B2, B3}},
	}}
	hand := Hand{Tiles: []Tile{Y1, Y2, Y3}}

	hasCheckmate, solution := SolveCheckmate(board, hand)
	if !hasCheckmate {
		t.Error("Expected checkmate, but got none")
	}
	if len(solution) != 3 {
		t.Errorf("Expected 3 melds in solution, got %d", len(solution))
	}
}
