package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

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
	status gameStatus

	snakeX, snakeY float64
}

type snakepart struct {
	i            *ebiten.Image
	x, y, px, py float64
	d            direction
}

var snake = make([]*snakepart, 5, 1000)

const snakesize = 4

func init() {
	var fx float64 = 100
	var fy float64 = 100
	imag := loadImage()
	for i := range snake {
		var s snakepart
		s.i = imag
		s.x = fx - snakesize*float64(i)
		s.y = fy
		s.d = right
		snake[i] = &s
	}
}

func loadImage() (eimage *ebiten.Image) {
	eimage, _ = ebiten.NewImage(snakesize, snakesize, ebiten.FilterDefault)
	eimage.Fill(color.White)
	return
}

var lastMoved = time.Now()
var lastInput direction = -1
var speed = time.Millisecond * 100

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	m := move()
	if time.Now().After(lastMoved.Add(speed)) {
		lastMoved = time.Now()
		if m != -1 {
			if (m == right && snake[0].d != left) ||
				(m == left && snake[0].d != right) ||
				(m == up && snake[0].d != down) ||
				(m == down && snake[0].d != up) {
				snake[0].d = m
			}
		}

		moveSnake()
	}
	drawSnake(screen)
	return nil
}

func moveSnake() {
	n := snake[len(snake)-1]
	n.d = snake[0].d
	switch n.d {
	case up:
		n.x = snake[0].x
		n.y = snake[0].y - snakesize
	case down:
		n.x = snake[0].x
		n.y = snake[0].y + snakesize
	case left:
		n.x = snake[0].x - snakesize
		n.y = snake[0].y
	case right:
		n.x = snake[0].x + snakesize
		n.y = snake[0].y
	}

	snake = append([]*snakepart{n}, snake[:len(snake)-1]...)

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
		lastInput = up
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		lastInput = left
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		lastInput = down
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		lastInput = right
	}
	return lastInput
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
