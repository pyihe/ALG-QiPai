package doudizhu

import (
	"ALG-QiPai/cards"
)

// ParseToGroup 解析指定牌组
func ParseToGroup(targetCards []cards.Card, ghosts ...cards.Card) (group *DDZGroup) {
	paiCountInfo := getCardsCountArray(targetCards, ghosts...)
	// 是否是火箭牌型(赖子不能当大小王)
	if rocketGroup := isRocket(paiCountInfo); rocketGroup.IsZero() == false {
		group = &rocketGroup
		return
	}
	// 是否是炸弹(包括赖子炸弹)
	if bombGroup := isBomb(paiCountInfo); bombGroup.IsZero() == false {
		group = &bombGroup
		return
	}
	// 是否是四带二对子
	if g := isQuadplexSetWithPair(paiCountInfo); g.IsZero() == false {
		group = &g
		return
	}
	// 是否是四带二单牌
	if g := isQuadplexSetWithSingle(paiCountInfo); g.IsZero() == false {
		group = &g
		return
	}
	// 是否是飞机带对子
	if g := isSequenceOfTripletWithPair(paiCountInfo); g.IsZero() == false {
		group = &g
		return
	}
	// 是否是飞机带单牌
	if g := isSequenceOfTripletWithSingle(paiCountInfo); g.IsZero() == false {
		group = &g
		return
	}
	// 是否是飞机不带牌
	if tripletSequence := isSequenceOfTriplet(paiCountInfo); tripletSequence.IsZero() == false {
		group = &tripletSequence
		return
	}
	// 是否是连对
	if seqPairGroup := isSequenceOfPair(paiCountInfo); seqPairGroup.IsZero() == false {
		group = &seqPairGroup
		return
	}
	// 是否是顺子
	if sequenceGroup := isSequence(paiCountInfo); sequenceGroup.IsZero() == false {
		group = &sequenceGroup
		return
	}
	// 是否是三带一对
	if triplePairGroup := isTripleWithPair(paiCountInfo); triplePairGroup.IsZero() == false {
		group = &triplePairGroup
		return
	}
	// 是否是三带一
	if tripleSingleGroup := isTripletWithSingle(paiCountInfo); tripleSingleGroup.IsZero() == false {
		group = &tripleSingleGroup
		return
	}
	// 是否是三不带
	if tripletGroup := isTriplet(paiCountInfo); tripletGroup.IsZero() == false {
		group = &tripletGroup
		return
	}
	//是否是对子
	if pairGroup := isPair(paiCountInfo); pairGroup.IsZero() == false {
		group = &pairGroup
		return
	}
	// 是否是单牌
	if singleGroup := isSingle(paiCountInfo); singleGroup.IsZero() == false {
		group = &singleGroup
		return
	}
	return
}

// isSingle 判断一组牌是否是单牌
func isSingle(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount != 1 {
		return
	}

	switch {
	case ghostCount == 1: // 赖子作为单牌使用时充当它自己
		group = DDZGroup{
			Type:       cards.GroupTypeDDZSingle,
			Length:     1,
			GhostCount: 1,
			Key:        ghosts[0].Value(),
			Cards:      ghosts,
		}
	default:
		for value, cs := range paiCountInfo.Cards {
			group = DDZGroup{
				Type:       cards.GroupTypeDDZSingle,
				Length:     1,
				GhostCount: 0,
				Key:        value,
				Cards:      cs,
			}
		}
	}
	return
}

// isPair 是否是对子
func isPair(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount != 2 {
		return
	}

	// 如果赖子牌数量刚好为2，证明赖子作为自己使用
	if ghostCount == 2 {
		group = DDZGroup{
			Type:       cards.GroupTypeDDZPair,
			Length:     1,
			GhostCount: 2,
			Cards:      ghosts,
		}
		// 因为可能有不同牌值的赖子牌，找最大的那张作为对子的key
		for _, c := range ghosts {
			if c.Value() >= group.Key {
				group.Key = c.Value()
			}
		}
		return
	}

	// 如果赖子数量不为2
loop:
	for i := 3; i <= 15; i++ {
		n := paiCountInfo.CountInfo[i]
		cs := paiCountInfo.GetCards(int32(i))
		switch {
		case n == 1:
			if n == paiCount && ghostCount == 1 {
				group = DDZGroup{
					Type:       cards.GroupTypeDDZPair,
					Length:     2,
					GhostCount: ghostCount,
					Key:        int32(i),
					Cards:      cs,
				}
				group.Cards = append(group.Cards, ghosts...)
				break loop
			}
		case n == 2:
			group = DDZGroup{
				Type:       cards.GroupTypeDDZPair,
				Length:     1,
				GhostCount: 0,
				Key:        int32(i),
				Cards:      cs,
			}
			break loop
		}
	}
	return
}

// isTriplet 是否是三不带
func isTriplet(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount != 3 {
		return
	}

	// 如果赖子牌刚好是三张，则赖子作为自己使用
	if ghostCount == 3 {
		group = DDZGroup{
			Type:       cards.GroupTypeDDZTriplet,
			Length:     1,
			GhostCount: ghostCount,
			Cards:      ghosts,
		}
		// 对于多个赖子值的，找到key最大的作为牌型的key
		for _, c := range ghosts {
			if c.Value() >= group.Key {
				group.Key = c.Value()
			}
		}
		return
	}

loop:
	for i := 3; i <= 15; i++ {
		cs := paiCountInfo.GetCards(int32(i))
		n := paiCountInfo.CountInfo[i]

		switch n {
		case 1:
			if ghostCount == 2 {
				group = DDZGroup{
					Type:       cards.GroupTypeDDZTriplet,
					Length:     1,
					GhostCount: ghostCount,
					Key:        int32(i),
					Cards:      cs,
				}
				group.Cards = append(group.Cards, ghosts...)
				break loop
			}
		case 2:
			if ghostCount == 1 {
				group = DDZGroup{
					Type:       cards.GroupTypeDDZTriplet,
					Length:     1,
					GhostCount: ghostCount,
					Key:        int32(i),
					Cards:      cs,
				}
				group.Cards = append(group.Cards, ghosts...)
				break loop
			}
		case 3:
			if ghostCount == 0 {
				group = DDZGroup{
					Type:       cards.GroupTypeDDZTriplet,
					Length:     1,
					GhostCount: 0,
					Key:        int32(i),
					Cards:      cs,
				}
				break
			}
		}
	}

	return
}

// isTripletWithSingle 是否是三带一张单牌
func isTripletWithSingle(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount != 4 {
		return
	}

	// 三带一非赖子牌的牌值数量一定大于1，即除了赖子牌外另有两种单牌
	// 三带一只需要找到三张相同的牌，剩下的一张不影响
	// 先找到三张相同的牌
	countInfo := paiCountInfo.CountInfo
loop:
	for i := 15; i >= 3; i-- {
		cs := paiCountInfo.GetCards(int32(i))
		n := countInfo[i]
		switch n {
		case 1: // 赖子数量只能为2，如果是一张
			if ghostCount == 2 {
				group.Type = cards.GroupTypeDDZTripletWithSingle
				group.Key = int32(i)
				group.Length = 1
				group.GhostCount = ghostCount
				group.Cards = append(group.Cards, cs...)
				countInfo[i] -= 1
				ghostCount -= 2
				paiCount -= 1
				break loop
			}
		case 2:
			if ghostCount == 1 {
				group.Type = cards.GroupTypeDDZTripletWithSingle
				group.Key = int32(i)
				group.Length = 1
				group.GhostCount = ghostCount
				group.Cards = append(group.Cards, cs...)
				countInfo[i] -= 2
				ghostCount -= 1
				paiCount -= 2
				break loop
			}
		case 3:
			if ghostCount == 0 {
				group.Type = cards.GroupTypeDDZTripletWithSingle
				group.Key = int32(i)
				group.Length = 1
				group.GhostCount = ghostCount
				group.Cards = append(group.Cards, cs...)
				countInfo[i] -= 3
				paiCount -= 3
				break loop
			}
		}
	}
	if group.Type == cards.GroupTypeDDZTripletWithSingle && ghostCount == 0 && paiCount == 1 {
		flag := false
		for i := 3; i <= 17; i++ {
			if n := countInfo[i]; n == 1 {
				cs := paiCountInfo.GetCards(int32(i))
				group.Cards = append(group.Cards, cs...)
				flag = true
				break
			}
		}
		if flag {
			group.Cards = append(group.Cards, ghosts...)
		} else {
			group.Reset()
		}
	}
	return
}

// isTripleWithPair 是否是三带一对
func isTripleWithPair(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount != 5 {
		return
	}
	// 先找到三张相同的牌
	countInfo := paiCountInfo.CountInfo
loop:
	for i := 15; i >= 3; i-- {
		cs := paiCountInfo.GetCards(int32(i))
		n := countInfo[i]
		switch n {
		case 1: // 赖子数量只能为2，如果是一张
			if ghostCount == 2 {
				group.Type = cards.GroupTypeDDZTripletWithPair
				group.Key = int32(i)
				group.Length = 1
				group.GhostCount = ghostCount
				group.Cards = append(group.Cards, cs...)
				countInfo[i] -= 1
				ghostCount -= 2
				paiCount -= 1
				break loop
			}
		case 2:
			if ghostCount == 1 {
				group.Type = cards.GroupTypeDDZTripletWithPair
				group.Key = int32(i)
				group.Length = 1
				group.GhostCount = ghostCount
				group.Cards = append(group.Cards, cs...)
				countInfo[i] -= 2
				ghostCount -= 1
				paiCount -= 2
				break loop
			}
		case 3:
			if ghostCount == 0 {
				group.Type = cards.GroupTypeDDZTripletWithPair
				group.Key = int32(i)
				group.Length = 1
				group.GhostCount = ghostCount
				group.Cards = append(group.Cards, cs...)
				countInfo[i] -= 3
				paiCount -= 3
				break loop
			}
		}
	}

	if group.Type == cards.GroupTypeDDZTripletWithPair && paiCount == 2 {
		flag := false
		for i := 3; i <= 17; i++ {
			n := countInfo[i]
			cs := paiCountInfo.GetCards(int32(i))
			if (n == 1 && ghostCount == 1) || (n == 2 && ghostCount == 0) {
				group.Cards = append(group.Cards, cs...)
				flag = true
				break
			}
		}
		if flag {
			group.Cards = append(group.Cards, ghosts...)
		} else {
			group.Reset()
		}
	}
	return
}

// isSequence 是否是顺子
func isSequence(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 5 {
		return
	}
	// 有大小王或者2
	if hasJoker(paiCountInfo) || len(paiCountInfo.GetCards(2)) > 0 {
		return
	}
	countInfo := paiCountInfo.CountInfo
	start, end := 0, 0 // 第一张有数量的牌的位置和最后一张有数量的牌的位置
	for i := 3; i <= 14; i++ {
		n := countInfo[i]
		if n > 1 {
			return
		}
		if n == 1 {
			if start == 0 {
				start = i
				continue
			}
			if start != 0 && i > end {
				end = i
				continue
			}
		}
	}
	// 如果是顺子牌型，必须有start和end，
	// start==0即paiCount==0
	// end==0可以解析成炸弹
	if start == 0 || end == 0 {
		return
	}
	// 如果有赖子，先补到start和end的中间，其次是end之后，最后再补到start之前
	if ghostCount > 0 {
		// 将赖子补充到start和end之间
		for i := start; i <= end; i++ {
			if countInfo[i] == 0 {
				if ghostCount == 0 {
					return
				} else {
					countInfo[i] = -1
					ghostCount -= 1
				}
			}
		}
		// 将赖子补充到end之后
		for i := end + 1; i <= 14; i++ {
			if countInfo[i] == 0 {
				if ghostCount == 0 {
					break
				} else {
					countInfo[i] = -1
					ghostCount -= 1
					end = i
				}
			}
		}

		// 将赖子补充到start之前
		for i := start - 1; i >= 3; i-- {
			if countInfo[i] == 0 {
				if ghostCount == 0 {
					break
				} else {
					countInfo[i] = -1
					ghostCount -= 1
					start = i
				}
			}
		}
	}

	// start,end 刚好是顺子的头和尾
	if end-start+1 == paiCount+paiCountInfo.GhostCount() && ghostCount == 0 {
		group = DDZGroup{
			Type:       cards.GroupTypeDDZSequence,
			Length:     paiCount + ghostCount,
			GhostCount: ghostCount,
			Key:        int32(start),
		}
		for i := start; i <= end; i++ {
			if countInfo[i] == 1 {
				cs := paiCountInfo.GetCards(int32(i))
				group.Cards = append(group.Cards, cs...)
			}
		}
		group.Cards = append(group.Cards, ghosts...)
	}
	return
}

// isSequenceOfPair 是否是连对
func isSequenceOfPair(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 6 || (paiCount+ghostCount)%2 != 0 {
		return
	}
	// 有大小王或者2不能形成连对
	if hasJoker(paiCountInfo) || len(paiCountInfo.GetCards(2)) > 0 {
		return
	}
	start, end := 0, 0
	countInfo := paiCountInfo.CountInfo
	for i := 3; i <= 14; i++ {
		n := countInfo[i]
		if n > 2 {
			return
		}
		if n != 0 {
			if start == 0 {
				start = i
				continue
			}
			if start != 0 && i > end {
				end = i
				continue
			}
		}
	}
	if start == 0 || end == 0 {
		return
	}

	if ghostCount > 0 {
		for i := start; i <= end; i++ {
			if countInfo[i] == 0 {
				if ghostCount < 2 {
					return
				} else {
					ghostCount -= 2
					countInfo[i] = -1
					continue
				}
			}
			if countInfo[i] == 1 {
				if ghostCount < 1 {
					return
				} else {
					ghostCount -= 1
					continue
				}
			}
		}
		for i := end + 1; i <= 14; i++ {
			if ghostCount%2 != 0 {
				return
			}
			if ghostCount == 0 {
				break
			} else {
				ghostCount -= 2
				countInfo[i] = -1
				end = i
			}
		}
		for i := start - 1; i >= 3; i-- {
			if ghostCount%2 != 0 {
				return
			}
			if ghostCount == 0 {
				break
			} else {
				ghostCount -= 2
				countInfo[i] = -1
				start = i
			}
		}
	}
	if ghostCount == 0 && (end-start+1)*2 == paiCount+paiCountInfo.GhostCount() {
		group = DDZGroup{
			Type:       cards.GroupTypeDDZSequenceOfPair,
			Length:     end - start + 1,
			GhostCount: ghostCount,
			Key:        int32(start),
		}
		for i := start; i <= end; i++ {
			if countInfo[i] > 0 {
				cs := paiCountInfo.GetCards(int32(i))
				group.Cards = append(group.Cards, cs...)
			}
		}
		group.Cards = append(group.Cards, ghosts...)
	}
	return
}

// isSequenceOfTriplet 是否是飞机不带牌
func isSequenceOfTriplet(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 6 || (paiCount+ghostCount)%3 != 0 {
		return
	}
	// 大小王和2不能组成飞机
	if hasJoker(paiCountInfo) || len(paiCountInfo.GetCards(2)) > 0 {
		return
	}

	start, end := 0, 0
	countInfo := paiCountInfo.CountInfo
	for i := 3; i <= 14; i++ {
		n := countInfo[i]
		if n > 3 {
			return
		}
		if n != 0 {
			if start == 0 {
				start = i
				continue
			}
			if start != 0 && i > end {
				end = i
				continue
			}
		}
	}
	if start == 0 || end == 0 {
		return
	}

	if ghostCount > 0 {
		for i := start; i <= end; i++ {
			if countInfo[i] == 0 {
				if ghostCount < 3 {
					return
				} else {
					ghostCount -= 3
					countInfo[i] = -1
					continue
				}
			}
			if countInfo[i] == 1 {
				if ghostCount < 2 {
					return
				} else {
					ghostCount -= 2
					continue
				}
			}
			if countInfo[i] == 2 {
				if ghostCount < 1 {
					return
				} else {
					ghostCount -= 1
					continue
				}
			}
		}
		for i := end + 1; i <= 14; i++ {
			if ghostCount%3 != 0 {
				return
			}
			if ghostCount == 0 {
				break
			} else {
				ghostCount -= 3
				countInfo[i] = -1
				end = i
			}
		}
		for i := start - 1; i >= 3; i-- {
			if ghostCount%3 != 0 {
				return
			}
			if ghostCount == 0 {
				break
			} else {
				ghostCount -= 3
				countInfo[i] = -1
				start = i
			}
		}
	}
	if ghostCount == 0 && (end-start+1)*3 == paiCount+paiCountInfo.GhostCount() {
		group = DDZGroup{
			Type:       cards.GroupTypeDDZSequenceOfTriplet,
			Length:     end - start + 1,
			GhostCount: ghostCount,
			Key:        int32(start),
		}
		for i := start; i <= end; i++ {
			if countInfo[i] > 0 {
				cs := paiCountInfo.GetCards(int32(i))
				group.Cards = append(group.Cards, cs...)
			}
		}
		group.Cards = append(group.Cards, ghosts...)
	}
	return
}

// isSequenceOfTripletWithSingle 是否是飞机带单牌
func isSequenceOfTripletWithSingle(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if cnt := paiCount + ghostCount; cnt < 8 || cnt%4 != 0 {
		return
	}
	// 飞机牌型只需要找到可以形成的三连顺即可, 翅膀是什么牌不用管, 翅膀只需要满足数量要求即可
	// 思路: 飞机带单牌牌型最短为2连，最长为5连(农名最长4连，地主理论上存在5连的可能性),
	// 遍历可能的飞机长度，判断长度和牌总数的数量关系是否满足
	// 如果满足，从A到3开始倒序遍历，看能否找到飞机(三连), 过程中数量不足的用赖子填充, 直到满足飞机带单牌, 否则不满足
outLoop:
	for i := 2; i <= 5; i++ { // i 为飞机长度
		// 牌总数和飞机长度为4倍关系
		if i*4 != paiCount+ghostCount {
			continue outLoop
		}
		// 从大到小的遍历是为了找到key最大的飞机
	innerLoop:
		for j := 14; j > i+1; j-- { // j为飞机最后的位置
			pn, gn := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
			for k := 0; k < i; k++ { // k为到飞机尾的距离, 判断三连顺中的每个位置是否满足数量要求
				n := paiCountInfo.GetCardCount(int32(j - k))
				if n+gn < 3 { // 牌数量和赖子数量必须>=3, 否则不能形成三连顺
					continue innerLoop
				}
				if n < 3 {
					gn -= 3 - n
					pn -= n
				} else {
					pn -= 3
				}
			}
			// 判断剩下的牌数量是否与飞机长度相同, 如果相同则是满足飞机带单牌的牌型
			if pn+gn == i {
				group = DDZGroup{
					Type:       cards.GroupTypeDDZSequenceOfTripletWithSingle,
					Length:     i,
					GhostCount: ghostCount,
					Key:        int32(j - i + 1),
				}
				for k := 3; k <= 17; k++ {
					if cs := paiCountInfo.GetCards(int32(k)); len(cs) > 0 {
						group.Cards = append(group.Cards, cs...)
					}
				}
				group.Cards = append(group.Cards, ghosts...)
				break outLoop // 找到了直接跳出并返回
			}
		}
	}
	return
}

// isSequenceOfTripletWithPair 是否是飞机带对子
func isSequenceOfTripletWithPair(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if cnt := paiCount + ghostCount; cnt < 10 || cnt%5 != 0 {
		return
	}
	// 飞机带对子不能有大小王
	if hasJoker(paiCountInfo) {
		return
	}

	// 思路: 筒飞机带单牌的判断，不同的是找到三连顺后还需要判断是否有飞机长度个对子
outLoop:
	for i := 2; i <= 4; i++ {
		if i*5 != paiCount+ghostCount {
			continue
		}
	innerLoop:
		for j := 14; j > i+1; j-- {
			countInfo := paiCountInfo.CountInfo
			pn, gn := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
			for k := 0; k < i; k++ {
				n := paiCountInfo.GetCardCount(int32(j - k))
				if n+gn < 3 {
					continue innerLoop
				}
				if n < 3 {
					gn -= 3 - n
					pn -= n
					countInfo[j-k] -= n
				} else {
					pn -= 3
					countInfo[j-k] -= 3
				}
			}
			// 剩余的牌和赖子牌数量必须满足i个对子的数量
			if pn+gn != 2*i {
				continue innerLoop
			}
			// 判断countInfo剩余的牌加上剩余的赖子牌能否构成i个对子
			pairCount := 0
			for k := 3; k <= 15; k++ {
				switch countInfo[k] {
				case 1:
					if gn >= 1 {
						gn -= 1
						pn -= 1
						pairCount += 1
					}
				case 2: // 两张牌刚好两对
					pn -= 2
					pairCount += 1
				case 3: // 三张牌加1张赖子组成两对
					if gn >= 1 {
						pn -= 3
						gn -= 1
						pairCount += 1
					}
				case 4: // 四张牌刚好组成两对
					pn -= 4
					pairCount = 2
				}
			}
			// 如果找完对子后，牌数量不为0或者赖子数量不为偶数，则不能形成飞机带对子的牌型
			if pn != 0 || gn%2 != 0 {
				continue innerLoop
			}
			if gn/2+pairCount == i {
				group = DDZGroup{
					Type:       cards.GroupTypeDDZSequenceOfTripletWithPair,
					Length:     i,
					GhostCount: ghostCount,
					Key:        int32(j - i + 1),
				}
				for k := 3; k <= 15; k++ {
					if cs := paiCountInfo.GetCards(int32(k)); len(cs) > 0 {
						group.Cards = append(group.Cards, cs...)
					}
				}
				group.Cards = append(group.Cards, ghosts...)
				break outLoop // 找到了直接跳出并返回
			}
		}
	}
	return
}

// isQuadplexSetWithSingle 是否是四带二(两张单牌)
func isQuadplexSetWithSingle(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount != 6 {
		return
	}
	// 先找到四张相同的牌
	for i := 15; i >= 3; i-- {
		pn, gn := paiCount, ghostCount
		n := paiCountInfo.GetCardCount(int32(i))
		if n+gn < 4 {
			continue
		}
		gn -= 4 - n
		pn -= n
		if pn+gn != 2 {
			continue
		}
		// TODO 如果n==0，那么牌型的key值取多少？取最大值还是赖子本身的值，如果赖子牌超过1种，那么又取哪个赖子牌本身的值？
		// 判断剩余的牌是否为2
		group = DDZGroup{
			Type:       cards.GroupTypeDDZQuadplexSetWithSingle,
			Length:     1,
			GhostCount: ghostCount,
			Key:        int32(i),
			Cards:      nil,
		}
		for j := 3; j <= 17; j++ {
			if cs := paiCountInfo.GetCards(int32(j)); len(cs) > 0 {
				group.Cards = append(group.Cards, cs...)
			}
		}
		group.Cards = append(group.Cards, ghosts...)
		break
	}
	return
}

// isQuadplexSetWithPair 是否是四带二(两个对子)
func isQuadplexSetWithPair(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount != 8 {
		return
	}
	// 四带二(两对)不能有大小王
	if hasJoker(paiCountInfo) {
		return
	}
	for i := 15; i >= 3; i-- {
		countInfo := paiCountInfo.CountInfo
		pn, gn := paiCount, ghostCount
		n := paiCountInfo.GetCardCount(int32(i))
		if n+gn < 4 {
			continue
		}
		gn -= 4 - n
		pn -= n
		countInfo[i] -= n
		// 判断剩余的牌能否构成两个对子
		if pn+gn != 4 {
			continue
		}
		// TODO 如果n==0，那么牌型的key值取多少？取最大值还是赖子本身的值，如果赖子牌超过1种，那么又取哪个赖子牌本身的值？
		// 找对子
		pairCount := 0
		for k := 3; k <= 15; k++ {
			switch countInfo[k] {
			case 1:
				if gn >= 1 {
					gn -= 1
					pn -= 1
					pairCount += 1
				}
			case 2: // 两张牌刚好两对
				pn -= 2
				pairCount += 1
			case 3: // 三张牌加1张赖子组成两对
				if gn >= 1 {
					pn -= 3
					gn -= 1
					pairCount += 1
				}
			case 4: // 四张牌刚好组成两对
				pn -= 4
				pairCount = 2
			}
		}
		if pn != 0 || gn%2 != 0 {
			continue
		}
		if gn/2+pairCount == 2 {
			group = DDZGroup{
				Type:       cards.GroupTypeDDZQuadplexSetWithPair,
				Length:     1,
				GhostCount: ghostCount,
				Key:        int32(i),
			}
			for j := 3; j <= 17; j++ {
				if cs := paiCountInfo.GetCards(int32(j)); len(cs) > 0 {
					group.Cards = append(group.Cards, cs...)
				}
			}
			group.Cards = append(group.Cards, ghosts...)
			break
		}
	}
	return
}

// isBomb 是否是普通炸弹
func isBomb(paiCountInfo *PaiCountData) (group DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	// 炸弹最短长度为4
	if paiCount+ghostCount < 4 {
		return
	}
	// 有大小王不能组成普通炸弹
	if hasJoker(paiCountInfo) {
		return
	}
	// 纯赖子炸弹
	if paiCount == 0 {
		group = DDZGroup{
			Type:       cards.GroupTypeDDZBomb,
			Length:     ghostCount,
			GhostCount: ghostCount,
			Key:        -1,
			Cards:      ghosts,
		}
		return
	}
	// 遍历每个牌值，找出炸弹(如果有的话)
	for i := 15; i >= 3; i-- {
		cs := paiCountInfo.GetCards(int32(i))
		n := paiCountInfo.CountInfo[i]
		if n == 0 {
			continue
		}
		// 非纯赖子炸弹只能有一种牌值
		if n != paiCount {
			break
		}
		group = DDZGroup{
			Type:       cards.GroupTypeDDZBomb,
			Length:     paiCount + ghostCount,
			GhostCount: ghostCount,
			Key:        int32(i),
			Cards:      cs,
		}
		group.Cards = append(group.Cards, ghosts...)
		break
	}
	return
}

// isRocket 是否是火箭
func isRocket(paiCountInfo *PaiCountData) (group DDZGroup) {
	paiCount := paiCountInfo.PaiCount()
	if paiCount != 2 {
		return
	}
	if paiCountInfo.GetCardCount(16) == 1 && paiCountInfo.GetCardCount(17) == 1 {
		group = DDZGroup{
			Type:       cards.GroupTypeDDZRocket,
			Key:        16,
			GhostCount: 0,
			Length:     1,
		}
		group.Cards = append(group.Cards, cards.CardBlackJoker, cards.CardRedJoker)
	}
	return
}
