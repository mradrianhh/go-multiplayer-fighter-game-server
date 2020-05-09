package models

// Players is the global collection of all players in-game.
var Players []Player

// Player represents the user in-game.
type Player struct {
	Username string
	Exp      int
}

// UpdateExp alters the player's exp.
func (player *Player) UpdateExp(exp int) {
	player.Exp += exp
}
