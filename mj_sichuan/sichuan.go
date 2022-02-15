package mj_sichuan

import (
	"ALG-QiPai/cards"
)

// MJGroup 麻将胡牌组合
type MJGroup struct {
	Type  cards.GroupType // 胡牌组合类型: 将对、刻子、顺子
	Suit  cards.CardSuit  // 组合的花色:　万条筒
	Cards []int           // 组合有哪些牌值
}

func (mg MJGroup) IsZero() bool {
	return mg.Suit == cards.CardSuitNone &&
		mg.Type == cards.GroupTypeNone &&
		len(mg.Cards) == 0
}

func getCardsCountArray(handCards []cards.Card) (paiCountInfo [][]int32, paiCount int) {
	paiCountInfo = make([][]int32, 3)
	for _, c := range handCards {
		paiCount += 1
		s, v := c.Suit(), c.Value()
		countInfo := paiCountInfo[s-1]
		if len(countInfo) == 0 {
			countInfo = make([]int32, 10)
			paiCountInfo[s-1] = countInfo
		}
		countInfo[v] += 1
	}
	return
}

// CheckHu 判断能否胡牌
func CheckHu(handCards []cards.Card, target cards.Card) (canHu bool, groups map[int][]MJGroup) {
	if target != cards.CardNone {
		handCards = append(handCards, target)
	}

	paiCountInfo, paiCount := getCardsCountArray(handCards)
	if paiCount%3 != 2 {
		return
	}
	groups = make(map[int][]MJGroup)
	if paiCount == 14 {
		isQiDui, qiDuiGroups := tryQiDui(paiCountInfo)
		if isQiDui {
			groups[-1] = qiDuiGroups
		}
	}
	huGroups := tryHu(paiCountInfo, paiCount)
	for gId, gs := range huGroups {
		groups[gId] = gs
	}
	canHu = len(groups) > 0
	return
}

func tryQiDui(paiCountInfo [][]int32) (canHu bool, groups []MJGroup) {
outLoop:
	for suit, countInfo := range paiCountInfo {
		for i, n := range countInfo {
			if i == 0 || n == 0 {
				continue
			}
			if n%2 != 0 {
				break outLoop
			}
			subGroup := MJGroup{
				Type:  cards.GroupTypeMJJiang,
				Suit:  cards.CardSuit(suit + 1),
				Cards: []int{i, i},
			}
			switch n {
			case 2:
				groups = append(groups, subGroup)
			case 4:
				groups = append(groups, subGroup, subGroup)
			}
		}
	}
	canHu = len(groups) == 7
	return
}

func tryHu(paiCountInfo [][]int32, paiCount int) (groups map[int][]MJGroup) {
	groups = make(map[int][]MJGroup)
	groupId := new(int)

	for suit, countInfo := range paiCountInfo {
		for i, cnt := range countInfo {
			if i == 0 || cnt < 2 {
				continue
			}
			countInfo[i] -= 2 // 扣除两张牌作为将牌
			subGroups := make(map[int][]MJGroup)
			canHu := tryKeAndShun(paiCountInfo, paiCount-2, 1, groupId, subGroups)
			if canHu {
				jiangGroup := MJGroup{
					Suit:  cards.CardSuit(suit + 1),
					Type:  cards.GroupTypeMJJiang,
					Cards: []int{i, i},
				}
				for gId, gs := range subGroups {
					gs = append(gs, jiangGroup)
					groups[gId] = gs
				}
				if len(groups) == 0 {
					*groupId += 1
					groups[*groupId] = []MJGroup{jiangGroup}
				}
			}
			countInfo[i] += 2
		}
	}
	return
}

func tryKeAndShun(paiCountInfo [][]int32, paiCount, level int, groupId *int, groups map[int][]MJGroup) (canHu bool) {
	if paiCount == 0 {
		canHu = true
		return
	}
	for suit, countInfo := range paiCountInfo {
		if len(countInfo) == 0 {
			continue
		}
		// 找刻子
		for i, cnt := range countInfo {
			if i == 0 || cnt < 3 {
				continue
			}
			countInfo[i] -= 3
			canHu = tryKeAndShun(paiCountInfo, paiCount-3, level+1, groupId, groups)
			if canHu {
				subGroup := MJGroup{
					Suit:  cards.CardSuit(suit + 1),
					Type:  cards.GroupTypeMJKe,
					Cards: []int{i, i, i},
				}
				groups[*groupId] = append(groups[*groupId], subGroup)
				countInfo[i] += 3
				if level == 1 {
					*groupId += 1
				} else {
					return
				}
			} else {
				countInfo[i] += 3
				if level == 1 {
					delete(groups, *groupId)
					*groupId += 1
				}
			}
		}

		// 找顺子
		for i := 1; i < 8; i++ {
			if countInfo[i] == 0 || countInfo[i+1] == 0 || countInfo[i+2] == 0 {
				continue
			}
			countInfo[i] -= 1
			countInfo[i+1] -= 1
			countInfo[i+2] -= 1
			canHu = tryKeAndShun(paiCountInfo, paiCount-3, level+1, groupId, groups)
			if canHu {
				subGroup := MJGroup{
					Type:  cards.GroupTypeMJShun,
					Suit:  cards.CardSuit(suit + 1),
					Cards: []int{i, i + 1, i + 2},
				}
				groups[*groupId] = append(groups[*groupId], subGroup)
				countInfo[i] += 1
				countInfo[i+1] += 1
				countInfo[i+2] += 1
				if level == 1 {
					*groupId += 1
				} else {
					return
				}
			} else {
				countInfo[i] += 1
				countInfo[i+1] += 1
				countInfo[i+2] += 1
				if level == 1 {
					delete(groups, *groupId)
					*groupId += 1
				}
			}
		}
	}
	return
}
