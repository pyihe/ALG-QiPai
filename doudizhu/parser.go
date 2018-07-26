package doudizhu

import (
	"sort"
)

/*
    @Create by GoLand
    @Author: hong
    @Time: 2018/7/26 18:08 
    @File: parser.go    
*/

//牌型解析：不带赖子牌

type PaiValueList []int32

func (l PaiValueList) Len() int {
	return len(l)
}

func (l PaiValueList) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l PaiValueList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

/*辅助方法*/
//从一组牌里找到指定张数的牌的牌值切片
func getPaiValueByCount(pais []*Poker, cnt int) (danZhangValue []int32) {
	if pais == nil || len(pais) <= 0 {
		return
	}

	paiValues := getPaiValueList(pais)
	for _, v := range paiValues {
		var count int
		for _, p := range pais {
			if p.GetValue() == v {
				count++
			}
		}
		if count == cnt {
			danZhangValue = append(danZhangValue, v)
		}
	}
	return
}

//获取一组牌去重后的牌值列表
func getPaiValueList(pais []*Poker) (values []int32) {
	if len(pais) <= 0 || pais == nil {
		return
	}
	for _, p := range pais {
		if p == nil {
			continue
		}
		values = append(values, p.GetValue())
	}
	values = removeValueRepetition(values)
	return
}

//[]int32{} 去重
func removeValueRepetition(values []int32) (nonRepetitionValues []int32) {
	for i := range values {
		flag := true
		for j := range nonRepetitionValues {
			if values[i] == nonRepetitionValues[j] {
				flag = false
				break
			}
		}
		if flag {
			nonRepetitionValues = append(nonRepetitionValues, values[i])
		}
	}
	return
}

/*******判断一组牌是否是指定的牌型， 返回判断结果和对应的key, 如果不是，则key为-1*******/
//是否是单牌
func isDanPai(pais []*Poker) (bool, int32) {
	if len(pais) == 1 {
		return true, pais[0].GetValue()
	}
	return false, -1
}

//是否是对子
func isDuiZi(pais []*Poker) (bool, int32) {
	if len(pais) == 2 && pais[0].GetValue() == pais[1].GetValue() {
		return true, pais[0].GetValue()
	}
	return false, -1
}

//是否是三不带
func isSanBuDai(pais []*Poker) (bool, int32) {
	if len(pais) == 3 {
		if pais[0].GetValue() == pais[1].GetValue() && pais[1].GetValue() == pais[2].GetValue() {
			return true, pais[0].GetValue()
		}
	}
	return false, -1
}

//是否是三带一
func isSanDaiYi(pais []*Poker) (bool, int32) {
	sanzhang := getPaiValueByCount(pais, 3)
	danzhang := getPaiValueByCount(pais, 1)
	if len(sanzhang) == 1 && len(danzhang) == 1 {
		return true, sanzhang[0]
	}
	return false, -1
}

//是否是三带对子
func isSanDaiDui(pais []*Poker) (bool, int32) {
	sanzhang := getPaiValueByCount(pais, 3)
	duizi := getPaiValueByCount(pais, 2)
	if len(sanzhang) == 1 && len(duizi) == 1 {
		return true, sanzhang[0]
	}
	return false, -1
}

//是否是顺子
func isShunZi(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 5 {
		return false, -1
	}
	if danzhang := getPaiValueByCount(pais, 1); len(danzhang) != paiCount {
		return false, -1
	}

	//对牌值从小到大排序
	paiValues := getPaiValueList(pais)
	sort.Sort(PaiValueList(paiValues))

	valueLen := len(paiValues)
	if paiValues[valueLen-1]-paiValues[0]+1 == int32(valueLen) {
		return true, paiValues[0]
	}
	return false, -1
}

//是否是连对
func isLianDui(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 6 || paiCount%2 != 0 {
		return false, -1
	}

	duizi := getPaiValueByCount(pais, 2)
	duiZiLen := len(duizi)
	if duiZiLen != paiCount/2 {
		return false, -1
	}

	//对子的牌值排序
	sort.Sort(PaiValueList(duizi))
	if duizi[duiZiLen-1]-duizi[0]+1 == int32(duiZiLen) {
		return true, duizi[0]
	}
	return false, -1
}

//是否是飞机不带牌
func isAirBuDai(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 6 || paiCount%3 != 0 {
		return false, -1
	}

	sanzhang := getPaiValueByCount(pais, 3)
	sanzhangLen := len(sanzhang)
	if sanzhangLen != paiCount/3 {
		return false, -1
	}

	//排序
	sort.Sort(PaiValueList(sanzhang))
	if sanzhang[sanzhangLen-1]-sanzhang[0]+1 == int32(sanzhangLen) {
		return true, sanzhang[0]
	}
	return false, -1
}

//是否是飞机带单牌
func isAirDaiDan(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 8 || paiCount%4 != 0 {
		return false, -1
	}
	sanzhang := getPaiValueByCount(pais, 3)
	danzhang := getPaiValueByCount(pais, 1)
	sanzhangLen := len(sanzhang)
	danzhangLen := len(danzhang)

	if sanzhangLen != danzhangLen || sanzhangLen*3+danzhangLen != paiCount {
		return false, -1
	}
	sort.Sort(PaiValueList(sanzhang))

	if sanzhang[sanzhangLen-1]-sanzhang[0]+1 == int32(sanzhangLen) {
		return true, sanzhang[0]
	}
	return false, -1
}

//是否是飞机带对子
func isAirDaiDui(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount < 10 || paiCount%5 != 0 {
		return false, -1
	}

	sanzhang := getPaiValueByCount(pais, 3)
	duizi := getPaiValueByCount(pais, 2)
	sanzhangLen := len(sanzhang)
	duiziLen := len(duizi)

	if sanzhangLen != duiziLen || sanzhangLen*3+duiziLen*2 != paiCount {
		return false, -1
	}
	sort.Sort(PaiValueList(sanzhang))

	if sanzhang[sanzhangLen-1]-sanzhang[0]+1 == int32(sanzhangLen) {
		return true, sanzhang[0]
	}
	return false, -1
}

//是否是四带单牌
func isSiDaiDan(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount != 6 {
		return false, -1
	}

	sizhang := getPaiValueByCount(pais, 4)
	danzhang := getPaiValueByCount(pais, 1)
	sizhangLen := len(sizhang)
	danzhangLen := len(danzhang)

	if sizhangLen == 1 && danzhangLen == 2 {
		return true, sizhang[0]
	}
	return false, -1
}

//是否是四带对子
func isSiDaiDui(pais []*Poker) (bool, int32) {
	paiCount := len(pais)
	if paiCount != 8 {
		return false, -1
	}
	sizhang := getPaiValueByCount(pais, 4)
	duizi := getPaiValueByCount(pais, 2)
	if len(sizhang) == 1 && len(duizi) == 2 {
		return true, sizhang[0]
	}
	return false, -1
}

//是否是炸弹
func isBoom(pais []*Poker) (bool, int32) {
	paiCount := len(pais)

	sizhang := getPaiValueByCount(pais, 4)
	if sizhangLen := len(sizhang); sizhangLen == 1 && sizhangLen == paiCount {
		return true, sizhang[0]
	}
	return false, -1
}

//是否是王炸
func isSuperBoom(pais []*Poker) (bool) {
	paiCount := len(pais)
	if paiCount == 2 {
		if pais[0].GetValue() == 16 && pais[1].GetValue() == 17 {
			return true
		}
		if pais[1].GetValue() == 16 && pais[0].GetValue() == 17 {
			return true
		}
	}
	return false
}

/*从一组牌中找出是否包含指定的牌型*/
//比指定key大的单牌
func largerDanPai(pais []*Poker, key int32) bool {
	var danPaiValue []int32
	values := getPaiValueList(pais)

	for _, v := range values {
		if v > key {
			danPaiValue = append(danPaiValue, v)
		}
	}

	return len(danPaiValue) > 0
}

//比指定key大的对子
func largerDuiZi(pais []*Poker, key int32) bool {
	var duiValue []int32

	duizi := getPaiValueByCount(pais, 2)
	for _, v := range duizi {
		if v > key {
			duiValue = append(duiValue, v)
		}
	}
	return len(duiValue) > 0
}

//比指定key大的三张
func largerSanZhang(pais []*Poker, key int32) bool {
	var sanValue []int32

	sanzhang := getPaiValueByCount(pais, 3)
	for _, v := range sanzhang {
		if v > key {
			sanValue = append(sanValue, v)
		}
	}
	return len(sanValue) > 0
}

//比指定key大的三带一
func largerSanDaiDan(pais []*Poker, key int32) bool {
	paiCount := len(pais)
	if paiCount < 4 {
		return false
	}

	return largerSanZhang(pais, key)
}

//比指定key大的三带对子
func largerSanDaiDui(pais []*Poker, key int32) bool {
	paiCount := len(pais)
	if paiCount < 5 {
		return false
	}

	duizi := getPaiValueByCount(pais, 2)
	sanzhang := getPaiValueByCount(pais, 3)
	if largerSanZhang(pais, key) && (len(duizi) > 0 || len(sanzhang) > 1) {
		return true
	}
	return false
}

//比指定key大的顺子
func largerShunZi(pais []*Poker, key int32, length int) bool {
	paiCount := len(pais)
	if paiCount < length {
		return false
	}

	valueList := getPaiValueList(pais)
	if len(valueList) < length {
		return false
	}
	sort.Sort(PaiValueList(valueList))
	if len(valueList) == length {
		if valueList[len(valueList)-1]-valueList[0]+1 != int32(length) {
			return false
		}
		if valueList[len(valueList)-1] > 14 {
			return false
		}
		return true
	}
	for i := 0; i < len(valueList)-length+1; i++ {
		if valueList[i+length-1]-valueList[i]+1 != int32(length) {
			continue
		}
		if valueList[i] > 14 || valueList[i+length-1] > 14 {
			continue
		}
		if valueList[i] <= key {
			continue
		}
		return true
	}
	return false
}

//比指定key大的连对
func largerLianDui(pais []*Poker, key int32, length int) bool {
	values := getPaiValueByCount(pais, 2)
	if len(values) < length {
		return false
	}

	sort.Sort(PaiValueList(values))
	if len(values) == length {
		if values[len(values)-1]-values[0] != int32(length) {
			return false
		}
		if values[len(values)-1] > 14 {
			return false
		}
		return true
	}
	for i := 0; i < len(values)-length+1; i++ {
		if values[i+length-1]-values[i]+1 != int32(length) {
			continue
		}
		if values[i] > 14 || values[i+length-1] > 14 {
			continue
		}
		if values[i] <= key {
			continue
		}
		return true
	}
	return false
}

//比指定key大的飞机不带牌
func largerAirBuDai(pais []*Poker, key int32, length int) bool {
	values := getPaiValueByCount(pais, 3)
	if len(values) < length {
		return false
	}

	sort.Sort(PaiValueList(values))
	if len(values) == length {
		if values[len(values)-1]-values[0] != int32(length) {
			return false
		}
		if values[len(values)-1] > 14 {
			return false
		}
		return true
	}
	for i := 0; i < len(values)-length+1; i++ {
		if values[i+length-1]-values[i]+1 != int32(length) {
			continue
		}
		if values[i] > 14 || values[i+length-1] > 14 {
			continue
		}
		if values[i] <= key {
			continue
		}
		return true
	}
	return false
}

//比指定key大的飞机带单牌
func largerAirDaiDan(pais []*Poker, key int32, length int) bool {
	paiCount := len(pais)
	if paiCount < length*4 {
		return false
	}

	return largerAirBuDai(pais, key, length)
}

//比指定key大的飞机带对子
func largerAireDaiDui(pais []*Poker, key int32, length int) bool {
	paiCount := len(pais)
	if paiCount < length*5 {
		return false
	}

	if largerAirBuDai(pais, key, length) {
		sanzhang := getPaiValueByCount(pais, 3)
		duizi := getPaiValueByCount(pais, 2)
		if len(sanzhang) > length+1 || len(duizi) >= length {
			return true
		}
	}
	return false
}

//比指定key大的炸弹
func largerBoom(pais []*Poker, key int32) bool {
	sizhang := getPaiValueByCount(pais, 4)
	for _, v := range sizhang {
		if v > key {
			return true
		}
	}
	return false
}
