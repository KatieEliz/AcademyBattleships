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
