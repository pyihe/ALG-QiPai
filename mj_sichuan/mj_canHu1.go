package mj_sichuan

//判断一手牌是否可以胡，这里的手牌包含了碰、杠到桌面的牌
func canHu(pais []*MjPai) (canHu, is19 bool, jiangPaiValue uint8) {
	paiCount := len(pais)
	//先判断牌的张数，胡牌必须满足14张牌
	if paiCount != 14 {
		return
	}

	canHu, is19, jiangIndex := tryHu(getPaisUniqueIndex(pais), paiCount)
	jiangPaiValue = getPaiValueByIndex(jiangIndex)

	return canHu, is19, jiangPaiValue
}

func getPaisUniqueIndex(pais []*MjPai) (indexs []uint8) {
	//	0-8   ——> 1-9筒
	//	9-18  ——> 1-9条
	//	19-26 ——> 1-9万
	indexs = make([]uint8, 27)
	for _, p := range pais {
		if p == nil {
			continue
		}
		indexs[p.GetUniqueIndexByValue()]++
	}
	return
}

func getPaiValueByIndex(index uint8) uint8 {
	var val uint8
	if index <= 8 {
		val = index + 1
	} else if index > 8 && index <= 17 {
		val = index - 8
	} else {
		val = index - 17
	}
	return val
}

//判断index对应的牌是否是1或9
func is19(index uint8) bool {
	val := getPaiValueByIndex(index)
	return val == 1 || val == 9
}

//根据牌的唯一索引判断是否组成了可以胡牌的牌型
//判断规则，采用递归
//有将牌的情况，除去将牌，然后判断剩下牌是否可以胡牌
//没有将牌的情况，判断是否组成顺子或坎
func tryHu(paiIndex []uint8, paiCount int) (canHu, is19Hu bool, jiangPai uint8) {
	canHu = false
	is19Hu = true
	jiangPai = 27
	if paiCount <= 0 {
		return true, is19Hu, jiangPai
	}

	if paiCount%3 == 2 {
		//包含将牌的情况
		//除去所有可能成为将牌的两张牌，继续判断剩下的牌是否组成了胡牌的牌型
		for i := 0; i < 27; i++ {
			if paiIndex[i] < 2 {
				continue
			}
			paiIndex[i] -= 2
			canHu, is19Hu, jiangPai = tryHu(paiIndex, paiCount-2)
			if canHu {
				if !is19(uint8(i)) {
					is19Hu = false
				}
				jiangPai = uint8(i)
				return canHu, is19Hu, jiangPai
			}
			paiIndex[i] += 2
		}
	} else {
		//TODO 四川麻将缺一门，所以实际项目中使用的时候可以根据玩家定缺的花色，减少递归的情况
		//没有将牌的情况，需要判断是否是顺子或者坎牌
		//每种花色单独组合判断
		//判断筒子是否有顺子组合
		for i := 0; i <= 6; i++ {
			if paiIndex[i] > 0 && paiIndex[i+1] > 0 && paiIndex[i+2] > 0 {
				//连续的三张牌都至少有一张，则分别扣除一张，然后判断是否可以胡牌
				paiIndex[i] -= 1
				paiIndex[i+1] -= 1
				paiIndex[i+2] -= 1
				canHu, is19Hu, jiangPai = tryHu(paiIndex, paiCount-3)
				if canHu {
					if !is19(uint8(i)) || !is19(uint8(i+2)) {
						is19Hu = false
					}
					return canHu, is19Hu, jiangPai
				}
				paiIndex[i] += 1
				paiIndex[i+1] += 1
				paiIndex[i+2] += 1
			}
		}

		//判断筒子是否有坎牌组合
		for i := 0; i <= 8; i++ {
			if paiIndex[i] >= 3 {
				paiIndex[i] -= 3
				canHu, is19Hu, jiangPai = tryHu(paiIndex, paiCount-3)
				if canHu {
					if !is19(uint8(i)) {
						is19Hu = false
					}
					return canHu, is19Hu, jiangPai
				}
				paiIndex[i] += 3
			}
		}

		//判断条子是否有顺子组合
		for i := 9; i <= 15; i++ {
			if paiIndex[i] > 0 && paiIndex[i+1] > 0 && paiIndex[i+2] > 0 {
				//连续的三张牌都至少有一张，则分别扣除一张，然后判断是否可以胡牌
				paiIndex[i] -= 1
				paiIndex[i+1] -= 1
				paiIndex[i+2] -= 1
				canHu, is19Hu, jiangPai = tryHu(paiIndex, paiCount-3)
				if canHu {
					if !is19(uint8(i)) || !is19(uint8(i+2)) {
						is19Hu = false
					}
					return canHu, is19Hu, jiangPai
				}
				paiIndex[i] += 1
				paiIndex[i+1] += 1
				paiIndex[i+2] += 1
			}
		}

		//判断条子是否有坎牌
		for i := 9; i <= 17; i++ {
			if paiIndex[i] >= 3 {
				paiIndex[i] -= 3
				canHu, is19Hu, jiangPai = tryHu(paiIndex, paiCount-3)
				if canHu {
					if !is19(uint8(i)) {
						is19Hu = false
					}
					return canHu, is19Hu, jiangPai
				}
				paiIndex[i] += 3
			}
		}

		//判断万子是否有顺子组合
		for i := 18; i <= 24; i++ {
			if paiIndex[i] > 0 && paiIndex[i+1] > 0 && paiIndex[i+2] > 0 {
				//连续的三张牌都至少有一张，则分别扣除一张，然后判断是否可以胡牌
				paiIndex[i] -= 1
				paiIndex[i+1] -= 1
				paiIndex[i+2] -= 1
				canHu, is19Hu, jiangPai = tryHu(paiIndex, paiCount-3)
				if canHu {
					if !is19(uint8(i)) || !is19(uint8(i+2)) {
						is19Hu = false
					}
					return canHu, is19Hu, jiangPai
				}
				paiIndex[i] += 1
				paiIndex[i+1] += 1
				paiIndex[i+2] += 1
			}
		}

		//判断万子是否有坎牌组合
		for i := 18; i <= 26; i++ {
			if paiIndex[i] >= 3 {
				paiIndex[i] -= 3
				canHu, is19Hu, jiangPai = tryHu(paiIndex, paiCount-3)
				if canHu {
					if !is19(uint8(i)) {
						is19Hu = false
					}
					return canHu, is19Hu, jiangPai
				}
				paiIndex[i] += 3
			}
		}

	}
	return canHu, is19Hu, jiangPai
}
