package doudizhu

import (
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
