package doudizhu

import "sort"

/*
    @Create by GoLand
    @Author: hong
    @Time: 2018/7/26 18:12 
    @File: parser_lai.go    
*/

/*牌型解析：带赖子牌*/
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
	//TODO
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
