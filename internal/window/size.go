package window

import "math"

type Size struct {
	Width  int
	Height int
}

func (s *Size) CalculateActual(size Size, style WindowStyle) {
	s.Width = int(math.Max(float64(style.Min.Width), float64(size.Width*style.Occupancy/100)))
	s.Height = int(math.Max(float64(style.Min.Height), float64(size.Height*style.Occupancy/100)))
}
