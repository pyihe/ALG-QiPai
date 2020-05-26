package doudizhu

import (
	"github.com/hong008/util/commonUtil"
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
		if paiValue := p.GetValue(); !p.IsLaiZi() && (paiValue == 15 || paiValue == 16 || paiValue == 17) {
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
		if paiValue := p.GetValue(); !p.IsLaiZi() && (paiValue == 15 || paiValue == 16 || paiValue == 17) {
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
	//TODO 地主出五连的情况为20张牌
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisByValue(danzhang[0], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//找的是赖子
			if laiZiCount > 0 {
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
						tempPais = delPokerFromPaisByValue(danzhang[i], tempPais)
						tempPais = delPokerFromPaisByValue(danzhang[j], tempPais)
						if isOk, key := isAirBuDai_lai(tempPais); isOk {
							return isOk, key
						}
					}
				}
			}
			if len(danzhang) == 1 && laiZiCount > 0 {
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisByValue(danzhang[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			if laiZiCount > 1 {
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			if len(danzhang) == 0 && len(liangzhang) > 0 {
				for _, v := range liangzhang {
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}

			//找赖子
			if laiZiCount > 0 {
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
							tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}

			//找赖子
			if laiZiCount > 0 {
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
							tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
								tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
							tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
							tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
	//TODO 地主出四连的情况为20张牌
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
			tempPais := commonUtil.DeepClone(newPais).([]*Poker)
			tempPais = delPokerFromPaisById(laiZi[0], tempPais)
			if isOk, key := isAirBuDai_lai(tempPais); isOk {
				return isOk, key
			}
		}
		//一对2, 再找一对作为带牌
		if len(daiPai) == 2 {
			//两个赖子
			if laiZiCount > 1 {
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//一个单张，一个赖子
			if len(danzhang) > 0 && laiZiCount > 0 {
				for _, v := range danzhang {
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					tempPais = delPokerFromPaisByValue(v, tempPais)
					if isOk, key := isAirBuDai_lai(tempPais); isOk {
						return isOk, key
					}
				}
			}
			//两个赖子
			if laiZiCount > 1 {
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
				tempPais = delPokerFromPaisById(laiZi[0], tempPais)
				tempPais = delPokerFromPaisById(laiZi[1], tempPais)
				if isOk, key := isAirBuDai_lai(tempPais); isOk {
					return isOk, key
				}
			}
			//一个赖子，一个单张
			if len(danzhang) > 0 && laiZiCount > 0 {
				for _, v := range danzhang {
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
				tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
							tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
							tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
					tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
							tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
						tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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
							tempPais := commonUtil.DeepClone(newPais).([]*Poker)
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

/*
	判断一组牌中是否有比指定key大的牌型，带赖子
*/

//比指定key大的对子
func largerDuiZi_lai(pais []*Poker, key int32) bool {
	paiCount := len(pais)
	if paiCount < 2 {
		return false
	}
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	if laiZiCount == 0 {
		return largerDuiZi(pais, key)
	} else {
		var largerValues []int32
		for _, v := range notLaiZi {
			if v.GetValue() > key {
				largerValues = append(largerValues, v.GetValue())
			}
		}
		return len(largerValues) > 0
	}
}

//是否有比指定key大的三不带
func largerSanBuDai_lai(pais []*Poker, key int32) bool {
	paiCount := len(pais)
	if paiCount < 3 {
		return false
	}
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	//sanzhang := getPaiValueByCount(notLaiZi, 3)

	if laiZiCount == 0 {
		return largerSanBuDai(pais, key)
	}

	if laiZiCount >= 1 {
		if len(liangzhang) > 0 {
			for _, v := range liangzhang {
				if v > key {
					return true
				}
			}
		}
	}

	if laiZiCount >= 2 {
		if len(danzhang) > 0 {
			for _, v := range danzhang {
				if v > key {
					return true
				}
			}
		}
	}
	if laiZiCount >= 3 {
		if laiZi[0].GetValue() > key {
			return true
		}
	}
	return false
}

//是否有比指定key大的三带一
func largerSanDaiDan_lai(pais []*Poker, key int32) bool {
	paiCount := len(pais)
	if paiCount < 4 {
		return false
	}
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)

	if laiZiCount == 0 {
		return largerSanDaiDan(pais, key)
	}

	if laiZiCount >= 1 {
		if len(liangzhang) > 0 {
			for _, v := range liangzhang {
				if v > key {
					return true
				}
			}
		}
	}

	if laiZiCount >= 2 {
		if len(danzhang) > 0 {
			for _, v := range danzhang {
				if v > key {
					return true
				}
			}
		}
	}

	if laiZiCount == 3 {
		if laiZi[0].GetValue() > key {
			return true
		}
	}
	return false
}

//比指定key大的三带对子
func largerSanDaiDui_lai(pais []*Poker, key int32) bool {
	paiCount := len(pais)
	if paiCount < 5 {
		return false
	}

	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)

	if laiZiCount == 0 {
		return largerSanDaiDui(pais, key)
	}
	//补一张赖子
	if laiZiCount >= 1 {
		//补成对子
		if len(sanzhang) > 0 {
			for _, v := range sanzhang {
				if v > key {
					if len(sanzhang) > 1 {
						return true
					}
					//找对子
					if len(liangzhang) > 0 {
						return true
					}
					if len(danzhang) > 0 {
						for _, dv := range danzhang {
							if dv != 16 && dv != 17 {
								return true
							}
						}
					}
				}
			}
		}
		//补成三张
		if liangzhangLen := len(liangzhang); liangzhangLen > 0 {
			for _, v := range liangzhang {
				if v > key {
					if len(sanzhang) > 0 {
						return true
					}
					if liangzhangLen > 1 {
						return true
					}
				}
			}
		}
	}
	//补两张赖子
	if laiZiCount >= 2 {
		//补成三张
		if len(danzhang) > 0 {
			for _, v := range danzhang {
				if v > key && (len(liangzhang) > 0 || len(sanzhang) > 0 ) {
					return true
				}
			}
		}

		//补成对子
		if len(sanzhang) > 0 {
			for _, v := range sanzhang {
				if v > key {
					return true
				}
			}
		}

		//一张赖子补成三张，一张赖子补成对子
		if len(liangzhang) > 0 {
			for _, v := range liangzhang {
				if v > key {
					if len(sanzhang) > 0 {
						return true
					}
					if len(liangzhang) > 1 {
						return true
					}
					if len(danzhang) > 0 {
						for _, dv := range danzhang {
							if dv != 16 && dv != 17 {
								return true
							}
						}
					}
				}
			}
		}
	}
	//补三张赖子
	if laiZiCount >= 3 {
		if laiZi[0].GetValue() > key {
			if len(liangzhang) > 0 || len(sanzhang) > 0 {
				return true
			}
		}
		if len(liangzhang) > 0 {
			//只要有一个两张满足大于key就return
			for _, v := range liangzhang {
				if v > key {
					return true
				}
			}
		}
		if len(danzhang) > 0 {
			for _, v := range danzhang {
				if v > key {
					if len(danzhang) > 1 {
						return true
					}
					if len(liangzhang) > 0 {
						return true
					}
					if len(sanzhang) > 0 {
						return true
					}
				}
			}
		}
	}
	return false
}

//比指定key大的单顺
func largerShunZi_lai(pais []*Poker, key int32, length int) bool {
	paiCount := len(pais)
	if paiCount < length {
		return false
	}

	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	values := getPaiValueList(notLaiZi)
	valueLen := len(values)
	if valueLen+laiZiCount < length {
		return false
	}

	sort.Sort(PaiValueList(values))

	if laiZiCount == 0 {
		return largerShunZi(pais, key, length)
	}

	//只取一个赖子
	if laiZiCount >= 1 {
		//找length长度的list， 不需要赖子的情况
		for i := 0; i <= valueLen-length; i++ {
			//刚好够，不需要赖子
			if values[i+length-1]-values[i]+1 == int32(length) {
				if values[i] > key && values[i+length-1] < 15 {
					return true
				}
			}
		}
		//找length-1长度的list，需要一个赖子
		for i := 0; i <= valueLen-length+1; i++ {
			//长度连续，赖子补在两边
			if values[i+length-2]-values[i]+1 == int32(length)-1 {
				if values[i] > key && values[i+length-2]+1 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-2] < 15 {
					return true
				}
			}
			//长度不连续，赖子补在中间
			if values[i+length-2]-values[i]+1 == int32(length) {
				if values[i] > key && values[i+length-2] < 15 {
					return true
				}
			}
		}
	}

	//取两个赖子,找length-2长度的list,
	if laiZiCount >= 2 {
		// 需要两个赖子
		for i := 0; i <= valueLen-length+2; i++ {
			//两个赖子都放在两边
			if values[i+length-3]-values[i]+1 == int32(length)-2 {
				if values[i] > key && values[i+length-3]+2 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-3]+1 < 15 {
					return true
				}
				if values[i]-2 > key && values[i+length-3] < 15 {
					return true
				}
			}
			//一个赖子在两边，一个在中间
			if values[i+length-3]-values[i]+1 == int32(length)-1 {
				if values[i] > key && values[i+length-3]+1 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-3] < 15 {
					return true
				}
			}
			//两个赖子都在中间
			if values[i+length-3]-values[i]+1 == int32(length) {
				if values[i] > key && values[i+length-3] < 15 {
					return true
				}
			}
		}
	}

	//取三个赖子,找length-3长度的list
	if laiZiCount >= 3 {
		for i := 0; i <= valueLen-length+3; i++ {
			//三个赖子在两边
			if values[i+length-4]-values[i]+1 == int32(length)-3 {
				if values[i] > key && values[i+length-4]+3 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-4]+2 < 15 {
					return true
				}
				if values[i]-2 > key && values[i+length-4]+1 < 15 {
					return true
				}
				if values[i]-3 > key && values[i+length-4] < 15 {
					return true
				}
			}
			//两个赖子在两边， 一个赖子在中间
			if values[i+length-4]-values[i]+1 == int32(length)-2 {
				if values[i] > key && values[i+length-4]+2 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-4]+1 < 15 {
					return true
				}
				if values[i]-2 > key && values[i+length-4] < 15 {
					return true
				}
			}
			//一个赖子在两边，两个赖子在中间
			if values[i+length-4]-values[i]+1 == int32(length)-1 {
				if values[i] > key && values[i+length-4]+1 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-3] < 15 {
					return true
				}
			}
			//三个赖子都在中间
			if values[i+length-4]-values[i]+1 == int32(length) {
				return true
			}
		}

	}

	//取四个赖子
	if laiZiCount >= 4 {
		for i := 0; i <= valueLen-length+4; i++ {
			//四个赖子都在两边
			if values[i+length-5]-values[i]+1 == int32(length)-4 {
				if values[i] > key && values[i+length-5]+4 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-5]+3 < 15 {
					return true
				}
				if values[i]-2 > key && values[i+length-5]+2 < 15 {
					return true
				}
				if values[i]-3 > key && values[i+length-5]+1 < 15 {
					return true
				}
				if values[i]-4 > key && values[i+length-5] < 15 {
					return true
				}
			}
			//一个赖子在中间，其余在两边
			if values[i+length-5]-values[i]+1 == int32(length)-3 {
				if values[i] > key && values[i+length-5]+3 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-5]+2 < 15 {
					return true
				}
				if values[i]-2 > key && values[i+length-5]+1 < 15 {
					return true
				}
				if values[i]-3 > key && values[i+length-5] < 15 {
					return true
				}
			}
			//两个赖子在中间，其余在两边
			if values[i+length-5]-values[i]+1 == int32(length)-2 {
				if values[i] > key && values[i+length-5]+2 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-5]+1 < 15 {
					return true
				}
				if values[i]-2 > key && values[i+length-5] < 15 {
					return true
				}
			}
			//三个赖子在中间，其余在两边
			if values[i+length-5]-values[i]+1 == int32(length)-1 {
				if values[i] > key && values[i+length-5]+1 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-5] < 15 {
					return true
				}
			}
			//四个赖子在中间
			if values[i+length-5]-values[i]+1 == int32(length) {
				if values[i] > key && values[i+length-5] < 15 {
					return true
				}
			}
		}
	}
	return false
}

//比指定key大的连对
func largerLianDui_lai(pais []*Poker, key int32, length int) bool {
	paiCount := len(pais)
	if paiCount < length*2 {
		return false
	}

	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	values := getPaiValueList(notLaiZi)
	valueLen := len(values)

	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)
	sizhang := getPaiValueByCount(notLaiZi, 4)

	if laiZiCount == 0 {
		return largerLianDui(pais, key, length)
	}

	//只取一个赖子
	if laiZiCount >= 1 {
		//不找赖子
		var duiZi []int32
		duiZi = append(duiZi, liangzhang...)
		duiZi = append(duiZi, sanzhang...)
		duiZi = append(duiZi, sizhang...)
		if len(duiZi) >= length {
			sort.Sort(PaiValueList(duiZi))
			for i := 0; i <= len(duiZi)-length; i++ {
				if duiZi[i+length-1]-duiZi[i]+1 == int32(length) {
					if duiZi[i] > key && duiZi[i+length-1] < 15 {
						return true
					}
				}
			}
		}

		//找一个赖子
		for i := 0; i <= valueLen-length; i++ {
			if values[i+length-1]-values[i]+1 == int32(length) {
				var danCount int32 //单张的数量,一张赖子只允许连队中有一个单牌
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 1 {
					continue
				}
				if values[i] > key && values[i+length-1] < 15 {
					return true
				}
			}
		}
	}
	//取两个赖子
	if laiZiCount >= 2 {
		//两个赖子分别补两个单张
		for i := 0; i <= valueLen-length; i++ {
			if values[i+length-1]-values[i]+1 == int32(length) {
				var danCount int32
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 2 {
					continue
				}
				if values[i] > key && values[i+length-1] < 15 {
					return true
				}
			}
		}
		//两个赖子补成一个对子
		for i := 0; i <= valueLen-length+1; i++ {
			//补的对子在两边
			if values[i+length-2]-values[i]+1 == int32(length) {
				var danCount int32
				for j := 0; j < length-1; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount > 0 {
					continue
				}
				if values[i] > key && values[i+length-2]+1 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-2] < 15 {
					return true
				}
			}
		}
		//两个赖子补成一个对子
		for i := 0; i < valueLen-length+1; i++ {
			//补的对子在中间
			if values[i+length-1]-values[i]+1 == int32(length) {
				var danCount int32
				for j := 0; j < length-1; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount > 0 {
					continue
				}
				if values[i] > key && values[i+length-1] < 15 {
					return true
				}
			}
		}
	}
	//取三个赖子
	if laiZiCount >= 3 {
		//三个赖子分别补三个单张，形成对子
		for i := 0; i <= valueLen-length; i++ {
			if values[i+length-1]-values[i]+1 == int32(length) {
				var danCount int32
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 3 {
					continue
				}
				if values[i] > key && values[i+length-1] < 15 {
					return true
				}
			}
		}
		//三个赖子补一个对子和一个单张
		for i := 0; i <= valueLen-length+1; i++ {
			//补的对子在两边
			if values[i+length-2]-values[i]+1 == int32(length)-1 {
				var danCount int32
				for j := 0; j < length-1; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 1 {
					continue
				}
				if values[i] > key && values[i+length-2]+1 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-2] < 15 {
					return true
				}
			}
		}
		//三个赖子补一个对子和一个单张
		for i := 0; i <= valueLen-length+1; i++ {
			//补的对子在中间
			if values[i+length-2]-values[i]+1 == int32(length) {
				var danCount int32
				for j := 0; j < length-1; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 1 {
					continue
				}
				if values[i] > key && values[i+length-2] < 15 {
					return true
				}
			}
		}
	}
	//取四个赖子
	if laiZiCount >= 4 {
		//补四个单张
		for i := 0; i <= valueLen-length; i++ {
			if values[i+length-1]-values[i]+1 == int32(length) {
				var danCount int32
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 4 {
					continue
				}
				if values[i] > key && values[i+length-1] < 15 {
					return true
				}
			}
		}
		//补两个单张，剩下一个对子
		for i := 0; i <= valueLen-length+1; i++ {
			//补的对子在两边
			if values[i+length-2]-values[i]+1 == int32(length)-1 {
				var danCount int32
				for j := 0; j < length-1; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 2 {
					continue
				}
				if values[i] > key && values[i+length-2]+1 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-2] < 15 {
					return true
				}
			}
		}
		//补两个单张，剩下一个对子
		for i := 0; i <= valueLen-length+1; i++ {
			//补的对子在中间
			if values[i+length-2]-values[i]+1 == int32(length) {
				var danCount int32
				for j := 0; j < length-1; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 2 {
					continue
				}
				if values[i] > key && values[i+length-2] < 15 {
					return true
				}
			}
		}
		//补两个对子
		for i := 0; i <= valueLen-length+2; i++ {
			//补的对子在两边
			if values[i+length-3]-values[i]+1 == int32(length)-2 {
				var danCount int32
				for j := 0; j < length-2; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount > 0 {
					continue
				}
				if values[i] > key && values[i+length-3]+2 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-3]+1 < 15 {
					return true
				}
				if values[i]-2 > key && values[i+length-3] < 15 {
					return true
				}
			}
		}
		//补两个对子
		for i := 0; i <= valueLen-length+2; i++ {
			//补的对子一个在中间一个在两边
			if values[i+length-3]-values[i]+1 == int32(length)-1 {
				var danCount int32
				for j := 0; j < length-2; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount > 0 {
					continue
				}
				if values[i] > key && values[i+length-3]+1 < 15 {
					return true
				}
				if values[i]-1 > key && values[i+length-3] < 15 {
					return true
				}
			}
		}
		//补两个对子
		for i := 0; i <= valueLen-length+2; i++ {
			//补的对子都在中间
			if values[i+length-3]-values[i]+1 == int32(length) {
				var danCount int32
				for j := 0; j < length-2; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount > 0 {
					continue
				}
				if values[i] > key && values[i+length-3] < 15 {
					return true
				}
			}
		}

	}
	return false
}

//比指定key大的飞机不带牌
func largerAirBuDai_lai(pais []*Poker, key int32, length int) bool {
	paiCount := len(pais)
	if paiCount < length*3 || paiCount > 15 {
		return false
	}
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	if laiZiCount == 0 {
		return largerAirBuDai(pais, key, length)
	}

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)
	sizhang := getPaiValueByCount(notLaiZi, 4)
	//找一个赖子
	if laiZiCount >= 1 {
		//不需要赖子的情况
		var sanValues []int32
		sanValues = append(sanValues, sanzhang...)
		sanValues = append(sanValues, sizhang...)
		if sanLen := len(sanValues); sanLen >= length {
			sort.Sort(PaiValueList(sanValues))
			for i := 0; i <= sanLen-length; i++ {
				if sanValues[i+length-1]-sanValues[i]+1 == int32(length) {
					if sanValues[i] > key && sanValues[i+length-1] < 15 {
						return true
					}
				}
			}
		}

		//需要一个赖子
		var values []int32
		values = append(values, liangzhang...)
		values = append(values, sanzhang...)
		values = append(values, sizhang...)
		sort.Sort(PaiValueList(values))
		if valueLen := len(values); valueLen >= length {
			for i := 0; i <= valueLen-length; i++ {
				var duiCount int32 //两张的数量
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 2 {
						duiCount++
					}
				}
				if duiCount != 1 {
					continue
				}
				if values[i+length-1]-values[i]+1 == int32(length) {
					if values[i] > key && values[i+length-1] < 15 {
						return true
					}
				}
			}
		}
	}

	//找两个赖子
	if laiZiCount >= 2 {
		//两个赖子补两个两张
		var values []int32
		values = append(values, liangzhang...)
		values = append(values, sanzhang...)
		values = append(values, sizhang...)
		sort.Sort(PaiValueList(values))
		if valueLen := len(values); valueLen >= length {
			for i := 0; i <= valueLen-length; i++ {
				var duiCount int32
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 2 {
						duiCount++
					}
				}
				if duiCount != 2 {
					continue
				}
				if values[i+length-1]-values[i]+1 == int32(length) {
					if values[i] > key && values[i+length-1] < 15 {
						return true
					}
				}
			}
		}

		//两个赖子补一个单张
		values = []int32{}
		values = append(values, danzhang...)
		values = append(values, sanzhang...)
		values = append(values, sizhang...)
		sort.Sort(PaiValueList(values))
		if valueLen := len(values); valueLen >= length {
			for i := 0; i <= valueLen-length; i++ {
				var danCount int32
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 1 {
					continue
				}
				if values[i+length-1]-values[i]+1 == int32(length) {
					if values[i] > key && values[i+length-1] < 15 {
						return true
					}
				}
			}
		}
	}

	//找三个赖子
	if laiZiCount >= 3 {
		//三个赖子补三个两张
		var values []int32
		values = append(values, liangzhang...)
		values = append(values, sanzhang...)
		values = append(values, sizhang...)
		sort.Sort(PaiValueList(values))
		if valueLen := len(values); valueLen >= length {
			for i := 0; i <= valueLen-length; i++ {
				var duiCount int32
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 2 {
						duiCount++
					}
				}
				if duiCount != 3 {
					continue
				}
				if values[i+length-1]-values[i]+1 == int32(length) {
					if values[i] > key && values[i+length-1] < 15 {
						return true
					}
				}
			}
		}

		//三个赖子补一个对子和一个单张
		values = append(values, danzhang...)
		sort.Sort(PaiValueList(values))
		if valueLen := len(values); valueLen >= length {
			for i := 0; i <= valueLen-length; i++ {
				var danCount, duiCount int32
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
					if getPaiCountByValue(notLaiZi, values[i+j]) == 2 {
						duiCount++
					}
				}
				if danCount != 1 || duiCount != 1 {
					continue
				}
				if values[i] > key && values[i+length-1] < 15 {
					return true
				}
			}
		}

		//三个赖子补一个三张
		values = []int32{}
		values = append(values, sanzhang...)
		values = append(values, sizhang...)
		sort.Sort(PaiValueList(values))
		if valueLen := len(values); valueLen >= length-1 {
			for i := 0; i <= valueLen-length+1; i++ {
				//补的三张在两边
				if values[i+length-2]-values[i]+1 == int32(length)-1 {
					if values[i] > key && values[i+length-2]+1 < 15 {
						return true
					}
					if values[i]-1 > key && values[i+length-2] < 15 {
						return true
					}
				}
				//补的三张在中间
				if values[i+length-2]-values[i]+1 == int32(length) {
					if values[i] > key && values[i+length-2] < 15 {
						return true
					}
				}
			}
		}
	}

	//找四个赖子
	if laiZiCount >= 4 {
		var values []int32
		values = append(values, sanzhang...)
		values = append(values, sizhang...)
		//四个赖子补四个两张
		if length >= 4 {
			values = append(values, liangzhang...)
			if valueLen := len(values); valueLen >= length {
				sort.Sort(PaiValueList(values))
				for i := 0; i <= valueLen-length; i++ {
					var duiCount int32
					for j := 0; j < length; j++ {
						if getPaiCountByValue(notLaiZi, values[i+j]) == 2 {
							duiCount++
						}
					}
					if duiCount != 4 {
						continue
					}
					if values[i+length-1]-values[i]+1 == int32(length) {
						if values[i] > key && values[i+length-1] < 15 {
							return true
						}
					}
				}
			}
		}
		//补两个两张，一个一张
		values = append(values, danzhang...)
		if length < 4 {
			values = append(values, liangzhang...)
		}
		if valueLen := len(values); valueLen >= length {
			sort.Sort(PaiValueList(values))
			for i := 0; i <= valueLen-length; i++ {
				var danCount, duiCount int32
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
					if getPaiCountByValue(notLaiZi, values[i+j]) == 2 {
						duiCount++
					}
				}
				if danCount != 1 || duiCount != 2 {
					continue
				}
				if values[i+length-1]-values[i]+1 == int32(length) {
					if values[i] > key && values[i+length-1] < 15 {
						return true
					}
				}
			}
		}
		//补一个两张，一个三张
		values = []int32{}
		values = append(values, sizhang...)
		values = append(values, sanzhang...)
		values = append(values, liangzhang...)
		if valueLen := len(values); valueLen >= length-1 {
			sort.Sort(PaiValueList(values))
			for i := 0; i <= valueLen-length+1; i++ {
				var duiCount int32
				for j := 0; j < length-1; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 2 {
						duiCount++
					}
				}
				if duiCount != 1 {
					continue
				}
				//补的三张在两边
				if values[i+length-2]-values[i]+1 == int32(length)-1 {
					if values[i] > key && values[i+length-2]+1 < 15 {
						return true
					}
					if values[i]-1 > key && values[i+length-2] < 15 {
						return true
					}
				}
				//补的三张在中间
				if values[i+length-2]-values[i]+1 == int32(length) {
					if values[i] > key && values[i+length-2] < 15 {
						return true
					}
				}
			}
		}

		//补两个单张
		values = []int32{}
		values = append(values, sizhang...)
		values = append(values, sanzhang...)
		values = append(values, danzhang...)
		if valueLen := len(values); valueLen >= length {
			sort.Sort(PaiValueList(values))
			for i := 0; i <= valueLen-length; i++ {
				var danCount int32
				for j := 0; j < length; j++ {
					if getPaiCountByValue(notLaiZi, values[i+j]) == 1 {
						danCount++
					}
				}
				if danCount != 2 {
					continue
				}
				if values[i+length-1]-values[i]+1 == int32(length) {
					if values[i] > key && values[i+length-1] < 15 {
						return true
					}
				}
			}
		}

	}
	return false
}

//比指定key大的飞机带单牌
func largerAirDaiDan_lai(pais []*Poker, key int32, length int) bool {
	paiCount := len(pais)
	if paiCount < length*4 || paiCount > 16 {
		return false
	}

	return largerAirBuDai_lai(pais, key, length)
}

//比指定key大的飞机带对子
func largerAirDaiDui_lai(pais []*Poker, key int32, length int) bool {
	paiCount := len(pais)
	if paiCount < length*5 || paiCount > 15 {
		return false
	}

	//找对子，找到对子后判断剩下的牌是否有比key大的飞机不带牌牌型
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	if laiZiCount == 0 {
		return largerAirDaiDui(pais, key, length)
	}

	danzhang := getPaiValueByCount(notLaiZi, 1)
	liangzhang := getPaiValueByCount(notLaiZi, 2)
	sanzhang := getPaiValueByCount(notLaiZi, 3)
	sizhang := getPaiValueByCount(notLaiZi, 4)

	var values []int32
	values = append(values, liangzhang...)
	values = append(values, sanzhang...)
	values = append(values, sizhang...)
	valueLen := len(values)
	danzhangLen := len(danzhang)

	if length == 2 {
		if laiZiCount >= 1 {
			//不找赖子
			if valueLen >= length {
				for i := 0; i < valueLen-1; i++ {
					for j := i + 1; j < valueLen; j++ {
						newPais := commonUtil.DeepClone(pais).([]*Poker)
						newPais = delPokerFromPaisByValue(values[i], newPais)
						newPais = delPokerFromPaisByValue(values[i], newPais)
						newPais = delPokerFromPaisByValue(values[j], newPais)
						newPais = delPokerFromPaisByValue(values[j], newPais)
						if largerAirBuDai_lai(newPais, key, length) {
							return true
						}
					}
				}
			}
			//取一个赖子作为带的对子
			if danzhangLen > 0 {
				for i := 0; i < valueLen; i++ {
					for j := 0; j < danzhangLen; j++ {
						if danzhang[j] == 16 || danzhang[j] == 17 {
							continue
						}
						newPais := commonUtil.DeepClone(pais).([]*Poker)
						newPais = delPokerFromPaisById(laiZi[0], newPais)
						newPais = delPokerFromPaisByValue(danzhang[j], newPais)
						newPais = delPokerFromPaisByValue(values[i], newPais)
						newPais = delPokerFromPaisByValue(values[i], newPais)
						if largerAirBuDai_lai(newPais, key, length) {
							return true
						}
					}
				}
			}
		}

		if laiZiCount >= 2 {
			//取两个赖子作为带牌
			//两个赖子补两个单张,作为两个对子
			if danzhangLen >= length {
				for i := 0; i < danzhangLen-1; i++ {
					for j := i + 1; j < danzhangLen; j++ {
						if danzhang[i] > 15 || danzhang[j] > 15 {
							continue
						}
						newPais := commonUtil.DeepClone(pais).([]*Poker)
						newPais = delPokerFromPaisById(laiZi[0], newPais)
						newPais = delPokerFromPaisById(laiZi[1], newPais)
						newPais = delPokerFromPaisByValue(danzhang[i], newPais)
						newPais = delPokerFromPaisByValue(danzhang[j], newPais)
						if largerAirBuDai_lai(newPais, key, length) {
							return true
						}
					}
				}
			}
			//两个赖子作为一个对子
			if valueLen >= 1 {
				for i := 0; i < valueLen; i++ {
					newPais := commonUtil.DeepClone(pais).([]*Poker)
					newPais = delPokerFromPaisById(laiZi[0], newPais)
					newPais = delPokerFromPaisById(laiZi[1], newPais)
					newPais = delPokerFromPaisByValue(values[i], newPais)
					newPais = delPokerFromPaisByValue(values[i], newPais)
					if largerAirBuDai_lai(newPais, key, length) {
						return true
					}
				}
			}
		}

		if laiZiCount >= 3 {
			//取三个赖子作为带牌
			for i := 0; i < danzhangLen; i++ {
				if danzhang[i] > 15 {
					continue
				}
				newPais := commonUtil.DeepClone(pais).([]*Poker)
				newPais = delPokerFromPaisById(laiZi[0], newPais)
				newPais = delPokerFromPaisById(laiZi[1], newPais)
				newPais = delPokerFromPaisById(laiZi[2], newPais)
				newPais = delPokerFromPaisByValue(danzhang[i], newPais)
				if largerAirBuDai_lai(newPais, key, length) {
					return true
				}
			}
		}

		if laiZiCount >= 4 {
			newPais := commonUtil.DeepClone(pais).([]*Poker)
			newPais = delPokerFromPaisById(laiZi[0], newPais)
			newPais = delPokerFromPaisById(laiZi[1], newPais)
			newPais = delPokerFromPaisById(laiZi[2], newPais)
			newPais = delPokerFromPaisById(laiZi[3], newPais)
			if largerAirBuDai_lai(newPais, key, length) {
				return true
			}
		}
	}

	if length == 3 {
		if laiZiCount >= 1 {
			//不找赖子
			if valueLen >= length {
				for i := 0; i < length-2; i++ {
					for j := i + 1; j < length-1; j++ {
						for k := j + 1; k < length; k++ {
							newPais := commonUtil.DeepClone(pais).([]*Poker)
							newPais = delPokerFromPaisByValue(values[i], newPais)
							newPais = delPokerFromPaisByValue(values[i], newPais)
							newPais = delPokerFromPaisByValue(values[j], newPais)
							newPais = delPokerFromPaisByValue(values[j], newPais)
							newPais = delPokerFromPaisByValue(values[k], newPais)
							newPais = delPokerFromPaisByValue(values[k], newPais)
							if largerAirBuDai_lai(newPais, key, length) {
								return true
							}
						}
					}
				}
			}

			//找一个赖子作为带牌
			if danzhangLen > 0 {
				if valueLen >= length-1 {
					for i := 0; i < valueLen-1; i++ {
						for j := i + 1; j < valueLen; j++ {
							for k := 0; k < danzhangLen; k++ {
								if danzhang[k] > 15 {
									continue
								}
								newPais := commonUtil.DeepClone(pais).([]*Poker)
								newPais = delPokerFromPaisById(laiZi[0], newPais)
								newPais = delPokerFromPaisByValue(danzhang[k], newPais)
								newPais = delPokerFromPaisByValue(values[i], newPais)
								newPais = delPokerFromPaisByValue(values[i], newPais)
								newPais = delPokerFromPaisByValue(values[j], newPais)
								newPais = delPokerFromPaisByValue(values[j], newPais)
								if largerAirBuDai_lai(newPais, key, length) {
									return true
								}
							}
						}
					}
				}
			}
		}
		if laiZiCount >= 2 {
			//找两个赖子作为带牌
			//两个赖子补两张单牌
			if valueLen > 0 && danzhangLen > 1 {
				for i := 0; i < valueLen; i++ {
					for j := 0; j < danzhangLen-1; j++ {
						for k := j + 1; k < danzhangLen; k++ {
							if danzhang[j] > 15 || danzhang[k] > 15 {
								continue
							}
							newPais := commonUtil.DeepClone(pais).([]*Poker)
							newPais = delPokerFromPaisById(laiZi[0], newPais)
							newPais = delPokerFromPaisById(laiZi[1], newPais)
							newPais = delPokerFromPaisByValue(values[i], newPais)
							newPais = delPokerFromPaisByValue(values[i], newPais)
							newPais = delPokerFromPaisByValue(danzhang[j], newPais)
							newPais = delPokerFromPaisByValue(danzhang[k], newPais)
							if largerAirBuDai_lai(newPais, key, length) {
								return true
							}
						}
					}
				}
			}
			//两个赖子补一个对子
			if valueLen >= 2 {
				for i := 0; i < valueLen-1; i++ {
					for j := i + 1; j < valueLen; j++ {
						newPais := commonUtil.DeepClone(pais).([]*Poker)
						newPais = delPokerFromPaisById(laiZi[0], newPais)
						newPais = delPokerFromPaisById(laiZi[1], newPais)
						newPais = delPokerFromPaisByValue(values[i], newPais)
						newPais = delPokerFromPaisByValue(values[i], newPais)
						newPais = delPokerFromPaisByValue(values[j], newPais)
						newPais = delPokerFromPaisByValue(values[j], newPais)
						if largerAirBuDai_lai(newPais, key, length) {
							return true
						}
					}
				}
			}
		}
		if laiZiCount >= 3 {
			//取三个赖子
			//三个赖子补三个单张
			if danzhangLen >= 3 {
				for i := 0; i < danzhangLen-2; i++ {
					for j := i + 1; j < danzhangLen-1; j++ {
						for k := j + 1; k < danzhangLen; k++ {
							if danzhang[i] > 15 || danzhang[j] > 15 || danzhang[k] > 15 {
								continue
							}
							newPais := commonUtil.DeepClone(pais).([]*Poker)
							newPais = delPokerFromPaisById(laiZi[0], newPais)
							newPais = delPokerFromPaisById(laiZi[1], newPais)
							newPais = delPokerFromPaisById(laiZi[2], newPais)
							newPais = delPokerFromPaisByValue(danzhang[i], newPais)
							newPais = delPokerFromPaisByValue(danzhang[j], newPais)
							newPais = delPokerFromPaisByValue(danzhang[k], newPais)
							if largerAirBuDai_lai(newPais, key, length) {
								return true
							}
						}
					}
				}
			}
			//三个赖子，补一个对子和一个单张
			if danzhangLen > 0 && valueLen > 0 {
				for i := 0; i < valueLen; i++ {
					for j := 0; j < danzhangLen; j++ {
						if danzhang[j] > 15 {
							continue
						}
						newPais := commonUtil.DeepClone(pais).([]*Poker)
						newPais = delPokerFromPaisById(laiZi[0], newPais)
						newPais = delPokerFromPaisById(laiZi[1], newPais)
						newPais = delPokerFromPaisById(laiZi[2], newPais)
						newPais = delPokerFromPaisByValue(danzhang[j], newPais)
						newPais = delPokerFromPaisByValue(values[i], newPais)
						newPais = delPokerFromPaisByValue(values[i], newPais)
						if largerAirBuDai_lai(newPais, key, length) {
							return true
						}
					}
				}
			}
		}
		if laiZiCount >= 4 {
			//四个赖子补一个对子和两个单张
			if valueLen > 0 && danzhangLen >= 2 {
				for j := 0; j < danzhangLen-1; j++ {
					for k := j + 1; k < danzhangLen; k++ {
						if danzhang[j] > 15 || danzhang[k] > 15 {
							continue
						}
						newPais := commonUtil.DeepClone(pais).([]*Poker)
						newPais = delPokerFromPaisById(laiZi[0], newPais)
						newPais = delPokerFromPaisById(laiZi[1], newPais)
						newPais = delPokerFromPaisById(laiZi[2], newPais)
						newPais = delPokerFromPaisById(laiZi[3], newPais)
						newPais = delPokerFromPaisByValue(danzhang[j], newPais)
						newPais = delPokerFromPaisByValue(danzhang[k], newPais)
						if largerAirBuDai_lai(newPais, key, length) {
							return true
						}
					}
				}
			}
			//四个赖子补两个对子
			if valueLen > 0 {
				for i := 0; i < valueLen; i++ {
					newPais := commonUtil.DeepClone(pais).([]*Poker)
					newPais = delPokerFromPaisById(laiZi[0], newPais)
					newPais = delPokerFromPaisById(laiZi[1], newPais)
					newPais = delPokerFromPaisById(laiZi[2], newPais)
					newPais = delPokerFromPaisById(laiZi[3], newPais)
					newPais = delPokerFromPaisByValue(values[i], newPais)
					newPais = delPokerFromPaisByValue(values[i], newPais)
					if largerAirBuDai_lai(newPais, key, length) {
						return true
					}
				}
			}
		}
	}
	return false
}

//比指定key大的炸弹，纯赖子炸弹比普通炸弹获取带赖子的炸弹大
func largerBoom_lai(pais []*Poker, key int32) bool {
	paiCount := len(pais)
	if paiCount < 4 {
		return false
	}
	laiZi, notLaiZi := getLaiZiFromPais(pais)
	laiZiCount := len(laiZi)

	if laiZiCount == 0 {
		return largerBoom(pais, key)
	}
	if laiZiCount >= 1 {
		if sanzhang := getPaiValueByCount(notLaiZi, 3); len(sanzhang) > 0 {
			for _, v := range sanzhang {
				if v > key {
					return true
				}
			}
		}

	}
	if laiZiCount >= 2 {
		if liangzhang := getPaiValueByCount(notLaiZi, 2); len(liangzhang) > 0 {
			for _, v := range liangzhang {
				if v > key {
					return true
				}
			}
		}
	}
	if laiZiCount >= 3 {
		if danzhang := getPaiValueByCount(notLaiZi, 1); len(danzhang) > 0 {
			for _, v := range danzhang {
				if v <= 15 && v > key {
					return true
				}
			}
		}
	}
	if laiZiCount >= 4 {
		return true
	}
	return false
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
