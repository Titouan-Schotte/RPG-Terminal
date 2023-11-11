package levelmanagement

func CloneLevel(level [][]map[string]string) [][]map[string]string {
	//Permet de dissocier les adresses mémoires en clonant les level pour garder toujours un level d'origine qui nous servira de référence
	levelBuff := make([][]map[string]string, len(level))
	for i := range level {
		levelBuff[i] = make([]map[string]string, len(level[i]))
		for j := range level[i] {
			levelBuff[i][j] = make(map[string]string)
			for key, value := range level[i][j] {
				levelBuff[i][j][key] = value
			}
		}
	}
	return levelBuff
}
