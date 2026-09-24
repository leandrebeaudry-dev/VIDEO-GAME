package worldsystem

import (
	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func StartMovementSession(p *playersystem.Player) {
}

type Vector2 struct {
	X int
	Y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func Move(pos Vector2, dir Direction) Vector2 {
	switch dir {
	case Up:
		pos.Y--
	case Down:
		pos.Y++
	case Left:
		pos.X--
	case Right:
		pos.X++
	}
	return pos
}
