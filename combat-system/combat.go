package combat

import "video-game/player"

func Attack(p player.Player, enemyHP int) int {
	return enemyHP - p.Attack
}
