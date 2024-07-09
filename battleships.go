package main

type Sea struct {
	// Not implemented
}

func NewSea() *Sea {
	return &Sea{}
}

func (s *Sea) HasNoShips() bool {
	return true // it's not lying, is it?
}
