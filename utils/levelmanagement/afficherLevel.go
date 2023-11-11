package levelmanagement

import (
	"fmt"
	"rpg/core"
	"rpg/utils"
	"rpg/utils/dialogsmanagement"
	"rpg/utils/jsonmanagement"
	"rpg/utils/questmanagement"
	"strings"
)

func AfficherLevel(name string, level [][]map[string]string, questlog []string, QuestIn questmanagement.Quest, player core.Player) {
	print(utils.Colorize(utils.Yellow, name))
	if player.QuestIn > 0 {
		for space := 0; space < 30-len(name); space++ {
			fmt.Print(" ")
		}
		fmt.Println(questlog[0])
	} else {
		fmt.Println()
	}
	objectifCompt := 0
	for i, row := range level {
		for _, cell := range row {
			for symbol, _ := range cell {
				colorize := ""
				switch symbol {
				case "🧱":
					colorize = utils.Red
				case "🟩":
					colorize = utils.Green
				case "🌳":
					colorize = utils.Green
				case "🌲":
					colorize = utils.Green
				case "🐊":
					colorize = utils.Green + utils.Bold
				case "▶ ":
					colorize = utils.Purple
				case "◀ ":
					colorize = utils.Purple
				case "🔺":
					colorize = utils.Purple
				case "🔻":
					colorize = utils.Purple
				case "🟫":
					colorize = utils.Gray
				case "🪨":
					colorize = utils.Gray + utils.Bold
				case "🍂":
					colorize = utils.Red + utils.Bold
				case "🟨":
					colorize = utils.Yellow
				case "💰":
					colorize = utils.Yellow + utils.Bold
				case "🪙":
					colorize = utils.Green + utils.Bold
				case "🛍️ ":
					colorize = utils.Cyan + utils.Bold
				case "🟧":
					colorize = utils.Red + utils.Bold
				case "🌋":
					colorize = utils.Red + utils.Bold
				case "🟦":
					colorize = utils.Blue
				case "🟥":
					colorize = utils.Red
				case "🟪":
					colorize = utils.Purple
				case "💧":
					colorize = utils.Blue
				case "🔰":
					colorize = utils.Yellow + utils.Bold
				case "🦆":
					colorize = utils.Yellow
				case "🌾":
					colorize = utils.Yellow + utils.Bold
				case "🌼":
					colorize = utils.White + utils.Bold
				case "🌻":
					colorize = utils.Yellow + utils.Bold
				case "🛌":
					colorize = utils.Purple
				case "🌹":
					colorize = utils.Red + utils.Bold
				case "🌱":
					colorize = utils.Green + utils.Bold
				case "🪴":
					colorize = utils.Green + utils.Bold
				case "🌺":
					colorize = utils.Purple + utils.Bold
				case "🔥":
					colorize = utils.Red
				case "⚪":
					colorize = utils.Gray
				case "😷":
					if !dialogsmanagement.CanReadDialog(4, player.DialogsCompleted) {
						colorize = utils.Yellow
					} else {
						colorize = utils.Gray
					}
				case "🫅":
					if !dialogsmanagement.CanReadDialog(7, player.DialogsCompleted) {
						colorize = utils.Yellow
					} else {
						colorize = utils.Gray
					}
				case "🧝":
					if !dialogsmanagement.CanReadDialog(10, player.DialogsCompleted) {
						colorize = utils.Yellow
					} else {
						colorize = utils.Gray
					}
				}
				fmt.Print(utils.Colorize(colorize, symbol))
			}
		}
		if player.QuestIn > 0 {
			for space := 0; space < 15-len(row); space++ {
				fmt.Print("  ")
			}
			if len(questlog) > i+1 && i+1 > 0 {
				fmt.Print("#")
				if strings.Contains(questlog[i+1], "%v") {
					if QuestIn.IsItemsQuest() {
						fmt.Printf(questlog[i+1], player.Inventaire[utils.GetKeys(QuestIn.Content.(questmanagement.ContentItems).Items)[objectifCompt]])
						objectifCompt++
					}
					if QuestIn.IsKillQuest() {
						fmt.Printf(questlog[i+1], QuestIn.Content.(questmanagement.ContentKill).TargetsIn[utils.GetKeys(QuestIn.Content.(questmanagement.ContentKill).TargetsIn)[objectifCompt]])
						objectifCompt++
					}
				} else {
					fmt.Print(questlog[i+1])
				}
			}
		}
		println()

	}
	if player.QuestIn > 0 {
		questLogOverflow := len(questlog) - len(level)
		for i := len(level); i < len(level)+questLogOverflow; i++ {
			for space := 0; space < 15; space++ {
				fmt.Print("  ")
			}

			if len(questlog) > i+1 && i+1 > 0 {
				fmt.Print("#")
				if strings.Contains(questlog[i+1], "%v") {
					if QuestIn.IsItemsQuest() {
						fmt.Printf(questlog[i+1], player.GetAmountOfItem(utils.GetKeys(QuestIn.Content.(questmanagement.ContentItems).Items)[objectifCompt]))
						objectifCompt++
					}
					if QuestIn.IsKillQuest() {
						fmt.Printf(questlog[i+1], QuestIn.Content.(questmanagement.ContentKill).TargetsIn[utils.GetKeys(QuestIn.Content.(questmanagement.ContentKill).TargetsIn)[objectifCompt]])
						objectifCompt++
					}
				} else {
					fmt.Print(questlog[i+1])
				}
			}
			println()
		}
	}
	jaugeLength := 20
	pourcentageVie := jsonmanagement.Get("parties/"+player.Pseudo, "health").(float64) / jsonmanagement.Get("parties/"+player.Pseudo, "maxhealth").(float64)
	nbCaracteresRemplis := int(pourcentageVie * float64(jaugeLength))
	nbCaracteresVides := jaugeLength - nbCaracteresRemplis

	jaugeRemplie := strings.Repeat("█", nbCaracteresRemplis)
	jaugeVide := strings.Repeat("░", nbCaracteresVides)

	jauge := jaugeRemplie + jaugeVide
	fmt.Printf("\n"+utils.Colorize(utils.Cyan, "🧝")+": [%s] %v/%v hp\n", utils.Colorize(utils.Green, jauge), utils.Colorize(utils.Yellow+utils.Bold, player.Health), utils.Colorize(utils.Green+utils.Bold, player.MaxHealth))
}
