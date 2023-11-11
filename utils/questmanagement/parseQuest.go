package questmanagement

import (
	"rpg/utils/jsonmanagement"
	"strconv"
)

type Rewards struct {
	DalisAmount int
	LevelToPass int
	Items       map[string]interface{}
}

type Quest struct {
	QuestEmp  string
	QuestNum  int
	Log       string
	TypeQuest string
	Content   interface{}
	Rewards   Rewards
}

func ParseQuest(questNumb int) Quest {
	questEmp := "quests/" + strconv.Itoa(questNumb)
	return Quest{
		QuestEmp:  questEmp,
		QuestNum:  questNumb,
		Log:       jsonmanagement.Get(questEmp, "log").(string),
		TypeQuest: jsonmanagement.Get(questEmp, "type").(string),
		Content:   jsonmanagement.Get(questEmp, "content"),
		Rewards: Rewards{
			DalisAmount: int(jsonmanagement.Get(questEmp, "dalis").(float64)),
			LevelToPass: int(jsonmanagement.Get(questEmp, "lvl").(float64)),
			Items:       jsonmanagement.Get(questEmp, "items").(map[string]interface{}),
		},
	}
}

func (quest *Quest) IsKillQuest() bool {
	return quest.TypeQuest == "kill"
}

func (quest *Quest) IsItemsQuest() bool {
	return quest.TypeQuest == "items"
}

func (quest *Quest) IsDialogQuest() bool {
	return quest.TypeQuest == "dialog"
}
