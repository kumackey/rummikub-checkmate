package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// JSON入力用の構造体
type GameStateJSON struct {
	Board [][]string `json:"board"`
	Hand  []string   `json:"hand"`
}

var tileMap = map[string]Tile{
	"R1": R1, "R2": R2, "R3": R3, "R4": R4, "R5": R5, "R6": R6, "R7": R7,
	"R8": R8, "R9": R9, "R10": R10, "R11": R11, "R12": R12, "R13": R13,

	"B1": B1, "B2": B2, "B3": B3, "B4": B4, "B5": B5, "B6": B6, "B7": B7,
	"B8": B8, "B9": B9, "B10": B10, "B11": B11, "B12": B12, "B13": B13,

	"Y1": Y1, "Y2": Y2, "Y3": Y3, "Y4": Y4, "Y5": Y5, "Y6": Y6, "Y7": Y7,
	"Y8": Y8, "Y9": Y9, "Y10": Y10, "Y11": Y11, "Y12": Y12, "Y13": Y13,

	"K1": K1, "K2": K2, "K3": K3, "K4": K4, "K5": K5, "K6": K6, "K7": K7,
	"K8": K8, "K9": K9, "K10": K10, "K11": K11, "K12": K12, "K13": K13,

	"JK": JK,
}

func parseTile(s string) (Tile, error) {
	s = strings.TrimSpace(s)
	if t, ok := tileMap[s]; ok {
		return t, nil
	}
	return Tile{}, fmt.Errorf("invalid tile: %s", s)
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
	for _, meldStrings := range gsj.Board {
		var tiles []Tile
		for _, s := range meldStrings {
			tile, err := parseTile(s)
			if err != nil {
				return nil, err
			}
			tiles = append(tiles, tile)
		}
		gs.Board.Melds = append(gs.Board.Melds, Meld(tiles))
	}

	// Hand変換
	for _, s := range gsj.Hand {
		tile, err := parseTile(s)
		if err != nil {
			return nil, err
		}
		gs.Hand.Tiles = append(gs.Hand.Tiles, tile)
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

	// 詰み判定
	fmt.Println("\nCheckmate Analysis:")
	hasCheckmate, solution := SolveCheckmate(gs.Board, gs.Hand)
	if hasCheckmate {
		fmt.Println("  Result: ✅ 詰みあり（手札を出し切れる）")
		fmt.Println("\n  Solution:")
		for i, meld := range solution {
			fmt.Printf("    %d: %s\n", i+1, meld.String())
		}
	} else {
		fmt.Println("  Result: ❌ 詰みなし（手札を出し切れない）")
	}
}
