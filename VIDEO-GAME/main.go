package main

import (
	enginecore "github.com/leandrebeaudry-dev/VIDEO-GAME/engine-core"
	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func main() {
	// 1. Animation d'introduction
	enginecore.PlayIntro()

	// 2. Création du personnage
	player := playersystem.CharacterCreation()

	// 3. Lancement du menu principal
	enginecore.MainMenu(player)
}
