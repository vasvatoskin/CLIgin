package shared

type FScreen struct {
	Pixels []Pixel
}

type GameObj struct {
	ID     uint64
	Pixels []Pixel
}

type Pixel struct {
	Coordinate Point
	Texture    rune
}
