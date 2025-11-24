package main

import "fmt"

// Meld は場に出ているタイルの組み合わせ（セットまたはラン）
type Meld struct {
	Tiles []Tile
}

// IsValidSet は同じ数字で異なる色の組み合わせかチェック
func (m *Meld) IsValidSet() bool {
	if len(m.Tiles) < 3 || len(m.Tiles) > 4 {
		return false
	}

	colors := make(map[Color]bool)
	var number TileNumber = -1

	for _, tile := range m.Tiles {
		if tile.IsJoker {
			continue
		}
		if number == -1 {
			number = tile.Number
		} else if tile.Number != number {
			return false
		}
		if colors[tile.Color] {
			return false
		}
		colors[tile.Color] = true
	}
	return true
}

// IsValidRun は同じ色で連続する数字の組み合わせかチェック
func (m *Meld) IsValidRun() bool {
	if len(m.Tiles) < 3 {
		return false
	}

	// ジョーカー以外のタイルを収集
	var nonJokers []Tile
	jokerCount := 0
	var runColor Color

	for _, tile := range m.Tiles {
		if tile.IsJoker {
			jokerCount++
		} else {
			nonJokers = append(nonJokers, tile)
			runColor = tile.Color
		}
	}

	if len(nonJokers) == 0 {
		return true // 全部ジョーカー
	}

	// 色が統一されているかチェック
	for _, tile := range nonJokers {
		if tile.Color != runColor {
			return false
		}
	}

	// 数字でソート（簡易版）
	for i := 0; i < len(nonJokers)-1; i++ {
		for j := i + 1; j < len(nonJokers); j++ {
			if nonJokers[i].Number > nonJokers[j].Number {
				nonJokers[i], nonJokers[j] = nonJokers[j], nonJokers[i]
			}
		}
	}

	// 連続性チェック（ジョーカーで埋められるか）
	gaps := 0
	for i := 0; i < len(nonJokers)-1; i++ {
		diff := int(nonJokers[i+1].Number - nonJokers[i].Number - 1)
		if diff < 0 {
			return false // 重複
		}
		gaps += diff
	}

	return gaps <= jokerCount
}

// IsValid はMeldが有効かどうかチェック
func (m *Meld) IsValid() bool {
	return m.IsValidSet() || m.IsValidRun()
}

func (m *Meld) String() string {
	result := "["
	for i, tile := range m.Tiles {
		if i > 0 {
			result += " "
		}
		result += tile.String()
	}
	return result + "]"
}

// Board は場に出ているすべてのMeld
type Board struct {
	Melds []Meld
}

func (b *Board) String() string {
	result := "Board:\n"
	for i, meld := range b.Melds {
		result += fmt.Sprintf("  %d: %s\n", i+1, meld.String())
	}
	return result
}

// Hand はプレイヤーの手札
type Hand struct {
	Tiles []Tile
}

func (h *Hand) String() string {
	result := "Hand: "
	for i, tile := range h.Tiles {
		if i > 0 {
			result += " "
		}
		result += tile.String()
	}
	return result
}

// GameState はゲームの状態を表す
type GameState struct {
	Board Board
	Hand  Hand
}

func (g *GameState) String() string {
	return g.Board.String() + g.Hand.String()
}
