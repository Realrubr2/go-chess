package Pieces

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

)
// represents the types of pieces
type PieceType int

const(
	Pawn PieceType = 0
	Queen
	Rook
	Bishop
	Knight
	King
)

// represents the piece
type Piece struct {
	Type PieceType
	color color.Color
	X, Y int
}

// represents the board
type Board struct {
	Pieces [8][8]*Piece
	ScreenWidth, ScreenHeight int
	TileSize int
}
// init stating point of the board
func (b *Board) Init() {
	b.Pieces[0][0] = &Piece{Rook, color.White, 0, 0}
	b.Pieces[0][1] = &Piece{Knight, color.White, 1, 0}
	b.Pieces[0][2] = &Piece{Bishop, color.White, 2, 0}
	b.Pieces[0][3] = &Piece{Queen, color.White, 3, 0}
	b.Pieces[0][4] = &Piece{King, color.White, 4, 0}
	b.Pieces[0][5] = &Piece{Bishop, color.White, 5, 0}
	b.Pieces[0][6] = &Piece{Knight, color.White, 6, 0}
	b.Pieces[0][7] = &Piece{Rook, color.White, 7, 0}

	b.Pieces[7][0] = &Piece{Rook, color.Black, 0, 7}
	b.Pieces[7][1] = &Piece{Knight, color.Black, 1, 7}
	b.Pieces[7][2] = &Piece{Bishop, color.Black, 2, 7}
	b.Pieces[7][3] = &Piece{Queen, color.Black, 3, 7}
	b.Pieces[7][4] = &Piece{King, color.Black, 4, 7}
	b.Pieces[7][5] = &Piece{Bishop, color.Black, 5, 7}
	b.Pieces[7][6] = &Piece{Knight, color.Black, 6, 7}
	b.Pieces[7][7] = &Piece{Rook, color.Black, 7, 7}

	for i := 0; i < 8; i++ {
		b.Pieces[1][i] = &Piece{Pawn, color.White, i, 1}
		b.Pieces[6][i] = &Piece{Pawn, color.Black, i, 6}
	}
}




func (b *Board)DrawChessBoard(screen *ebiten.Image) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			// Alternate colors for the chessboard pattern
			if (x+y)%2 == 0 {
				vector.DrawFilledRect(screen, float32(x*b.TileSize), float32(y*b.TileSize), float32(b.TileSize), float32(b.TileSize), color.RGBA{255, 255, 255, 255}, false) // White tile
			} else {
				vector.DrawFilledRect(screen, float32(x*b.TileSize), float32(y*b.TileSize), float32(b.TileSize), float32(b.TileSize), color.RGBA{0, 0, 0, 255}, false) // Black tile
			}
		}
	}
}

func (b *Board)DrawPieces(screen *ebiten.Image) {

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if piece := b.Pieces[y][x]; piece != nil {
				// Draw a simple circle as a placeholder for each piece
				vector.DrawFilledRect(screen, float32(x*b.TileSize+10), float32(y*b.TileSize+10), float32(b.TileSize)-20, float32(b.TileSize)-20, piece.color, false)
			}
		}
	}
}