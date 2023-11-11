package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"rpg/core"
	"rpg/utils"
	"rpg/utils/dialogsmanagement"
	"rpg/utils/jsonmanagement"
	"rpg/utils/levelmanagement"
	"rpg/utils/marketmanagement"
	"rpg/utils/questmanagement"
	"slices"
	"strconv"
	"strings"
)

func main() {
	utils.ClearConsole()
	utils.Timeout(0.3)
	title := []string{" ", " ", "█", "█", "╗", " ", " ", " ", " ", " ", "█", "█", "╗", "█", "█", "╗", " ", " ", " ", "█", "█", "╗", "█", "█", "╗", "█", "█", "█", "█", "█", "█", "╗", " ", "█", "█", "╗", " ", " ", " ", "█", "█", "╗", "█", "█", "█", "█", "█", "█", "█", "╗", "\n", " ", " ", "█", "█", "║", " ", " ", " ", " ", " ", "█", "█", "║", "█", "█", "║", " ", " ", " ", "█", "█", "║", "█", "█", "║", "█", "█", "╔", "═", "═", "█", "█", "╗", "█", "█", "║", " ", " ", " ", "█", "█", "║", "█", "█", "╔", "═", "═", "═", "═", "╝", "\n", " ", " ", "█", "█", "║", " ", " ", " ", " ", " ", "█", "█", "║", "█", "█", "║", " ", " ", " ", "█", "█", "║", "█", "█", "║", "█", "█", "║", " ", " ", "█", "█", "║", "█", "█", "║", " ", " ", " ", "█", "█", "║", "█", "█", "█", "█", "█", "█", "█", "╗", "\n", " ", " ", "█", "█", "║", " ", " ", " ", " ", " ", "█", "█", "║", "╚", "█", "█", "╗", " ", "█", "█", "╔", "╝", "█", "█", "║", "█", "█", "║", " ", " ", "█", "█", "║", "█", "█", "║", " ", "  ", "█", "█", "║", "╚", "═", "═", "═", "═", "█", "█", "║", "\n", " ", " ", "█", "█", "█", "█", "█", "█", "█", "╗", "█", "█", "║", " ", "╚", "█", "█", "█", "█", "╔", "╝", " ", "█", "█", "║", "█", "█", "█", "█", "█", "█", "╔", "╝", "╚", "█", "█", "█", "█", "█", "█", "╔", "╝", "█", "█", "█", "█", "█", "█", "█", "║", "\n", " ", " ", "╚", "═", "═", "═", "═", "═", "═", "╝", "╚", "═", "╝", " ", " ", "╚", "═", "═", "═", "╝", " ", " ", "╚", "═", "╝", "╚", "═", "═", "═", "═", "═", "╝", " ", " ", "╚", "═", "═", "═", "═", "═", "╝", " ", "╚", "═", "═", "═", ")", "\n"}
	for i := 0; i < len(title); i++ {
		fmt.Print(utils.Colorize(utils.Blue, title[i]))
		utils.Timeout(0.000000001)
	}
	utils.Timeout(1)
	fmt.Println(utils.Colorize(utils.Yellow, "1") + " > Charger une partie")
	fmt.Println(utils.Colorize(utils.Yellow, "2") + " > Créer une nouvelle partie")
	fmt.Println(utils.Colorize(utils.Yellow, "3") + " > Quitter le jeu\n")

	for inp := utils.Input(">>"); true; inp = utils.Input(">>") {
		switch inp {
		case "1":
			ChargerMenuPartie()
		case "2":
			CreerPartie()
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Veuillez réessayer !")
		}
	}
}

func CreerPartie() {
	utils.ClearConsole()
	fmt.Println("######################################")
	fmt.Print("\n\tPseudo : ")
	pseudo := utils.Input("")
	fmt.Println("\n######################################")
	utils.Timeout(1)
	utils.ClearConsole()
	fmt.Print(utils.Colorize(utils.Yellow, "Partie en cours de création "))
	partie := "/parties/" + pseudo
	jsonmanagement.CreateJson(partie)
	utils.Timeout(0.5)
	fmt.Print(utils.Colorize(utils.Yellow, "."))
	jsonmanagement.Add(partie, "damages", 5)
	jsonmanagement.Add(partie, "dialogsCompleted", []string{})
	jsonmanagement.Add(partie, "health", 100)
	jsonmanagement.Add(partie, "maxhealth", 100)
	jsonmanagement.Add(partie, "levelIn", 0)
	jsonmanagement.Add(partie, "money", 0)
	utils.Timeout(0.5)
	fmt.Print(utils.Colorize(utils.Yellow, "."))
	jsonmanagement.Add(partie, "mapIn", "spawn")
	jsonmanagement.Add(partie, "pseudo", pseudo)
	jsonmanagement.Add(partie, "questIn", 0)
	jsonmanagement.Add(partie, "inventory", map[string]int{
		"armure-cendre":   0,
		"armure-lividus":  0,
		"epee-cendre":     0,
		"epee-lividus":    0,
		"marteau-siorwel": 0,
		"potion-vie":      0,
	})
	jsonmanagement.Add(partie, "x", 5)
	jsonmanagement.Add(partie, "y", 5)
	utils.Timeout(0.5)
	fmt.Print(utils.Colorize(utils.Yellow, "."))
	Tuto(pseudo)
}

func Tuto(pseudo string) {
	utils.ClearConsole()
	fmt.Println("Bienvenue " + pseudo + " dans " + utils.Colorize(utils.Blue+utils.Bold, "Lividus") + " !")
	utils.Input("\n>>> (Appuyez sur la touche \"Entrée\")")
	fmt.Println("Pour te déplacer dans ce monde merveilleux il te suffit d'appuyer sur ...")
	fmt.Println(utils.Colorize(utils.Red, "Z -> Haut"))
	fmt.Println(utils.Colorize(utils.Red, "S -> Bas"))
	fmt.Println(utils.Colorize(utils.Red, "D -> Droite"))
	fmt.Println(utils.Colorize(utils.Red, "Q -> Gauche"))
	utils.Input("\n>>>")
	fmt.Println("Vous incarnez un voyageur venu d'ailleurs qui a été retrouvé inconscient au Bosquet de Lividus.")
	utils.Input("\n>>>")
	fmt.Println("Ramené par les soldats de la garde de Lividus, vous avez été placé à l'infirmerie royale à cause d'une mystérieuse marque se trouvant sur votre bras !")
	utils.Input("\n>>>")
	fmt.Println("A votre réveil parlez au " + utils.Colorize(utils.Purple, "Docteur Mundo") + " !")
	utils.Input("\n>>>")
	utils.ClearConsole()
	z := []string{"Z", "Z", "Z", "z", "z", "z", "z", "z", "z", "z", "z"}
	for _, v := range z {
		fmt.Print(utils.Colorize(utils.Cyan, v))
		utils.Timeout(0.2)
	}
	utils.Timeout(0.5)
	fmt.Println()
	for _, v := range z[:6] {
		fmt.Print(utils.Colorize(utils.Cyan, v))
		utils.Timeout(0.1)
	}
	utils.Timeout(0.3)
	fmt.Println()
	for _, v := range z[:4] {
		fmt.Print(utils.Colorize(utils.Cyan, v))
		utils.Timeout(0.05)
	}
	utils.Timeout(0.1)
	fmt.Println(utils.Colorize(utils.Cyan, "\n\n<Se réveil>"))
	utils.Timeout(2)
	ChargerPartie(pseudo)
}

func ChargerMenuPartie() {
	utils.ClearConsole()
	files, _ := ioutil.ReadDir("./storage/saves/parties/")
	var fileNames []string
	if len(files) == 0 {
		fmt.Println("Vous avez actuellement aucune partie à charger !\nNous allons donc créer votre première partie !")
		CreerPartie()
		return
	}
	fmt.Println(utils.Colorize(utils.Underline+utils.Bold, "Quelle partie voulez vous charger !\n"))
	for i, v := range files {
		fileName := strings.Split(v.Name(), ".")[0]
		fmt.Printf("%v > %v", utils.Colorize(utils.Yellow, i+1), fileName)
		fmt.Println()
		fileNames = append(fileNames, fileName)
	}
	println()
	for inp := utils.Input(">>"); true; inp = utils.Input(">>") {
		if len(inp) == 0 {
			fmt.Println("Veuillez réessayer !")
			continue
		}
		inpInt, err := strconv.Atoi(string(inp[0]))
		if err == nil {
			if len(fileNames) >= inpInt && inpInt > 0 {
				utils.ClearConsole()
				fmt.Print(utils.Colorize(utils.Yellow, "Partie en cours de chargement "))
				utils.Timeout(0.5)
				fmt.Print(utils.Colorize(utils.Yellow, "."))
				utils.Timeout(0.5)
				fmt.Print(utils.Colorize(utils.Yellow, "."))
				utils.Timeout(0.5)
				fmt.Print(utils.Colorize(utils.Yellow, "."))
				ChargerPartie(fileNames[inpInt-1])
			}
		}
	}
}

func ChargerPartie(partiePseudo string) {
	Game("parties/" + partiePseudo)
}

// //// GAME //////
var PlayerIn = core.Player{}
var PartieIn = ""
var QuestLogIn []string
var QuestIn questmanagement.Quest

func Game(partiePseudo string) {
	PartieIn = partiePseudo
	PlayerIn = PlayerIn.Load(PartieIn)
	QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, int(jsonmanagement.Get(partiePseudo, "questIn").(float64)))
	PlayerIn.Reload(PartieIn)
	// Boucle de jeu
	player, name, nameFile, level, levelBuff, ennemis := LoadLevel(PlayerIn.MapIn, PlayerIn.X, PlayerIn.Y)
	for {
		utils.ClearConsole()

		//NEW QUEST MANAGER
		if QuestIn.QuestNum == 0 && slices.Contains(PlayerIn.DialogsCompleted, 4) {
			QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, 1)
			PlayerIn.QuestIn = QuestIn.QuestNum
			PlayerIn.Reload(PartieIn)
			continue
		} else if QuestIn.QuestNum == 1 && slices.Contains(PlayerIn.DialogsCompleted, 6) {
			PlayerIn = questmanagement.ShowRewards(PlayerIn, QuestIn)
			QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, 2)
			PlayerIn.QuestIn = QuestIn.QuestNum
			PlayerIn.Reload(PartieIn)
			continue
		} else if QuestIn.QuestNum == 2 {
			if questmanagement.AllObjectivesItems(QuestIn, PlayerIn) {
				PlayerIn = questmanagement.ShowRewards(PlayerIn, QuestIn)
				QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, 3)
				PlayerIn.QuestIn = QuestIn.QuestNum
				PlayerIn.Reload(PartieIn)
				continue
			}
		} else if QuestIn.QuestNum == 3 {
			if questmanagement.AllObjectivesKills(QuestIn, PlayerIn) {
				PlayerIn = questmanagement.ShowRewards(PlayerIn, QuestIn)
				QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, 4)
				PlayerIn.QuestIn = QuestIn.QuestNum
				PlayerIn.Reload(PartieIn)
				continue
			}
		} else if QuestIn.QuestNum == 4 && slices.Contains(PlayerIn.DialogsCompleted, 10) {
			PlayerIn = questmanagement.ShowRewards(PlayerIn, QuestIn)
			QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, 5)
			PlayerIn.QuestIn = QuestIn.QuestNum
			PlayerIn.Reload(PartieIn)
			continue
		} else if QuestIn.QuestNum == 5 {
			if questmanagement.AllObjectivesItems(QuestIn, PlayerIn) {
				PlayerIn = questmanagement.ShowRewards(PlayerIn, QuestIn)
				QuestLogIn, QuestIn = questmanagement.GiveQuest(PartieIn, 6)
				PlayerIn.QuestIn = QuestIn.QuestNum
				PlayerIn.Reload(PartieIn)
				continue
			}
		} else if QuestIn.QuestNum == 6 {
			if questmanagement.AllObjectivesKills(QuestIn, PlayerIn) {
				utils.ClearConsole()
				fmt.Println("Vous avez terminé le jeu !")
				utils.Timeout(2)
				fmt.Println("Merci d'avoir joué")
				utils.Timeout(1)
				fmt.Println("\n\nJeu réalisé par Manel, Tom, Clément et Titouan")
				os.Exit(0)
			}
		}
		for i := range ennemis {
			level[ennemis[i].Y][ennemis[i].X] = map[string]string{utils.Colorize(utils.Red, ennemis[i].Smiley): "mob/" + ennemis[i].FileName}
		}
		level[player.Y][player.X] = map[string]string{utils.Colorize(utils.Cyan, "🧝"): "player"}
		levelmanagement.AfficherLevel(name, level, QuestLogIn, QuestIn, PlayerIn)
		// Demander au joueur de saisir une direction
		input := utils.Input(">>> ")
		// Effacer la position actuelle du personnage
		level[player.Y][player.X] = levelBuff[player.Y][player.X]
		for i := range ennemis {
			level[ennemis[i].Y][ennemis[i].X] = levelBuff[ennemis[i].Y][ennemis[i].X]
		}
		isMarketShowned := false
		// Déplacer le personnage en fonction de la direction
		switch input {
		case "q":
			if player.X-1 >= 0 {
				if !levelmanagement.IsBlockWay(levelBuff, player.X-1, player.Y) {
					isNPC, dialogId := levelmanagement.IsNpc(level[player.Y][player.X-1])
					isMarket, marketName := levelmanagement.IsMarket(level[player.Y][player.X-1])
					if isNPC {
						if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
							PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
							PlayerIn.Reload(PartieIn)
						}
					} else if isMarket {
						player = marketmanagement.MarketShow(marketName, player, PartieIn)
						player.Reload(PartieIn)
						isMarketShowned = true
					} else {
						player.X--
					}
				}
			}
			break
		case "z":
			if player.Y-1 >= 0 {
				if !levelmanagement.IsBlockWay(levelBuff, player.X, player.Y-1) {
					isNPC, dialogId := levelmanagement.IsNpc(level[player.Y-1][player.X])
					isMarket, marketName := levelmanagement.IsMarket(level[player.Y-1][player.X])
					if isNPC {
						if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
							PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
							PlayerIn.Reload(PartieIn)
						}
					} else if isMarket {
						player = marketmanagement.MarketShow(marketName, player, PartieIn)
						player.Reload(PartieIn)
						isMarketShowned = true
					} else {
						player.Y--
					}
				}
			}
			break
		case "d":
			if player.X+1 < len(level[0]) {
				if !levelmanagement.IsBlockWay(levelBuff, player.X+1, player.Y) {
					isNPC, dialogId := levelmanagement.IsNpc(level[player.Y][player.X+1])
					isMarket, marketName := levelmanagement.IsMarket(level[player.Y][player.X+1])
					if isNPC {
						if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
							PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
							PlayerIn.Reload(PartieIn)
						}
					} else if isMarket {
						player = marketmanagement.MarketShow(marketName, player, PartieIn)
						player.Reload(PartieIn)
						isMarketShowned = true
					} else {
						player.X++
					}
				}
			}
			break
		case "s":
			if player.Y+1 < len(level) {
				if !levelmanagement.IsBlockWay(levelBuff, player.X, player.Y+1) {
					isNPC, dialogId := levelmanagement.IsNpc(level[player.Y+1][player.X])
					isMarket, marketName := levelmanagement.IsMarket(level[player.Y+1][player.X])
					if isNPC {
						if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
							PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
							PlayerIn.Reload(PartieIn)
						}
					} else if isMarket {
						player = marketmanagement.MarketShow(marketName, player, PartieIn)
						player.Reload(PartieIn)
						isMarketShowned = true
					} else {
						player.Y++
					}
				}
			}
			break
		case "quitter":
			jsonmanagement.Update(PartieIn, "mapIn", nameFile)
			jsonmanagement.Update(PartieIn, "x", player.X)
			jsonmanagement.Update(PartieIn, "y", player.Y)
			os.Exit(0)
		case "heal":
			player.Heal(PartieIn)
		case "level":
			utils.ClearConsole()
			fmt.Println("Vous êtes au niveau " + utils.Colorize(utils.Yellow, player.Level) + " !")
			utils.Input(">>>")
		case "inventaire":
			player.ShowInventory()
			utils.Input(">>>")
		case "pause":
			utils.ClearConsole()
			utils.Input(">>> (Appuyez sur la touche Entrée pour lever la pause)")
		case "help":
			utils.ClearConsole()
			fmt.Println(utils.Colorize(utils.Underline, "Liste des commandes :"))
			fmt.Println("\n- z, q, s, d => déplacements")
			fmt.Println("- heal	      => utilise une potion de vie")
			fmt.Println("- level	  => affiche le niveau actuel")
			fmt.Println("- pause	  => mets en pause le jeu")
			fmt.Println("- quitter	  => quitte le jeu")
			utils.Input(">>> (Appuyez sur la touche Entrée pour lever la pause)")
		}
		//Gérer le type de structure/ présence ennemis, npc
		//PROCHAIN NIVEAU
		isNextLevel, levelName, xNextSpawn, yNextSpawn := levelmanagement.IsNextPiece(levelBuff, level[player.Y][player.X])
		if isNextLevel {
			player, name, nameFile, level, levelBuff, ennemis = LoadLevel(levelName, xNextSpawn, yNextSpawn)
			continue
		}
		//NPC
		isNPC, dialogId := levelmanagement.IsNpc(level[player.Y][player.X])
		if isNPC {
			if dialogsmanagement.CanReadDialog(dialogId, PlayerIn.DialogsCompleted) {
				PlayerIn = dialogsmanagement.ReadDialog(dialogId, PlayerIn, PartieIn)
				PlayerIn.Reload(PartieIn)
			}
			continue
		}
		isMarket, marketName := levelmanagement.IsMarket(level[player.Y][player.X])
		if isMarket {
			player = marketmanagement.MarketShow(marketName, player, PartieIn)
			player.Reload(PartieIn)
			isMarketShowned = true
		}
		if isMarketShowned {
			for k, v := range player.Inventaire {
				switch k {
				case "epee-cendre":
					if v > 0 {
						jsonmanagement.Update(PartieIn, "damages", 50)
						player.Damages = 50
						PlayerIn.Damages = 50
					}
				case "epee-lividus":
					if v > 0 {
						jsonmanagement.Update(PartieIn, "damages", 30)
						player.Damages = 30
						PlayerIn.Damages = 30
					}
				}
				switch k {
				case "armure-cendre":
					if v > 0 {
						jsonmanagement.Update(PartieIn, "health", 225)
						jsonmanagement.Update(PartieIn, "maxhealth", 225)
						player.Health = 400
						player.MaxHealth = 400
						PlayerIn.Health = 400
						PlayerIn.MaxHealth = 400
					}
				case "armure-lividus":
					if v > 0 {
						jsonmanagement.Update(PartieIn, "health", 150)
						jsonmanagement.Update(PartieIn, "maxhealth", 150)
						player.Health = 150
						player.MaxHealth = 150
						PlayerIn.Health = 150
						PlayerIn.MaxHealth = 150
					}
				}
			}
		}

		player.Reload(PartieIn)

		//ENNEMIS DEPLACEMENTS
		for i := range ennemis {
			if rand.Intn(2) == 1 {
				ennemis[i] = ennemis[i].WalkToTarget(player, levelBuff)
			}
			if ennemis[i].X == player.X && ennemis[i].Y == player.Y {
				isEnnemyDead := false
				ennemis[i], player, isEnnemyDead = core.Attack(ennemis[i], player, PartieIn)
				if isEnnemyDead {
					if QuestIn.IsKillQuest() {
						if utils.ContainsKey(QuestIn.Content.(questmanagement.ContentKill).Targets, ennemis[i].FileName) {
							QuestIn.Content.(questmanagement.ContentKill).TargetsIn[ennemis[i].FileName]++
						}
					}
					ennemis = core.RemoveEnemy(ennemis, ennemis[i])
					break
				} else {
					player.Reload(PartieIn)
					player, name, nameFile, level, levelBuff, ennemis = LoadLevel("spawn", 5, 5)
					continue
				}
			}
			level[ennemis[i].Y][ennemis[i].X] = map[string]string{utils.Colorize(utils.Red, ennemis[i].Smiley): "mob/" + ennemis[i].FileName}
		}
		// Mettre à jour la nouvelle position du personnage
		level[player.Y][player.X] = map[string]string{utils.Colorize(utils.Cyan, "🧝"): "player"}
	}
}
func LoadLevel(levelName string, xSpawn int, ySpawn int) (core.Player, string, string, [][]map[string]string, [][]map[string]string, []core.Enemy) {
	// Chargement du level en cours
	name, level, _ := levelmanagement.ReadLevel(levelName)
	// Pour garder toujours une version orginelle de la carte on clone dans une autre variable la carte
	levelBase := levelmanagement.CloneLevel(level)
	// Création du personnage
	PlayerIn.X = xSpawn
	PlayerIn.Y = ySpawn
	if PlayerIn.Y >= len(level) {
		PlayerIn.Y = len(level) / 2
	}
	if PlayerIn.X >= len(level[0]) {
		PlayerIn.X = len(level) / 2
	}
	level[PlayerIn.Y][PlayerIn.X] = map[string]string{utils.Colorize(utils.Cyan, "🧝"): "player"}
	jsonmanagement.Update(PartieIn, "mapIn", levelName)
	jsonmanagement.Update(PartieIn, "x", xSpawn)
	jsonmanagement.Update(PartieIn, "y", ySpawn)
	return PlayerIn, name, levelName, level, levelBase, core.LoadMobs(levelName)
}
