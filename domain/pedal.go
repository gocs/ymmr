package domain

import (
	"github.com/hajimehoshi/ebiten"
)

// Pedal ..
type Pedal struct {
	x, y    int

	w, h    int
	swidth  int
	sheight int

	score   int
	img     *ebiten.Image
	op      *ebiten.DrawImageOptions
}

// NewPedal initializes
func NewPedal(x, y, swidth, sheight int, img *ebiten.Image) *Pedal {

	w, h := img.Size()
	return &Pedal{
		x:       x,
		y:       y,
		w:       w,
		h:       h,
		swidth:  swidth,
		sheight: sheight,
		img:     img,
	}
}

// Scores use this if pedal scored a point
func (pedal *Pedal) Scores() {
	pedal.score++
}

// Move moves the pedal by units
func (pedal *Pedal) Move(units int, left, right ebiten.Key) {
	w, _ := pedal.img.Size()
	if ebiten.IsKeyPressed(left) && pedal.x > 0 {
		pedal.x -= units
	} else if ebiten.IsKeyPressed(right) && pedal.x < pedal.swidth-w {
		pedal.x += units
	}
}

// Draw draws the screen
//	IsDrawingSkipped needs to before this draw
func (pedal *Pedal) Draw(screen *ebiten.Image) {
	pedal.op = &ebiten.DrawImageOptions{}
	pedal.op.GeoM.Translate(float64(pedal.x), float64(pedal.y))
	screen.DrawImage(pedal.img, pedal.op)
}

// Pos returns the position perimeter idk whats called
func (pedal *Pedal) Pos() (int, int, int, int) {
	return pedal.x, pedal.y, pedal.w, pedal.h
}
