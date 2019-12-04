package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}, color.RGBA{0xfd, 0xf1, 0x00, 0xff}}

const (
	blackIndex = iota
	redIndex
	blueIndex
	greenIndex
	yelloIndex
)

func main() {
	rand.Seed(time.Now().Unix())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			px := size + int(x*size+0.5)
			py := size + int(y*size+0.5)
			img.SetColorIndex(px, py, getColor(px, py))

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func getColor(x, y int) uint8 {
	if x < 100 && y < 100 {
		return 1
	}
	if x < 100 && y >= 100 {
		return 2
	}
	if x >= 100 && y < 100 {
		return 3
	}
	return 4
}
