package texas

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	//牌初始长度
	AllPokerNum = 52

	//花色
	FlowerHei     = iota //黑桃
	FlowerHong           //红桃
	FlowerMei            //梅花
	FlowerFang           //方块
	FlowerUnknown = -1   //错误花色
)

//poker map
var pokerMap map[int32]string

func init() {
	pokerMap = make(map[int32]string, AllPokerNum)

	//描述：id_value_flower_name
	//黑桃
	pokerMap[1] = "1_2_黑桃_2"
	pokerMap[2] = "2_3_黑桃_3"
	pokerMap[3] = "3_4_黑桃_4"
	pokerMap[4] = "4_5_黑桃_5"
	pokerMap[5] = "5_6_黑桃_6"
	pokerMap[6] = "6_7_黑桃_7"
	pokerMap[7] = "7_8_黑桃_8"
	pokerMap[8] = "8_9_黑桃_9"
	pokerMap[9] = "9_10_黑桃_10"
	pokerMap[10] = "10_11_黑桃_J"
	pokerMap[11] = "11_12_黑桃_Q"
	pokerMap[12] = "12_13_黑桃_K"
	pokerMap[13] = "13_14_黑桃_A"

	//红桃
	pokerMap[14] = "14_2_红桃_2"
	pokerMap[15] = "15_3_红桃_3"
	pokerMap[16] = "16_4_红桃_4"
	pokerMap[17] = "17_5_红桃_5"
	pokerMap[18] = "18_6_红桃_6"
	pokerMap[19] = "19_7_红桃_7"
	pokerMap[20] = "20_8_红桃_8"
	pokerMap[21] = "21_9_红桃_9"
	pokerMap[22] = "22_10_红桃_10"
	pokerMap[23] = "23_11_红桃_J"
	pokerMap[24] = "24_12_红桃_Q"
	pokerMap[25] = "25_13_红桃_K"
	pokerMap[26] = "26_14_红桃_A"

	//梅花
	pokerMap[27] = "27_2_梅花_2"
	pokerMap[28] = "28_3_梅花_3"
	pokerMap[29] = "29_4_梅花_4"
	pokerMap[30] = "30_5_梅花_5"
	pokerMap[31] = "31_6_梅花_6"
	pokerMap[32] = "32_7_梅花_7"
	pokerMap[33] = "33_8_梅花_8"
	pokerMap[34] = "34_9_梅花_9"
	pokerMap[35] = "35_10_梅花_10"
	pokerMap[36] = "36_11_梅花_J"
	pokerMap[37] = "37_12_梅花_Q"
	pokerMap[38] = "38_13_梅花_K"
	pokerMap[39] = "39_14_梅花_A"

	//方块
	pokerMap[40] = "40_2_方块_2"
	pokerMap[41] = "41_3_方块_3"
	pokerMap[42] = "42_4_方块_4"
	pokerMap[43] = "43_5_方块_5"
	pokerMap[44] = "44_6_方块_6"
	pokerMap[45] = "45_7_方块_7"
	pokerMap[46] = "46_8_方块_8"
	pokerMap[47] = "47_9_方块_9"
	pokerMap[48] = "48_10_方块_10"
	pokerMap[49] = "49_11_方块_J"
	pokerMap[50] = "50_12_方块_Q"
	pokerMap[51] = "51_13_方块_K"
	pokerMap[52] = "52_14_方块_A"

}

//牌的结构
type Poker struct {
	Id     int32  //牌的ID：1~52
	Value  int32  //牌值：牌面的值
	Flower int32  //花色
	Name   string //牌面描述
}

func (p *Poker) GetId() int32 {
	return p.Id
}

func (p *Poker) GetValue() int32 {
	return p.Value
}

func (p *Poker) GetFlower() int32 {
	return p.Flower
}

func (p *Poker) GetName() string {
	return p.Name
}

func (p *Poker) Desc() (desc string) {
	var flower string
	switch p.GetFlower() {
	case FlowerHei:
		flower = "黑桃"
	case FlowerHong:
		flower = "红桃"
	case FlowerMei:
		flower = "梅花"
	case FlowerFang:
		flower = "方块"
	case FlowerUnknown:
		flower = ""
	}
	desc = fmt.Sprintf("%v%v ", flower, p.GetName())
	return
}

//根据map index解析一张牌
func parsePokerByIndex(index int32) (id int32, value int32, flower int32, name string) {
	if index < 1 || index > AllPokerNum {
		return -1, -1, -1, ""
	}
	pokerStr := strings.Split(pokerMap[index], "_")

	id = index
	v1, _ := strconv.Atoi(pokerStr[1])
	value = int32(v1)
	name = pokerStr[3]

	switch pokerStr[2] {
	case "黑桃":
		flower = FlowerHei
	case "红桃":
		flower = FlowerHong
	case "梅花":
		flower = FlowerMei
	case "方块":
		flower = FlowerFang
	}
	return
}

//增加shuffle洗牌算法
func shuffleXiPai(paiCount int32) []int32 {
	var result = make([]int32, paiCount)
	for i := 0; i < int(paiCount); i++ {
		result[i] = int32(i + 1)
	}
	for i := paiCount - 1; i > 0; i-- {
		randIndex := rand.Intn(int(i))
		result[i], result[randIndex] = result[randIndex], result[i]
	}
	return result
}

//根据牌的张数洗牌
func shuffleXiPai2(startIndex int, paiCount int32) []int32 {
	result := make([]int32, paiCount)

	//生成原生索引slice
	original := make([]int32, paiCount)
	for i := 0; i < int(paiCount); i++ {
		original[i] = int32(i + startIndex)
	}

	rand.Seed(time.Now().UnixNano())
	for i := range result {
		index := rand.Intn(len(original))
		result[i] = original[index]
		original = append(original[:index], original[index+1:]...)
	}
	return result
}

func shuffleXiPai3() []int32 {
	var ids = make([]int32, AllPokerNum)
	for i := 0; i < AllPokerNum; i++ {
		ids[i] = int32(i + 1)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(52, func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})
	return ids
}

//根据ID初始化一张牌
func InitPokerById(id int32) *Poker {
	poker := &Poker{}
	poker.Id, poker.Value, poker.Flower, poker.Name = parsePokerByIndex(id)
	return poker
}
