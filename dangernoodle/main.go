package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

type gameStatus int

const (
	over gameStatus = iota
	playing
	notStarted
)

type game struct {
	status         gameStatus
	snakeX, snakeY float64
}

type snakepart struct {
	i            *ebiten.Image
	x, y, px, py float64
	d            direction
}

var snake = make([]*snakepart, 200, 1000)

func init() {
	var fx float64 = 100
	var fy float64 = 100
	imag := loadImage()
	for i := range snake {
		var s snakepart
		s.i = imag
		s.x = fx - 2*float64(i)
		s.y = fy
		s.d = right
		snake[i] = &s
	}
}

func loadImage() (eimage *ebiten.Image) {
	eimage, _ = ebiten.NewImage(2, 2, ebiten.FilterDefault)
	eimage.Fill(color.White)
	return
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))

	m := move()
	if m != -1 {
		if (m == right && snake[0].d != left) ||
			(m == left && snake[0].d != right) ||
			(m == up && snake[0].d != down) ||
			(m == down && snake[0].d != up) {
			snake[0].d = m
		}
	}

	moveSnake()
	drawSnake(screen)

	return nil
}

func moveSnake() {
	for i := range snake {
		switch snake[i].d {
		case up:
			snake[i].y = snake[i].y - 2
		case down:
			snake[i].y = snake[i].y + 2
		case left:
			snake[i].x = snake[i].x - 2
		case right:
			snake[i].x = snake[i].x + 2
		}
	}

	for i := len(snake) - 1; i >= 1; i-- {
		snake[i].d = snake[i-1].d
	}

}

func drawSnake(screen *ebiten.Image) {
	for i := range snake {
		drawSnakePart(snake[i].x, snake[i].y, snake[i].i, screen)
	}
}

func drawSnakePart(x float64, y float64, sp *ebiten.Image, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(sp, op)
}

func move() direction {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		return up
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		return left
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return down
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		return right
	}
	return -1
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
