package questmanagement

import (
	"fmt"
)

type ContentKill struct {
	Targets   map[string]int
	TargetsIn map[string]int
}

func (c *ContentKill) Init(contentIn interface{}) ContentKill {
	targets := map[string]int{}
	targetsIn := map[string]int{}
	fmt.Println("C", contentIn.(map[string]interface{}))
	for k, v := range contentIn.(map[string]interface{}) {
		targets[k] = int(v.(float64))
		targetsIn[k] = 0
	}
	c.Targets = targets
	c.TargetsIn = targetsIn
	return *c
}

type ContentDialog struct {
	dialog int
}

func (c *ContentDialog) Init(contentIn interface{}) ContentDialog {
	c.dialog = int(contentIn.(float64))
	return *c
}

type ContentItems struct {
	Items map[string]int
}

func (c *ContentItems) Init(contentIn interface{}) ContentItems {
	items := map[string]int{}
	for k, v := range contentIn.(map[string]interface{}) {
		items[k] = int(v.(float64))
	}
	c.Items = items
	return *c
}

func (quest *Quest) TransformQuestByType() {
	if quest.IsKillQuest() {
		temp := ContentKill{}
		quest.Content = temp.Init(quest.Content)
	} else if quest.IsDialogQuest() {
		temp := ContentDialog{}
		quest.Content = temp.Init(quest.Content)
	} else if quest.IsItemsQuest() {
		temp := ContentItems{}
		quest.Content = temp.Init(quest.Content)
	}
}
