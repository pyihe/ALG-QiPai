package texas

import "sort"

type texasPokerType int

const (
	unknown texasPokerType = iota
	highCard
	onePair
	twoPair
	threeKind
	straight
	flush
	fullHouse
	fourKind
	straightFlush
	royalStraightFlush
)

type PokerList []*Poker

//排序
func (pokerList PokerList) Len() int {
	return len(pokerList)
}

func (pokerList PokerList) Swap(i, j int) {
	pokerList[i], pokerList[j] = pokerList[j], pokerList[i]
}

func (pokerList PokerList) Less(i, j int) bool {
	return pokerList[i].GetValue() > pokerList[j].GetValue()
}

func (pokerList PokerList) GetPokerIds() (ids []int32) {
	if len(pokerList) == 0 {
		return
	}
	for _, p := range pokerList {
		if p != nil {
			ids = append(ids, p.GetId())
		}
	}
	return
}

func (pokerList PokerList) Desc(args ...interface{}) (desc string) {
	if len(args) > 0 {
		switch b := args[0].(type) {
		case bool:
			if b {
				sort.Sort(pokerList)
			}
		}
	}
	for _, p := range pokerList {
		if p == nil {
			continue
		}
		desc += p.Desc()
	}
	return
}

func (pokerList PokerList) GetValue() (value int32) {
	for _, p := range pokerList {
		if p != nil {
			value += p.GetValue()
		}
	}
	return
}

//牌型解析
func (pokerList PokerList) GetTopType() (PokerList, texasPokerType, int32) {
	switch len(pokerList) {
	case 2:
		return pokerList.twoPoker()
	case 5:
		return pokerList.fivePoker()
	case 6:
		return pokerList.sixPoker()
	case 7:
		return pokerList.sevenPoker()
	default:
		return PokerList{}, unknown, 0
	}
}

//解析两张牌的牌型
func (pokerList PokerList) twoPoker() (PokerList, texasPokerType, int32) {
	key := pokerList[0].GetValue() + pokerList[1].GetValue()
	if pokerList[0].GetValue() == pokerList[1].GetValue() {
		return pokerList, onePair, key
	}
	if pokerList[0].GetFlower() == pokerList[1].GetFlower() {
		return pokerList, flush, key
	}
	return pokerList, highCard, key
}

//解析5张牌
func (pokerList PokerList) fivePoker() (pokers PokerList, pokerType texasPokerType, value int32) {
	pokers = pokerList
	value = pokerList.getValues()
	pokerType = highCard
	if ok, isRoyal, _ := pokerList.GetStraightFlush(); ok {
		if isRoyal {
			pokerType = royalStraightFlush
		} else {
			pokerType = straightFlush
		}
		return
	}

	if ok, _ := pokerList.GetFourKind(); ok {
		pokerType = fourKind
		return
	}

	if ok, _ := pokerList.GetFullHouse(); ok {
		pokerType = fullHouse
		return
	}

	if ok, _ := pokerList.GetFlush(); ok {
		pokerType = flush
		return
	}

	if ok, _ := pokerList.GetStraight(); ok {
		pokerType = straight
		return
	}
	if ok, _ := pokerList.GetThreeKind(); ok {
		pokerType = threeKind
		return
	}
	if ok, tPair := pokerList.GetPairs(); ok {
		if len(tPair) >= 4 {
			pokerType = twoPair
			return
		} else {
			pokerType = onePair
			return
		}
	}
	return
}

//解析6张牌
func (pokerList PokerList) sixPoker() (pokers PokerList, pokerType texasPokerType, value int32) {
	for i := 0; i < len(pokerList); i++ {
		tempPokers := pokerList.DeletePokers(pokerList[i])
		list, pType, pValue := tempPokers.fivePoker()
		if pType > pokerType {
			pokers = list
			pokerType = pType
			value = pValue
		}
		if pType == pokerType && pValue > value {
			pokers = list
			pokerType = pType
			value = pValue
		}
	}
	return
}

//解析7张牌
func (pokerList PokerList) sevenPoker() (pokers PokerList, pokerType texasPokerType, value int32) {
	for i := 0; i < len(pokerList); i++ {
		for j := i + 1; j < len(pokerList); j++ {
			tempPokers := pokerList.DeletePokers(pokerList[i], pokerList[j])
			list, pType, pValue := tempPokers.fivePoker()
			if pType > pokerType {
				pokers = list
				pokerType = pType
				value = pValue
			}
			if pType == pokerType && pValue > value {
				pokers = list
				pokerType = pType
				value = pValue
			}
		}
	}
	return
}

func (pokerList PokerList) Copy() PokerList {
	var list = make(PokerList, len(pokerList))
	copy(list, pokerList)
	return list
}

func DeleteOnePoker(p *Poker, pokerList PokerList) PokerList {
	for i, poker := range pokerList {
		if poker.GetId() == p.GetId() {
			pokerList = append(pokerList[:i], pokerList[i+1:]...)
		}
	}
	return pokerList
}

//从牌组里删除指定ID的牌
func (pokerList PokerList) DeletePokers(pokers ...*Poker) (copyList PokerList) {
	copyList = pokerList.Copy()
	for _, p := range pokers {
		copyList = DeleteOnePoker(p, copyList)
	}
	return
}

//获取所有对子
func (pokerList PokerList) GetPairs() (ok bool, pairs PokerList) {
	values := pokerList.getValueAndCount()
	//后面开始遍历，从大的牌值开始获取对子
	for i := len(values) - 1; i >= 0; i-- {
		if values[i] >= 2 {
			ok = true
			pairs = append(pairs, pokerList.getPokerByValue(int32(i))[:2]...)
		}
	}
	return
}

//获取三条
func (pokerList PokerList) GetThreeKind() (ok bool, threeKind PokerList) {
	values := pokerList.getValueAndCount()
	for i := len(values) - 1; i >= 0; i-- {
		if values[i] >= 3 {
			ok = true
			threeKind = append(threeKind, pokerList.getPokerByValue(int32(i))[:3]...)
		}
	}
	return
}

//获取最大的顺子
func (pokerList PokerList) GetStraight() (ok bool, straight PokerList, ) {
	if straights := pokerList.getStraight(); len(straights) > 0 {
		ok = true
		straight = straights[0]
	}
	return
}

//获取最大的同花
func (pokerList PokerList) GetFlush() (ok bool, flush PokerList) {
	flowerCount := pokerList.getFlowerCount()
	for flower, count := range flowerCount {
		if count >= 5 {
			ok = true
			flush = append(flush, pokerList.getPokerByFlower(int32(flower))[:5]...)
			return
		}
	}
	return
}

//获取最大的葫芦
func (pokerList PokerList) GetFullHouse() (ok bool, house PokerList) {
	yes, threeKind := pokerList.GetThreeKind()
	if !yes {
		return
	}

	list := pokerList.DeletePokers(threeKind...)
	yes, pairs := list.GetPairs()
	if !yes {
		return
	}
	house = append(house, threeKind...)
	house = append(house, pairs[:2]...)
	ok = true
	return
}

//获取四条
func (pokerList PokerList) GetFourKind() (ok bool, fourKind PokerList) {
	values := pokerList.getValueAndCount()
	for value, cnt := range values {
		if cnt == 4 {
			ok = true
			fourKind = append(fourKind, pokerList.getPokerByValue(int32(value))[:4]...)
		}
	}
	return
}

//获取最大的同花顺
func (pokerList PokerList) GetStraightFlush() (ok bool, isRoyal bool, straightFlush PokerList) {
	straights := pokerList.getStraight()
	if len(straights) == 0 {
		return
	}
	for _, s := range straights {
		ok, straightFlush = s.GetFlush()
		if ok {
			if straightFlush[0].GetValue() == 14 {
				isRoyal = true
			}
			break
		}
	}
	return
}

//根据牌值获取所有对应的牌
func (pokerList PokerList) getPokerByValue(value int32) (list PokerList) {
	for _, p := range pokerList {
		if p.GetValue() == value {
			list = append(list, p)
		}
	}
	return
}

//获取一组牌的牌值和数量
func (pokerList PokerList) getValueAndCount() (values []int32) {
	values = make([]int32, 15)
	for _, p := range pokerList {
		values[p.GetValue()]++
	}
	return
}

//获取一组牌同花色的牌的数量
func (pokerList PokerList) getFlowerCount() (flower []int32) {
	flower = make([]int32, 5)
	for _, p := range pokerList {
		flower[p.GetFlower()]++
	}
	return
}

//获取同花色的牌
func (pokerList PokerList) getPokerByFlower(flower int32) (list PokerList) {
	for _, p := range pokerList {
		if p.GetFlower() == flower {
			list = append(list, p)
		}
	}
	sort.Sort(list)
	return
}

//获取所有的顺子
func (pokerList PokerList) getStraight() (straights []PokerList) {
	values := pokerList.getValueAndCount()
	for i := len(values) - 1; i >= 5; i-- {
		var tempPokers PokerList
		for j := i; j > i-5; j-- {
			if values[j] == 0 {
				break
			}
			tempPokers = append(tempPokers, pokerList.getPokerByValue(int32(j))[:1]...)
		}
		if len(tempPokers) == 5 {
			straights = append(straights, tempPokers)
		}
	}
	//判断A2345
	if values[2] > 0 && values[3] > 0 && values[4] > 0 && values[5] > 0 && values[14] > 0 {
		valueList := []int32{2, 3, 4, 5, 14}
		var pais PokerList
		for _, v := range valueList {
			pais = append(pais, pokerList.getPokerByValue(v)[:1]...)
		}
		straights = append(straights, pais)
	}
	return
}

//获取一组牌的牌值总和
func (pokerList PokerList) getValues() (value int32) {
	sort.Sort(pokerList)
	if pokerList[0].GetValue() == 14 && pokerList[4].GetValue() == 2 && pokerList[1].GetValue()-pokerList[4].GetValue() == 3 {
		value = 15 //1 + 2 + 3 + 4 + 5
		return
	}
	for _, p := range pokerList {
		value += p.GetValue()
	}
	return
}


//根据一组ID初始化一组牌
func InitPokersByIds(indexs []int32) PokerList {
	var pokers []*Poker
	for _, id := range indexs {
		p := InitPokerById(id)
		pokers = append(pokers, p)
	}
	return pokers
}

//洗牌
func XiPai() PokerList {
	return InitPokersByIds(shuffleXiPai(AllPokerNum))
}
