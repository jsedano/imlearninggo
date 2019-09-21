package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"

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

var snake = make([]*snakepart, 10, 1000)

const snakesize = 8

var gameControl game

var mplusNormalFont font.Face

var fontx int
var fonty int

func init() {

	gameControl.status = notStarted

	var fx float64 = 100
	var fy float64 = 200
	imag := loadImage()
	for i := range snake {
		var s snakepart
		s.i = imag
		s.x = fx - snakesize*float64(i)
		s.y = fy
		s.d = right
		snake[i] = &s
	}

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	bound, _ := font.BoundString(mplusNormalFont, "Press spacebar to play")
	w := (bound.Max.X - bound.Min.X).Ceil()
	h := (bound.Max.Y - bound.Min.Y).Ceil()
	fontx = (320 - w) / 2
	fonty = (240 - h) / 2
}

func loadImage() (eimage *ebiten.Image) {
	eimage, _ = ebiten.NewImage(snakesize, snakesize, ebiten.FilterDefault)
	eimage.Fill(color.White)
	return
}

var lastMoved = time.Now()
var lastInput direction = -1
var speed = time.Millisecond * 250

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))

	switch gameControl.status {
	case notStarted:

		text.Draw(screen, "Press spacebar to play", mplusNormalFont, fontx, fonty, color.White)
		drawSnake(screen)
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			gameControl.status = playing
		}

	case playing:
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
	case over:
	}

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
	if err := ebiten.Run(update, 320, 240, 2, "danger noodle"); err != nil {
		log.Fatal(err)
	}
}
