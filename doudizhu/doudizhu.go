package doudizhu

import (
	"ALG-QiPai/cards"
	"errors"
	"fmt"
	"reflect"
	"sort"
)

func PrintCards(handCards []cards.Card) (desc string) {
	for _, c := range handCards {
		sd, vd := "", ""
		suit, value := c.Suit(), c.Value()

		switch suit {
		case cards.CardSuitDiamond:
			sd = "方块"
		case cards.CardSuitClub:
			sd = "梅花"
		case cards.CardSuitHeart:
			sd = "红桃"
		case cards.CardSuitSpade:
			sd = "黑桃"
		}

		switch value {
		case 11:
			vd = "J"
		case 12:
			vd = "Q"
		case 14:
			vd = "A"
		case 13:
			vd = "K"
		case 15:
			vd = "2"
		case 16:
			vd = "小王"
		case 17:
			vd = "大王"
		default:
			vd = fmt.Sprintf("%d", value)
		}
		desc += sd + vd + " "
	}
	return
}

type ddzCards []cards.Card

func (cs ddzCards) Len() int {
	return len(cs)
}

func (cs ddzCards) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs ddzCards) Less(i, j int) bool {
	return cs[i].Value() < cs[j].Value()
}

func sortCards(targetCards []cards.Card) {
	sort.Sort(ddzCards(targetCards))
}

// PaiCountData 解析一副牌的数量信息
type PaiCountData struct {
	// 索引0为非赖子牌总数量，索引1为赖子总数量，索引2预留，其余每个牌值对应的牌数量，下标表示牌值，元素为数量
	CountInfo [18]int
	// 牌值对应的牌的集合
	Cards map[int32][]cards.Card
}

func (pd *PaiCountData) Clone() *PaiCountData {
	data := &PaiCountData{
		CountInfo: pd.CountInfo,
		Cards:     make(map[int32][]cards.Card),
	}

	for k, pais := range pd.Cards {
		data.Cards[k] = make([]cards.Card, len(pais))
		copy(data.Cards[k], pais)
	}
	return data
}

func (pd *PaiCountData) PaiCount() int {
	return pd.CountInfo[0]
}

func (pd *PaiCountData) GhostCount() int {
	return pd.CountInfo[1]
}

func (pd *PaiCountData) Ghosts() []cards.Card {
	ghosts := make([]cards.Card, pd.CountInfo[1])
	copy(ghosts, pd.Cards[-1])
	return ghosts
}

func (pd *PaiCountData) GetCards(v int32) []cards.Card {
	cs := make([]cards.Card, pd.CountInfo[v])
	copy(cs, pd.Cards[v])
	return cs
}

func (pd *PaiCountData) GetCardCount(v int32) int {
	return pd.CountInfo[v]
}

// DDZGroup 斗地主牌型组合
type DDZGroup struct {
	Type       cards.GroupType // 牌型
	Length     int             // 牌型长度
	GhostCount int             // 牌型包含的赖子牌数量
	Key        int32           // 用于相同牌型之间的比较
	Cards      []cards.Card    // 组成牌型的牌
}

func (dg DDZGroup) String() (desc string) {
	switch dg.Type {
	case cards.GroupTypeDDZSingle:
		desc = "单牌"
	case cards.GroupTypeDDZPair:
		desc = "对子"
	case cards.GroupTypeDDZTriplet:
		desc = "三不带"
	case cards.GroupTypeDDZTripletWithSingle:
		desc = "三带一"
	case cards.GroupTypeDDZTripletWithPair:
		desc = "三带对"
	case cards.GroupTypeDDZSequence:
		desc = "顺子"
	case cards.GroupTypeDDZSequenceOfPair:
		desc = "连对"
	case cards.GroupTypeDDZSequenceOfTriplet:
		desc = "飞机不带牌"
	case cards.GroupTypeDDZSequenceOfTripletWithSingle:
		desc = "飞机带单牌"
	case cards.GroupTypeDDZSequenceOfTripletWithPair:
		desc = "飞机带对子"
	case cards.GroupTypeDDZQuadplexSetWithSingle:
		desc = "四带二(单牌)"
	case cards.GroupTypeDDZQuadplexSetWithPair:
		desc = "四带二(对子)"
	case cards.GroupTypeDDZBomb:
		desc = "炸弹"
	case cards.GroupTypeDDZRocket:
		desc = "火箭"
	}
	desc = fmt.Sprintf("Type:%s, Length:%v, GhostCount:%v, Key:%v, Cards:%v", desc, dg.Length, dg.GhostCount, dg.Key, PrintCards(dg.Cards))
	return
}

func (dg DDZGroup) IsZero() bool {
	return dg.GhostCount == 0 &&
		dg.Key == 0 &&
		dg.Type == cards.GroupTypeNone &&
		dg.Length == 0 &&
		len(dg.Cards) == 0
}

func (dg DDZGroup) Reset() {
	if dg.IsZero() {
		return
	}
	dg.Type = cards.GroupTypeNone
	dg.Length = 0
	dg.GhostCount = 0
	dg.Key = 0
	dg.Cards = nil
}

// 对一组牌型进行去重
func deduplicationGroups(srcGroups []DDZGroup) (groups []DDZGroup) {
	for _, sg := range srcGroups {
		flag := true
		for _, dg := range groups {
			if reflect.DeepEqual(sg, dg) {
				flag = false
				break
			}
		}
		if flag {
			groups = append(groups, sg)
		}
	}
	return
}

// 比较两个牌型的大小
func compareGroups(g1, g2 DDZGroup) (bigger int, err error) {
	switch {
	case g1.Type == g2.Type: // 两个类型相同, 比较key和长度
		switch g1.Type {
		case cards.GroupTypeDDZBomb: // 都是炸弹
			if g1.Length == g2.Length {
				if g2.Key == -1 {
					bigger = -1
					return
				}
				if g1.Key == -1 {
					bigger = 1
					return
				}
				if g1.Key == g2.Key {
					return
				}
				if g1.Key > g2.Key {
					bigger = 1
					return
				} else {
					bigger = -1
					return
				}
			} else if g1.Length > g2.Length {
				bigger = 1
				return
			} else {
				bigger = -1
				return
			}
		default: // 非炸弹
			if g1.Length != g1.Length {
				err = errors.New("牌型相同, 牌数量不同, 不能做比较")
				return
			}
			if g1.Key == g2.Key {
				return
			} else if g1.Key > g2.Key {
				bigger = 1
				return
			} else {
				bigger = -1
				return
			}
		}
	case g1.Type >= cards.GroupTypeDDZBomb && g2.Type < cards.GroupTypeDDZBomb: // 一个是炸弹，一个非炸弹
		bigger = 1
	case g1.Type < cards.GroupTypeDDZBomb && g2.Type >= cards.GroupTypeDDZBomb: // 一个是炸弹，一个是火箭
		bigger = -1
	case g1.Type == cards.GroupTypeDDZRocket || g2.Type == cards.GroupTypeDDZRocket: // 有一个是火箭
		if g1.Type == cards.GroupTypeDDZRocket {
			bigger = 1
		} else {
			bigger = -1
		}
	default:
		err = errors.New("不同牌型不能比较")
	}
	return
}

// 比较c1, c2的大小, 大于返回1, 小于返回-1, 相等返回0
func compareCards(c1, c2 cards.Card) (bigger int) {
	v1, v2 := c1.Value(), c2.Value()
	if v1 > v2 {
		bigger = 1
	}
	if v1 < v2 {
		bigger = -1
	}
	return
}

func getCardsCountArray(targetCards []cards.Card, ghosts ...cards.Card) (paiCountData *PaiCountData) {
	// 下标为牌值, 元素为对应牌值的牌的数量
	// 位置0存放非赖子牌数量，位置1存放赖子牌数量，位置2存放赖子牌值
	paiCountData = &PaiCountData{
		CountInfo: [18]int{},
		Cards:     make(map[int32][]cards.Card),
	}
	v := int32(0)
	for _, c := range targetCards {
		if !IsGhost(c, ghosts...) {
			v = c.Value()
			paiCountData.CountInfo[v] += 1
			paiCountData.CountInfo[0] += 1
			paiCountData.Cards[v] = append(paiCountData.Cards[v], c)
		} else {
			paiCountData.CountInfo[1] += 1
			paiCountData.Cards[-1] = append(paiCountData.Cards[-1], c)
		}
	}
	// 赖子牌根据牌值从小到大排序
	sortCards(paiCountData.Cards[-1])
	return
}

// IsGhost 判断一张牌是否是赖子
func IsGhost(target cards.Card, ghosts ...cards.Card) (ok bool) {
	if len(ghosts) == 0 {
		return
	}
	for _, g := range ghosts {
		ghostValue, targetValue := g.Value(), target.Value()
		switch {
		case ghostValue > 15:
			ok = targetValue == 16 || targetValue == 17
		default:
			ok = ghostValue == targetValue
		}
		if ok {
			break
		}
	}
	return
}

// hasSingle 找出所有的单牌
func hasSingle(paiCountInfo *PaiCountData) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount == 0 {
		return
	}
	// 先找单张赖子
	if ghostCount > 0 {
		for _, c := range ghosts {
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
	// 从非赖子中找出单牌
	for i := 3; i <= 17; i++ {
		n := paiCountInfo.GetCardCount(int32(i))
		pais := paiCountInfo.GetCards(int32(i))
		for j := 0; j < n; j++ {
			g := DDZGroup{
				Type:       cards.GroupTypeDDZSingle,
				Length:     1,
				GhostCount: 0,
				Key:        int32(i),
				Cards:      []cards.Card{pais[j]},
			}
			groups = append(groups, g)
		}
	}
	return
}

// hasPair 找出所有的对子
func hasPair(paiCountInfo *PaiCountData) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	if paiCount+ghostCount < 2 {
		return
	}

	// 找纯赖子的对子
	if ghostCount >= 2 {
		for _, c := range ghosts {
			if c.Value() > 15 {
				continue
			}
			g := DDZGroup{
				Type:       cards.GroupTypeDDZPair,
				Length:     1,
				GhostCount: 2,
				Key:        c.Value(),
			}
			g.Cards = append(g.Cards, ghosts[:2]...)
			groups = append(groups, g)
		}
	}

	for i := 3; i <= 15; i++ {
		n := paiCountInfo.GetCardCount(int32(i))
		if n == 0 || n+ghostCount < 2 {
			continue
		}
		g := DDZGroup{
			Type:   cards.GroupTypeDDZPair,
			Length: 1,
			Key:    int32(i),
		}
		pais := paiCountInfo.GetCards(int32(i))
		if n == 1 {
			paiCountInfo.CountInfo[0] -= n
			paiCountInfo.CountInfo[i] -= n
			paiCountInfo.Cards[int32(i)] = paiCountInfo.Cards[int32(i)][:1]
			paiCountInfo.CountInfo[1] -= 1
			paiCountInfo.Cards[-1] = paiCountInfo.Cards[-1][1:]
			g.GhostCount = 1
			g.Cards = append(g.Cards, pais[0], ghosts[0])
			groups = append(groups, g)
			gs := hasPair(paiCountInfo)
			if len(gs) > 0 {
				groups = append(groups, gs...)
			}
			paiCountInfo.CountInfo[0] += n
			paiCountInfo.CountInfo[i] += n
			paiCountInfo.Cards[int32(i)] = pais
			paiCountInfo.CountInfo[1] += 1
			paiCountInfo.Cards[-1] = ghosts
		} else {
			paiCountInfo.CountInfo[0] -= 2
			paiCountInfo.CountInfo[i] -= 2
			paiCountInfo.Cards[int32(i)] = paiCountInfo.Cards[int32(i)][2:]
			g.Cards = append(g.Cards, pais[:2]...)
			groups = append(groups, g)
			gs := hasPair(paiCountInfo)
			if len(gs) > 0 {
				groups = append(groups, gs...)
			}
			paiCountInfo.CountInfo[0] += 2
			paiCountInfo.CountInfo[i] += 2
			paiCountInfo.Cards[int32(i)] = pais
		}
	}
	return
}

// hasBomb 是否有炸弹: 找出所有的炸弹组合
func hasBomb(paiCountInfo *PaiCountData) (groups []DDZGroup) {
	ghosts := paiCountInfo.Ghosts()
	paiCount, ghostCount := paiCountInfo.PaiCount(), paiCountInfo.GhostCount()
	// 炸弹张数必须满足大于4
	if paiCount+ghostCount < 4 {
		return
	}

	// 先找纯赖子炸弹
	if ghostCount >= 4 {
		// 纯赖子的炸弹，赖子张数不同，大小也不同，张数越多的炸弹越大
		// 纯赖子的炸弹key值设为-1
		for i := 4; i <= ghostCount; i++ {
			g := DDZGroup{
				Type:       cards.GroupTypeDDZBomb,
				Length:     i,
				GhostCount: i,
				Key:        -1,
			}
			g.Cards = append(g.Cards, ghosts[:i]...)
			groups = append(groups, g)
		}
	}

	// 纯赖子炸弹找完后，剩下只考虑硬炸弹和软炸弹
	for i := 3; i <= 15; i++ {
		cs := paiCountInfo.GetCards(int32(i))
		n := paiCountInfo.CountInfo[i]
		switch n {
		case 1:
			if ghostCount < 3 {
				break
			}
			for j := 3; j <= ghostCount; j++ {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZBomb,
					Length:     1 + j,
					GhostCount: j,
					Key:        int32(i),
					Cards:      cs,
				}
				g.Cards = append(g.Cards, ghosts[:j]...)
				groups = append(groups, g)          // 将当前找到的炸弹加入结果
				paiCountInfo.CountInfo[i] -= 1      // 扣除一张非赖子牌
				paiCountInfo.CountInfo[0] -= 1      // 牌总数量减1
				paiCountInfo.CountInfo[1] -= j      // 扣除j张赖子
				paiCountInfo.Cards[-1] = ghosts[j:] // 移除被用了的赖子
				bs := hasBomb(paiCountInfo)         // 进行下一次递归
				if len(bs) > 0 {
					groups = append(groups, bs...)
				}
				paiCountInfo.CountInfo[i] += 1                                         // 归还扣除的那张牌
				paiCountInfo.CountInfo[0] += 1                                         // 归还牌总数
				paiCountInfo.CountInfo[1] += j                                         // 归还赖子
				paiCountInfo.Cards[-1] = append(paiCountInfo.Cards[-1], ghosts[:j]...) // 归还赖子
			}
		case 2:
			if ghostCount < 2 {
				break
			}
			// 出两张i
			for j := 2; j <= ghostCount; j++ {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZBomb,
					Length:     2 + j,
					GhostCount: j,
					Key:        int32(i),
					Cards:      cs,
				}
				g.Cards = append(g.Cards, ghosts[:j]...)
				groups = append(groups, g)
				paiCountInfo.CountInfo[i] -= 2
				paiCountInfo.CountInfo[0] -= 2
				paiCountInfo.CountInfo[1] -= j
				paiCountInfo.Cards[-1] = ghosts[j:]
				bs := hasBomb(paiCountInfo)
				if len(bs) > 0 {
					groups = append(groups, bs...)
				}
				paiCountInfo.CountInfo[i] += 2
				paiCountInfo.CountInfo[0] += 2
				paiCountInfo.CountInfo[1] += j
				paiCountInfo.Cards[-1] = append(paiCountInfo.Cards[-1], ghosts[:j]...)
			}
		case 3:
			if ghostCount < 1 {
				break
			}
			// 出三张i
			for j := 1; j <= ghostCount; j++ {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZBomb,
					Length:     3 + j,
					GhostCount: j,
					Key:        int32(i),
					Cards:      cs,
				}
				g.Cards = append(g.Cards, ghosts[:j]...)
				paiCountInfo.CountInfo[i] -= 3
				paiCountInfo.CountInfo[0] -= 3
				paiCountInfo.CountInfo[1] -= j
				paiCountInfo.Cards[-1] = ghosts[j:]
				groups = append(groups, g)
				bomb1s := hasBomb(paiCountInfo)
				if len(bomb1s) > 0 {
					groups = append(groups, bomb1s...)
				}
				paiCountInfo.CountInfo[i] += 3
				paiCountInfo.CountInfo[0] += 3
				paiCountInfo.CountInfo[1] += j
				paiCountInfo.Cards[-1] = append(paiCountInfo.Cards[-1], ghosts[:j]...)
			}
		case 4:
			for j := 0; j <= ghostCount; j++ {
				g := DDZGroup{
					Type:       cards.GroupTypeDDZBomb,
					Length:     4 + j,
					GhostCount: j,
					Key:        int32(i),
					Cards:      cs,
				}
				g.Cards = append(g.Cards, ghosts[:j]...)
				groups = append(groups, g)
				paiCountInfo.CountInfo[i] -= 4
				paiCountInfo.CountInfo[0] -= 4
				paiCountInfo.CountInfo[1] -= j
				bs := hasBomb(paiCountInfo)
				if len(bs) > 0 {
					groups = append(groups, bs...)
				}
				paiCountInfo.CountInfo[i] += 4
				paiCountInfo.CountInfo[0] += 4
				paiCountInfo.CountInfo[1] += j
			}
		}
	}
	return
}

// hasRocket 是否有火箭
func hasRocket(paiCountInfo *PaiCountData) (group DDZGroup) {
	if paiCountInfo.CountInfo[16] == 1 && paiCountInfo.CountInfo[17] == 1 {
		group = DDZGroup{
			Type:   cards.GroupTypeDDZRocket,
			Key:    16,
			Length: 2,
			Cards:  []cards.Card{paiCountInfo.Cards[16][0], paiCountInfo.Cards[17][0]},
		}
	}
	return
}

// hasJoker 是否有大小王
func hasJoker(paiCountInfo *PaiCountData) (ok bool) {
	return paiCountInfo.CountInfo[16] > 0 || paiCountInfo.CountInfo[17] > 0
}
