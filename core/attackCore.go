package core

import (
	"fmt"
	"math/rand"
	"rpg/utils"
	"rpg/utils/jsonmanagement"
	"strconv"
	"time"
)

func Attack(enemy Enemy, player Player, partie string) (Enemy, Player, bool) {
	player.Damages = int(jsonmanagement.Get(partie, "damages").(float64))
	utils.ClearConsole()
	PrintCombatVS(enemy)
	utils.Input(">>>")
	utils.ClearConsole()
	PrintEnnemyStats(enemy)
	PrintPlayerStats(player)
	utils.Input(">>>")
	utils.ClearConsole()
	isEnnemyDead := false
	for true {
		enemy, player = TourEnnemy(enemy, player)
		utils.Input(">>>")
		if player.Health <= 0 {
			utils.ClearConsole()
			fmt.Println(utils.Colorize(utils.Purple+utils.RedBackground, "GAME OVER !"))
			player.MapIn = "spawn"
			player.X = 5
			player.Y = 5
			player.Health = player.MaxHealth
			utils.Input(">>>")
			break
		}
		utils.ClearConsole()
		PrintEnnemyStats(enemy)
		PrintPlayerStats(player)
		utils.Input(">>>")
		utils.ClearConsole()
		enemy, player = TourPlayer(enemy, player)
		if enemy.HP <= 0 {
			utils.ClearConsole()
			fmt.Println(utils.Colorize(utils.Yellow+utils.GreenBackground, "Combat réussi !"))
			money := 0
			switch enemy.FileName {
			case "gobelin":
				money += 10
			case "ingenieur":
				money += 25
			case "ogre":
				money += 70
			case "siorwel":
				money += 1000000000000
			}
			player.Money += money
			player.Reload(partie)
			fmt.Println("\n-> Vous gagnez " + utils.Colorize(utils.Yellow, money) + " dalis !\n")
			isEnnemyDead = true
			utils.Input(">>>")
			break
		}
		utils.Input(">>>")
		utils.ClearConsole()
	}

	return enemy, player, isEnnemyDead

}

func TourPlayer(enemy Enemy, player Player) (Enemy, Player) {
	fmt.Println("C'est le tour de " + utils.Colorize(utils.Cyan, player.Pseudo))
	utils.Timeout(2)
	var score int
	if rand.Intn(2) == 1 {
		fmt.Println(utils.Colorize(utils.Cyan, player.Pseudo) + " vous tentez d'invoquer un sortilège ! 🔮")
		utils.Input(">>>")
		score = JeuxRetyping()
	} else {
		fmt.Println(utils.Colorize(utils.Cyan, player.Pseudo) + " vous tentez de donner un coups d'épée ! ⚔")
		utils.Input(">>>")
		score = JeuxCalculMental()
	}
	if score > 0 {
		damages := score * player.Damages
		fmt.Println(utils.Colorize(utils.GreenBackground, "\nVous infligez "+utils.Colorize(utils.Black, strconv.Itoa(damages))) + utils.Colorize(utils.GreenBackground, " dégats à "+utils.Colorize(utils.Black, enemy.RealName)))
		enemy.HP -= damages
		utils.Input(">>>")
	} else {
		fmt.Println("Vous n'avez infligé aucun dégats à " + utils.Colorize(utils.Red+utils.Bold, enemy.RealName))
	}
	return enemy, player
}

func JeuxCalculMental() int {
	for i := 3; i > 0; i-- {
		utils.ClearConsole()
		fmt.Println("Réussissez le plus de calculs pour augmenter votre combo !")
		fmt.Printf("Début du jeux dans %v secondes...", utils.Colorize(utils.Yellow, i))
		utils.Timeout(1)
	}
	utils.ClearConsole()
	score := 0
	timer := time.NewTimer(10 * time.Second)
	startTime := time.Now()
	for {
		select {
		case <-timer.C:
			return score
		default:
			num1 := rand.Intn(10+score*10) + 1
			num2 := rand.Intn(10+score*10) + 1

			elapsed := time.Since(startTime)
			remainingTime := int(10 - elapsed.Seconds())
			if remainingTime < 0 {
				return score
			}
			fmt.Printf("Il vous reste %v secondes.\n\nCalculez %v + %v : ", utils.Colorize(utils.Blue, remainingTime), utils.Colorize(utils.Purple, num1), utils.Colorize(utils.Purple, num2))

			inp, _ := strconv.Atoi(utils.Input("\n>>>"))
			utils.ClearConsole()
			if inp == num1+num2 {
				score++
				fmt.Printf("Vous réussissez vos coups d'épée. ✅\nCombo %v", utils.Colorize(utils.Yellow, score)+" coups ! ⚔\n\n")
			} else {
				fmt.Println("Vous ratez votre coups d'épée !\nVous avez fait un combo de " + utils.Colorize(utils.Yellow, score) + " coups ! ⚔")
				return score
			}
		}
	}
}

func JeuxRetyping() int {
	for i := 3; i > 0; i-- {
		utils.ClearConsole()
		fmt.Println("Prononcez le plus d'incantations pour augmenter votre combo !")
		fmt.Printf("Début du jeux dans %v secondes...", utils.Colorize(utils.Yellow, i))
		utils.Timeout(1)
	}
	utils.ClearConsole()
	score := 0
	timer := time.NewTimer(15 * time.Second)
	startTime := time.Now()
	for {
		select {
		case <-timer.C:
			return score
		default:
			word := generateRandomWord(score)
			elapsed := time.Since(startTime)
			remainingTime := int(10 - elapsed.Seconds())
			if remainingTime < 0 {
				return score
			}
			fmt.Printf("Il vous reste %v secondes.\n\nPrononcez le mot suivant : %s\n", utils.Colorize(utils.Blue, remainingTime), utils.Colorize(utils.Purple, word))

			inp := utils.Input(">>>")

			if inp == word {
				score++
				utils.ClearConsole()
				fmt.Printf("Vous réussissez votre incantation. ✅\nCombo %v\n", utils.Colorize(utils.Yellow, score)+" incantations ! 🔮")
			} else {
				fmt.Println("Vous avez mal prononcé l'incantation !\nVous avez fait un combo de " + utils.Colorize(utils.Yellow, score) + " incantations ! 🔮")
				return score
			}
		}
	}
}
func generateRandomWord(scoreIn int) string {
	letters := "abcdefghijklmnopqrstuvwxyz1234567890^$*./?&"
	wordLength := rand.Intn(scoreIn+3) + 1
	var word string
	for i := 0; i < wordLength; i++ {
		randIndex := rand.Intn(len(letters))
		word += string(letters[randIndex])
	}
	return word
}

func TourEnnemy(enemy Enemy, player Player) (Enemy, Player) {
	fmt.Println("C'est le tour de " + utils.Colorize(utils.Red, enemy.RealName))
	utils.Timeout(2)
	fmt.Println(utils.Colorize(utils.Red, enemy.RealName) + " vous attaque ! 🤺")
	utils.Timeout(1)
	if rand.Intn(100) > enemy.RateAttack {
		fmt.Println(utils.Colorize(utils.Red, enemy.RealName) + " a raté son coups !")
	} else {
		degats := enemy.Damages - 10 + rand.Intn(20)
		fmt.Println(utils.Colorize(utils.Red, enemy.RealName) + " vous a infligé " + utils.Colorize(utils.Red+utils.Bold, strconv.Itoa(degats)) + " dégats")
		player.Health -= degats
	}
	return enemy, player
}

func PrintEnnemyStats(enemy Enemy) {
	fmt.Println("##########################")
	fmt.Println(utils.Colorize(utils.Red, enemy.RealName) + "\n")
	fmt.Println("-     Dégats : " + utils.Colorize(utils.Yellow, strconv.Itoa(enemy.Damages)))
	fmt.Println("-     HP     : " + utils.Colorize(utils.Green, strconv.Itoa(enemy.HP)) + "\n")
	fmt.Println("##########################")
}
func PrintCombatVS(enemy Enemy) {
	fmt.Println("########- " + utils.Colorize(utils.Red, "COMBAT") + " -########")
	fmt.Println("#                        #")
	fmt.Println("#     " + utils.Colorize(utils.Cyan, "🧝") + "    VS     " + utils.Colorize(utils.Red, enemy.Smiley) + "    #")
	fmt.Println("#                        #")
	fmt.Println("##########################")
}

func PrintPlayerStats(player Player) {
	fmt.Println("##########################")
	fmt.Println(utils.Colorize(utils.Cyan, player.Pseudo) + "\n")
	fmt.Println("-     Dégats : " + utils.Colorize(utils.Yellow, strconv.Itoa(player.Damages)))
	fmt.Println("-     HP     : " + utils.Colorize(utils.Green, strconv.Itoa(player.Health)) + "\n")
	fmt.Println("##########################")
}
