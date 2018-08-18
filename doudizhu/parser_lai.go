package doudizhu

import (
	"github.com/name5566/leaf/util"
	"sort"
)

/*
    @Create by GoLand
    @Author: hong
    @Time: 2018/7/26 18:12 
    @File: parser_lai.go    
*/

/*牌型解析：带赖子牌,赖子为4个，如果赖子多的话，可以考虑通过发牌对手中赖子的控制*/
//判断一组牌是否是指定牌型
func isDanPai_lai(pais []*Poker) (bool, int32) {
	return isDanPai(pais)
}

//判断一组牌是否是对子
func isDuiZi_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)

	if paiCount != 2 {
		return false, -1
	}

	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	if laiZiCount == 0 || laiZiCount == 2 {
		return isDuiZi(pais)
	}
	if laiZiCount == 1 {
		return true, notLaiZi[0].GetValue()
	}
	return false, -1
}

//判断一组牌是否是三张
func isSanBuDai_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount != 3 {
		return false, -1
	}

	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	if laiZiCount == 0 || laiZiCount == 3 {
		return isSanBuDai(pais)
	}
	if laiZiCount == 1 {
		if notLaiZi[0].GetValue() == notLaiZi[1].GetValue() {
			return true, notLaiZi[0].GetValue()
		}
	}
	if laiZiCount == 2 {
		return true, notLaiZi[0].GetValue()
	}
	return false, -1
}

//判断一组牌是否是三带一
func isSanDaiYi_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount != 4 {
		return false, -1
	}

	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	if laiZiCount == 0 {
		return isSanDaiYi(pais)
	}

	liangzhang := getPaiValueByCount(notLaiZi, 2)
	danzhang := getPaiValueByCount(notLaiZi, 1)
	if laiZiCount == 1 {
		if len(liangzhang) == 1 && len(danzhang) == 1 {
			return true, liangzhang[0]
		}
	}
	if laiZiCount == 2 {
		if len(danzhang) == 2 {
			sort.Sort(PaiValueList(danzhang))
			return true, danzhang[1]
		}
	}
	return false, -1
}

//判断一组牌是否是三带对子
func isSanDaiDui_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount != 5 {
		return false, -1
	}

	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)
	if laiZiCount == 0 {
		return isSanDaiDui(pais)
	}
	if laiZiCount == 1 {
		if len(sanzhang) == 1 && len(danzhang) == 1 && danzhang[0] < 16 {
			return true, sanzhang[0]
		}
		if len(liangzhang) == 2 {
			sort.Sort(PaiValueList(liangzhang))
			return true, liangzhang[1]
		}
	}

	if laiZiCount == 2 {
		if len(sanzhang) == 1 {
			return true, sanzhang[0]
		}
		if len(liangzhang) == 1 && len(danzhang) == 1 {
			if liangzhang[0] > danzhang[0] {
				return true, liangzhang[0]
			}
			if liangzhang[0] < danzhang[0] && danzhang[0] < 16 {
				return true, danzhang[0]
			}
		}
	}

	if laiZiCount == 3 {
		if len(liangzhang) == 1 {
			if liangzhang[0] > laiZi[0].GetValue() {
				return true, liangzhang[0]
			}
			if liangzhang[0] < laiZi[0].GetValue() {
				return true, laiZi[0].GetValue()
			}
		}
		if len(danzhang) == 2 {
			sort.Sort(PaiValueList(danzhang))
			if danzhang[0] < 16 {
				return true, danzhang[1]
			}
		}
	}

	if laiZiCount == 4 {
		if len(danzhang) == 1 {
			if danzhang[0] > laiZi[0].GetValue() && danzhang[0] < 16 {
				return true, danzhang[0]
			}
			if danzhang[0] < laiZi[0].GetValue() {
				return true, laiZi[0].GetValue()
			}
		}
	}
	return false, -1
}

//判断一组牌是否是顺子
func isShunZi_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 5 || paiCount > 12 {
		return false, -1
	}

	for _, p := range pais {
		if p == nil {
			continue
		}
		if paiValue := p.GetValue(); paiValue == 15 || paiValue == 16 || paiValue == 17 {
			return false, -1
		}
	}
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)
	values := getPaiValueList(notLaiZi)
	valueCount := len(values)

	if len(notLaiZi) != valueCount {
		return false, -1
	}

	//牌值排序
	sort.Sort(PaiValueList(values))
	if laiZiCount == 0 {
		return isShunZi(pais)
	}

	//1.赖子全放两边
	if values[valueCount-1]-values[0]+1 == int32(valueCount) {
		//全放左边
		if values[0]-int32(laiZiCount) > 3 {
			return true, values[0] - int32(laiZiCount)
		}
		//全放右边
		if values[valueCount-1]+int32(laiZiCount) < 15 {
			return true, values[0]
		}
		//左右都放
		if laiZiCount == 2 {
			if values[0] >= 4 && values[valueCount-1] < 14 {
				return true, values[0] - 1
			}
		}
		if laiZiCount == 3 {
			//1+2
			if values[0] >= 4 && values[valueCount-1] < 13 {
				return true, values[0] - 1
			}
			//2+1
			if values[0] >= 5 && values[valueCount-1] < 14 {
				return true, values[0] - 2
			}
		}
		if laiZiCount == 4 {
			//1+3
			if values[0] >= 4 && values[valueCount-1] < 12 {
				return true, values[0] - 1
			}
			//2+2
			if values[0] >= 5 && values[valueCount-1] < 13 {
				return true, values[0] - 2
			}
			//3+1
			if values[0] >= 6 && values[valueCount-1] < 14 {
				return true, values[0] - 3
			}
		}
	}

	//2.赖子全在中间
	if values[valueCount-1]-values[0]+1 == int32(valueCount+laiZiCount) {
		return true, values[0]
	}

	//3.部分赖子在中间，其余在两边
	if values[valueCount-1]-values[0]+1 < int32(valueCount+laiZiCount) {
		if laiZiCount == 2 {
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+1 {
				if values[valueCount-1] < 14 {
					return true, values[0]
				}
				if values[0] >= 4 {
					return true, values[0] - 1
				}
			}
		}
		if laiZiCount == 3 {
			//中间一个
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+1 {
				if values[valueCount-1] < 13 {
					return true, values[0]
				}
				if values[0] >= 5 {
					return true, values[0] - 2
				}
				if values[0] >= 4 && values[valueCount-1] < 14 {
					return true, values[0] - 1
				}
			}

			//中间两个
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+2 {
				if values[valueCount-1] < 14 {
					return true, values[0]
				}
				if values[0] >= 4 {
					return true, values[0] - 1
				}
			}
		}

		if laiZiCount == 4 {
			//中间一个
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+1 {
				if values[valueCount-1] < 12 {
					return true, values[0]
				}
				if values[valueCount-1] < 13 && values[0] >= 4 {
					return true, values[0] - 1
				}
				if values[valueCount-1] < 14 && values[0] >= 5 {
					return true, values[0] - 2
				}
				if values[0] >= 6 {
					return true, values[0] - 3
				}
			}
			//中间两个
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+2 {
				if values[valueCount-1] < 13 {
					return true, values[0]
				}
				if values[valueCount-1] < 14 && values[0] >= 4 {
					return true, values[0] - 1
				}
				if values[0] >= 5 {
					return true, values[0] - 2
				}
			}
			//中间三个
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+3 {
				if values[valueCount-1] < 14 {
					return true, values[0]
				}
				if values[0] >= 4 {
					return true, values[0] - 1
				}
			}
		}
	}
	return false, -1
}

func isLianDui_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 6 || paiCount > 16 || paiCount%2 != 0 {
		return false, -1
	}

	for _, p := range pais {
		if p == nil {
			continue
		}
		if paiValue := p.GetValue(); paiValue == 15 || paiValue == 16 || paiValue == 17 {
			return false, -1
		}
	}
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)
	values := getPaiValueList(notLaiZi)
	valueCount := len(values)

	//牌值排序
	sort.Sort(PaiValueList(values))
	if laiZiCount == 0 {
		return isLianDui(pais)
	}

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	if laiZiCount == 1 {
		if values[valueCount-1]-values[0]+1 == int32(valueCount) {
			if len(danzhang) == 1 && len(liangzhang)*2+len(danzhang) == len(notLaiZi) {
				return true, values[0]
			}
		}
	}

	if laiZiCount == 2 {
		//刚好差一对
		if len(liangzhang)*2 == len(notLaiZi) {
			//差的这一对在两边
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				if values[valueCount-1] < 14 {
					return true, values[0]
				}
				if values[0] >= 4 {
					return true, values[0] - 1
				}
			}

			//差的这一对在中间
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+1 {
				return true, values[0]
			}
		}

		//两张赖子组成两个对子
		if len(liangzhang)*2+len(danzhang) == len(notLaiZi) && len(danzhang) == 2 {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}
	}

	if laiZiCount == 3 {
		//三张单牌跟赖子牌组成三个对子
		if len(danzhang) == 3 && len(liangzhang)*2+len(danzhang) == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}
		//三张赖子组成一个对子和一张单牌
		if len(danzhang) == 1 && len(liangzhang)*2+len(danzhang) == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				if values[valueCount-1] < 14 {
					return true, values[0]
				}
				if values[0] >= 4 {
					return true, values[0] - 1
				}
			}
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+1 {
				return true, values[0]
			}
		}
	}

	if laiZiCount == 4 {
		//四张单牌和四个赖子组成四个对子
		if len(danzhang) == 4 && len(danzhang)+len(liangzhang)*2 == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}

		//四张赖子组成三个对子
		if len(danzhang) == 2 && len(danzhang)+len(liangzhang)*2 == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				if values[valueCount-1] < 14 {
					return true, values[0]
				}
				if values[0] >= 4 {
					return true, values[0] - 1
				}
			}
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+1 {
				return true, values[0]
			}
		}
		//四张赖子组成两个对子
		if len(danzhang) == 0 && len(liangzhang)*2 == len(notLaiZi) {
			//两个对子在两边
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				if values[valueCount-1] < 13 {
					return true, values[0]
				}
				if values[valueCount-1] < 14 {
					return true, values[0] - 1
				}
				if values[0] >= 5 {
					return true, values[0] - 2
				}
			}
			//两个对子在中间
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+2 {
				return true, values[0]
			}
			//一个在两边，一个在中间
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+1 {
				if values[valueCount-1] < 14 {
					return true, values[0]
				}
				if values[0] >= 4 {
					return true, values[0] - 1
				}
			}
		}
	}
	return false, -1
}

//判断是否是飞机不带牌
func isAirBuDai_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 6 || paiCount > 15 || paiCount%3 != 0 {
		return false, -1
	}

	for _, p := range pais {
		if p == nil {
			continue
		}
		if paiValue := p.GetValue(); paiValue == 15 || paiValue == 16 || paiValue == 17 {
			return false, -1
		}
	}
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)
	values := getPaiValueList(notLaiZi)
	valueCount := len(values)

	//牌值排序
	sort.Sort(PaiValueList(values))

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)

	if laiZiCount == 0 {
		return isAirBuDai(pais)
	}

	if laiZiCount == 1 {
		if len(liangzhang) == 1 && len(sanzhang)*3+len(liangzhang)*2 == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}
	}

	if laiZiCount == 2 {
		if len(liangzhang) == 2 && len(sanzhang)*3+len(liangzhang)*2 == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}
		if len(danzhang) == 1 && len(sanzhang)*3+len(danzhang) == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}
	}

	if laiZiCount == 3 {
		if len(sanzhang)*3 == len(notLaiZi) {
			//444555xxx
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				if values[valueCount-1] < 14 {
					return true, values[0]
				}
				if values[0] >= 4 {
					return true, values[0] - 1
				}
			}
			//333xxx555
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+1 {
				return true, values[0]
			}
		}
		//33x44x55x666
		if len(liangzhang) == 3 && len(sanzhang)*3+len(liangzhang)*2 == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}
		//33x4xx555
		if len(liangzhang) == 1 && len(danzhang) == 1 && len(sanzhang)*3+len(liangzhang)*2+len(danzhang) == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}
	}

	if laiZiCount == 4 {
		if len(liangzhang) == 1 && len(sanzhang)*3+len(liangzhang)*2 == len(notLaiZi) {
			//33444xxx
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				if values[valueCount-1] < 14 {
					return true, values[0]
				}
				if values[0] >= 4 {
					return true, values[0] - 1
				}
			}
			//33xxx555
			if values[valueCount-1]-values[0]+1 == int32(valueCount)+1 {
				return true, values[0]
			}
		}
		//33445566777
		if len(liangzhang) == 4 && len(liangzhang)*2+len(sanzhang)*3 == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}
		//33445666
		if len(liangzhang) == 2 && len(danzhang) == 1 && len(sanzhang)*3+len(liangzhang)*2+len(danzhang) == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}
		//34555
		if len(danzhang) == 2 && len(sanzhang)*3+len(danzhang) == len(notLaiZi) {
			if values[valueCount-1]-values[0]+1 == int32(valueCount) {
				return true, values[0]
			}
		}

	}
	return false, -1
}

//是否是飞机带单牌
func isAirDaiDan_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 8 || paiCount > 16 || paiCount%4 != 0 {
		return false, -1
	}

	var daiPai []*Poker                 //如果pais中有2、大小王，则一定是被带的牌
	newPais := make([]*Poker, paiCount) //可以组成三张的牌
	copy(newPais, pais)

	for _, p := range pais {
		if p == nil {
			continue
		}
		if p.GetValue() == 16 || p.GetValue() == 17 {
			daiPai = append(daiPai, p)
			newPais = delPokerFromPaisById(p, newPais)
		}
		if p.GetValue() == 15 && !p.IsLaiZi() {
			daiPai = append(daiPai, p)
			newPais = delPokerFromPaisById(p, newPais)
		}
	}

	laiZi, notLaiZi := getLaiZiFromPais(newPais)
	laiZiCount := len(laiZi)
	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	//sanzhang := getPaiValueByCount(notLaiZi, 3)

	if laiZiCount == 0 {
		return isAirDaiDan(pais)
	}

	//规则：从牌堆中剔除单牌，然后判断剩余的牌是否是飞机不带牌
	//两连的飞机
	if paiCount == 8 {
		//刚好带了两张2或者大小王，只需要判断剩下的牌是否是飞机不带牌
		if len(daiPai) == 2 {
			if isOk, key := isAirBuDai_lai(newPais); isOk {
				return isOk, key
			}
		}
		//带牌只有一张，需要在从剩余的牌里找一张带牌
		if len(daiPai) == 1 {
			//找的是单张
			if len(danzhang) > 0 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisByValue(danzhang[0], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//找的是赖子
			if laiZiCount > 0 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
		}
		if len(daiPai) == 0 {
			if len(danzhang) >= 2 {
				for i := 0; i < len(danzhang)-1; i++ {
					for j := i + 1; j < len(danzhang); j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			if len(danzhang) == 1 && laiZiCount > 0 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisByValue(danzhang[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			if laiZiCount > 1 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			if len(danzhang) == 0 && len(liangzhang) > 0 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
		}
	}

	//三连的飞机
	if paiCount == 12 {
		//带牌刚好够
		if len(daiPai) == 3 {
			if isOk, key := isAirBuDai_lai(newPais); isOk {
				return isOk, key
			}
		}
		//带牌少一张, 需要找一张带牌
		if len(daiPai) == 2 {
			//找单牌
			if len(danzhang) > 0 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}

			//找赖子
			if laiZiCount > 0 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
		}

		//如果带牌少两张，需要从牌堆里找两张牌作为带的牌
		if len(daiPai) == 1 {
			//找的全是单牌
			if len(danzhang) >= 2 {
				for i := 0; i < len(danzhang)-1; i++ {
					for j := i + 1; j < len(danzhang); j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//找的全是赖子
			if laiZiCount >= 2 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//找的是单牌和赖子混合
			if len(danzhang) > 0 && laiZiCount > 0 {
				//1张单张，1张赖子
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//找的是两张
			if len(liangzhang) > 0 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
		}
		//没有带2或者大小王，需要找三张带牌
		if len(daiPai) == 0 {
			//全是单张
			if len(danzhang) >= 3 {
				for i := 0; i < len(danzhang)-2; i++ {
					for j := i + 1; j < len(danzhang)-1; j++ {
						for k := j + 1; k < len(danzhang); k++ {
							tempPais := util.DeepClone(newPais).([]*Poker)
							tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[k], tempPais)
							if isOk, key := isAirBuDai_lai(tempPais); isOk {
								return isOk, key
							}
						}
					}
				}
			}
			//全是赖子
			if laiZiCount >= 3 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				tempPais = delPokerFromPaisById(laiZi[2], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}

			//单张和赖子的混合
			//两张单牌，一张赖子
			if len(danzhang) >= 2 && laiZiCount >= 1 {
				for i := 0; i < len(danzhang)-1; i++ {
					for j := i + 1; j < len(danzhang); j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//两张赖子，一张单牌
			if len(danzhang) >= 1 && laiZiCount >= 2 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//找一个两张，一个单张
			if len(liangzhang) > 0 && len(danzhang) > 0 {
				for _, v := range liangzhang {
					for _, dv := range danzhang {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(v, tempPais)
						tempPais = delPokerFromPaisByValue(v, tempPais)
						tempPais = delPokerFromPaisByValue(dv, tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//找一个两张，一个赖子
			if len(liangzhang) > 0 && laiZiCount > 0 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
		}
	}

	//三连的飞机
	if paiCount == 16 {
		if len(daiPai) == 4 {
			if isOk, key := isAirBuDai_lai(newPais); isOk {
				return isOk, key
			}
		}

		if len(daiPai) == 3 {
			//找单牌
			if len(danzhang) > 0 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}

			//找赖子
			if laiZiCount > 0 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
		}

		if len(daiPai) == 2 {
			//找的全是单牌
			if len(danzhang) >= 2 {
				for i := 0; i < len(danzhang)-1; i++ {
					for j := i + 1; j < len(danzhang); j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//找的全是赖子
			if laiZiCount >= 2 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//找的是单牌和赖子混合
			if len(danzhang) > 0 && laiZiCount > 0 {
				//1张单张，1张赖子
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//找的是两张
			if len(liangzhang) > 0 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
		}

		if len(daiPai) == 1 {
			//全是单张
			if len(danzhang) >= 3 {
				for i := 0; i < len(danzhang)-2; i++ {
					for j := i + 1; j < len(danzhang)-1; j++ {
						for k := j + 1; k < len(danzhang); k++ {
							tempPais := util.DeepClone(newPais).([]*Poker)
							tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[k], tempPais)
							if isOk, key := isAirBuDai_lai(tempPais); isOk {
								return isOk, key
							}
						}
					}
				}
			}
			//全是赖子
			if laiZiCount >= 3 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				tempPais = delPokerFromPaisById(laiZi[2], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}

			//单张和赖子的混合
			//两张单牌，一张赖子
			if len(danzhang) >= 2 && laiZiCount >= 1 {
				for i := 0; i < len(danzhang)-1; i++ {
					for j := i + 1; j < len(danzhang); j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//两张赖子，一张单牌
			if len(danzhang) >= 1 && laiZiCount >= 2 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//找一个两张，一个单张
			if len(liangzhang) > 0 && len(danzhang) > 0 {
				for _, v := range liangzhang {
					for _, dv := range danzhang {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(v, tempPais)
						tempPais = delPokerFromPaisByValue(v, tempPais)
						tempPais = delPokerFromPaisByValue(dv, tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//找一个亮张，一个赖子
			if len(liangzhang) > 0 && laiZiCount > 0 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
		}

		//找四张带牌
		if len(daiPai) == 0 {
			//四张单牌
			if len(danzhang) >= 4 {
				for i := 0; i < len(danzhang)-3; i++ {
					for j := i + 1; j < len(danzhang)-2; j++ {
						for k := j + 1; j < len(danzhang)-1; k++ {
							for m := k + 1; m < len(danzhang); m++ {
								tempPais := util.DeepClone(newPais).([]*Poker)
								tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
								tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
								tempPais = delPokerFromPaisByValue(danzhang[k], tempPais)
								tempPais = delPokerFromPaisByValue(danzhang[m], tempPais)
								if isOk, key := isAirBuDai_lai(tempPais); isOk {
									return isOk, key
								}
							}
						}
					}
				}
			}
			//四张赖子
			if laiZiCount == 4 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				tempPais = delPokerFromPaisById(laiZi[2], tempPais)
				tempPais = delPokerFromPaisById(laiZi[3], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//单牌和赖子混合
			if len(danzhang) >= 3 && laiZiCount >= 1 {
				for i := 0; i < len(danzhang)-2; i++ {
					for j := i + 1; j < len(danzhang)-1; j++ {
						for k := j + 1; k < len(danzhang); k++ {
							tempPais := util.DeepClone(newPais).([]*Poker)
							tempPais = delPokerFromPaisById(laiZi[0], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[k], tempPais)
							if isOk, key := isAirBuDai_lai(tempPais); isOk {
								return isOk, key
							}
						}
					}
				}
			}
			//两张赖子，2张单牌
			if len(danzhang) >= 2 && laiZiCount >= 2 {
				for i := 0; i < len(danzhang)-1; i++ {
					for j := i + 1; j < len(danzhang); j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisById(laiZi[1], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//三张赖子，1张单牌
			if len(danzhang) >= 1 && laiZiCount >= 3 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					tempPais = delPokerFromPaisById(laiZi[2], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//两张
			if len(liangzhang) > 1 {
				for i := 0; i < len(liangzhang)-1; i++ {
					for j := i + 1; j < len(liangzhang); j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//两张和单牌混合
			if len(liangzhang) > 0 && len(danzhang) > 1 {
				for i := 0; i < len(liangzhang); i++ {
					for j := 0; j < len(danzhang)-1; j++ {
						for k := j + 1; k < len(danzhang); k++ {
							tempPais := util.DeepClone(newPais).([]*Poker)
							tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[k], tempPais)
							if isOk, key := isAirBuDai_lai(tempPais); isOk {
								return isOk, key
							}
						}
					}
				}
			}
			//两张和赖子混合
			if len(liangzhang) > 0 && laiZiCount > 1 {
				for i := 0; i < len(liangzhang); i++ {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
					tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
		}
	}
	return false, -1
}

//是否是飞机带对子
func isAirDaiDui_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 10 || paiCount > 15 || paiCount%5 != 0 {
		return false, -1
	}

	var daiPai []*Poker                 //如果pais中有2、大小王，则一定是被带的牌
	newPais := make([]*Poker, paiCount) //可以组成三张的牌
	copy(newPais, pais)

	for _, p := range pais {
		if p == nil {
			continue
		}
		if p.GetValue() == 16 || p.GetValue() == 17 {
			return false, -1
		}
		if p.GetValue() == 15 && !p.IsLaiZi() {
			daiPai = append(daiPai, p)
			newPais = delPokerFromPaisById(p, newPais)
		}
	}

	laiZi, notLaiZi := getLaiZiFromPais(newPais)
	laiZiCount := len(laiZi)
	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)

	if laiZiCount == 0 {
		return isAirDaiDui(pais)
	}

	//两连的飞机
	if paiCount == 10 {
		//不能带双王，
		if len(daiPai) == 4 {
			if isOk, key := isAirBuDai_lai(newPais); isOk {
				return isOk, key
			}
		}

		//三个2
		if len(daiPai) == 3 {
			tempPais := util.DeepClone(newPais).([]*Poker)
			tempPais = delPokerFromPaisById(laiZi[0], tempPais)
			if isOk, key := isAirBuDai_lai(tempPais); isOk {
				return isOk, key
			}
		}
		//一对2, 再找一对作为带牌
		if len(daiPai) == 2 {
			//两个赖子
			if laiZiCount > 1 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//一个单张，一个赖子
			if len(danzhang) > 0 && laiZiCount > 0 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//一个两张
			if len(liangzhang) > 0 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
		}

		//带了一个2, 肯定需要一个赖子
		if len(daiPai) == 1 {
			if laiZiCount < 1 {
				return false, -1
			}
			//一个赖子，一个两张
			if len(liangzhang) > 0 && laiZiCount >= 1 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//两个赖子，一个单张
			if len(danzhang) > 0 && laiZiCount >= 2 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//三个赖子
			if laiZiCount >= 3 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				tempPais = delPokerFromPaisById(laiZi[2], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
		}

		//没有带牌，需要找两对
		if len(daiPai) == 0 {
			//找两个两张
			if liangzhangLen := len(liangzhang); liangzhangLen >= 2 {
				for i := 0; i < liangzhangLen-1; i++ {
					for j := i + 1; j < liangzhangLen; j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}

			//找四个赖子
			if laiZiCount == 4 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				tempPais = delPokerFromPaisById(laiZi[2], tempPais)
				tempPais = delPokerFromPaisById(laiZi[3], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}

			//找一个两张，两个赖子
			if len(liangzhang) >= 1 && laiZiCount >= 2 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}

			//找一个两张，一个赖子， 一个单张
			if len(liangzhang) > 0 && laiZiCount >= 1 && len(danzhang) > 0 {
				for _, lv := range liangzhang {
					for _, dv := range danzhang {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(lv, tempPais)
						tempPais = delPokerFromPaisByValue(lv, tempPais)
						tempPais = delPokerFromPaisByValue(dv, tempPais)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}

			//找两个单张，两个赖子
			if danzhangLen := len(danzhang); danzhangLen >= 1 && laiZiCount >= 1 {
				for i := 0; i < danzhangLen-1; i++ {
					for j := i + 1; j < danzhangLen; j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisById(laiZi[1], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
		}
	}

	//三连的飞机
	if paiCount == 15 {
		//带了4个2， 还需要一对
		if len(daiPai) == 4 {
			//一个两张
			if liangzhangLen := len(liangzhang); liangzhangLen > 0 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//两个赖子
			if laiZiCount > 1 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//一个赖子，一个单张
			if len(danzhang) > 0 && laiZiCount > 0 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
		}

		//带三个2，至少需要一个2
		if len(daiPai) == 3 {
			//一个两张，一个赖子
			if len(liangzhang) > 0 && laiZiCount > 0 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//一个赖子， 两个赖子
			if laiZiCount >= 3 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				tempPais = delPokerFromPaisById(laiZi[2], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//一个赖子， 一个赖子，一个单张
			if len(danzhang) > 0 && laiZiCount >= 2 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
		}

		//带了两个2， 还需要2对
		if len(daiPai) == 2 {
			//两个两张
			if liangzhangLen := len(liangzhang); liangzhangLen > 1 {
				for i := 0; i < liangzhangLen-1; i++ {
					for j := i + 1; j < liangzhangLen; j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//一个两张，两个赖子
			if len(liangzhang) > 0 && laiZiCount > 1 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//一个两张， 一个单张， 一个赖子
			if len(liangzhang) > 0 && len(danzhang) > 0 && laiZiCount > 0 {
				for _, lv := range liangzhang {
					for _, dv := range danzhang {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(lv, tempPais)
						tempPais = delPokerFromPaisByValue(lv, tempPais)
						tempPais = delPokerFromPaisByValue(dv, tempPais)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//四个赖子
			if laiZiCount == 4 {
				tempPais := util.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				tempPais = delPokerFromPaisById(laiZi[2], tempPais)
				tempPais = delPokerFromPaisById(laiZi[3], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}

			//一个单张， 三个赖子
			if len(danzhang) > 0 && laiZiCount >= 3 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					tempPais = delPokerFromPaisById(laiZi[2], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//两个单张， 两个赖子
			if danzhangLen := len(danzhang); danzhangLen > 1 && laiZiCount > 1 {
				for i := 0; i < danzhangLen-1; i++ {
					for j := i + 1; j < danzhangLen; j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisById(laiZi[1], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
		}

		//带了一个2
		if len(daiPai) == 1 {
			//一个三张，三个赖子
			if len(sanzhang) > 0 && laiZiCount >= 3 {
				for _, v := range sanzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					tempPais = delPokerFromPaisById(laiZi[2], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}

			//两个两张，一个赖子
			if liangzhangLen := len(liangzhang); liangzhangLen >= 2 {
				for i := 0; i < liangzhangLen-1; i++ {
					for j := i + 1; j < liangzhangLen; j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//一个两张，一个赖子， 两个赖子
			if len(liangzhang) > 0 && laiZiCount >= 3 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					tempPais = delPokerFromPaisById(laiZi[2], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//一个两张， 一个赖子，一个单张， 一个赖子
			if len(liangzhang) > 0 && len(danzhang) > 0 && laiZiCount >= 2 {
				for _, lv := range liangzhang {
					for _, dv := range danzhang {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(lv, tempPais)
						tempPais = delPokerFromPaisByValue(lv, tempPais)
						tempPais = delPokerFromPaisByValue(dv, tempPais)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisById(laiZi[1], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//一个赖子， 一个单张，三个赖子
			if len(danzhang) > 0 && laiZiCount == 4 {
				for _, v := range danzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					tempPais = delPokerFromPaisById(laiZi[2], tempPais)
					tempPais = delPokerFromPaisById(laiZi[3], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//一个赖子， 两个单张， 两个赖子
			if danzhangLen := len(danzhang); danzhangLen >= 2 && laiZiCount >= 3 {
				for i := 0; i < danzhangLen-1; i++ {
					for j := i + 1; j < danzhangLen; j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisById(laiZi[1], tempPais)
						tempPais = delPokerFromPaisById(laiZi[2], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
		}

		//不带2，需要找三对
		if len(daiPai) == 0 {
			//三个两张
			if liangzhangLen := len(liangzhang); liangzhangLen >= 3 {
				for i := 0; i < liangzhangLen-2; i++ {
					for j := i + 1; j < liangzhangLen-1; j++ {
						for k := j + 1; k < liangzhangLen; k++ {
							tempPais := util.DeepClone(newPais).([]*Poker)
							tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
							tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
							tempPais = delPokerFromPaisByValue(liangzhang[k], tempPais)
							tempPais = delPokerFromPaisByValue(liangzhang[k], tempPais)
							if isOk, key := isAirBuDai_lai(tempPais); isOk {
								return isOk, key
							}
						}
					}
				}
			}
			//两个两张，两个赖子
			if liangzhangLen := len(liangzhang); liangzhangLen >= 2 && laiZiCount >= 2 {
				for i := 0; i < liangzhangLen-1; i++ {
					for j := i + 1; j < liangzhangLen; j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisById(laiZi[1], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//两个两张， 一个单张，一个赖子
			if len(liangzhang) >= 2 && len(danzhang) > 0 && laiZiCount >= 1 {
				for i := 0; i < len(liangzhang)-1; i++ {
					for j := i + 1; j < len(liangzhang); j++ {
						for _, v := range danzhang {
							tempPais := util.DeepClone(newPais).([]*Poker)
							tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(liangzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
							tempPais = delPokerFromPaisByValue(liangzhang[j], tempPais)
							tempPais = delPokerFromPaisByValue(v, tempPais)
							tempPais = delPokerFromPaisById(laiZi[0], tempPais)
							if isOk, key := isAirBuDai_lai(tempPais); isOk {
								return isOk, key
							}
						}
					}
				}
			}
			//一个两张，四个赖子
			if len(liangzhang) > 0 && laiZiCount == 4 {
				for _, v := range liangzhang {
					tempPais := util.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisById(laiZi[0], tempPais)
					tempPais = delPokerFromPaisById(laiZi[1], tempPais)
					tempPais = delPokerFromPaisById(laiZi[2], tempPais)
					tempPais = delPokerFromPaisById(laiZi[3], tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//一个两张， 一个单张，三个赖子
			if len(liangzhang) > 0 && len(danzhang) > 0 && laiZiCount >= 3 {
				for _, lv := range liangzhang {
					for _, dv := range danzhang {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(lv, tempPais)
						tempPais = delPokerFromPaisByValue(lv, tempPais)
						tempPais = delPokerFromPaisByValue(dv, tempPais)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisById(laiZi[1], tempPais)
						tempPais = delPokerFromPaisById(laiZi[2], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//一个两张， 两个单张，两个赖子
			if danzhangLen := len(danzhang); len(liangzhang) > 0 && danzhangLen >= 2 && laiZiCount >= 2 {
				for _, v := range liangzhang {
					for i := 0; i < danzhangLen-1; i++ {
						for j := i + 1; j < danzhangLen; j++ {
							tempPais := util.DeepClone(newPais).([]*Poker)
							tempPais = delPokerFromPaisByValue(v, tempPais)
							tempPais = delPokerFromPaisByValue(v, tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
							tempPais = delPokerFromPaisById(laiZi[0], tempPais)
							tempPais = delPokerFromPaisById(laiZi[1], tempPais)
							if isOk, key := isAirBuDai_lai(tempPais); isOk {
								return isOk, key
							}
						}
					}
				}
			}
			//两个单张，四个赖子
			if danzhangLen := len(danzhang); danzhangLen >= 2 && laiZiCount == 4 {
				for i := 0; i < danzhangLen-1; i++ {
					for j := i + 1; j < danzhangLen; j++ {
						tempPais := util.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						tempPais = delPokerFromPaisById(laiZi[0], tempPais)
						tempPais = delPokerFromPaisById(laiZi[1], tempPais)
						tempPais = delPokerFromPaisById(laiZi[2], tempPais)
						tempPais = delPokerFromPaisById(laiZi[3], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			//三个单张，三个赖子
			if danzhangLen := len(danzhang); danzhangLen >= 3 && laiZiCount >= 3 {
				for i := 0; i < danzhangLen-2; i++ {
					for j := i + 1; j < danzhangLen-1; j++ {
						for k := j + 1; k < danzhangLen; k++ {
							tempPais := util.DeepClone(newPais).([]*Poker)
							tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
							tempPais = delPokerFromPaisByValue(danzhang[k], tempPais)
							tempPais = delPokerFromPaisById(laiZi[0], tempPais)
							tempPais = delPokerFromPaisById(laiZi[1], tempPais)
							tempPais = delPokerFromPaisById(laiZi[2], tempPais)
							if isOk, key := isAirBuDai_lai(tempPais); isOk {
								return isOk, key
							}
						}
					}
				}
			}
		}
	}

	return false, -1
}

//是否是四带两张单牌
func isSiDaiDan_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount != 6 {
		return false, -1
	}

	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)
	sizhang := getPaiValueByCount(notLaiZi, 4)

	if laiZiCount == 0 {
		return isSiDaiDan(pais)
	}
	if laiZiCount == 1 {
		if len(sizhang) == 1 && len(danzhang) == 1 && len(sizhang)*4+len(danzhang) == len(notLaiZi) {
			return true, sizhang[0]
		}
		if len(sanzhang) == 1 && len(danzhang) == 2 && len(sanzhang)*3+len(danzhang) == len(notLaiZi) {
			return true, sanzhang[0]
		}
		if len(sanzhang) == 1 && len(liangzhang) == 1 && len(sanzhang)*3+len(liangzhang)*2 == len(notLaiZi) {
			return true, sanzhang[0]
		}
	}
	if laiZiCount == 2 {
		if len(sizhang) == 1 && sizhang[0] != laiZi[0].GetValue() {
			return true, sizhang[0]
		}
		if len(liangzhang) == 1 && len(danzhang) == 2 && len(liangzhang)*2+len(danzhang) == len(notLaiZi) {
			return true, liangzhang[0]
		}
		if len(liangzhang) == 2 && len(liangzhang)*2 == len(notLaiZi) {
			sort.Sort(PaiValueList(liangzhang))
			return true, liangzhang[1]
		}
	}
	if laiZiCount == 3 {
		if len(danzhang) == 3 && len(danzhang) == len(notLaiZi) {
			sort.Sort(PaiValueList(danzhang))
			return true, danzhang[2]
		}
		if len(danzhang) == 1 && len(liangzhang) == 1 && len(danzhang)+len(liangzhang)*2 == len(notLaiZi) {
			return true, danzhang[0]
		}
	}
	if laiZiCount == 4 {
		return true, laiZi[0].GetValue()
	}

	return false, -1
}

//是否是四带两对牌
func isSiDaiDui_lai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount != 6 {
		return false, -1
	}

	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)
	sizhang := getPaiValueByCount(notLaiZi, 4)

	if laiZiCount == 0 {
		return isSiDaiDui(pais)
	}

	if laiZiCount == 1 {
		if len(sizhang) == 1 && len(danzhang) == 1 {
			return true, sizhang[0]
		}
		if len(sanzhang) == 1 && len(liangzhang) == 1 {
			return true, sanzhang[0]
		}
	}
	if laiZiCount == 2 {
		if len(sizhang) == 1 && sizhang[0] != laiZi[0].GetValue() {
			return true, sizhang[0]
		}
		if len(sanzhang) == 1 && len(danzhang) == 1 {
			return true, sanzhang[0]
		}
		if len(liangzhang) == 2 {
			if liangzhang[0] > liangzhang[1] {
				return true, liangzhang[0]
			} else {
				return true, liangzhang[1]
			}
		}
	}
	if laiZiCount == 3 {
		if len(sanzhang) == 1 {
			return true, sanzhang[0]
		}
		if len(liangzhang) == 1 && len(danzhang) == 1 {
			if danzhang[0] > liangzhang[0] {
				return true, danzhang[0]
			} else {
				return true, liangzhang[0]
			}
		}
	}
	if laiZiCount == 4 {
		if len(liangzhang) == 1 {
			if laiZi[0].GetValue() > liangzhang[0] {
				return true, laiZi[0].GetValue()
			} else {
				return true, liangzhang[0]
			}
		}
	}
	return false, -1
}

//是否是普通炸弹
func isNormalBoom(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount != 4 {
		return false, -1
	}
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)
	sizhang := getPaiValueByCount(notLaiZi, 4)

	if laiZiCount == 0 {
		if len(sizhang) == 1 {
			return true, sizhang[0]
		}
	}
	if laiZiCount == 1 {
		if len(sanzhang) == 1 && len(sanzhang)*3+laiZiCount == paiCount {
			return true, sanzhang[0]
		}
	}
	if laiZiCount == 2 {
		if len(liangzhang) == 2 && len(liangzhang)*2+laiZiCount == paiCount {
			return true, liangzhang[0]
		}
	}
	if laiZiCount == 3 {
		if len(danzhang) == 1 && len(danzhang)+laiZiCount == paiCount {
			return true, danzhang[0]
		}
	}
	//全赖子的炸弹比普通炸弹大
	if laiZiCount == 4 {
		return true, 16
	}
	return false, -1
}

//获取一组牌中的赖子牌
func getLaiZiFromPais(pais []*Poker) (laizi []*Poker, notLaiZi []*Poker) {
	for _, p := range pais {
		if p == nil {
			continue
		}
		if p.IsLaiZi() {
			laizi = append(laizi, p)
		}
		if !p.IsLaiZi() {
			notLaiZi = append(notLaiZi, p)
		}
	}
	return
}

//从牌堆中删除一张指定的牌
func delPokerFromPaisById(delP *Poker, pais []*Poker) []*Poker {
	var ret []*Poker
	for i, p := range pais {
		if p == nil {
			continue
		}
		if p.GetId() == delP.GetId() {
			ret = append(pais[:i], pais[i+1:]...)
			return ret
		}
	}
	return ret
}

func delPokerFromPaisByValue(v int32, pais []*Poker) []*Poker {
	var ret []*Poker
	for i, p := range pais {
		if p == nil {
			continue
		}
		if p.GetValue() == v {
			ret = append(pais[:i], pais[i+1:]...)
			return ret
		}
	}
	return ret
}
