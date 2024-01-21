package shared

import "github.com/gdamore/tcell/v2"

type FScreen struct {
	Pixels []Pixel `json:"pixels"`
}

type GameObj struct {
	Pixels []Pixel `json:"pixels"`
}

type Pixel struct {
	Point
	Symbol rune        `json:"texture"`
	Style  tcell.Style `json:"style"`
}
