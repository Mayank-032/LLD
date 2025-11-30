package lsp

// LSP: Liskov Substitution Principle

// LSP - Violated
type RectangleViolation struct {
	Width  float64
	Height float64
}

func (r *RectangleViolation) SetWidth(width float64) {
	r.Width = width
}

func (r *RectangleViolation) SetHeight(height float64) {
	r.Height = height
}

func (r *RectangleViolation) Area() float64 {
	return r.Width * r.Height
}

type SquareViolation struct {
	RectangleViolation
}

func (s *SquareViolation) SetWidth(width float64) {
	s.Width = width
	s.Height = width
}

func (s *SquareViolation) SetHeight(height float64) {
	s.Height = height
	s.Width = height
}

// LSP - Followed
type IShape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}