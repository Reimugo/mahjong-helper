package util

// 根据玩家的副露、手切来判断其听牌率 (0-100)
func CalcTenpaiRate(meldCount int, discardTiles []int, meldDiscardsAt []int) float64 {
	if meldCount == 0 {
		return 0
	}
	if meldCount == 4 {
		return 100
	}

	_tenpaiRate := tenpaiRate[meldCount]

	turn := MinInt(len(discardTiles), len(_tenpaiRate)-1)
	_tenpaiRateWithTurn := _tenpaiRate[turn]

	// 计算上一次副露后的手切数
	countTedashi := 0
	latestDiscardAt := meldDiscardsAt[len(meldDiscardsAt)-1]
	if len(discardTiles) > latestDiscardAt {
		for _, disTile := range discardTiles[latestDiscardAt+1:] {
			if disTile >= 0 {
				countTedashi++
			}
		}
	}
	countTedashi = MinInt(countTedashi, len(_tenpaiRateWithTurn)-1)

	return _tenpaiRateWithTurn[countTedashi]
}
