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

type TileID int
type TileNumber int

// Tile はラミィキューブのタイルを表す
type Tile struct {
	Number  TileNumber // 1-13 (ジョーカーの場合は0)
	Color   Color      // タイルの色
	IsJoker bool       // ジョーカーかどうか
	ID      TileID
}

// NewTile は通常のタイルを作成する
func NewTile(color Color, number TileNumber) Tile {
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

var (
	R1  = NewTile(Red, 1)
	R2  = NewTile(Red, 2)
	R3  = NewTile(Red, 3)
	R4  = NewTile(Red, 4)
	R5  = NewTile(Red, 5)
	R6  = NewTile(Red, 6)
	R7  = NewTile(Red, 7)
	R8  = NewTile(Red, 8)
	R9  = NewTile(Red, 9)
	R10 = NewTile(Red, 10)
	R11 = NewTile(Red, 11)
	R12 = NewTile(Red, 12)
	R13 = NewTile(Red, 13)

	B1  = NewTile(Blue, 1)
	B2  = NewTile(Blue, 2)
	B3  = NewTile(Blue, 3)
	B4  = NewTile(Blue, 4)
	B5  = NewTile(Blue, 5)
	B6  = NewTile(Blue, 6)
	B7  = NewTile(Blue, 7)
	B8  = NewTile(Blue, 8)
	B9  = NewTile(Blue, 9)
	B10 = NewTile(Blue, 10)
	B11 = NewTile(Blue, 11)
	B12 = NewTile(Blue, 12)
	B13 = NewTile(Blue, 13)

	Y1  = NewTile(Yellow, 1)
	Y2  = NewTile(Yellow, 2)
	Y3  = NewTile(Yellow, 3)
	Y4  = NewTile(Yellow, 4)
	Y5  = NewTile(Yellow, 5)
	Y6  = NewTile(Yellow, 6)
	Y7  = NewTile(Yellow, 7)
	Y8  = NewTile(Yellow, 8)
	Y9  = NewTile(Yellow, 9)
	Y10 = NewTile(Yellow, 10)
	Y11 = NewTile(Yellow, 11)
	Y12 = NewTile(Yellow, 12)
	Y13 = NewTile(Yellow, 13)

	K1  = NewTile(Black, 1)
	K2  = NewTile(Black, 2)
	K3  = NewTile(Black, 3)
	K4  = NewTile(Black, 4)
	K5  = NewTile(Black, 5)
	K6  = NewTile(Black, 6)
	K7  = NewTile(Black, 7)
	K8  = NewTile(Black, 8)
	K9  = NewTile(Black, 9)
	K10 = NewTile(Black, 10)
	K11 = NewTile(Black, 11)
	K12 = NewTile(Black, 12)
	K13 = NewTile(Black, 13)

	JK = NewJoker()
)
