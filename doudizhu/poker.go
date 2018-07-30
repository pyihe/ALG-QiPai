package doudizhu

import (
	"fmt"
	"strconv"
	"strings"
)

/*
    @Create by GoLand
    @Author: hong
    @Time: 2018/7/26 17:04 
    @File: poker.go    
*/

/*定义牌结构以及辅助方法*/

const (
	//牌初始长度
	ALL_POKER_NUM = 54

	//花色
	FLOWER_RED_JOKER   = 6
	FLOWER_BLACK_JOKER = 5
	FLOWER_HEI         = 4
	FLOWER_HONG        = 3
	FLOWER_MEI         = 2
	FLOWER_FANG        = 1
	FLOWER_UNKNOWN     = -1
)

//poker map
var pokerMap map[int32]string

func init() {
	pokerMap = make(map[int32]string, ALL_POKER_NUM)

	//描述：id_value_flower_name
	//黑桃
	pokerMap[1] = "1_3_黑桃_3"
	pokerMap[2] = "2_4_黑桃_4"
	pokerMap[3] = "3_5_黑桃_5"
	pokerMap[4] = "4_6_黑桃_6"
	pokerMap[5] = "5_7_黑桃_7"
	pokerMap[6] = "6_8_黑桃_8"
	pokerMap[7] = "7_9_黑桃_9"
	pokerMap[8] = "8_10_黑桃_10"
	pokerMap[9] = "9_11_黑桃_J"
	pokerMap[10] = "10_12_黑桃_Q"
	pokerMap[11] = "11_13_黑桃_K"
	pokerMap[12] = "12_14_黑桃_A"
	pokerMap[13] = "13_15_黑桃_2"

	//红桃
	pokerMap[14] = "14_3_红桃_3"
	pokerMap[15] = "15_4_红桃_4"
	pokerMap[16] = "16_5_红桃_5"
	pokerMap[17] = "17_6_红桃_6"
	pokerMap[18] = "18_7_红桃_7"
	pokerMap[19] = "19_8_红桃_8"
	pokerMap[20] = "20_9_红桃_9"
	pokerMap[21] = "21_10_红桃_10"
	pokerMap[22] = "22_11_红桃_J"
	pokerMap[23] = "23_12_红桃_Q"
	pokerMap[24] = "24_13_红桃_K"
	pokerMap[25] = "25_14_红桃_A"
	pokerMap[26] = "26_15_红桃_2"

	//梅花
	pokerMap[27] = "27_3_梅花_3"
	pokerMap[28] = "28_4_梅花_4"
	pokerMap[29] = "29_5_梅花_5"
	pokerMap[30] = "30_6_梅花_6"
	pokerMap[31] = "31_7_梅花_7"
	pokerMap[32] = "32_8_梅花_8"
	pokerMap[33] = "33_9_梅花_9"
	pokerMap[34] = "34_10_梅花_10"
	pokerMap[35] = "35_11_梅花_J"
	pokerMap[36] = "36_12_梅花_Q"
	pokerMap[37] = "37_13_梅花_K"
	pokerMap[38] = "38_14_梅花_A"
	pokerMap[39] = "39_15_梅花_2"

	//方块
	pokerMap[40] = "40_3_方块_3"
	pokerMap[41] = "41_4_方块_4"
	pokerMap[42] = "42_5_方块_5"
	pokerMap[43] = "43_6_方块_6"
	pokerMap[44] = "44_7_方块_7"
	pokerMap[45] = "45_8_方块_8"
	pokerMap[46] = "46_9_方块_9"
	pokerMap[47] = "47_10_方块_10"
	pokerMap[48] = "48_11_方块_J"
	pokerMap[49] = "49_12_方块_Q"
	pokerMap[50] = "50_13_方块_K"
	pokerMap[51] = "51_14_方块_A"
	pokerMap[52] = "52_15_方块_2"

	//大小王
	pokerMap[53] = "53_16_小_Joker"
	pokerMap[54] = "54_17_大_Joker"
}

//牌的结构
type Poker struct {
	Id     int32  //牌的ID：1~54
	Value  int32  //牌值：牌面的值
	Flower int32  //花色
	Name   string //牌面描述
	Lai    bool   //是否是赖子牌
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

func (p *Poker) IsLaiZi() bool {
	return p.Lai
}

//根据map index解析一张牌
func parsePokerByIndex(index int32) (id int32, value int32, flower int32, name string) {
	if index < 1 || index > ALL_POKER_NUM {
		return -1, -1, -1, ""
	}
	pokerStr := strings.Split(pokerMap[index], "_")

	id = index
	v1, _ := strconv.Atoi(pokerStr[1])
	value = int32(v1)
	name = pokerStr[3]

	switch pokerStr[2] {
	case "黑桃":
		flower = FLOWER_HEI
	case "红桃":
		flower = FLOWER_HONG
	case "梅花":
		flower = FLOWER_MEI
	case "方块":
		flower = FLOWER_FANG
	case "大":
		flower = FLOWER_RED_JOKER
	case "小":
		flower = FLOWER_BLACK_JOKER
	}
	return
}

//根据ID初始化一张牌
func initPaiById(id int32) *Poker {
	poker := &Poker{}
	poker.Id, poker.Value, poker.Flower, poker.Name = parsePokerByIndex(id)
	return poker
}

//根据一组ID初始化一组牌
func initPaiByArray(indexs []int32) (pais []*Poker) {
	for _, id := range indexs {
		pais = append(pais, initPaiById(id))
	}
	return
}

//将牌以字符串形式打印出来，增加易读性
func printPoker(p *Poker) string {
	if p == nil {
		return ""
	}
	var flower string
	switch p.GetFlower() {
	case FLOWER_BLACK_JOKER:
		flower = "小"
	case FLOWER_RED_JOKER:
		flower = "大"
	case FLOWER_HEI:
		flower = "黑桃"
	case FLOWER_HONG:
		flower = "红桃"
	case FLOWER_MEI:
		flower = "梅花"
	case FLOWER_FANG:
		flower = "方块"
	case FLOWER_UNKNOWN:
		flower = ""
	}
	str := fmt.Sprintf("%v%v ", flower, p.GetName())
	return str
}

func printPokers(pais []*Poker) string {
	if pais == nil || len(pais) <= 0 {
		return ""
	}
	var str string
	for _, p := range pais {
		if p == nil {
			continue
		}
		str += printPoker(p)
	}
	return str
}
