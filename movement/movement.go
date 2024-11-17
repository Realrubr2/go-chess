package movement

import (
	"github.com/Realrubr2/go-chess/pieces"
	"image/color"
)

type Move struct {
	board *Pieces.Board
	piece *Pieces.Piece
	fromX, fromY, toX, toY int
}

func (m *Move)MovePiece()bool{
piece := m.board.Pieces[m.fromX][m.fromY]
if piece == nil{
	return false
}
if !m.IsMoveValid(){
	return false
}

m.board.Pieces[m.toX][m.toY] = piece
m.board.Pieces[m.fromX][m.fromY] = nil
m.piece.X = m.toX
m.piece.Y = m.toY
return true


}

func (m *Move)IsMoveValid()bool{
	switch m.piece.Type{
	case Pieces.Pawn:
		m.ValidatePawnMove()
	}
	return false
}

func (m *Move)ValidatePawnMove()bool{
	    // Check if pawn is moving one step forward
		direction := 1
		if m.piece.Color == color.Black {
			direction = -1
		}
	
		// Check if moving forward by one tile
		if m.toY == m.fromY+direction && m.toX == m.fromX && m.board.Pieces[m.toY][m.toX] == nil {
			return true
		}
		
		// Add more pawn rules (e.g., captures, two-square first move) here
		return false
}