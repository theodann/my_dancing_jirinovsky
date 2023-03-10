package main

import (
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type window struct{
	jirik Jirik
	triangle Triangle
}

type Jirik struct {
	x, y *image.Point
}

type Triangle struct {
	way int
	point *image.Point
}

func newWindow() *window {
	p1 := &image.Point{100, 100}
	p2 := &image.Point{200, 200}
	return &window{
		jirik: p1
		triangle: 1, p2
	}
}

func (w *window) Update() {
}

func (w *window) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{239, 207, 227, 1})
	drawJirik(screen)
	drawTriangle(screen)
}