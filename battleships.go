package main

type Sea struct {
	player1Grid [][]bool
	player2Grid [][]bool
	playerTurn  int
}

func NewSea() *Sea {

	grid := make([][]bool, 7)
	for i := range grid {
		grid[i] = make([]bool, 7)
	}
	return &Sea{
		player1Grid: grid,
		player2Grid: grid,
		playerTurn:  0,
	}
}
func (s *Sea) getCurrentPlayerGrid() [][]bool {
	if s.playerTurn == 0 {
		return s.player2Grid
	}
	return s.player1Grid
}

func (s *Sea) isValidLocation(row, col int) bool {
	return row >= 0 && row < len(s.player1Grid) && col >= 0 && col < len(s.player1Grid[0])
}

func (s *Sea) AddShip(row, col int) {
	if s.playerTurn == 0 {
		if s.isValidLocation(row, col) {
			s.player1Grid[row][col] = true
		}
	} else {
		if s.isValidLocation(row, col) {
			s.player2Grid[row][col] = true
		}
	}
}

func (s *Sea) HasShips(row, col int) bool {
	if s.isValidLocation(row, col) {
		return s.getCurrentPlayerGrid()[row][col]
	}
	return false
}

func (s *Sea) HasNoShips() bool {
	gridToCheck := s.getCurrentPlayerGrid()

	for row := range gridToCheck {
		for col := range gridToCheck[row] {
			if gridToCheck[row][col] {
				return false
			}
		}
	}
	return true

}

func (s *Sea) FireShot(row, col int) string {
	if !s.isValidLocation(row, col) {
		return "invalid"
	}

	gridToCheck := s.getCurrentPlayerGrid()

	if gridToCheck[row][col] {
		gridToCheck[row][col] = true
		return "hit"
	}
	return "miss"
}

func (s *Sea) WasShipSunk(row, col int) string {
	if !s.isValidLocation(row, col) {
		return "invalid"
	}
	shotResult := s.FireShot(row, col)
	if shotResult == "hit" {
		return "sunk"
	}
	return shotResult
}

func (s *Sea) SwitchPlayer() string {
	s.playerTurn = (s.playerTurn + 1) % 2
	return "player switched"
}

func (s *Sea) HasACtivePlayerWon() string {
	gridToCheck := s.getCurrentPlayerGrid()

	for row := range gridToCheck {
		for col := range gridToCheck[row] {
			if gridToCheck[row][col] {
				return "game not over"
			}
		}
	}

	if s.playerTurn == 0 {
		return "player 1 won"
	} else {
		return "player 2 won"
	}
}
