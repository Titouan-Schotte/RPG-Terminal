package questmanagement

import (
	"fmt"
	"rpg/core"
	"rpg/utils"
	"strconv"
)

func ShowRewards(player core.Player, quest Quest) core.Player {
	fmt.Println("######## " + utils.Colorize(utils.Green+utils.Bold, "Récompense de Quête") + " ########\n")
	if quest.Rewards.DalisAmount > 0 {
		player.Money += quest.Rewards.DalisAmount

		fmt.Println("         - " + strconv.Itoa(quest.Rewards.DalisAmount) + utils.Colorize(utils.Green, " dalis"))
	}
	if player.Level < quest.Rewards.LevelToPass {
		player.Level = quest.Rewards.LevelToPass
	}
	for k, v := range quest.Rewards.Items {
		fmt.Println("         - " + strconv.Itoa(int(v.(float64))) + " " + utils.Colorize(utils.Red, k))
	}
	fmt.Println("         - " + utils.Colorize(utils.Yellow, "Level ") + strconv.Itoa(quest.Rewards.LevelToPass))
	utils.Input(">>>")
	return player
}
