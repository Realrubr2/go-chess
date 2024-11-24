package game

import (
	"image/color"
	"github.com/Realrubr2/go-chess/movement"
	"github.com/Realrubr2/go-chess/pieces"
	"github.com/hajimehoshi/ebiten/v2"
    "fmt"
)
type Game struct{
	Board Pieces.Board
	CurrentMove *movement.Move
}

func(g *Game)Layout(outsideWidth, outsideHeight int)(screenWidth, screenHeight int){
	return g.Board.ScreenWidth, g.Board.ScreenHeight
}

//  !TODO the update function needs to be refactored atrying to see why my piece is filled in this function but in move piece is nil!
func (g *Game) Update() error {
    x, y := ebiten.CursorPosition()
    boardX, boardY := x/g.Board.TileSize, y/g.Board.TileSize

    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
        if g.CurrentMove == nil {
            if boardX >= 0 && boardX < 8 && boardY >= 0 && boardY < 8 {
                piece := g.Board.Pieces[boardY][boardX]
                if piece != nil {
                    g.CurrentMove = &movement.Move{
                        Board: &g.Board,
                        Piece: piece,
                        FromX: boardX,
                        FromY: boardY,
                        ToX:   boardX,
                        ToY:   boardY,
                    }
                }
            }
        } else {
            g.CurrentMove.ToX = boardX
            g.CurrentMove.ToY = boardY
        }
    } else if g.CurrentMove != nil {
        if boardX >= 0 && boardX < 8 && boardY >= 0 && boardY < 8 {
            g.CurrentMove.ToX = boardX
            g.CurrentMove.ToY = boardY

            if g.CurrentMove.MovePiece() {
                fmt.Println("Move successful!")
                g.CurrentMove = nil
            } else {
                fmt.Println("Invalid move")
                g.CurrentMove.ToX = g.CurrentMove.FromX
                g.CurrentMove.ToY = g.CurrentMove.FromY
                g.CurrentMove = nil
            }
        }
    }
    if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
        g.CurrentMove = nil
    }

    return nil
}



// draw on the screen with the selected piece getting highligted and lifted up[ to show that it is selected]
func (g *Game) Draw(screen *ebiten.Image) {
    g.Board.DrawChessBoard(screen)
    g.Board.DrawPieces(screen)


    if g.CurrentMove != nil {
        x, y := g.CurrentMove.FromX, g.CurrentMove.FromY


        highlight := ebiten.NewImage(g.Board.TileSize, g.Board.TileSize)
        highlight.Fill(color.RGBA{255, 255, 0, 100}) 
        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(float64(x*g.Board.TileSize), float64(y*g.Board.TileSize))
        screen.DrawImage(highlight, op)
        piece := g.CurrentMove.Piece
        if piece != nil {
            op := &ebiten.DrawImageOptions{}
            mouseX, mouseY := ebiten.CursorPosition()
            op.GeoM.Translate(float64(mouseX-g.Board.TileSize/2), float64(mouseY-g.Board.TileSize/2))
            screen.DrawImage(piece.Image, op)
        }
    }
}
