package geo

type Rect struct {
	TL, BR Point
}

func NewRect(tl Point, width, height int) Rect {
	return Rect{TL: tl, BR: Point{X: tl.X + width - 1, Y: tl.Y + height - 1}}
}

// Check if intersects given rect
func (r Rect) Intersects(rect Rect) bool {
	return !(r.BR.X < rect.TL.X || r.TL.X > rect.BR.X ||
		r.TL.Y > rect.BR.Y || r.BR.Y < rect.TL.Y)
}

// Add offset to all sides of rect
func (r Rect) Add(d Distance) Rect {
	return Rect{TL: r.TL.Sub(d), BR: r.BR.Add(d)}
}

// Get width
func (r Rect) Width() int {
	return r.BR.X - r.TL.X + 1
}

// Get height
func (r Rect) Height() int {
	return r.BR.Y - r.TL.Y + 1
}
