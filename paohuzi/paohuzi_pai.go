package paohuzi

/*
    @Create by GoLand
    @Author: hong
    @Time: 2018/10/11 11:42 
    @File: paohuzi_pai.go    
*/

const PHZ_ALLPOKER_NUM int32 = 80 //所有牌的数量

var PokerMap map[int32]string

func init() {
	PokerMap = make(map[int32]string, PHZ_ALLPOKER_NUM)
	PokerMap[0] = ""
	PokerMap[1] = "Black_Y_11"
	PokerMap[2] = "Black_Y_11"
	PokerMap[3] = "Black_Y_11"
	PokerMap[4] = "Black_Y_11"
	PokerMap[5] = "Black_N_1"
	PokerMap[6] = "Black_N_1"
	PokerMap[7] = "Black_N_1"
	PokerMap[8] = "Black_N_1"
	PokerMap[9] = "Red_Y_12"
	PokerMap[10] = "Red_Y_12"
	PokerMap[11] = "Red_Y_12"
	PokerMap[12] = "Red_Y_12"
	PokerMap[13] = "Red_N_2"
	PokerMap[14] = "Red_N_2"
	PokerMap[15] = "Red_N_2"
	PokerMap[16] = "Red_N_2"
	PokerMap[17] = "Black_Y_13"
	PokerMap[18] = "Black_Y_13"
	PokerMap[19] = "Black_Y_13"
	PokerMap[20] = "Black_Y_13"
	PokerMap[21] = "Black_N_3"
	PokerMap[22] = "Black_N_3"
	PokerMap[23] = "Black_N_3"
	PokerMap[24] = "Black_N_3"
	PokerMap[25] = "Black_Y_14"
	PokerMap[26] = "Black_Y_14"
	PokerMap[27] = "Black_Y_14"
	PokerMap[28] = "Black_Y_14"
	PokerMap[29] = "Black_N_4"
	PokerMap[30] = "Black_N_4"
	PokerMap[31] = "Black_N_4"
	PokerMap[32] = "Black_N_4"
	PokerMap[33] = "Black_Y_15"
	PokerMap[34] = "Black_Y_15"
	PokerMap[35] = "Black_Y_15"
	PokerMap[36] = "Black_Y_15"
	PokerMap[37] = "Black_N_5"
	PokerMap[38] = "Black_N_5"
	PokerMap[39] = "Black_N_5"
	PokerMap[40] = "Black_N_5"
	PokerMap[41] = "Black_Y_16"
	PokerMap[42] = "Black_Y_16"
	PokerMap[43] = "Black_Y_16"
	PokerMap[44] = "Black_Y_16"
	PokerMap[45] = "Black_N_6"
	PokerMap[46] = "Black_N_6"
	PokerMap[47] = "Black_N_6"
	PokerMap[48] = "Black_N_6"
	PokerMap[49] = "Red_Y_17"
	PokerMap[50] = "Red_Y_17"
	PokerMap[51] = "Red_Y_17"
	PokerMap[52] = "Red_Y_17"
	PokerMap[53] = "Red_N_7"
	PokerMap[54] = "Red_N_7"
	PokerMap[55] = "Red_N_7"
	PokerMap[56] = "Red_N_7"
	PokerMap[57] = "Black_Y_18"
	PokerMap[58] = "Black_Y_18"
	PokerMap[59] = "Black_Y_18"
	PokerMap[60] = "Black_Y_18"
	PokerMap[61] = "Black_N_8"
	PokerMap[62] = "Black_N_8"
	PokerMap[63] = "Black_N_8"
	PokerMap[64] = "Black_N_8"
	PokerMap[65] = "Black_Y_19"
	PokerMap[66] = "Black_Y_19"
	PokerMap[67] = "Black_Y_19"
	PokerMap[68] = "Black_Y_19"
	PokerMap[69] = "Black_N_9"
	PokerMap[70] = "Black_N_9"
	PokerMap[71] = "Black_N_9"
	PokerMap[72] = "Black_N_9"
	PokerMap[73] = "Red_Y_20"
	PokerMap[74] = "Red_Y_20"
	PokerMap[75] = "Red_Y_20"
	PokerMap[76] = "Red_Y_20"
	PokerMap[77] = "Red_N_10"
	PokerMap[78] = "Red_N_10"
	PokerMap[79] = "Red_N_10"
	PokerMap[80] = "Red_N_10"
}

type PengPai struct {
	PengType  int32 //碰牌的类型：偎牌，碰，臭偎
	Pais      []*PHZPoker
	HuXi      int32
}

type PaoPai struct {
	Pais      []*PHZPoker
	PaoType   int32 //跑牌类型
	HuXi      int32
}

type TiPai struct {
	TiType    int32 //提牌的类型
	Pais      []*PHZPoker
	HuXi      int32
}

type ChiPai struct {
	Pais      []*PHZPoker
	Pai       *PHZPoker
	HuXi      int32
}

type WeiPai struct {
	Pais []*PHZPoker
	Pai  *PHZPoker
	HuXi int32
}

type TiInfo struct {
}

type PHZPoker struct {
	Id      int32  //每张牌唯一的ID
	Value   int32  //牌值
	Flower  FLOWER //花色
	BigWord bool   //true：大字  false：小字
	Des     string //描述
}

type FLOWER int32

var (
	FLOWER_ERROR FLOWER = 0
	FLOWER_R     FLOWER = 1 //红字
	FLOWER_B     FLOWER = 2 //黑字
)

func (p *PHZPoker) GetId() int32 {
	return p.Id
}

func (p *PHZPoker) GetValue() int32 {
	return p.Value
}

func (p *PHZPoker) GetFlower() int32 {
	return int32(p.Flower)
}

func (p *PHZPoker) GetDes() string {
	return p.Des
}

func (p *PHZPoker) IsBig() bool {
	return p.BigWord
}

func (p *PHZPoker) GetPaiIndexByValue() int32 {
	//根据牌的value和花色返回index，此处的index用于统计牌的张数
	return p.GetValue()
}

func CountHandPais(pais []*PHZPoker) []int {
	counts := make([]int, TotalPaiValueCount+1) //牌值1——10:小字  11——20:大字
	if pais == nil || len(pais) == 0 {
		return nil
	}
	for _, p := range pais {
		if p == nil {
			continue
		}
		if p.GetPaiIndexByValue() == 0 {
			return nil
		}
		counts[p.GetPaiIndexByValue()]++
	}
	return counts
}

//根据牌值找到一组牌里几张相同牌值的牌
func GetSameValuePaisByValue(handPokers []*PHZPoker, value int32) []*PHZPoker {
	retPais := []*PHZPoker{}
	for _, poker := range handPokers {
		if poker == nil {
			continue
		}
		if poker.GetValue() == value {
			retPais = append(retPais, poker)
		}
	}
	return retPais
}

//从一组牌中删除指定ID的一张牌
func DelPaiFromPokersByID(handPokers []*PHZPoker, poker *PHZPoker) []*PHZPoker {
	for i, p := range handPokers {
		if p.GetId() == poker.GetId() {
			handPokers = append(handPokers[:i], handPokers[i+1:]...)
			break
		}
	}
	return handPokers
}

//根据牌值在手牌中找相应的poker
func GetPaiByValue(handPokers []*PHZPoker, value int32) *PHZPoker {
	for _, poker := range handPokers {
		if poker == nil {
			continue
		}
		if poker.GetValue() == value {
			return poker
		}
	}
	return nil
}

//删除一张指定牌值的牌
func DelPaiFromPokers(handPokers []*PHZPoker, poker *PHZPoker) []*PHZPoker {
	if poker == nil {
		return handPokers
	}
	for i, p := range handPokers {
		if p == nil {
			continue
		}
		if p.GetValue() == poker.GetValue() {
			handPokers = append(handPokers[:i], handPokers[i+1:]...)
			break
		}
	}
	return handPokers
}