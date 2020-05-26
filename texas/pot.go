package texas

import "sort"

const (
	sidePot potType = 1
	mainPot potType = 2
)

type node struct {
	userId   int32
	amount   int32
	allInTag bool //是否allin，有玩家all
}

type (
	betNodes []*node
	potType  int
)

func (b betNodes) Len() int {
	return len(b)
}

func (b betNodes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b betNodes) Less(i, j int) bool {
	return b[i].amount < b[j].amount
}

func (b *betNodes) AddNode(userId, amount int32, allIn bool) {
	for i := range *b {
		if (*b)[i] != nil && (*b)[i].userId == userId {
			(*b)[i].amount += amount
			(*b)[i].allInTag = allIn
			return
		}
	}
	n := &node{
		userId:   userId,
		amount:   amount,
		allInTag: allIn,
	}
	*b = append(*b, n)
	//if allIn {
	//	calPot(*b)
	//}
}

//奖池
type pot struct {
	pType   potType //奖池类型
	amount  int32   //奖池金额
	userIds []int32 //奖池参与者
}

//计算奖池
func calPot(nodes betNodes) (result []*pot) {
	sort.Sort(nodes)

	for i, b := range nodes {
		if b.amount <= 0 {
			continue
		}
		s := nodes[i:]
		amount := b.amount
		pot := &pot{pType: sidePot, amount: b.amount * int32(len(s))}
		for j := range s {
			s[j].amount -= amount
			pot.userIds = append(pot.userIds, s[j].userId)
		}
		result = append(result, pot)
	}
	var index = -1
	var maxIds = -1
	for i, p := range result {
		userLen := len(p.userIds)
		if p != nil && userLen > maxIds {
			maxIds = userLen
			index = i
		}
	}
	result[index].pType = mainPot
	return result
}
