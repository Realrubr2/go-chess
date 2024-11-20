package Pieces

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// represents the types of pieces
type PieceType int

// iota is a keyword in Go that is used to make the others increment by 1
const(
	Pawn PieceType = iota
	Queen
	Rook
	Bishop
	Knight
	King
)

// represents the piece
type Piece struct {
	Type PieceType
	Color color.Color
	X, Y int
	image *ebiten.Image
	alive bool
}

// represents the board
type Board struct {
	Pieces [8][8]*Piece
	ScreenWidth, ScreenHeight int
	TileSize int
}
// init stating point of the board
func (b *Board) Init() {
// White pieces
	b.Pieces[0][0] = &Piece{Rook, color.White, 0, 0, loadImage("assets/rook_white.png"), true}
	b.Pieces[0][1] = &Piece{Knight, color.White, 1, 0,loadImage("assets/knight_white.png"), true}
	b.Pieces[0][2] = &Piece{Bishop, color.White, 2, 0,loadImage("assets/Bishop_white.png"), true}
	b.Pieces[0][4] = &Piece{Queen, color.White, 3, 0,loadImage("assets/queen_white.png"), true}
	b.Pieces[0][3] = &Piece{King, color.White, 4, 0,loadImage("assets/king_white.png"), true}
	b.Pieces[0][5] = &Piece{Bishop, color.White, 5, 0,loadImage("assets/bishop_white.png"), true}
	b.Pieces[0][6] = &Piece{Knight, color.White, 6, 0,loadImage("assets/knight_white.png"), true}
	b.Pieces[0][7] = &Piece{Rook, color.White, 7, 0,loadImage("assets/rook_white.png"), true}
	// White pawns
	for i := 0; i < 8; i++ {
		b.Pieces[1][i] = &Piece{Pawn, color.White, i, 1, loadImage("assets/pawn_white.png"), true}
	}

	// Black pieces
	b.Pieces[7][0] = &Piece{Rook, color.Black, 0, 7, loadImage("assets/rook_black.png"),true }
	b.Pieces[7][1] = &Piece{Knight, color.Black, 1, 7, loadImage("assets/knight_black.png"),true}
	b.Pieces[7][2] = &Piece{Bishop, color.Black, 2, 7, loadImage("assets/Bishop_black.png"),true}
	b.Pieces[7][4] = &Piece{Queen, color.Black, 3, 7, loadImage("assets/queen_black.png"),true}
	b.Pieces[7][3] = &Piece{King, color.Black, 4, 7, loadImage("assets/king_black.png"),true }
	b.Pieces[7][5] = &Piece{Bishop, color.Black, 5, 7, loadImage("assets/bishop_black.png"),true}
	b.Pieces[7][6] = &Piece{Knight, color.Black, 6, 7, loadImage("assets/knight_black.png"),true}
	b.Pieces[7][7] = &Piece{Rook, color.Black, 7, 7, loadImage("assets/rook_black.png"),true }
	// Black pawns
	for i := 0; i < 8; i++ {
		b.Pieces[6][i] = &Piece{Pawn, color.Black, i, 6,loadImage("assets/pawn_black.png"), true}
	}
}

// Remove a piece from the board if it has been hit
func (b *Board) RemovePieceAtPosition(x, y int) {
    if b.Pieces[y][x] != nil {
        b.Pieces[y][x].alive = false
        b.Pieces[y][x] = nil
    }
}


// Helper function to load image 
func loadImage(filePath string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return img
}


// draws the chess board
func (b *Board)DrawChessBoard(screen *ebiten.Image) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			// Alternate colors for the chessboard pattern
			if (x+y)%2 == 0 {
				vector.DrawFilledRect(screen, float32(x*b.TileSize), float32(y*b.TileSize), float32(b.TileSize), float32(b.TileSize), color.RGBA{255, 203, 161, 255}, false) // White tile
			} else {
				vector.DrawFilledRect(screen, float32(x*b.TileSize), float32(y*b.TileSize), float32(b.TileSize), float32(b.TileSize), color.RGBA{61, 29, 2, 255}, false) // Black tile
			}
		}
	}
}

func (b *Board) DrawPieces(screen *ebiten.Image) {
	// Loop through all the positions on the board (8x8 grid)
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if piece := b.Pieces[y][x]; piece != nil {
				img := piece.image
				//here we center the image of the piece by getting the bounds of the image and calculating the offset
				offset := (b.TileSize - img.Bounds().Dx()) / 2
				// Draw the piece on the board
				op := &ebiten.DrawImageOptions{}
				// Translate the image to the correct position on the board
				op.GeoM.Translate(float64(x*b.TileSize + offset), float64(y*b.TileSize + offset))
				// Draw the image on the screen
				screen.DrawImage(img, op)
			}
		}
	}
}

func (b *Board)DeletePiece(){
	return
}