package movement

import (
	"fmt"
	"image/color"
	"math"

	"github.com/Realrubr2/go-chess/pieces"
)

type Move struct {
	Board *Pieces.Board
	Piece *Pieces.Piece
	FromX, FromY, ToX, ToY int
}

func (m *Move)MovePiece()bool{
piece := m.Board.Pieces[m.FromX][m.FromY]
if piece == nil{
	fmt.Println("no piece selected")
	return false
}

if !m.IsMoveValid(){
	fmt.Println("Invalid move (movepiuece)")
	return false
}



	// Simulate the move
	m.Board.Pieces[m.ToX][m.ToY] = piece
	m.Board.Pieces[m.FromX][m.FromY] = nil
	m.Piece.X = m.ToX
	m.Piece.Y = m.ToY

	// Check if the move puts the current player's king in check
	if m.IsKingInCheck(m.Board.Pieces, m.Piece.Color) {
		fmt.Println("Invalid move (king in check)")
		m.Board.Pieces[m.FromX][m.FromY] = piece
		m.Board.Pieces[m.ToX][m.ToY] = nil
		m.Piece.X = m.FromX
		m.Piece.Y = m.FromY
		return false
	}

	return true
}



func (m *Move)IsMoveValid()bool{
	switch m.Piece.Type{
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
		if m.Piece.Color == color.Black {
			direction = -1
		}
	
		if m.ToY == m.FromY+direction && m.ToX == m.FromX && m.Board.Pieces[m.ToY][m.ToX] == nil {
			return true
		}
		
		// Add more pawn rules (e.g., captures, two-square first move) here
		return false
}

// Check if rook is moving vertically or horizontally
func (m *Move)ValidateRookMove()bool{
	if m.ToX == m.FromX || m.ToY == m.FromY {
		return true
	}
	return false
}
// Check if knight is moving in an L-shape
func(m*Move)ValidateKnightMove()bool{
	if (m.ToX == m.FromX+2 || m.ToX == m.FromX-2) && (m.ToY == m.FromY+1 || m.ToY == m.FromY-1) {
		return true
	}
	if (m.ToY == m.FromY+2 || m.ToY == m.FromY-2) && (m.ToX == m.FromX+1 || m.ToX == m.FromX-1) {
		return true
	}
	return false
}

// checks if move is diagonal by roundingg it to an absole value and comparing if the x and y both have the same value
func (m *Move)ValidateBishopMove()bool{
	return math.Abs(float64(m.ToX-m.FromX)) == math.Abs(float64(m.ToY-m.FromY))
}
// valiudates the king move by checking if the move is within 1 tile of the current position
func (m *Move)ValidateKingMove()bool{
	if math.Abs(float64(m.ToX-m.FromX)) <= 1 && math.Abs(float64(m.ToY-m.FromY)) <= 1 {
		return true
	}
	return false
}