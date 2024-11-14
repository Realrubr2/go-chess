package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/Realrubr2/go-chess/pieces"
	"github.com/Realrubr2/go-chess/game"
)


func main() {
	board := Pieces.Board{
		ScreenWidth: 640,
		ScreenHeight: 640,
		TileSize: 80,
	}
	board.Init()

	game := game.Game{Board: board}
	ebiten.SetWindowSize(board.ScreenWidth, board.ScreenHeight)
	ebiten.SetWindowTitle("Chess")
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}