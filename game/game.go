package game

import (
	"github.com/Realrubr2/go-chess/pieces"
	"github.com/hajimehoshi/ebiten/v2"
)
type Game struct{
	Board Pieces.Board
}

func(g *Game)Layout(outsideWidth, outsideHeight int)(screenWidth, screenHeight int){
	return g.Board.ScreenWidth, g.Board.ScreenHeight
}


func (g *Game) Update()error {
	// here we need to update the game

	return nil
}


func (g *Game)Draw(screen *ebiten.Image){
	// here we need to draw the game
	g.Board.DrawChessBoard(screen)
	g.Board.DrawPieces(screen)

}

