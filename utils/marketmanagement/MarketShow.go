package marketmanagement

import (
	"fmt"
	"rpg/core"
	"rpg/utils"
	"rpg/utils/jsonmanagement"
	"strconv"
)

func MarketShow(marketName string, player core.Player, partie string) core.Player {
	utils.ClearConsole()
	var itemsToSell map[string]int
	switch marketName {
	case "alchimiste":
		itemsToSell = map[string]int{
			"potion-vie": 2,
		}
		fmt.Println(utils.Colorize(utils.Yellow, "** Bienvenue chez l'Alchimiste de Lividus **"))
	case "forgeron":
		itemsToSell = map[string]int{
			"epee-lividus":   15,
			"armure-lividus": 25,
		}
		fmt.Println(utils.Colorize(utils.Yellow, "** Bienvenue au Forgeron de Lividus **"))
	case "ziak":
		itemsToSell = map[string]int{
			"epee-cendre":   60,
			"armure-cendre": 100,
		}
		fmt.Println(utils.Colorize(utils.Yellow, "** Bienvenue chez Ziak le Forgeron Militaire **"))
	}
	println()
	nItem := 1
	for _, v := range utils.GetKeys(itemsToSell) {
		fmt.Printf("%v > %v : %v dalis\n", utils.Colorize(utils.Yellow, nItem), utils.Colorize(utils.Red, v), utils.Colorize(utils.Yellow+utils.Bold, itemsToSell[v]))
		nItem++
	}
	fmt.Println("\n\n-> Selectionnez le numéro associé pour acheter l'item")
	fmt.Println("-> Appuyez sur \"l\" pour quitter le commerce.")
	inp := utils.Input(">>")
	if len(inp) > 0 {
		if inp[0] == 'l' {
			return player
		}

		inpInt, err := strconv.Atoi(string(inp[0]))
		if err == nil {
			if len(itemsToSell) >= inpInt && inpInt > 0 {
				amoutOfDalisNeed := itemsToSell[utils.GetKeys(itemsToSell)[inpInt-1]]
				if player.Money >= amoutOfDalisNeed {
					player.Money -= amoutOfDalisNeed

					player.AddItem(partie, utils.GetKeys(itemsToSell)[inpInt-1], 1)
					utils.ClearConsole()
					fmt.Println(utils.Colorize(utils.Yellow, "-> Vous avez acheté 1 "+utils.GetKeys(itemsToSell)[inpInt-1]) + ".")
					fmt.Println(utils.Colorize(utils.Green, "-> Vous avez en votre possession "+strconv.Itoa(player.Inventaire[utils.GetKeys(itemsToSell)[inpInt-1]])+" "+utils.GetKeys(itemsToSell)[inpInt-1]) + ".")
					fmt.Println(utils.Colorize(utils.Green, "-> Il vous reste "+strconv.Itoa(player.Money)+" dalis."))

					switch utils.GetKeys(itemsToSell)[inpInt-1] {
					case "armure-lividus":
						fmt.Println("-> Vous avez désormais " + utils.Colorize(utils.Green, "150 points de vie") + " !")
						jsonmanagement.Update(partie, "health", 150)
						jsonmanagement.Update(partie, "maxhealth", 150)
					case "epee-lividus":
						fmt.Println("-> Vous faites désormais " + utils.Colorize(utils.Red, "30 dégats") + " !")
						jsonmanagement.Update(partie, "damages", 30)
					case "armure-cendre":
						fmt.Println("-> Vous avez désormais " + utils.Colorize(utils.Green, "225 points de vie") + " !")
						jsonmanagement.Update(partie, "health", 225)
						jsonmanagement.Update(partie, "maxhealth", 225)
					case "epee-cendre":
						fmt.Println("-> Vous faites désormais " + utils.Colorize(utils.Red, "50 dégats") + " !")
						jsonmanagement.Update(partie, "damages", 50)
					case "potion-vie":
						fmt.Println("-> " + utils.Colorize(utils.Yellow+utils.Bold, "/heal") + " pour utiliser une potion hors combat !")
					}
					utils.Input(">>>")
				} else {
					utils.ClearConsole()
					fmt.Println(utils.Colorize(utils.Red, "-> Vous n'avez pas assez de dalis."))
					utils.Input(">>>")
				}
			}
		}
	}
	MarketShow(marketName, player, partie)
	return player
}
