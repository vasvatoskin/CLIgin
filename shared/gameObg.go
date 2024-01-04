package shared

type FScreen struct {
	Pixels []Pixel `json:"pixels"`
}

type GameObj struct {
	ID     uint64  `json:"id"`
	Pixels []Pixel `json:"pixels"`
}

type Pixel struct {
	Point
	Texture rune `json:"texture"`
}
