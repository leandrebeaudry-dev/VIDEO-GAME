package main

import (
	enginecore "github.com/leandrebeaudry-dev/VIDEO-GAME/engine-core"
	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func main() {
	enginecore.PlayIntro()

	player := playersystem.CharacterCreation()

	enginecore.MainMenu(player)
}
