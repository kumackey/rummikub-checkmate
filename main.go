package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON入力用の構造体
type TileJSON struct {
	Number int    `json:"number,omitempty"`
	Color  string `json:"color,omitempty"`
	Joker  bool   `json:"joker,omitempty"`
}

type GameStateJSON struct {
	Board [][]TileJSON `json:"board"`
	Hand  []TileJSON   `json:"hand"`
}

func parseColor(s string) Color {
	switch s {
	case "R":
		return Red
	case "B":
		return Blue
	case "Y":
		return Yellow
	case "K":
		return Black
	default:
		return Red
	}
}

func (tj TileJSON) ToTile() Tile {
	if tj.Joker {
		return NewJoker()
	}
	return NewTile(tj.Number, parseColor(tj.Color))
}

func LoadGameState(filename string) (*GameState, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var gsj GameStateJSON
	if err := json.Unmarshal(data, &gsj); err != nil {
		return nil, err
	}

	gs := &GameState{}

	// Board変換
	for _, meldJSON := range gsj.Board {
		var tiles []Tile
		for _, tj := range meldJSON {
			tiles = append(tiles, tj.ToTile())
		}
		gs.Board.Melds = append(gs.Board.Melds, Meld{Tiles: tiles})
	}

	// Hand変換
	for _, tj := range gsj.Hand {
		gs.Hand.Tiles = append(gs.Hand.Tiles, tj.ToTile())
	}

	return gs, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rummikub-checkmate <json-file>")
		os.Exit(1)
	}

	gs, err := LoadGameState(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(gs)

	// バリデーション表示
	fmt.Println("\nValidation:")
	for i, meld := range gs.Board.Melds {
		valid := "✅"
		if !meld.IsValid() {
			valid = "❌"
		}
		fmt.Printf("  Meld %d: %s %s\n", i+1, meld.String(), valid)
	}
}
