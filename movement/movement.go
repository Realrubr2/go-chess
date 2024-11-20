package movement

import (
	"github.com/Realrubr2/go-chess/pieces"
	"image/color"
	"math"
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



	// Simulate the move
	m.board.Pieces[m.toX][m.toY] = piece
	m.board.Pieces[m.fromX][m.fromY] = nil
	m.piece.X = m.toX
	m.piece.Y = m.toY

	// Check if the move puts the current player's king in check
	if m.IsKingInCheck(m.board.Pieces, m.piece.Color) {
		m.board.Pieces[m.fromX][m.fromY] = piece
		m.board.Pieces[m.toX][m.toY] = nil
		m.piece.X = m.fromX
		m.piece.Y = m.fromY
		return false
	}

	return true
}



func (m *Move)IsMoveValid()bool{
	switch m.piece.Type{
	case Pieces.Pawn:
		m.ValidatePawnMove()
	case Pieces.Rook:
		m.ValidateRookMove()
	case Pieces.Knight:
		m.ValidateKnightMove()
	case Pieces.Bishop:
		m.ValidateBishopMove()
	case Pieces.Queen:
	return m.ValidateRookMove() || m.ValidateBishopMove()
	case Pieces.King:
		m.ValidateKingMove()
	}
	return false
}

// Check if pawn is moving one step forward
// Check if moving forward by one tile
func (m *Move)ValidatePawnMove()bool{
		direction := 1
		if m.piece.Color == color.Black {
			direction = -1
		}
	
		if m.toY == m.fromY+direction && m.toX == m.fromX && m.board.Pieces[m.toY][m.toX] == nil {
			return true
		}
		
		// Add more pawn rules (e.g., captures, two-square first move) here
		return false
}

// Check if rook is moving vertically or horizontally
func (m *Move)ValidateRookMove()bool{
	if m.toX == m.fromX || m.toY == m.fromY {
		return true
	}
	return false
}
// Check if knight is moving in an L-shape
func(m*Move)ValidateKnightMove()bool{
	if (m.toX == m.fromX+2 || m.toX == m.fromX-2) && (m.toY == m.fromY+1 || m.toY == m.fromY-1) {
		return true
	}
	if (m.toY == m.fromY+2 || m.toY == m.fromY-2) && (m.toX == m.fromX+1 || m.toX == m.fromX-1) {
		return true
	}
	return false
}

// checks if move is diagonal by roundingg it to an absole value and comparing if the x and y both have the same value
func (m *Move)ValidateBishopMove()bool{
	return math.Abs(float64(m.toX-m.fromX)) == math.Abs(float64(m.toY-m.fromY))
}
// valiudates the king move by checking if the move is within 1 tile of the current position
func (m *Move)ValidateKingMove()bool{
	if math.Abs(float64(m.toX-m.fromX)) <= 1 && math.Abs(float64(m.toY-m.fromY)) <= 1 {
		return true
	}
	return false
}