package questmanagement

import (
	"rpg/utils"
	"rpg/utils/dialogsmanagement"
	"rpg/utils/jsonmanagement"
	"strconv"
	"strings"
)

func GiveQuest(partieIn string, questNum int) ([]string, Quest) {
	if questNum <= 0 {
		return make([]string, 0), Quest{}
	}
	utils.ClearConsole()
	var questLog []string
	questLog = append(questLog, "#########> "+utils.Colorize(utils.Black+utils.YellowBackground, "Quête "+strconv.Itoa(questNum))+" <##########")
	quest := ParseQuest(questNum)

	for _, sQuestLog := range strings.Split(quest.Log, "|") {
		questLog = append(questLog, " "+utils.Colorize(utils.Purple, sQuestLog)+"#")
	}
	questLog = append(questLog, "#############################")

	quest.TransformQuestByType()
	if quest.IsKillQuest() {
		targets := quest.Content.(ContentKill).Targets
		for _, v := range utils.GetKeys(targets) {
			phraseIn := utils.Colorize(utils.Red, jsonmanagement.Get("mobs/"+v, "pseudo").(string)) + " : " + utils.Colorize(utils.Bold, "%v / "+utils.Colorize(utils.Yellow, strconv.Itoa(targets[v])))
			for i := 0; i < 27-len(jsonmanagement.Get("mobs/"+v, "pseudo").(string)+" : 0 / "+strconv.Itoa(targets[v])); i++ {
				phraseIn += " "
			}
			phraseIn += "#"
			questLog = append(questLog, " "+phraseIn)
		}
	} else if quest.IsItemsQuest() {
		items := quest.Content.(ContentItems).Items
		for _, v := range utils.GetKeys(items) {
			actualAmount := strconv.Itoa(int(jsonmanagement.Get(partieIn, "inventory").(map[string]interface{})[v].(float64)))
			phraseIn := utils.Colorize(utils.Yellow, v) + " : " + utils.Colorize(utils.Bold, "%v") + " / " + utils.Colorize(utils.Bold, strconv.Itoa(items[v]))
			for i := 0; i < 27-len(v+" : "+actualAmount+" / "+strconv.Itoa(items[v])); i++ {
				phraseIn += " "
			}
			phraseIn += "#"
			questLog = append(questLog, " "+phraseIn)
		}
	} else {
		phraseIn := utils.Colorize(utils.Cyan, "Discuter : "+dialogsmanagement.GetDialog(quest.Content.(ContentDialog).dialog).NpcName)
		for i := 0; i < 27-len("Discuter : "+dialogsmanagement.GetDialog(quest.Content.(ContentDialog).dialog).NpcName); i++ {
			phraseIn += " "
		}
		phraseIn += "#"
		questLog = append(questLog, " "+utils.Colorize(utils.Cyan, phraseIn))
	}
	questLog = append(questLog, "#############################")
	jsonmanagement.Update(partieIn, "questIn", questNum)
	return questLog, quest
}
