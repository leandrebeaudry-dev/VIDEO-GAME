package main

import (
	enginecore "github.com/leandrebeaudry-dev/VIDEO-GAME/engine-core"
	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func main() {
	// 1. Lancement de l'introduction narrative
	enginecore.PlayIntro()

	// 2. Création et initialisation du personnage
	player := playersystem.CharacterCreation()

	// 3. Affichage du menu principal
	enginecore.MainMenu(player)
}
