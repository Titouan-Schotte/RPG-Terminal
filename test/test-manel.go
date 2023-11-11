package main

import (
	"rpg/utils/jsonmanagement"
)

func main() {
	jsonmanagement.CreateJson("lividus")
	jsonmanagement.Add("lividus", "nombreHabitants", 50)
	nombreHabitant := jsonmanagement.Get("lividus", "nombreHabitants").(float64)
	nombreHabitant = nombreHabitant - 30
	jsonmanagement.Update("lividus", "nombreHabitants", nombreHabitant)
}
