package util

// 计算各张待牌的和率
func CalculateAgariRateOfEachTile(waits Waits, selfDiscards []int) map[int]float64 {
	// 根据自家舍牌，确定各个牌的类型（无筋、半筋、筋、两筋）
	// 站在其他玩家的视角，自家各个牌的安全度

	tileAgariRate := map[int]float64{}
	for tile, left := range waits {
		if left == 0 {
			continue
		}
		var rate float64
		if tile < 27 { // 数牌
			t := tile % 9
			if t > 4 {
				t = 8 - t
			}
			// 判断是否为筋牌
			var targetTile int
			if t<=2 {
				targetTile = t+3
			}
			rate = agariTable[t][left-1]
		} else { // 字牌
			rate = honorTileAgariTable[left-1]
		}
		tileAgariRate[tile] = rate
	}
	return tileAgariRate
}

// 计算平均和率
// TODO: selfDiscards: 自家舍牌，用于分析骗筋时的和率
func CalculateAvgAgariRate(waits Waits, selfDiscards []int) float64 {
	tileAgariRate := CalculateAgariRateOfEachTile(waits, selfDiscards)
	agariRate := 0.0
	for _, rate := range tileAgariRate {
		agariRate = agariRate + rate - agariRate*rate/100
	}
	return agariRate
}
