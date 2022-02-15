package doudizhu

import (
	"ALG-QiPai/cards"
	"ALG-QiPai/pkg"
)

// BiggerGroup 找到比指定牌型更大的牌型
func BiggerGroup(targetCards []cards.Card, ghost cards.Card, srcGroup *DDZGroup) (groups []DDZGroup) {
	if srcGroup == nil || srcGroup.IsZero() == true {
		return
	}
	paiCountInfo := getCardsCountArray(targetCards, ghost)
	switch srcGroup.Type {
	case cards.GroupTypeDDZSingle:
		singleGroups := biggerSingle(paiCountInfo, *srcGroup)
		groups = append(groups, singleGroups...)
	case cards.GroupTypeDDZPair:
		pairGroups := biggerPair(paiCountInfo, *srcGroup)
		groups = append(groups, pairGroups...)
	case cards.GroupTypeDDZTriplet:
		tripleGroups := biggerTriplet(paiCountInfo, *srcGroup)
		groups = append(groups, tripleGroups...)
	case cards.GroupTypeDDZTripletWithSingle:
		gs := biggerTripletWithSingle(paiCountInfo, *srcGroup)
		groups = append(groups, gs...)
	case cards.GroupTypeDDZTripletWithPair:
		gs := biggerTripleWithPair(paiCountInfo, *srcGroup)
		groups = append(groups, gs...)
	case cards.GroupTypeDDZSequence:
		gs := biggerSequence(paiCountInfo, *srcGroup)
		groups = append(groups, gs...)
	case cards.GroupTypeDDZSequenceOfPair:
		gs := biggerSequenceOfPair(paiCountInfo, *srcGroup)
		groups = append(groups, gs...)
	case cards.GroupTypeDDZSequenceOfTriplet:
		gs := biggerSequenceOfTriplet(paiCountInfo, *srcGroup)
		groups = append(groups, gs...)
	case cards.GroupTypeDDZSequenceOfTripletWithSingle:
		gs := biggerSequenceOfTripletWithSingle(paiCountInfo, *srcGroup)
		groups = append(groups, gs...)
	case cards.GroupTypeDDZSequenceOfTripletWithPair:
		gs := biggerSequenceOfTripletWithPair(paiCountInfo, *srcGroup)
		groups = append(groups, gs...)
	case cards.GroupTypeDDZQuadplexSetWithSingle:
		gs := biggerQuadplexSetWithSingle(paiCountInfo, *srcGroup)
		groups = append(groups, gs...)
	case cards.GroupTypeDDZQuadplexSetWithPair:
		gs := biggerQuadplexSetWithPair(paiCountInfo, *srcGroup)
		groups = append(groups, gs...)
	case cards.GroupTypeDDZBomb:
		bombs := biggerBomb(paiCountInfo, *srcGroup)
		groups = append(groups, bombs...)
	case cards.GroupTypeDDZRocket:
		// 火箭为最大牌型，直接返回
		return
	}
	// 判断是否要得起时, 对于非炸弹类牌型还需要找到所有炸弹
	if srcGroup.Type < cards.GroupTypeDDZBomb {
		if bombGroups := hasBomb(paiCountInfo); len(bombGroups) > 0 {
			groups = append(groups, bombGroups...)
		}
	}
	if rocketGroup := hasRocket(paiCountInfo); rocketGroup.IsZero() == false {
		groups = append(groups, rocketGroup)
	}
	groups = deduplicationGroups(groups)
	return
}

// biggerSingle 是否有比指定牌大的单牌
func biggerSingle(paiCountInfo *PaiCountData, single DDZGroup) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount == 0 {
		return
	}
	// 判断赖子作为单牌是否比single大
	if ghostCount > 0 {
		for _, c := range ghosts {
			if compareCards(c, single.Cards[0]) == 1 {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZSingle,
					Length:     1,
					GhostCount: 1,
					Key:        c.Value(),
					Cards:      []cards.Card{c},
				}
				groups = append(groups, g)
			}
		}
	}

	if paiCount > 0 {
		for i := single.Key + 1; i <= 17; i++ {
			cs := paiCountInfo.GetCards(i)
			n := paiCountInfo.CountInfo[i]
			if n == 0 {
				continue
			}
			g := DDZGroup{
				Type:       cards.GroupTypeDDZSingle,
				Length:     1,
				GhostCount: 0,
				Key:        i,
				Cards:      []cards.Card{cs[0]},
			}
			groups = append(groups, g)
		}
	}
	return
}

// biggerPair 是否有比指定key大的对子
func biggerPair(paiCountInfo *PaiCountData, pair DDZGroup) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 2 {
		return
	}
	// 纯赖子组成的对子
	if ghostCount >= 2 {
		for _, c := range ghosts {
			if compareCards(c, pair.Cards[0]) == 1 {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZPair,
					Length:     1,
					GhostCount: 2,
					Key:        c.Value(),
					Cards:      []cards.Card{c},
				}
				for _, cc := range ghosts {
					if c != cc {
						g.Cards = append(g.Cards, cc)
						break
					}
				}
				groups = append(groups, g)
			}
		}
	}

	for i := pair.Key + 1; i <= 15; i++ {
		n := paiCountInfo.GetCardCount(i)
		if n == 0 || n+ghostCount < 2 {
			continue
		}
		g := DDZGroup{
			Type:   cards.GroupTypeDDZPair,
			Length: 1,
			Key:    i,
		}
		if n <= 2 {
			g.GhostCount = 2 - n
			g.Cards = paiCountInfo.GetCards(i)
			g.Cards = append(g.Cards, ghosts[:2-n]...)
		} else {
			g.Cards = append(g.Cards, paiCountInfo.GetCards(i)[:2]...)
		}
		groups = append(groups, g)
	}
	return
}

// biggerTriplet 比指定key大的三不带
func biggerTriplet(paiCountInfo *PaiCountData, triplet DDZGroup) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 3 {
		return
	}

	// 纯赖子组成的三不带
	if ghostCount >= 3 {
		for _, c := range ghosts {
			if compareCards(c, triplet.Cards[0]) == 1 {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZTriplet,
					Length:     1,
					GhostCount: 3,
					Key:        c.Value(),
					Cards:      []cards.Card{c},
				}
				for _, cc := range ghosts {
					if len(g.Cards) >= 3 {
						break
					}
					if cc != c {
						g.Cards = append(g.Cards, cc)
					}
				}
				groups = append(groups, g)
			}
		}
	}

	for i := triplet.Key + 1; i <= 15; i++ {
		n := paiCountInfo.GetCardCount(i)
		if n == 0 || n+ghostCount < 3 {
			continue
		}
		g := DDZGroup{
			Type:       cards.GroupTypeDDZTriplet,
			Length:     1,
			GhostCount: 0,
			Key:        i,
			Cards:      nil,
		}
		if n <= 3 {
			g.GhostCount = 3 - n
			g.Cards = paiCountInfo.GetCards(i)
			g.Cards = append(g.Cards, ghosts[:3-n]...)
		} else {
			g.Cards = append(g.Cards, paiCountInfo.GetCards(i)[:3]...)
		}
		groups = append(groups, g)
	}
	return
}

// biggerTripletWithSingle 是否有比指定key大的三带一
func biggerTripletWithSingle(paiCountInfo *PaiCountData, triplet DDZGroup) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 4 {
		return
	}

	// 不考虑三张相同的牌都有赖子构成，如果这样，那直接解析成炸弹
	for i := triplet.Key + 1; i <= 15; i++ {
		countInfo := paiCountInfo.CountInfo
		n := paiCountInfo.GetCardCount(i)
		// 不考虑数量为0的情况
		if n == 0 || n+ghostCount < 3 {
			continue
		}
		// n张全部用来做三张
		countInfo[i] -= n
		g := DDZGroup{
			Type:   cards.GroupTypeDDZTripletWithSingle,
			Length: 1,
			Key:    i,
		}
		if n < 3 {
			g.GhostCount = 3 - n
			g.Cards = paiCountInfo.GetCards(i)
			g.Cards = append(g.Cards, ghosts[:3-n]...)
		} else {
			g.Cards = append(g.Cards, paiCountInfo.GetCards(i)[:3]...)
		}
		// 找到了三张过后，找被带的那张单牌
		for j, cnt := range countInfo {
			if j < 3 || cnt == 0 || j == int(i) {
				continue
			}
			g.Cards = append(g.Cards, paiCountInfo.GetCards(int32(j))[0])
			groups = append(groups, g)
		}
	}
	return
}

// biggerTripleWithPair 比指定key大的三带对
func biggerTripleWithPair(paiCountInfo *PaiCountData, triplet DDZGroup) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 5 {
		return
	}
	for i := triplet.Key; i <= 15; i++ {
		countInfo := paiCountInfo.CountInfo
		n := paiCountInfo.GetCardCount(i)
		gn := ghostCount
		if n == 0 || n+gn < 3 {
			continue
		}
		countInfo[i] -= n
		g := DDZGroup{
			Type:   cards.GroupTypeDDZTripletWithPair,
			Length: 1,
			Key:    i,
		}
		if n < 3 {
			g.GhostCount = 3 - n
			gn -= g.GhostCount
			g.Cards = append(g.Cards, ghosts[:3-n]...)
		} else {
			g.Cards = append(g.Cards, paiCountInfo.GetCards(i)[:3]...)
		}
		// 找对子
		for j := 3; j <= 15; j++ {
			m := countInfo[j]
			if m == 0 || j == int(i) {
				continue
			}
			if m+gn >= 2 {
				if m == 1 {
					g.GhostCount += 1
					g.Cards = append(g.Cards, paiCountInfo.GetCards(int32(j))...)
					g.Cards = append(g.Cards, ghosts[3-n:3-n+1]...)
				} else {
					g.Cards = append(g.Cards, paiCountInfo.GetCards(int32(j))[:2]...)
				}
				groups = append(groups, g)
			}
		}
	}
	return
}

// biggerSequence 是否有比指定key和长度更大的顺子
func biggerSequence(paiCountInfo *PaiCountData, sequence DDZGroup) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < sequence.Length {
		return
	}
	// 如果顺子已经封顶了，直接返回
	if sequence.Key+int32(sequence.Length)-1 == 14 {
		return
	}
	// 如果key+1到A长度不足，直接返回
	if 14-sequence.Key+1 < int32(sequence.Length) {
		return
	}

outLoop:
	for i := int(sequence.Key) + 1; i <= 14-sequence.Length+1; i++ {
		var cs []cards.Card
		gn := ghostCount
		countInfo := paiCountInfo.CountInfo
		for j := 0; j < sequence.Length; j++ {
			n := countInfo[i+j]
			if n+gn == 0 {
				continue outLoop
			}
			if n == 0 {
				gn -= 1
			} else {
				cs = append(cs, paiCountInfo.GetCards(int32(i + j))[0])
			}
		}
		g := DDZGroup{
			Type:       cards.GroupTypeDDZSequence,
			Length:     sequence.Length,
			GhostCount: ghostCount - gn,
			Key:        int32(i),
			Cards:      cs,
		}
		g.Cards = append(g.Cards, ghosts[:g.GhostCount]...)
		groups = append(groups, g)
	}
	return
}

// biggerSequenceOfPair 是否有比指定key和长度更大的连对
func biggerSequenceOfPair(paiCountInfo *PaiCountData, sequence DDZGroup) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < sequence.Length*2 {
		return
	}
	// 如果顺子已经封顶了，直接返回
	if sequence.Key+int32(sequence.Length)-1 == 14 {
		return
	}
	// 如果key+1到A长度不足，直接返回
	if 14-sequence.Key+1 < int32(sequence.Length) {
		return
	}

outLoop:
	for i := int(sequence.Key) + 1; i <= 14-sequence.Length+1; i++ {
		var cs []cards.Card
		gn := ghostCount
		countInfo := paiCountInfo.CountInfo
		for j := 0; j < sequence.Length; j++ {
			n := countInfo[i+j]
			if n+gn < 2 {
				continue outLoop
			}
			pais := paiCountInfo.GetCards(int32(i + j))
			if n <= 2 {
				gn -= 2 - n
				cs = append(cs, pais[:n]...)
			} else {
				cs = append(cs, pais[:2]...)
			}
		}
		g := DDZGroup{
			Type:       cards.GroupTypeDDZSequenceOfPair,
			Length:     sequence.Length,
			GhostCount: ghostCount - gn,
			Key:        int32(i),
			Cards:      cs,
		}
		g.Cards = append(g.Cards, ghosts[:g.GhostCount]...)
		groups = append(groups, g)
	}
	return
}

// biggerSequenceOfTriplet 是否有比指定key和长度更大的三连顺
func biggerSequenceOfTriplet(paiCountInfo *PaiCountData, sequence DDZGroup) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < sequence.Length*3 {
		return
	}
	// 如果顺子已经封顶了，直接返回
	if sequence.Key+int32(sequence.Length)-1 == 14 {
		return
	}
	// 如果key+1到A长度不足，直接返回
	if 14-sequence.Key+1 < int32(sequence.Length) {
		return
	}
outLoop:
	for i := int(sequence.Key) + 1; i <= 14-sequence.Length+1; i++ {
		var cs []cards.Card
		gn := ghostCount
		countInfo := paiCountInfo.CountInfo
		for j := 0; j < sequence.Length; j++ {
			n := countInfo[i+j]
			if n+gn < 3 {
				continue outLoop
			}
			pais := paiCountInfo.GetCards(int32(i + j))
			if n == 4 {
				cs = append(cs, pais[:3]...)
			} else {
				gn -= 3 - n
				cs = append(cs, pais[:n]...)
			}
		}
		g := DDZGroup{
			Type:       cards.GroupTypeDDZSequenceOfPair,
			Length:     sequence.Length,
			GhostCount: ghostCount - gn,
			Key:        int32(i),
			Cards:      cs,
		}
		g.Cards = append(g.Cards, ghosts[:g.GhostCount]...)
		groups = append(groups, g)
	}
	return
}

// biggerSequenceOfTripletWithSingle 是否有比指定key和长度更大的飞机带单牌
func biggerSequenceOfTripletWithSingle(paiCountInfo *PaiCountData, sequence DDZGroup) (groups []DDZGroup) {
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < sequence.Length*4 {
		return
	}
	// 如果顺子已经封顶了，直接返回
	if sequence.Key+int32(sequence.Length)-1 == 14 {
		return
	}
	// 如果key+1到A长度不足，直接返回
	if 14-sequence.Key+1 < int32(sequence.Length) {
		return
	}

outLoop:
	for i := int(sequence.Key) + 1; i <= 14-sequence.Length+1; i++ {
		var cs []cards.Card
		var newPaiCountInfo = paiCountInfo.Clone()
		var ghosts = newPaiCountInfo.Ghosts()
		var gn = newPaiCountInfo.GhostCount()
		for j := 0; j < sequence.Length; j++ {
			key := int32(i + j)
			n := newPaiCountInfo.GetCardCount(key)
			if n+gn < 3 {
				continue outLoop
			}

			pais := newPaiCountInfo.GetCards(key)
			if n == 4 {
				cs = append(cs, pais[:3]...)
				newPaiCountInfo.CountInfo[key] -= 3
				newPaiCountInfo.CountInfo[0] -= 3
				newPaiCountInfo.Cards[key] = newPaiCountInfo.Cards[key][3:]
			} else {
				cs = append(cs, pais[:n]...)
				newPaiCountInfo.CountInfo[key] -= n
				newPaiCountInfo.CountInfo[0] -= n
				newPaiCountInfo.Cards[key] = newPaiCountInfo.Cards[key][n:]

				gn -= 3 - n
				cs = append(cs, ghosts[:3-n]...)
				newPaiCountInfo.CountInfo[1] -= 3 - n
				newPaiCountInfo.Cards[-1] = newPaiCountInfo.Cards[-1][3-n:]
			}
		}
		// 找出所有的单牌组合
		singleGroups := deduplicationGroups(hasSingle(newPaiCountInfo))
		if len(singleGroups) < sequence.Length {
			continue outLoop
		}
		// 从单牌的所有组合里找出length个组合
		combs := pkg.Combination(len(singleGroups), sequence.Length)
		singles := make([][]DDZGroup, len(combs))
		for k, ss := range combs {
			for m, v := range ss {
				if v == 1 {
					singles[k] = append(singles[k], singleGroups[m])
				}
			}
		}
		for _, gs := range singles {
			g := DDZGroup{
				Type:       cards.GroupTypeDDZSequenceOfTripletWithSingle,
				Length:     sequence.Length,
				GhostCount: ghostCount - gn,
				Key:        int32(i),
			}
			g.Cards = append(g.Cards, cs...)
			for _, gg := range gs {
				g.Cards = append(g.Cards, gg.Cards...)
				g.GhostCount += gg.GhostCount
			}
			groups = append(groups, g)
		}
	}
	return
}

// biggerSequenceOfTripletWithPair 是否有比指定key和长度更大的飞机带对子
func biggerSequenceOfTripletWithPair(paiCountInfo *PaiCountData, sequence DDZGroup) (groups []DDZGroup) {
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < sequence.Length*5 {
		return
	}
	// 如果顺子已经封顶了，直接返回
	if sequence.Key+int32(sequence.Length)-1 == 14 {
		return
	}
	// 如果key+1到A长度不足，直接返回
	if 14-sequence.Key+1 < int32(sequence.Length) {
		return
	}

outLoop:
	for i := int(sequence.Key) + 1; i <= 14-sequence.Length+1; i++ {
		var cs []cards.Card
		var newPaiCountInfo = paiCountInfo.Clone()
		var ghosts = newPaiCountInfo.Ghosts()
		var gn = newPaiCountInfo.GhostCount()
		for j := 0; j < sequence.Length; j++ {
			key := int32(i + j)
			n := newPaiCountInfo.GetCardCount(key)
			if n+gn < 3 {
				continue outLoop
			}

			pais := newPaiCountInfo.GetCards(key)
			if n == 4 {
				cs = append(cs, pais[:3]...)
				newPaiCountInfo.CountInfo[key] -= 3
				newPaiCountInfo.CountInfo[0] -= 3
				newPaiCountInfo.Cards[key] = newPaiCountInfo.Cards[key][3:]
			} else {
				cs = append(cs, pais[:n]...)
				newPaiCountInfo.CountInfo[key] -= n
				newPaiCountInfo.CountInfo[0] -= n
				newPaiCountInfo.Cards[key] = newPaiCountInfo.Cards[key][n:]

				gn -= 3 - n
				cs = append(cs, ghosts[:3-n]...)
				newPaiCountInfo.CountInfo[1] -= 3 - n
				newPaiCountInfo.Cards[-1] = newPaiCountInfo.Cards[-1][3-n:]
			}
		}
		// 找出所有的单牌组合
		pairGroups := deduplicationGroups(hasPair(newPaiCountInfo))
		if len(pairGroups) < sequence.Length {
			continue outLoop
		}
		// 从单牌的所有组合里找出length个组合
		combs := pkg.Combination(len(pairGroups), sequence.Length)
		singles := make([][]DDZGroup, len(combs))
		for k, ss := range combs {
			for m, v := range ss {
				if v == 1 {
					singles[k] = append(singles[k], pairGroups[m])
				}
			}
		}
		for _, gs := range singles {
			g := DDZGroup{
				Type:       cards.GroupTypeDDZSequenceOfTripletWithPair,
				Length:     sequence.Length,
				GhostCount: ghostCount - gn,
				Key:        int32(i),
			}
			g.Cards = append(g.Cards, cs...)
			for _, gg := range gs {
				g.Cards = append(g.Cards, gg.Cards...)
				g.GhostCount += gg.GhostCount
			}
			groups = append(groups, g)
		}
	}
	return
}

// biggerQuadplexSetWithSingle 是否有比指定key更大的四带二单
func biggerQuadplexSetWithSingle(paiCountInfo *PaiCountData, quadplex DDZGroup) (groups []DDZGroup) {
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 5 {
		return
	}
outLoop:
	for i := quadplex.Key + 1; i <= 15; i++ {
		newPaiCountInfo := paiCountInfo.Clone()
		n := newPaiCountInfo.GetCardCount(i)
		if n+ghostCount < 4 {
			continue
		}
		cs := paiCountInfo.GetCards(i)
		newPaiCountInfo.CountInfo[i] -= n
		if n < 4 {
			newPaiCountInfo.CountInfo[0] -= n
			newPaiCountInfo.CountInfo[1] -= 4 - n
			newPaiCountInfo.CountInfo[i] -= n
			newPaiCountInfo.Cards[i] = newPaiCountInfo.Cards[i][n:]
			cs = append(cs, newPaiCountInfo.Cards[-1][:4-n]...)
			newPaiCountInfo.Cards[-1] = newPaiCountInfo.Cards[-1][4-n:]
		}
		// 找出所有的单牌组合
		singleGroups := deduplicationGroups(hasSingle(newPaiCountInfo))
		if len(singleGroups) < 2 {
			continue outLoop
		}
		// 从单牌的所有组合里找出length个组合
		combs := pkg.Combination(len(singleGroups), 2)
		singles := make([][]DDZGroup, len(combs))
		for k, ss := range combs {
			for m, v := range ss {
				if v == 1 {
					singles[k] = append(singles[k], singleGroups[m])
				}
			}
		}
		for _, gs := range singles {
			g := DDZGroup{
				Type:       cards.GroupTypeDDZQuadplexSetWithSingle,
				Length:     1,
				GhostCount: 4 - n,
				Key:        i,
			}
			g.Cards = append(g.Cards, cs...)
			for _, gg := range gs {
				g.Cards = append(g.Cards, gg.Cards...)
				g.GhostCount += gg.GhostCount
			}
			groups = append(groups, g)
		}
	}
	return
}

// biggerQuadplexSetWithPair 是否有比指定key更大的四带二对
func biggerQuadplexSetWithPair(paiCountInfo *PaiCountData, quadplex DDZGroup) (groups []DDZGroup) {
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 5 {
		return
	}
outLoop:
	for i := quadplex.Key + 1; i <= 15; i++ {
		newPaiCountInfo := paiCountInfo.Clone()
		n := newPaiCountInfo.GetCardCount(i)
		if n+ghostCount < 4 {
			continue
		}
		cs := paiCountInfo.GetCards(i)
		newPaiCountInfo.CountInfo[i] -= n
		if n < 4 {
			newPaiCountInfo.CountInfo[0] -= n
			newPaiCountInfo.CountInfo[1] -= 4 - n
			newPaiCountInfo.CountInfo[i] -= n
			newPaiCountInfo.Cards[i] = newPaiCountInfo.Cards[i][n:]
			cs = append(cs, newPaiCountInfo.Cards[-1][:4-n]...)
			newPaiCountInfo.Cards[-1] = newPaiCountInfo.Cards[-1][4-n:]
		}
		// 找出所有的单牌组合
		pairGroups := deduplicationGroups(hasPair(newPaiCountInfo))
		if len(pairGroups) < 2 {
			continue outLoop
		}
		// 从单牌的所有组合里找出length个组合
		combs := pkg.Combination(len(pairGroups), 2)
		pairs := make([][]DDZGroup, len(combs))
		for k, ss := range combs {
			for m, v := range ss {
				if v == 1 {
					pairs[k] = append(pairs[k], pairGroups[m])
				}
			}
		}
		for _, gs := range pairs {
			g := DDZGroup{
				Type:       cards.GroupTypeDDZQuadplexSetWithPair,
				Length:     1,
				GhostCount: 4 - n,
				Key:        i,
			}
			g.Cards = append(g.Cards, cs...)
			for _, gg := range gs {
				g.Cards = append(g.Cards, gg.Cards...)
				g.GhostCount += gg.GhostCount
			}
			groups = append(groups, g)
		}
	}
	return
}

// biggerBomb 是否有比指定key大的普通炸弹
// 四软炸<四硬炸<四癞子炸<五软炸<五癞子炸<…<12软炸<王炸
func biggerBomb(paiCountInfo *PaiCountData, srcGroup DDZGroup) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < srcGroup.Length {
		return
	}

	bombLength := srcGroup.Length
	bombGhost := srcGroup.GhostCount
	for i := 3; i <= 15; i++ {
		cs := paiCountInfo.GetCards(int32(i))
		n := paiCountInfo.CountInfo[i]
		switch {
		case bombGhost > 0 && bombLength > bombGhost: // 软炸弹(级别最低的炸弹): 找到相同长度的纯赖子炸弹或者长度相同但key更大的软炸弹或者长度key都相同的硬炸弹，或者长度更长的炸弹
			// 找长度相同但key更大的软炸弹
			if n > 0 && i > int(srcGroup.Key) && n+ghostCount >= bombLength {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZBomb,
					Length:     bombLength,
					GhostCount: bombLength - n,
					Key:        int32(i),
					Cards:      cs, // 将牌值为i的牌添加进group中
				}
				g.Cards = append(g.Cards, ghosts[:bombLength-n]...)
				groups = append(groups, g)
			}
			fallthrough
		case bombGhost == 0: // 硬炸弹(长度一定为4)：只有找长度相同的纯赖子炸弹或者长度相同但key更大的炸弹或者长度更长的炸弹
			// 长度相同但key更大的硬炸弹
			if i > int(srcGroup.Key) && n == 4 && bombLength == 4 {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZBomb,
					Length:     4,
					GhostCount: 0,
					Key:        int32(i),
					Cards:      cs,
				}
				groups = append(groups, g)
			}
			// 找长度相同的纯赖子炸弹
			if ghostCount >= bombLength {
				// 从bomb长度开始，没多一个赖子就多一种纯赖子炸弹的情况
				for j := bombLength; j <= ghostCount; j++ {
					g := DDZGroup{
						Type:       cards.GroupTypeDDZBomb,
						Length:     j,
						GhostCount: j,
						Key:        -1,
						Cards:      ghosts[:j],
					}
					groups = append(groups, g)
				}
			}
			fallthrough
		case bombLength == bombGhost: // 纯赖子炸弹：只有找长度更长的炸弹才能大过
			if n+ghostCount <= bombLength {
				break
			}
			// 找长度更长的炸弹
			for j := bombLength - n + 1; j <= ghostCount; j++ {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZBomb,
					Length:     n + j,
					GhostCount: j,
					Key:        int32(i),
					Cards:      cs,
				}
				g.Cards = append(g.Cards, ghosts[:j]...)
				groups = append(groups, g)
			}
		}
	}
	return
}
