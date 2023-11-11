package dialogsmanagement

import (
	"rpg/utils/jsonmanagement"
	"strconv"
)

func GetDialog(dialogNum int) Dialog {
	dialogEmp := "dialogs/" + strconv.Itoa(dialogNum)
	return Dialog{
		NpcName:    jsonmanagement.Get(dialogEmp, "npcName").(string),
		Texts:      jsonmanagement.Get(dialogEmp, "texts").([]interface{}),
		Reply:      jsonmanagement.Get(dialogEmp, "reply").(string),
		NextDialog: int(jsonmanagement.Get(dialogEmp, "nextDialog").(float64)),
	}
}
