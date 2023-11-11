package questmanagement

import "rpg/core"

func AllObjectivesItems(QuestIn Quest, PlayerIn core.Player) bool {
	allOk := true
	items := QuestIn.Content.(ContentItems).Items
	for k, v := range items {
		if PlayerIn.Inventaire[k] < v {
			allOk = false
		}
	}
	return allOk
}

func AllObjectivesKills(QuestIn Quest, PlayerIn core.Player) bool {
	allOk := true
	targetsObj := QuestIn.Content.(ContentKill).Targets
	targetsIn := QuestIn.Content.(ContentKill).TargetsIn
	for k, v := range targetsObj {
		if targetsIn[k] < v {
			allOk = false
		}
	}
	return allOk
}
