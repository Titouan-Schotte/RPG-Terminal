package dialogsmanagement

import (
	"rpg/core"
	"rpg/utils"
)

type Dialog struct {
	NpcName    string
	Texts      []interface{}
	Reply      string
	NextDialog int
}

func ReadDialog(dialogNum int, playerIn core.Player, partieIn string) core.Player {
	if !CanReadDialog(dialogNum, playerIn.DialogsCompleted) {
		return playerIn
	}
	dialogIn := GetDialog(dialogNum)
	utils.ClearConsole()
	for i, text := range dialogIn.Texts {
		println(utils.Colorize(utils.Purple, dialogIn.NpcName) + " >>> " + text.(string))
		utils.Input("")
		if i+1 == len(dialogIn.Texts) {
			println(utils.Colorize(utils.Cyan, playerIn.Pseudo) + " >>> " + dialogIn.Reply)
			utils.Input("")
			utils.ClearConsole()
		}
	}
	playerIn.DialogsCompleted = append(playerIn.DialogsCompleted, dialogNum)
	if dialogIn.NextDialog > 0 {
		playerIn = ReadDialog(dialogIn.NextDialog, playerIn, partieIn)
	}
	return playerIn
}
