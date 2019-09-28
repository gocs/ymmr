package main

import (
	"fmt"
	"log"

	"github.com/gocs/ymmr/domain"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var leftX int
var rightX int

var ballX int
var ballY int

var leftPedal *domain.Pedal
var rightPedal *domain.Pedal
var ball *domain.Ball

func init() {
	leftX = screenWidth/2 - 40
	rightX = screenWidth/2 - 40

	ballX = screenWidth/2 - 5
	ballY = screenHeight/2 - 5

	var err error
	leftPedalImg, _, err := ebitenutil.NewImageFromFile("pedal.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	leftPedal = domain.NewPedal(leftX, 170, screenWidth, screenHeight, leftPedalImg)

	rightPedalImg, _, err := ebitenutil.NewImageFromFile("pedal.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	rightPedal = domain.NewPedal(rightX, 60, screenWidth, screenHeight, rightPedalImg)

	ballImg, _, err := ebitenutil.NewImageFromFile("ball.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	ball = domain.NewBall(float64(ballX), float64(ballY), screenWidth, screenHeight, ballImg)
}

func update(screen *ebiten.Image) (err error) {
	// Does pedal needs to know ball's location or vice versa like here?
	leftPedal.Move(3, ebiten.KeyA, ebiten.KeyD)
	rightPedal.Move(3, ebiten.KeyLeft, ebiten.KeyRight)
	ball.Bounce(leftPedal, rightPedal)
	ball.Move()

	// ------------------------------------------------------------------------
	if ebiten.IsDrawingSkipped() { ////////////////////////////////////////////
		return nil ////////////////////////////////////////////////////////////
	} /////////////////////////////////////////////////////////////////////////
	// ------------------------------------------------------------------------

	rightPedal.Draw(screen)
	leftPedal.Draw(screen)
	ball.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprint("FPS:", ebiten.FPS))

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
