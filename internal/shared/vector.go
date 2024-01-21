package shared

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Vector struct {
	Dx int `json:"dx"`
	Dy int `json:"dy"`
}

func (v *Vector) SetDx(newDx int) {
	v.Dx = newDx
	v.Dy = 0
}

func (v *Vector) SetDy(newDy int) {
	v.Dy = newDy
	v.Dx = 0
}
