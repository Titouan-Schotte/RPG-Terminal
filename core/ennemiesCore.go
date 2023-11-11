package core

import (
	"container/heap"
	"rpg/utils"
	"rpg/utils/jsonmanagement"
)

type Enemy struct {
	X, Y                       int
	FileName, Smiley, RealName string
	Damages                    int
	HP                         int
	RateAttack                 int
}

func (e Enemy) WalkToTarget(player Player, mapBuff [][]map[string]string) Enemy {
	shortestPath := findShortestPathWithAStar(mapBuff, e.X, e.Y, player.X, player.Y)
	if len(shortestPath) <= 1 {
		return e
	}

	e.X = shortestPath[1].X
	e.Y = shortestPath[1].Y
	return e
}

type Coord struct {
	X, Y int
}

type Node struct {
	Coord   Coord
	Parent  *Node
	G, H, F int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func findShortestPathWithAStar(mapData [][]map[string]string, startX, startY, targetX, targetY int) []Coord {
	movements := []Coord{
		{0, -1}, // Haut
		{0, 1},  // Bas
		{-1, 0}, // Gauche
		{1, 0},  // Droite
	}

	openList := make(PriorityQueue, 0)
	closedList := make(map[Coord]bool)
	nodeMap := make(map[Coord]*Node)

	startNode := &Node{Coord: Coord{startX, startY}}
	startNode.G = 0
	startNode.H = utils.Abs(targetX-startX) + utils.Abs(targetY-startY)
	startNode.F = startNode.G + startNode.H

	heap.Push(&openList, startNode)
	nodeMap[startNode.Coord] = startNode

	for len(openList) > 0 {
		currentNode := heap.Pop(&openList).(*Node)

		if currentNode.Coord.X == targetX && currentNode.Coord.Y == targetY {
			path := make([]Coord, 0)
			for currentNode != nil {
				path = append(path, currentNode.Coord)
				currentNode = currentNode.Parent
			}
			reversePath(path)
			return path
		}

		closedList[currentNode.Coord] = true

		for _, move := range movements {
			newX := currentNode.Coord.X + move.X
			newY := currentNode.Coord.Y + move.Y
			newCoord := Coord{newX, newY}

			if newX >= 0 && newX < len(mapData) && newY >= 0 && newY < len(mapData[0]) {
				emoji := utils.GetFirstKey(mapData[newX][newY])

				if emoji != "b" && !closedList[newCoord] {
					newG := currentNode.G + 1
					inOpenList := false

					existingNode, exists := nodeMap[newCoord]
					if exists {
						inOpenList = true
					}

					if !inOpenList || newG < existingNode.G {
						if !inOpenList {
							newNode := &Node{
								Coord:  newCoord,
								Parent: currentNode,
								G:      newG,
								H:      utils.Abs(targetX-newX) + utils.Abs(targetY-newY),
							}
							newNode.F = newNode.G + newNode.H
							heap.Push(&openList, newNode)
							nodeMap[newCoord] = newNode
						} else {
							existingNode.G = newG
							existingNode.F = newG + existingNode.H
							existingNode.Parent = currentNode

							// Retirer le nœud existant de la file de priorité
							// et ajouter le nœud mis à jour
							for i, n := range openList {
								if n == existingNode {
									openList = append(openList[:i], openList[i+1:]...)
									break
								}
							}
							heap.Push(&openList, existingNode)
						}
					}
				}
			}
		}
	}

	return nil // Aucun chemin trouvé
}

func LoadMobs(mapIn string) []Enemy {
	mobsReading := jsonmanagement.Get("../levels/"+mapIn, "mobs")
	mobs := []Enemy{}
	if mobsReading != nil {
		for k, v := range mobsReading.(map[string]interface{}) {
			mobs = append(mobs, Enemy{
				X:          int(v.(map[string]interface{})["x"].(float64)),
				Y:          int(v.(map[string]interface{})["y"].(float64)),
				FileName:   k[:len(k)-1],
				Smiley:     jsonmanagement.Get("mobs/"+k[:len(k)-1], "smiley").(string),
				RealName:   jsonmanagement.Get("mobs/"+k[:len(k)-1], "pseudo").(string),
				Damages:    int(jsonmanagement.Get("mobs/"+k[:len(k)-1], "damages").(float64)),
				HP:         int(jsonmanagement.Get("mobs/"+k[:len(k)-1], "health").(float64)),
				RateAttack: int(jsonmanagement.Get("mobs/"+k[:len(k)-1], "rateFail").(float64)),
			})
		}
		return mobs
	}
	return []Enemy{}

}

func reversePath(path []Coord) {
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - 1 - i
		path[i], path[j] = path[j], path[i]
	}
}
func RemoveEnemy(slice []Enemy, element interface{}) []Enemy {
	index := -1
	for i, item := range slice {
		if item == element {
			index = i
			break
		}
	}
	if index != -1 {
		slice = append(slice[:index], slice[index+1:]...)
	}

	return slice
}
