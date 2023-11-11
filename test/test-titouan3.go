package main

import (
	"fmt"
	"os"
	"rpg/core"
	"rpg/utils"
	"rpg/utils/dialogsmanagement"
	"rpg/utils/jsonmanagement"
	"rpg/utils/levelmanagement"
	"rpg/utils/questmanagement"
	"slices"
)

var PlayerIn = core.Player{}
var PartieIn = "Partie1"
var QuestLogIn []string
var QuestIn questmanagement.Quest

func main() {
	PlayerIn = PlayerIn.Load(PartieIn)
	QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, 0)
	PlayerIn.Reload(PartieIn)
	// Boucle de jeu

	player, name, nameFile, level, levelBuff := LoadLevel(PlayerIn.MapIn, PlayerIn.X, PlayerIn.Y)
	var ennemis = []core.Enemy{
		//{
		//	X:          7,
		//	Y:          6,
		//	FileName:   "gobelin",
		//	RealName:   "Gobelin Cendré",
		//	Damages:    10,
		//	HP:         60,
		//	RateAttack: 70,
		//	Smiley:     "🧌 ",
		//},
	}

	for {
		utils.ClearConsole()
		// Enemy

		for _, enemy := range ennemis {
			fmt.Println(enemy.Y, enemy.X)
			fmt.Println(level[enemy.Y][enemy.X])
		}

		//NEW QUEST MANAGER
		if QuestIn.QuestNum == 0 && slices.Contains(PlayerIn.DialogsCompleted, 4) {
			QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, 1)
			PlayerIn.QuestIn = QuestIn.QuestNum
			PlayerIn.Reload(PartieIn)
			continue
		}
		if QuestIn.QuestNum == 1 && slices.Contains(PlayerIn.DialogsCompleted, 6) {
			QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, 2)
			PlayerIn.QuestIn = QuestIn.QuestNum
			PlayerIn.Reload(PartieIn)
			continue
		}

		levelmanagement.AfficherLevel(name, level, QuestLogIn, QuestIn, PlayerIn)

		// Demander au joueur de saisir une direction
		input := utils.Input(">>> ")

		// Effacer la position actuelle du personnage
		level[player.Y][player.X] = levelBuff[player.Y][player.X]
		for i := range ennemis {
			level[ennemis[i].Y][ennemis[i].X] = levelBuff[ennemis[i].Y][ennemis[i].X]
		}
		// Déplacer le personnage en fonction de la direction
		switch input {
		case "q":
			if player.X-1 >= 0 {
				if !levelmanagement.IsBlockWay(levelBuff, player.X-1, player.Y) {
					isNPC, dialogId := levelmanagement.IsNpc(levelBuff, level[player.Y][player.X-1])
					if isNPC {
						if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
							PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
							PlayerIn.Reload(PartieIn)
						}
					} else {
						player.X--
					}
				}
			}
			break
		case "z":
			if player.Y-1 >= 0 {
				if !levelmanagement.IsBlockWay(levelBuff, player.X, player.Y-1) {
					isNPC, dialogId := levelmanagement.IsNpc(levelBuff, level[player.Y-1][player.X])
					if isNPC {
						if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
							PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
							PlayerIn.Reload(PartieIn)
						}
					} else {
						player.Y--
					}
				}
			}
			break
		case "d":
			if player.X+1 < len(level[0]) {
				if !levelmanagement.IsBlockWay(levelBuff, player.X+1, player.Y) {
					isNPC, dialogId := levelmanagement.IsNpc(levelBuff, level[player.Y][player.X+1])
					if isNPC {
						if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
							PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
							PlayerIn.Reload(PartieIn)
						}
					} else {
						player.X++
					}
				}
			}
			break
		case "s":
			if player.Y+1 < len(level) {
				if !levelmanagement.IsBlockWay(levelBuff, player.X, player.Y+1) {
					isNPC, dialogId := levelmanagement.IsNpc(levelBuff, level[player.Y+1][player.X])
					if isNPC {
						if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
							PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
							PlayerIn.Reload(PartieIn)
						}
					} else {
						player.Y++
					}
				}

			}
			break
		case "l":
			jsonmanagement.Update(PartieIn, "mapIn", nameFile)
			jsonmanagement.Update(PartieIn, "x", player.X)
			jsonmanagement.Update(PartieIn, "y", player.Y)
			os.Exit(0)
		}

		//Gérer le type de structure/ présence ennemis, npc
		//PROCHAIN NIVEAU
		isNextLevel, levelName, xNextSpawn, yNextSpawn := levelmanagement.IsNextPiece(levelBuff, level[player.Y][player.X])
		if isNextLevel {
			player, name, nameFile, level, levelBuff = LoadLevel(levelName, xNextSpawn, yNextSpawn)
			continue
		}

		//NPC
		isNPC, dialogId := levelmanagement.IsNpc(levelBuff, level[player.Y][player.X])
		if isNPC {
			if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
				PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
				PlayerIn.Reload(PartieIn)
			}
			continue
		}

		for i := range ennemis {
			ennemis[i] = ennemis[i].WalkToTarget(player, levelBuff)
			if ennemis[i].X == player.X && ennemis[i].Y == player.Y {
				ennemis[i], PlayerIn = core.Attack(ennemis[i], PlayerIn)
				PlayerIn.Reload(PartieIn)
			}
			level[ennemis[i].Y][ennemis[i].X] = map[string]string{utils.Colorize(utils.Red, ennemis[i].Smiley): ennemis[i].FileName}
		}
		// Mettre à jour la nouvelle position du personnage
		level[player.Y][player.X] = map[string]string{utils.Colorize(utils.Cyan, "🧝"): "player"}
	}
}
func LoadLevel(levelName string, xSpawn int, ySpawn int) (core.Player, string, string, [][]map[string]string, [][]map[string]string) {
	// Chargement du level en cours

	name, level, _ := levelmanagement.ReadLevel(levelName)

	// Pour garder toujours une version orginelle de la carte on clone dans une autre variable la carte
	levelBase := levelmanagement.CloneLevel(level)

	// Création du personnage
	PlayerIn.X = xSpawn
	PlayerIn.Y = ySpawn
	level[PlayerIn.Y][PlayerIn.X] = map[string]string{utils.Colorize(utils.Cyan, "🧝"): "player"}
	jsonmanagement.Update(PartieIn, "mapIn", levelName)
	jsonmanagement.Update(PartieIn, "x", xSpawn)
	jsonmanagement.Update(PartieIn, "y", ySpawn)
	return PlayerIn, name, levelName, level, levelBase
}
