package go_utils

import "image"

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	UpLeft
	UpRight
	DownLeft
	DownRight
	None
)

func GetNextDir(p image.Point, d Direction) (image.Point, bool) {
	switch d {
	case Up:
		return image.Pt(p.X, p.Y-1), false
	case Down:
		return image.Pt(p.X, p.Y+1), false
	case Left:
		return image.Pt(p.X-1, p.Y), false
	case Right:
		return image.Pt(p.X+1, p.Y), false
	case UpLeft:
		return image.Pt(p.X-1, p.Y-1), false
	case UpRight:
		return image.Pt(p.X+1, p.Y-1), false
	case DownLeft:
		return image.Pt(p.X-1, p.Y+1), false
	case DownRight:
		return image.Pt(p.X+1, p.Y+1), false
	case None:
		return image.Pt(p.X, p.Y), false
	default:
		return p, true
	}
}
