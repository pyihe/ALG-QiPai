package mj_sichuan

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
	定义牌的结构
*/

const (
	_MjPaiCount = 108

	_MjFlowerTong = iota
	_MjFlowerTiao
	_MjFlowerWan
)

var (
	paiMap = make(map[uint8]string, _MjPaiCount)
)

//初始化牌结构
func init() {
	//1筒
	paiMap[1] = "1_T_1"
	paiMap[2] = "2_T_1"
	paiMap[3] = "3_T_1"
	paiMap[4] = "4_T_1"

	//2筒
	paiMap[5] = "5_T_2"
	paiMap[6] = "6_T_2"
	paiMap[7] = "7_T_2"
	paiMap[8] = "8_T_2"

	//3筒
	paiMap[9] = "9_T_3"
	paiMap[10] = "10_T_3"
	paiMap[11] = "11_T_3"
	paiMap[12] = "12_T_3"

	//4筒
	paiMap[13] = "13_T_4"
	paiMap[14] = "14_T_4"
	paiMap[15] = "15_T_4"
	paiMap[16] = "16_T_4"

	//5筒
	paiMap[17] = "17_T_5"
	paiMap[18] = "18_T_5"
	paiMap[19] = "19_T_5"
	paiMap[20] = "20_T_5"

	//6筒
	paiMap[21] = "21_T_6"
	paiMap[22] = "22_T_6"
	paiMap[23] = "23_T_6"
	paiMap[24] = "24_T_6"

	//7筒
	paiMap[25] = "25_T_7"
	paiMap[26] = "26_T_7"
	paiMap[27] = "27_T_7"
	paiMap[28] = "28_T_7"

	//8筒
	paiMap[29] = "29_T_8"
	paiMap[30] = "30_T_8"
	paiMap[31] = "31_T_8"
	paiMap[32] = "32_T_8"

	//9筒
	paiMap[33] = "33_T_9"
	paiMap[34] = "34_T_9"
	paiMap[35] = "35_T_9"
	paiMap[36] = "36_T_9"

	//1条
	paiMap[37] = "37_S_1"
	paiMap[38] = "38_S_1"
	paiMap[39] = "39_S_1"
	paiMap[40] = "40_S_1"

	//2条
	paiMap[41] = "41_S_2"
	paiMap[42] = "42_S_2"
	paiMap[43] = "43_S_2"
	paiMap[44] = "44_S_2"

	//3条
	paiMap[45] = "45_S_3"
	paiMap[46] = "46_S_3"
	paiMap[47] = "47_S_3"
	paiMap[48] = "48_S_3"

	//4条
	paiMap[49] = "49_S_4"
	paiMap[50] = "50_S_4"
	paiMap[51] = "51_S_4"
	paiMap[52] = "52_S_4"

	//5条
	paiMap[53] = "53_S_5"
	paiMap[54] = "54_S_5"
	paiMap[55] = "55_S_5"
	paiMap[56] = "56_S_5"

	//6条
	paiMap[57] = "57_S_6"
	paiMap[58] = "58_S_6"
	paiMap[59] = "59_S_6"
	paiMap[60] = "60_S_6"

	//7条
	paiMap[61] = "61_S_7"
	paiMap[62] = "62_S_7"
	paiMap[63] = "63_S_7"
	paiMap[64] = "64_S_7"

	//8条
	paiMap[65] = "65_S_8"
	paiMap[66] = "66_S_8"
	paiMap[67] = "67_S_8"
	paiMap[68] = "68_S_8"

	//9条
	paiMap[69] = "69_S_9"
	paiMap[70] = "70_S_9"
	paiMap[71] = "71_S_9"
	paiMap[72] = "72_S_9"

	//1万
	paiMap[73] = "73_W_1"
	paiMap[74] = "74_W_1"
	paiMap[75] = "75_W_1"
	paiMap[76] = "76_W_1"

	//2万
	paiMap[77] = "77_W_2"
	paiMap[78] = "78_W_2"
	paiMap[79] = "79_W_2"
	paiMap[80] = "80_W_2"

	//3万
	paiMap[81] = "81_W_3"
	paiMap[82] = "82_W_3"
	paiMap[83] = "83_W_3"
	paiMap[84] = "84_W_3"

	//4万
	paiMap[85] = "85_W_4"
	paiMap[86] = "86_W_4"
	paiMap[87] = "87_W_4"
	paiMap[88] = "88_W_4"

	//5万
	paiMap[89] = "89_W_5"
	paiMap[90] = "90_W_5"
	paiMap[91] = "91_W_5"
	paiMap[92] = "92_W_5"

	//6万
	paiMap[93] = "93_W_6"
	paiMap[94] = "94_W_6"
	paiMap[95] = "95_W_6"
	paiMap[96] = "96_W_6"

	//7万
	paiMap[97] = "97_W_7"
	paiMap[98] = "98_W_7"
	paiMap[99] = "99_W_7"
	paiMap[100] = "100_W_7"

	//8万
	paiMap[101] = "101_W_8"
	paiMap[102] = "102_W_8"
	paiMap[103] = "103_W_8"
	paiMap[104] = "104_W_8"

	//9万
	paiMap[105] = "105_W_9"
	paiMap[106] = "106_W_9"
	paiMap[107] = "107_W_9"
	paiMap[108] = "108_W_9"
}

type MjPai struct {
	id     uint8 //牌ID
	flower uint8 //花色: 筒条万
	value  uint8 //牌值: 1-9
}

func (m *MjPai) GetId() uint8 {
	return m.id
}

func (m *MjPai) GetFlower() uint8 {
	return m.flower
}

func (m *MjPai) GetValue() uint8 {
	return m.value
}

//	根据牌值和花色得出每张牌的唯一标识
//	0-8   ——> 1-9筒
//	9-18  ——> 1-9条
//	19-26 ——> 1-9万
func (m *MjPai) GetUniqueIndexByValue() uint8 {
	return m.value + (m.flower-1)*9 - 1
}

func initPaiByIndex(index uint8) (*MjPai, error) {
	paiDes := paiMap[index]
	paiArray := strings.Split(paiDes, "_")

	value, err := strconv.ParseInt(paiArray[2], 10, 8)
	if err != nil {
		fmt.Printf("parse pai value fail: [%v] \n", err)
		return nil, err
	}
	pai := &MjPai{
		id:    index,
		value: uint8(value),
	}

	switch paiArray[1] {
	case "T":
		pai.flower = _MjFlowerTong
	case "S":
		pai.flower = _MjFlowerTiao
	case "W":
		pai.flower = _MjFlowerWan
	default:
		fmt.Printf("unknown flower des: [%v]", paiArray[1])
		return nil, fmt.Errorf("unknown flower des: [%v]", paiArray[1])
	}
	return pai, nil
}

func initPaisByArray(array []uint8) ([]*MjPai, error) {
	if len(array) <= 0 || len(array) > 108 {
		fmt.Printf("invalid array length: [%v]", len(array))
		return nil, fmt.Errorf("invalid index length")
	}
	var pais []*MjPai
	for _, v := range array {
		p, err := initPaiByIndex(v)
		if err != nil {
			return nil, err
		}
		pais = append(pais, p)
	}
	return pais, nil
}

func printMjPais(pais []*MjPai) string {
	desc := ""
	for _, p := range pais {
		if p == nil {
			continue
		}
		flower := ""
		switch p.flower {
		case _MjFlowerTiao:
			flower = "条"
		case _MjFlowerTong:
			flower = "筒"
		case _MjFlowerWan:
			flower = "万"
		}
		pStr := fmt.Sprintf("%d%s ", p.value, flower)
		desc += pStr
	}
	return desc
}

//根据牌的张数洗牌
func xiPai(startIndex int, paiCount int) []uint8 {
	result := make([]uint8, paiCount)

	//生成原生索引slice
	original := make([]uint8, paiCount)
	for i := 0; i < paiCount; i++ {
		original[i] = uint8(i + startIndex)
	}

	rand.Seed(time.Now().UnixNano())
	for i := range result {
		index := rand.Intn(len(original))
		result[i] = original[index]
		original = append(original[:index], original[index+1:]...)
	}
	return result
}

//增加shuffle洗牌算法
func shuffleXiPai(paiCount int) []int {
	var result = make([]int, paiCount)
	for i := 0; i < paiCount; i++ {
		result[i] = i + 1
	}
	for i := paiCount - 1; i > 0; i-- {
		randIndex := rand.Intn(i)
		result[i], result[randIndex] = result[randIndex], result[i]
	}
	return result
}
