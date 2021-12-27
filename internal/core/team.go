package core

type TeamID uint

type Team struct {
	ID       TeamID
	Name     string
	Category CategoryID
}
