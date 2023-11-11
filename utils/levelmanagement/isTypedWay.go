package levelmanagement

import (
	"fmt"
	"strconv"
)
import "strings"

func IsTypedWay(level [][]map[string]string, X int, Y int, typeIn string) bool {
	for _, action := range level[Y][X] {
		return action == typeIn
	}
	return false
}

func IsBlockWay(level [][]map[string]string, X int, Y int) bool {
	return IsTypedWay(level, X, Y, "b")
}
func IsNpc(nextEmplacement map[string]string) (bool, int) {
	for _, action := range nextEmplacement {
		if len(action) > 3 {
			fmt.Println(action[:3])

			if action[:3] == "npc" {
				actionRow := strings.Split(action, "/")
				dialogId, _ := strconv.Atoi(actionRow[1])
				return true, dialogId
			}
		}
	}
	return false, 0
}
func IsMarket(nextEmplacement map[string]string) (bool, string) {
	for _, action := range nextEmplacement {
		if len(action) > 6 {

			if action[:6] == "market" {
				actionRow := strings.Split(action, "/")
				marketName := actionRow[1]
				return true, marketName
			}
		}
	}
	return false, ""
}
func IsEnemy(level [][]map[string]string, X int, Y int) bool {
	return IsTypedWay(level, X, Y, "e")
}
func IsNextPiece(level [][]map[string]string, nextEmplacement map[string]string) (bool, string, int, int) {
	for _, action := range nextEmplacement {
		if len(action) > 3 {
			if action[:2] == "go" {
				actionRow := strings.Split(action, "/")
				x, _ := strconv.Atoi(actionRow[2])
				y, _ := strconv.Atoi(actionRow[3])
				return true, actionRow[1], x, y
			}
		}
	}
	return false, "", 0, 0
}
