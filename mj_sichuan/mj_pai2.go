package mj_sichuan

import (
	"fmt"
	"sort"
)

//这里是利用位操作

/*
麻将牌一共108张，三种花色，每个花色9种牌值，每个牌值4张牌，可以使用一个字节（8位）来表示所有的麻将牌
低四位表示牌值，0～15
5至6位表示花色：0～3
剩下两位待定

表示花色的两位：
00: 不用
01: 筒
10: 条
11: 万

所以，最终的麻将组合为：
00 01 0001 ——> 17(1筒)
00 01 0010 ——> 18(2筒)
00 01 0011 ——> 19(3筒)
00 01 0100 ——> 20(4筒)
00 01 0101 ——> 21(5筒)
00 01 0110 ——> 22(6筒)
00 01 0111 ——> 23(7筒)
00 01 1000 ——> 24(8筒)
00 01 1001 ——> 25(9筒)

00 10 0001 ——> 33(1条)
00 10 0010 ——> 34(2条)
00 10 0011 ——> 35(3条)
00 10 0100 ——> 36(4条)
00 10 0101 ——> 37(5条)
00 10 0110 ——> 38(6条)
00 10 0111 ——> 39(7条)
00 10 1000 ——> 40(8条)
00 10 1001 ——> 41(9条)

00 11 0001 ——> 49(1万)
00 11 0010 ——> 50(2万)
00 11 0011 ——> 51(3万)
00 11 0100 ——> 52(4万)
00 11 0101 ——> 53(5万)
00 11 0110 ——> 54(6万)
00 11 0111 ——> 55(7万)
00 11 1000 ——> 56(8万)
00 11 1001 ——> 57(9万)
*/

//花色
const (
	None = iota
	Tong //筒
	Tiao //条
	Wan  //万
)

//牌定义
type (
	Pai     uint8
	PaiList []Pai
)

var AllPais = []Pai{
	17, 18, 19, 20, 21, 22, 23, 24, 25,
	17, 18, 19, 20, 21, 22, 23, 24, 25,
	17, 18, 19, 20, 21, 22, 23, 24, 25,
	17, 18, 19, 20, 21, 22, 23, 24, 25,

	33, 34, 35, 36, 37, 38, 39, 40, 41,
	33, 34, 35, 36, 37, 38, 39, 40, 41,
	33, 34, 35, 36, 37, 38, 39, 40, 41,
	33, 34, 35, 36, 37, 38, 39, 40, 41,

	49, 50, 51, 52, 53, 54, 55, 56, 57,
	49, 50, 51, 52, 53, 54, 55, 56, 57,
	49, 50, 51, 52, 53, 54, 55, 56, 57,
	49, 50, 51, 52, 53, 54, 55, 56, 57,
}

func (p Pai) GetFlower() uint8 {
	return uint8(p & 48 >> 4) //与110000位运算得到花色
}

func (p Pai) GetValue() uint8 {
	return uint8(p & 15) //与"1111"位运算得到牌值
}

func (p Pai) String() (des string) {
	switch p.GetFlower() {
	case Tong:
		des = fmt.Sprintf("%d%s", p.GetValue(), "筒")
	case Tiao:
		des = fmt.Sprintf("%d%s", p.GetValue(), "条")
	case Wan:
		des = fmt.Sprintf("%d%s", p.GetValue(), "万")
	default:
		des = "invalid flower"
	}
	return
}

func (ps PaiList) String(args ...interface{}) (des string) {
	if len(args) > 0 {
		switch t := args[0].(type) {
		case bool:
			if t {
				sort.Sort(ps)
			}
		}
	}
	for i, p := range ps {
		if i == 0 {
			des += p.String()
		} else {
			des += ", " + p.String()
		}
	}
	return
}

//排序
func (ps PaiList) Len() int {
	return len(ps)
}

func (ps PaiList) Less(i, j int) bool {
	return ps[i] < ps[j]
}

func (ps PaiList) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps PaiList) Copy() PaiList {
	cp := make(PaiList, len(ps))
	copy(cp, ps)
	return cp
}

func (ps PaiList) GetValueCount() [][]uint8 {
	//获取每个花色对应每个牌值的数量,花色值减一作为下标
	result := make([][]uint8, 3)
	for i := range result {
		//牌值减一作为下标
		result[i] = make([]uint8, 10)
	}
	for i := range ps {
		result[ps[i].GetFlower()-1][0]++
		result[ps[i].GetFlower()-1][ps[i].GetValue()]++
	}
	return result
}

//判断是否可以胡牌
func (ps PaiList) CanHu() (canHu bool, jiangPai uint8) {
	//胡牌数最多14张
	if ps.Len() > 14 || ps.Len() <= 1 {
		return
	}

	valueList := ps.GetValueCount()
	fmt.Println(valueList)
	//必须缺一门花色才能胡
	if valueList[0][0]*valueList[1][0]*valueList[2][0] != 0 {
		return
	}

	canHu, jiangPai = tryHu1(valueList, ps.Len())

	return
}

func GetPaiByFlowerAndValue(flower uint8, value uint8) Pai {
	flower = flower << 4
	return Pai(flower + value)
}

//试图胡牌，如果可以胡牌则同时返回将牌
//1. 从数量大于1的牌值里找出两张作为将牌
//2. 判断其余牌能否组成砍或者杠或者顺子
func tryHu1(data [][]uint8, paiCount int) (canHu bool, jiangPai uint8) {
	if paiCount == 0 {
		canHu = true
		return
	}
	if paiCount%3 == 2 {
		for i := range data {
			for j := range data[i] {
				if j == 0 {
					continue
				}
				if data[i][j] < 2 {
					continue
				}
				data[i][j] -= 2
				canHu, jiangPai = tryHu1(data, paiCount-2)
				if canHu {
					jiangPai = uint8(GetPaiByFlowerAndValue(uint8(i+1), uint8(j)))
					return
				}
				data[i][j] += 2
			}
		}
	} else {
		//找顺子、砍、杠
		for i := range data {
			for j := 1; j <= 7; j++ {
				if data[i][j] > 0 && data[i][j+1] > 0 && data[i][j+2] > 0 {
					data[i][j] -= 1
					data[i][j+1] -= 1
					data[i][j+2] -= 1
					canHu, jiangPai = tryHu1(data, paiCount-3)
					if canHu {
						return
					}
					data[i][j] += 1
					data[i][j+1] += 1
					data[i][j+2] += 1
				}
			}
		}
		for i := range data {
			for j := 1; j <= 9; j++ {
				if data[i][j] >= 3 {
					data[i][j] -= 3
					canHu, jiangPai = tryHu1(data, paiCount-3)
					if canHu {
						return
					}
					data[i][j] += 3
				}
			}
		}
	}
	return
}
