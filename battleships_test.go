package main

import "testing"

func TestSeaStartsEmpty(t *testing.T) {
	// Arrange
	sea := NewSea()

	// Act
	// Not required

	// Assert
	want := true
	got := sea.HasNoShips()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSeaNotEmpty(t *testing.T) {
	// Arrange
	sea := NewSea()

	// Act
	sea.AddShip(1, 1) // ??? Your design decision goes here

	// Assert
	want := false
	got := sea.HasNoShips()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPlaceShipAtLocation(t *testing.T) {
	sea := NewSea()
	wantedRow := 1
	wantedCol := 1

	sea.AddShip(1, 1)

	if !sea.HasShips(wantedRow, wantedCol) {
		t.Errorf("Expected a ship at (%d, %d) but found none", wantedRow, wantedCol)
	}

}

func TestPlacingNineShips(t *testing.T) {
	sea := NewSea()

	shipLocations := []struct {
		row, col int
	}{
		{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6},
		{2, 1}, {3, 1}, {4, 1},
	}

	for _, location := range shipLocations {
		sea.AddShip(location.row, location.col)
	}
	for _, location := range shipLocations {
		if !sea.HasShips(location.row, location.col) {
			t.Errorf("Expected a ship at (%d, %d) but found none", location.row, location.col)
		}
	}

}

func TestReportMissFromShotAtOpponent(t *testing.T) {
	sea := NewSea()
	sea.AddShip(1, 1)

	got := sea.FireShot(2, 2)

	want := "miss"

	if got != want {
		t.Errorf("Shoudl be miss")
	}
}

func TestReportHitFromShotAtOpponent(t *testing.T) {
	sea := NewSea()
	sea.AddShip(1, 1)

	got := sea.FireShot(1, 1)

	want := "hit"

	if got != want {
		t.Errorf("Should be hit")
	}
}

func TestReportErrorFromShotOutOfBounds(t *testing.T) {
	sea := NewSea()
	sea.AddShip(1, 1)

	got := sea.FireShot(8, 8)

	want := "invalid"

	if got != want {
		t.Errorf("Shot at out-of-bounds (8, 8) should return 'invalid', got %v", got)
	}

}

func TestSinkShipAfterBeingHit(t *testing.T) {
	sea := NewSea()
	sea.AddShip(1, 1)
	sea.FireShot(1, 1)

	got := sea.WasShipSunk(1, 1)

	want := "sunk"

	if got != want {
		t.Error("Ship was not sunk")
	}
}

func TestSwitchActivePlayerToOtherPlayer(t *testing.T) {
	sea := NewSea()

	got := sea.SwitchPlayer()

	want := "player switched"

	if got != want {
		t.Error("Player was not switched")
	}
}

func TestHasActivePlayerWon(t *testing.T) {
	sea := NewSea()

	got := sea.HasACtivePlayerWon()

	want := "player 1 won"

	if got != want {
		t.Error("Active player has not won the game")
	}
}
