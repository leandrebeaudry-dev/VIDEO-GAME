package worldsystem

import (
	"fmt"

	playersystem "github.com/leandrebeaudry-dev/VIDEO-GAME/player-system"
)

func StartNpcsSession(p *playersystem.Player) {
}

type NPC struct {
	Name        string
	Description string
	Dialogue    []string
}

func (n NPC) Talk() {
	fmt.Println("------------------------------------------------")
	fmt.Println("Vous parlez à :", n.Name)
	fmt.Println(n.Description)
	fmt.Println("------------------------------------------------")

	for _, line := range n.Dialogue {
		fmt.Println(line)
	}

	fmt.Println("------------------------------------------------")
}

var Achim = NPC{
	Name:        "Achim",
	Description: "Un homme mystérieux, porteur d'une lourde vérité.",
	Dialogue: []string{
		"Je t'ai cherché partout… le domaine de Dana est tombé.",
		"Les Simériens ont ravagé nos terres. Rien n'a pu les arrêter.",
		"Ton voyage dans la vallée de Dana commence maintenant. Tu n'es pas seul.",
	},
}

var DanaSurvivor = NPC{
	Name:        "Survivant de Dana",
	Description: "Un habitant marqué par la destruction de son domaine.",
	Dialogue: []string{
		"Ils sont arrivés sans prévenir… les Simériens…",
		"Nos maisons ont brûlé, nos familles ont fui.",
		"Si tu vas dans la vallée, fais attention… ils rôdent encore.",
	},
}

var VillageGuard = NPC{
	Name:        "Garde Armoricain",
	Description: "Un garde qui protège les derniers villages encore debout.",
	Dialogue: []string{
		"Bienvenue dans notre village. Nous tenons encore debout… pour l’instant.",
		"La vallée de Dana est dangereuse, mais tu dois t’y rendre.",
		"Achim t’attend plus loin, ne traîne pas.",
	},
}

var MerchantNPC = NPC{
	Name:        "Marchand Itinérant",
	Description: "Un marchand qui semble connaître beaucoup de secrets.",
	Dialogue: []string{
		"Les Simériens ont coupé nos routes commerciales… tout est plus rare.",
		"Si tu vas dans la vallée, tu auras besoin de potions et d’équipement.",
		"Le forgeron peut t’aider à survivre là-bas.",
	},
}

func GetNPCByZone(zone string) *NPC {
	switch zone {
	case "village":
		return &VillageGuard
	case "camp_achim":
		return &Achim
	case "ruines_de_dana":
		return &DanaSurvivor
	case "route_commerciale":
		return &MerchantNPC
	default:
		return nil
	}
}
