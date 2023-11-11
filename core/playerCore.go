package core

import (
	"fmt"
	"rpg/utils"
	"rpg/utils/jsonmanagement"
	"strconv"
)

type Player struct {
	X, Y             int
	Pseudo           string
	MapIn            string
	Level            int
	QuestIn          int
	DialogsCompleted []int
	Health           int
	MaxHealth        int
	Damages          int
	Money            int
	Inventaire       map[string]int
}

func (p *Player) Load(partie string) Player {
	dialogsId := make([]int, 0)
	for _, v := range jsonmanagement.Get(partie, "dialogsCompleted").([]interface{}) {
		dialogsId = append(dialogsId, int(v.(float64)))
	}
	inventaire := map[string]int{}
	for k, v := range jsonmanagement.Get(partie, "inventory").(map[string]interface{}) {
		inventaire[k] = int(v.(float64))
	}
	return Player{
		X:                int(jsonmanagement.Get(partie, "x").(float64)),
		Y:                int(jsonmanagement.Get(partie, "y").(float64)),
		Pseudo:           jsonmanagement.Get(partie, "pseudo").(string),
		MapIn:            jsonmanagement.Get(partie, "mapIn").(string),
		Level:            int(jsonmanagement.Get(partie, "levelIn").(float64)),
		QuestIn:          int(jsonmanagement.Get(partie, "questIn").(float64)),
		Health:           int(jsonmanagement.Get(partie, "health").(float64)),
		MaxHealth:        int(jsonmanagement.Get(partie, "maxhealth").(float64)),
		Damages:          int(jsonmanagement.Get(partie, "damages").(float64)),
		Money:            int(jsonmanagement.Get(partie, "money").(float64)),
		DialogsCompleted: dialogsId,
		Inventaire:       inventaire,
	}
}

func (p *Player) Reload(partie string) {
	jsonmanagement.Update(partie, "x", p.X)
	jsonmanagement.Update(partie, "y", p.Y)
	jsonmanagement.Update(partie, "pseudo", p.Pseudo)
	jsonmanagement.Update(partie, "pseudo", p.Pseudo)
	jsonmanagement.Update(partie, "mapIn", p.MapIn)
	jsonmanagement.Update(partie, "levelIn", p.Level)
	jsonmanagement.Update(partie, "questIn", p.QuestIn)
	jsonmanagement.Update(partie, "health", p.Health)
	jsonmanagement.Update(partie, "maxhealth", p.MaxHealth)
	jsonmanagement.Update(partie, "damages", p.Damages)
	jsonmanagement.Update(partie, "money", p.Money)
	jsonmanagement.Update(partie, "dialogsCompleted", p.DialogsCompleted)
	jsonmanagement.Update(partie, "inventory", p.Inventaire)
}

func (p *Player) AddItem(partie string, nameItem string, amount int) {
	valeurNouvelle := p.Inventaire[nameItem] + amount
	p.Inventaire[nameItem] = valeurNouvelle
	jsonmanagement.Update(partie, "inventory", p.Inventaire)
}
func (p *Player) RemoveItem(partie string, nameItem string, amount int) {
	valeurNouvelle := p.Inventaire[nameItem] - amount
	p.Inventaire[nameItem] = valeurNouvelle
	jsonmanagement.Update(partie, "inventory", p.Inventaire)
}
func (p *Player) ShowInventory() {
	utils.ClearConsole()
	fmt.Println(utils.Colorize(utils.Cyan+utils.Underline, "Inventaire :"))
	for key, values := range p.Inventaire {
		if values > 0 {
			fmt.Println(key + " => " + utils.Colorize(utils.Yellow, strconv.Itoa(values)))
		}
	}
	utils.Input(">>>")
}
func (p *Player) GetAmountOfItem(keyIn string) int {
	for key, values := range p.Inventaire {
		if key == keyIn {
			return values
		}
	}
	return 0
}

func (p *Player) Heal(partie string) {
	utils.ClearConsole()
	if p.Inventaire["potion-vie"] > 0 {
		if p.Health == p.MaxHealth {
			fmt.Println(utils.Colorize(utils.Green, "Vous avez déjà 100% d'hp"))
		} else if p.Health+70 > p.MaxHealth {
			p.Health = p.MaxHealth
			fmt.Println(utils.Colorize(utils.Green, "+ 70 hp"))
		} else {
			p.Health += 70
			fmt.Println(utils.Colorize(utils.Green, "+ 70 hp"))
		}
	}
	utils.Input(">>>")
}
