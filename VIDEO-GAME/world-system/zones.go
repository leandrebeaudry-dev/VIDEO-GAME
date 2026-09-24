package worldsystem

import (
	"fmt"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func StartZonesSession(p *playersystem.Player) {
	for {
		fmt.Println("\n==========================================")
		fmt.Println("       Bienvenue dans l'exploration")
		fmt.Println("==========================================")

		fmt.Println("1. Village Armoricain")
		fmt.Println("2. Camp d'Achim")
		fmt.Println("3. Ruines du Domaine de Dana")
		fmt.Println("4. Route Commerciale")
		fmt.Println("5. Vallée de Dana")
		fmt.Println("6. Retour au menu principal")
		fmt.Print("Entrez votre choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			DisplayZone("village")
		case 2:
			DisplayZone("camp_achim")
		case 3:
			DisplayZone("ruines_de_dana")
		case 4:
			DisplayZone("route_commerciale")
		case 5:
			DisplayZone("vallee_de_dana")
		case 6:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

type Zone struct {
	Name        string
	Description string
	HasNPC      bool
}

var Zones = map[string]Zone{
	"village": {
		Name:        "Village Armoricain",
		Description: "Un petit village encore debout malgré la menace des Simériens.",
		HasNPC:      true,
	},

	"camp_achim": {
		Name:        "Camp d'Achim",
		Description: "Un campement isolé où Achim vous attend pour commencer votre voyage.",
		HasNPC:      true,
	},

	"ruines_de_dana": {
		Name:        "Ruines du Domaine de Dana",
		Description: "Les vestiges du domaine de Dana, détruit par les Simériens.",
		HasNPC:      true,
	},

	"route_commerciale": {
		Name:        "Ancienne Route Commerciale",
		Description: "Une route autrefois prospère, maintenant abandonnée et dangereuse.",
		HasNPC:      true,
	},

	"vallee_de_dana": {
		Name:        "Vallée de Dana",
		Description: "La vallée sacrée, lieu de votre voyage initiatique.",
		HasNPC:      false,
	},
}

func DisplayZone(zoneName string) {
	zone, exists := Zones[zoneName]
	if !exists {
		fmt.Println("Zone inconnue :", zoneName)
		return
	}

	fmt.Println("------------------------------------------------")
	fmt.Println("Vous êtes dans :", zone.Name)
	fmt.Println(zone.Description)
	fmt.Println("------------------------------------------------")

	if zone.HasNPC {
		npc := GetNPCByZone(zoneName)
		if npc != nil {
			npc.Talk()
		}
	}
}
