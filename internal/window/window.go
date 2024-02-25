package window

import "github.com/vasvatoskin/CLIgin/internal/shared"

type Window struct {
	TopLeft shared.Point
	Actual  Size
	WindowStyle
}

func New(start shared.Point, style WindowStyle, screenSize Size) *Window {
	var actual Size
	actual.CalculateActual(screenSize, style)

	return &Window{
		TopLeft:     start,
		Actual:      actual,
		WindowStyle: style,
	}
}
