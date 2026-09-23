package worldsystem

import "fmt"

// Structure d’une zone du monde
type Zone struct {
	Name        string
	Description string
	HasNPC      bool
}

// Liste des zones du jeu
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

// Fonction pour afficher une zone
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
