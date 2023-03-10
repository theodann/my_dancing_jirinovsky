package main

import (
	//"fmt"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Game struct{
	window *window
}

func New() *Game {
	return &Game{
		window: newWindow(),
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.window.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.window.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1080, 720
}

func main() {
	g := New()
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("My dancing president")
	ebiten.RunGame(g)
}
