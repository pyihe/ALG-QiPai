package wenling

import (
	"ALG-QiPai/cards"
)

type Player struct {
	ZWFX          cards.ZWFX     // 玩家座位方向
	HandCards     []cards.Card   // 手牌
	HuaCards      []cards.Card   // 补花的牌
	PengCards     [][]cards.Card // 碰了的牌
	ChiCards      [][]cards.Card // 吃了的牌
	MingGangCards [][]cards.Card // 明杠的牌
	AnGangCards   [][]cards.Card // 暗杠的牌
	BuGangCards   [][]cards.Card // 补杠的牌
}

// IsCaiShen 普通闹闹情况下，判断target是否是财神牌
func IsCaiShen(caiShen, target cards.Card) bool {
	if !caiShen.Valid() || !target.Valid() {
		return false
	}
	cSuit := caiShen.Suit()
	switch cSuit {
	case cards.CardSuitCharacter, cards.CardSuitBamboo, cards.CardSuitDot, cards.CardSuitWind, cards.CardSuitDragon:
		return caiShen == target
	case cards.CardSuitSeason, cards.CardSuitFlower:
		tSuit := target.Suit()
		return tSuit == cSuit
	default:
		return false
	}
}

// CheckLaZiHu 判断是否可以辣子胡
// 辣子胡作为温岭麻将的特殊胡法，直接根据手牌、花牌、碰牌、杠牌进行各种判断即可
// 1. 集齐春夏秋冬或者梅兰竹菊
// 2. 乱花推土: 集齐牌值为1、2、3、4的花牌
// 3. 集齐四本门风
// 4. 集齐四红中
// 5. 集齐四白板
// 6. 集齐四发财
// 7. 碰了东南西北
// 8. 三财神
func CheckLaZiHu(player *Player, caiShen cards.Card) (canHu bool) {
	var (
		ownWindCard    cards.Card       // 本门风牌
		caiShenCnt     int              // 财神牌数量
		seasonCnt      int              // 春夏秋冬数量
		flowerCnt      int              // 梅兰竹菊数量
		redCnt         int              // 红中数量
		greenCnt       int              // 发财数量
		whiteCnt       int              // 白板数量
		ownWindCnt     int              // 本门风数量
		pengWindCnt    int              // 碰东南西北的数量
		bulldozingInfo = make([]int, 4) // 乱花推土牌值与数量信息
	)
	switch player.ZWFX {
	case cards.ZWFXDong:
		ownWindCard = cards.CardEastWind
	case cards.ZWFXNan:
		ownWindCard = cards.CardSouthWind
	case cards.ZWFXXi:
		ownWindCard = cards.CardWestWind
	case cards.ZWFXBei:
		ownWindCard = cards.CardNorthWind
	}

	// 判断手牌
	for _, c := range player.HandCards {
		if IsCaiShen(caiShen, c) {
			caiShenCnt += 1
			continue
		}
		s, v := c.Suit(), c.Value()
		switch s {
		case cards.CardSuitSeason: // 正常情况下手牌里不会有花牌，但是考虑到最后一张摸起来是花牌，此时不能再补花，所以手牌里可能出现花牌
			seasonCnt += 1
			bulldozingInfo[v-1] += 1
		case cards.CardSuitFlower:
			flowerCnt += 1
			bulldozingInfo[v-1] += 1
		case cards.CardSuitWind:
			if c == ownWindCard {
				ownWindCnt += 1
			}
		case cards.CardSuitDragon:
			switch c {
			case cards.CardRedDragon:
				redCnt += 1
			case cards.CardGreenDragon:
				greenCnt += 1
			case cards.CardWhiteDragon:
				whiteCnt += 1
			}
		}
	}

	// 判断花牌
	for _, c := range player.HuaCards {
		s, v := c.Suit(), c.Value()
		switch s {
		case cards.CardSuitSeason:
			seasonCnt += 1
			bulldozingInfo[v-1] += 1
		case cards.CardSuitFlower:
			flowerCnt += 1
			bulldozingInfo[v-1] += 1
		case cards.CardSuitDragon: // 闹闹时，中发白也需要补花到桌面
			switch c {
			case cards.CardRedDragon:
				redCnt += 1
			case cards.CardGreenDragon:
				greenCnt += 1
			case cards.CardWhiteDragon:
				whiteCnt += 1
			}
		}
	}

	// 是否碰了东南西北
	for _, data := range player.PengCards {
		if data[0].Suit() == cards.CardSuitWind {
			pengWindCnt += 1
		}
	}

	// 暗杠里是否有东南西北、中发白
	for _, data := range player.AnGangCards {
		s := data[0].Suit()
		switch s {
		case cards.CardSuitWind:
			if data[0] == ownWindCard {
				ownWindCnt += 4
			}
		case cards.CardSuitDragon:
			switch data[0] {
			case cards.CardRedDragon:
				redCnt += 1
			case cards.CardGreenDragon:
				greenCnt += 1
			case cards.CardWhiteDragon:
				whiteCnt += 1
			}
		}
	}

	switch {
	case caiShenCnt == 3: // 三财神
		return true
	case seasonCnt == 4: // 集齐春夏秋冬
		return true
	case flowerCnt == 4: // 集齐梅兰竹菊
		return true
	case redCnt == 4: // 四红中
		return true
	case greenCnt == 4: // 四发财
		return true
	case whiteCnt == 4: // 四白板
		return true
	case ownWindCnt == 4: // 四本门风
		return true
	case pengWindCnt == 4: // 碰了东南西北
		return true
	default:
		// 是否乱花推土
		for _, n := range bulldozingInfo {
			if n == 0 {
				return false
			}
		}
		return true
	}
}

// CheckHu 判断是否可以胡牌
// handCards: 手牌
// caiShen: 财神牌
// target: 自己摸的牌或者别人打的牌
func CheckHu(handCards []cards.Card, caiShen, target cards.Card) (canHu bool, groups map[int][]cards.MJGroup) {
	if target.Valid() {
		handCards = append(handCards, target)
	}
	paiCountInfo, paiCount, guiCount := getCardsCountArray(handCards, caiShen)
	// 通用胡牌规则牌数量必须满足count = 3*n + 2
	if (paiCount+guiCount)%3 != 2 {
		return
	}
	// 如果非财神牌数量为0，证明手里只有辣子牌了，可以胡牌
	switch paiCount == 0 {
	case true:
		groups = make(map[int][]cards.MJGroup)
		sb := cards.MJGroup{
			Suit:  caiShen.Suit(),
			Type:  cards.GroupTypeMJJiang,
			Cards: []int{-1, -1},
		}
		groups[0] = append(groups[0], sb)
	default:
		groups = tryHu(paiCountInfo, paiCount, guiCount)
	}

	canHu = len(groups) > 0
	return
}

// 获取一组牌每张牌的数量情况以及鬼牌和非牌数量
// cards: 目标牌组
// caiShen: 本局财神
func getCardsCountArray(handCards []cards.Card, caiShen cards.Card) (paiCountInfo [][]int32, paiCount, guiCount int) {
	paiCountInfo = make([][]int32, 6)
	for _, c := range handCards {
		switch {
		case IsCaiShen(caiShen, c):
			guiCount += 1

		case c.Valid():
			paiCount += 1
			suit, value := c.Suit(), c.Value()
			countInfo := paiCountInfo[suit-1]
			if len(countInfo) == 0 {
				switch suit {
				case cards.CardSuitWind, cards.CardSuitDragon:
					countInfo = make([]int32, 5)
				default:
					countInfo = make([]int32, 10)
				}
				paiCountInfo[suit-1] = countInfo
			}
			countInfo[value] += 1
		}
	}
	return
}

// 尝试胡牌
func tryHu(paiCountInfo [][]int32, paiCount, guiCount int) (groups map[int][]cards.MJGroup) {
	groupId := new(int)
	groups = make(map[int][]cards.MJGroup)
	for suit, countInfo := range paiCountInfo {
		for i, cnt := range countInfo {
			// 找到可能的将牌
			switch {
			case i == 0:
			case cnt >= 2:
				countInfo[i] -= 2
				subGroups := make(map[int][]cards.MJGroup)
				canHu := tryKeAndShun(paiCountInfo, 1, paiCount-2, guiCount, groupId, subGroups)
				if len(subGroups) > 0 || canHu { // 可以胡牌
					jiangGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJJiang,
						Cards: []int{i, i},
					}
					// 把将牌加入每种胡牌组合里
					for gId, g := range subGroups {
						g = append(g, jiangGroup)
						groups[gId] = g
					}
					// 可以胡牌，但是没有胡牌组合，证明只剩一组将对
					if len(groups) == 0 {
						*groupId += 1
						groups[*groupId] = []cards.MJGroup{jiangGroup}
					}
				}
				countInfo[i] += 2
			case cnt >= 1 && guiCount > 0:
				countInfo[i] -= 1
				subGroups := make(map[int][]cards.MJGroup)
				canHu := tryKeAndShun(paiCountInfo, 1, paiCount-1, guiCount-1, groupId, subGroups)
				if len(subGroups) > 0 || canHu {
					jiangGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJJiang,
						Cards: []int{-1, i},
					}
					for gId, g := range subGroups {
						g = append(g, jiangGroup)
						groups[gId] = g
					}
					if len(groups) == 0 {
						*groupId += 1
						groups[*groupId] = []cards.MJGroup{jiangGroup}
					}
				}
				countInfo[i] += 1
			case cnt >= 0 && guiCount > 1:
				subGroups := make(map[int][]cards.MJGroup)
				canHu := tryKeAndShun(paiCountInfo, 1, paiCount, guiCount-2, groupId, subGroups)
				if len(subGroups) > 0 || canHu {
					jiangGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJJiang,
						Cards: []int{-1, -1},
					}
					for gId, g := range subGroups {
						g = append(g, jiangGroup)
						groups[gId] = g
					}
					if len(groups) == 0 {
						*groupId += 1
						groups[*groupId] = []cards.MJGroup{jiangGroup}
					}
				}
			}
		}
	}
	return
}

// 找出所有的顺子和刻子
func tryKeAndShun(paiCountInfo [][]int32, level, paiCount, guiCount int, groupId *int, groups map[int][]cards.MJGroup) (canHu bool) {
	if paiCount+guiCount == 0 {
		canHu = true
		return
	}
	for suit, countInfo := range paiCountInfo {
		for i, cnt := range countInfo {
			if i == 0 || cnt == 0 { // 温岭麻将规则约定手里最多只有2个财神，超过2个则直接胡牌了
				continue
			}
			// 找到可能的刻子
			switch {
			case cnt >= 3: // 现成的刻子
				countInfo[i] -= 3
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-3, guiCount, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJKe,
						Cards: []int{i, i, i},
					}
					groups[*groupId] = append(groups[*groupId], subGroup)
					countInfo[i] += 3 // 这里需要返还递归前扣除的牌
					if level == 1 {   // 如果已经回退到第一层了，需要重置groupId，继续后续的递归
						*groupId += 1
					} else { // 否则直接返回
						return
					}
				} else {
					countInfo[i] += 3
					if level == 1 {
						delete(groups, *groupId)
						*groupId += 1
					}
				}
			case cnt >= 2 && guiCount > 0: // 补一张财神组成刻子
				countInfo[i] -= 2
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-2, guiCount-1, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJKe,
						Cards: []int{-1, i, i},
					}
					groups[*groupId] = append(groups[*groupId], subGroup)
					countInfo[i] += 2
					if level == 1 { // 如果已经回退到第一层了，需要重置groupId，继续后续的递归
						*groupId += 1
					} else { // 否则直接返回
						return
					}
				} else {
					countInfo[i] += 2
					if level == 1 {
						delete(groups, *groupId)
						*groupId += 1
					}
				}
			case cnt >= 1 && guiCount > 1: // 补两张财神组成刻子
				countInfo[i] -= 1
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-1, guiCount-2, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJKe,
						Cards: []int{-1, -1, i},
					}
					groups[*groupId] = append(groups[*groupId], subGroup)
					countInfo[i] += 1
					if level == 1 {
						*groupId += 1
					} else {
						return
					}
				} else {
					countInfo[i] += 1
					if level == 1 {
						delete(groups, *groupId)
						*groupId += 1
					}
				}
			}
		}

		// 找到可能的顺子
		// 风牌和字牌不能组成顺子，忽略掉
		if cards.CardSuit(suit+1) > cards.CardSuitDot || len(countInfo) == 0 {
			continue
		}
		for i := 1; i < 8; i++ {
			// 如果连着的三个牌值的数量都为零的话，直接继续下一轮循环
			if countInfo[i] == 0 && countInfo[i+1] == 0 && countInfo[i+2] == 0 {
				continue
			}
			// 现成的顺子
			if countInfo[i] > 0 && countInfo[i+1] > 0 && countInfo[i+2] > 0 { // 1, 2, 3
				countInfo[i] -= 1
				countInfo[i+1] -= 1
				countInfo[i+2] -= 1
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-3, guiCount, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJShun,
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
			// 补一张赖子成为顺子
			if countInfo[i] > 0 && countInfo[i+1] > 0 && countInfo[i+2] >= 0 && guiCount > 0 { // 1, 2, _
				countInfo[i+1] -= 1
				countInfo[i] -= 1
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-2, guiCount-1, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJShun,
						Cards: []int{i, i + 1, -1},
					}
					groups[*groupId] = append(groups[*groupId], subGroup)
					countInfo[i+1] += 1
					countInfo[i] += 1
					if level == 1 {
						*groupId += *groupId
					} else {
						return
					}
				} else {
					countInfo[i+1] += 1
					countInfo[i] += 1
					if level == 1 {
						delete(groups, *groupId)
						*groupId += 1
					}
				}

			}
			if countInfo[i] > 0 && countInfo[i+1] >= 0 && countInfo[i+2] > 0 && guiCount > 0 { // 1, _, 3
				countInfo[i+2] -= 1
				countInfo[i] -= 1
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-2, guiCount-1, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJShun,
						Cards: []int{i, -1, i + 2},
					}
					groups[*groupId] = append(groups[*groupId], subGroup)
					countInfo[i+2] += 1
					countInfo[i] += 1
					if level == 1 {
						*groupId += 1
					} else {
						return
					}
				} else {
					countInfo[i+2] += 1
					countInfo[i] += 1
					if level == 1 {
						delete(groups, *groupId)
						*groupId += 1
					}
				}

			}
			if countInfo[i] >= 0 && countInfo[i+1] > 0 && countInfo[i+2] > 0 && guiCount > 0 { // _, 2, 3
				countInfo[i+1] -= 1
				countInfo[i+2] -= 1
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-2, guiCount-1, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJShun,
						Cards: []int{-1, i + 1, i + 2},
					}
					groups[*groupId] = append(groups[*groupId], subGroup)
					countInfo[i+1] += 1
					countInfo[i+2] += 1
					if level == 1 {
						*groupId += 1
					} else {
						return
					}
				} else {
					countInfo[i+1] += 1
					countInfo[i+2] += 1
					if level == 1 {
						delete(groups, *groupId)
						*groupId += 1
					}
				}

			}
			// 补两张赖子成为顺子
			if countInfo[i] > 0 && countInfo[i+1] >= 0 && countInfo[i+2] >= 0 && guiCount > 1 { // 1, _, _
				countInfo[i] -= 1
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-1, guiCount-2, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJShun,
						Cards: []int{i, -1, -1},
					}
					groups[*groupId] = append(groups[*groupId], subGroup)
					countInfo[i] += 1
					if level == 1 {
						*groupId += 1
					} else {
						return
					}
				} else {
					countInfo[i] += 1
					if level == 1 {
						delete(groups, *groupId)
						*groupId += 1
					}
				}
			}
			if countInfo[i] >= 0 && countInfo[i+1] > 0 && countInfo[i+2] >= 0 && guiCount > 1 { // _, 2, _
				countInfo[i+1] -= 1
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-1, guiCount-2, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJShun,
						Cards: []int{-1, i + 1, -1},
					}
					groups[*groupId] = append(groups[*groupId], subGroup)
					countInfo[i+1] += 1
					if level == 1 {
						*groupId += 1
					} else {
						return
					}
				} else {
					countInfo[i+1] += 1
					if level == 1 {
						delete(groups, *groupId)
						*groupId += 1
					}
				}
			}
			if countInfo[i] >= 0 && countInfo[i+1] >= 0 && countInfo[i+2] > 0 && guiCount > 1 { // _, _, 3
				countInfo[i+2] -= 1
				canHu = tryKeAndShun(paiCountInfo, level+1, paiCount-1, guiCount-2, groupId, groups)
				if canHu {
					subGroup := cards.MJGroup{
						Suit:  cards.CardSuit(suit + 1),
						Type:  cards.GroupTypeMJShun,
						Cards: []int{-1, -1, i + 2},
					}
					groups[*groupId] = append(groups[*groupId], subGroup)
					countInfo[i+2] += 1
					if level == 1 {
						*groupId += 1
					} else {
						return
					}
				} else {
					countInfo[i+2] += 1
					if level == 1 {
						delete(groups, *groupId)
						*groupId += 1
					}
				}
			}
		}
	}
	return
}
