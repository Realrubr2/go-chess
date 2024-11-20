package movement

import (
	"github.com/Realrubr2/go-chess/pieces"
	"math"
	"image/color"
)

// IsKingInCheck checks if the given player's king is in check after a move.
func (m *Move)IsKingInCheck(board [8][8]*Pieces.Piece, kingColor color.Color) bool {
	var kingX, kingY int
	// Find the position of the king for the given color
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			piece := board[y][x]
			if piece != nil && piece.Type == Pieces.King && piece.Color == kingColor {
				kingX, kingY = x, y
				break
			}
		}
	}

	// check if any opposing piece can attack the king
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			opponentPiece := board[y][x]
			if opponentPiece != nil && opponentPiece.Color != kingColor {
				if m.isPieceThreateningKing(board, *opponentPiece, x, y, kingX, kingY) {
					return true // King is in check
				}
			}
		}
	}

	
	return false
}

// isPieceThreateningKing checks if a piece can attack the king's position.
func (m *Move)isPieceThreateningKing(board [8][8]*Pieces.Piece, attackingPiece Pieces.Piece, fromX, fromY, toX, toY int) bool {
	switch attackingPiece.Type {
	case Pieces.Pawn:
		 if m.ValidatePawnMove() && m.isPathClear(board, fromX, fromY, toX, toY) {
			 return true
		 }
	case Pieces.Rook:
		if m.ValidateRookMove() && m.isPathClear(board, fromX, fromY, toX, toY) {
			return true
		}
	case Pieces.Knight:
		if m.ValidateKnightMove() && m.isPathClear(board, fromX, fromY, toX, toY){
			return true
		}
	case Pieces.Bishop:
		if m.ValidateBishopMove() && m.isPathClear(board, fromX, fromY, toX, toY) {
			return true
		}
	case Pieces.Queen:
		if m.ValidateRookMove() || m.ValidateBishopMove() && m.isPathClear(board, fromX, fromY, toX, toY){
			return true
		}
	}
	return false
}

// isPathClear checks if there are no Pieces blocking the path between two positions
func (m *Move)isPathClear(board [8][8]*Pieces.Piece, fromX, fromY, toX, toY int) bool {
	dx := 0
	dy := 0

	if fromX != toX {
		dx = (toX - fromX) / int(math.Abs(float64(toX-fromX)))
	}
	if fromY != toY {
		dy = (toY - fromY) / int(math.Abs(float64(toY-fromY)))
	}

	x, y := fromX+dx, fromY+dy
	for x != toX || y != toY {
		if board[y][x] != nil {
			return false // There is a piece blocking the path
		}
		x += dx
		y += dy
	}

	return true
}
