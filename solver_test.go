package main

import "testing"

func TestSolveCheckmate_SimpleRun(t *testing.T) {
	board := Board{Melds: []Meld{
		{R1, R2, R3},
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
		{R1, R2, R3},
		{R7, B7, Y7},
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
		{R1, R2, R3},
	}}
	hand := Hand{Tiles: []Tile{B5, B6, Y7}}

	hasCheckmate, _ := SolveCheckmate(board, hand)
	if hasCheckmate {
		t.Error("Expected no checkmate, but got one")
	}
}

func TestSolveCheckmate_EmptyHand(t *testing.T) {
	board := Board{Melds: []Meld{
		{R1, R2, R3},
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
		{R1, R2, R3},
		{B1, B2, B3},
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

// 盤面の組み換えが必要な複雑なケース
// 既存のメルドを崩さないと手札を出し切れない
func TestSolveCheckmate_HeavyRearrangement(t *testing.T) {
	// https://twitter.com/syura9999/status/1764145671274373470/photo/1
	board := Board{Melds: []Meld{
		{K1, K2, K3},
		{K4, Y4, B4},
		{R4, R5, R6},
		{R7, K7, B7},
		{R13, B13, Y13},
		{B10, B11, B12},
		{K10, R10, Y10},
		{Y7, Y8, Y9},	
	}}
	hand := Hand{Tiles: []Tile{B1, Y1, B13}}

	hasCheckmate, solution := SolveCheckmate(board, hand)
	if !hasCheckmate {
		t.Error("Expected checkmate with heavy rearrangement, but got none")
	}

	for i, meld := range solution {
		t.Logf("%d: %s", i+1, meld.String())
	}
}
