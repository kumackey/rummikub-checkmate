package main

import "fmt"

// Color はタイルの色を表す
type Color int

const (
	Red Color = iota
	Blue
	Yellow
	Black
)

func (c Color) String() string {
	switch c {
	case Red:
		return "R"
	case Blue:
		return "B"
	case Yellow:
		return "Y"
	case Black:
		return "K"
	default:
		return "?"
	}
}

// ColorCode はANSIカラーコードを返す
func (c Color) ColorCode() string {
	switch c {
	case Red:
		return "\033[31m" // 赤
	case Blue:
		return "\033[34m" // 青
	case Yellow:
		return "\033[33m" // 黄
	case Black:
		return "\033[90m" // グレー
	default:
		return "\033[0m"
	}
}

const resetColor = "\033[0m"

// Tile はラミィキューブのタイルを表す
type Tile struct {
	Number  int   // 1-13 (ジョーカーの場合は0)
	Color   Color // タイルの色
	IsJoker bool  // ジョーカーかどうか
}

// NewTile は通常のタイルを作成する
func NewTile(number int, color Color) Tile {
	return Tile{
		Number:  number,
		Color:   color,
		IsJoker: false,
	}
}

// NewJoker はジョーカータイルを作成する
func NewJoker() Tile {
	return Tile{
		Number:  0,
		Color:   0,
		IsJoker: true,
	}
}

func (t Tile) String() string {
	if t.IsJoker {
		return "\033[35mJK" + resetColor // 紫
	}
	return fmt.Sprintf("%s%s%d%s", t.Color.ColorCode(), t.Color, t.Number, resetColor)
}
