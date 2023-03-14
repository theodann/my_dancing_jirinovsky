package main

import (
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type window struct {
	jirik Jirik
	triangles Triangles
}

type Jirik struct {
	pic *ebiten.Image
}

type Triangles struct {
	pic []ebiten.Image
}

type O struct {
	pic *ebiten.Image
	x, y int
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

func drawJirik(screen *ebiten.Image) {
	screen.
}

func (w *window) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{239, 207, 227, 1})
	drawJirik(screen)
	drawTriangle(screen)
}