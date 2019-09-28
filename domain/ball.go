package domain

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

// Ball the 2d moving object
type Ball struct {
	x, y   float64
	dx, dy float64

	r       float64
	swidth  int
	sheight int

	speed float64
	img   *ebiten.Image
	op    *ebiten.DrawImageOptions
}

// NewBall creates a new ball instance
func NewBall(x, y float64, swidth, sheight int, img *ebiten.Image) *Ball {

	w, h := img.Size()
	return &Ball{
		x:       x,
		y:       y,
		dx:      rand.Float64() - rand.Float64(),
		dy:      rand.Float64() - rand.Float64(),
		r:       math.Sqrt(float64(w*w/3 + h*h/3)),
		swidth:  swidth,
		sheight: sheight,
		speed:   5,
		img:     img,
		op:      &ebiten.DrawImageOptions{},
	}
}

// Move moves the ball by units
func (ball *Ball) Move() {

	// if ball.dx or ball.dy changes, give new value
	if ball.dx == 0 && ball.dy == 0 {
		return
	}
	if ball.x < 0 || ball.x > float64(ball.swidth)-ball.r {
		ball.dx *= -1
	}
	if ball.y < 0 || ball.y > float64(ball.sheight)-ball.r {
		ball.dy *= -1
	}
	ball.x += ball.dx / (math.Sqrt(ball.dx*ball.dx + ball.dy*ball.dy)) * ball.speed
	ball.y += ball.dy / (math.Sqrt(ball.dx*ball.dx + ball.dy*ball.dy)) * ball.speed
}

// Draw draws the screen
//	IsDrawingSkipped needs to before this draw
func (ball *Ball) Draw(screen *ebiten.Image) {
	ball.op = &ebiten.DrawImageOptions{}
	ball.op.GeoM.Translate(ball.x, ball.y)
	screen.DrawImage(ball.img, ball.op)
}

// Poser struct must give us their location
type Poser interface {
	Pos() (x, y, w, h int)
}

// Bounce bounce on the colliders
func (ball *Ball) Bounce(p ...Poser) {

	var x, y, w, h int
	var px, py, pw, ph float64
	for _, v := range p {
		x, y, w, h = v.Pos()
		px, py, pw, ph = float64(x), float64(y), float64(w), float64(h)

		// if ball hits all sides of the pedal
		// x <= X+R && X <= x+w
		if px <= ball.x+ball.r && ball.x <= px+pw {
			if py <= ball.y+ball.r && ball.y <= py+ph {
				// which ball's wall hit the pedals
				// ball's wall is 2 points
				ball.dy *= -1
				ball.dx *= -1
			}
		}
	}
}
