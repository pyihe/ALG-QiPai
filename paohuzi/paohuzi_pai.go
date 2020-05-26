package paohuzi

const PHZ_ALLPOKER_NUM int32 = 80 //所有牌的数量

var pokerMap map[int32]string

func init() {
	pokerMap = make(map[int32]string, PHZ_ALLPOKER_NUM)
	pokerMap[0] = ""
	pokerMap[1] = "Black_Y_11"
	pokerMap[2] = "Black_Y_11"
	pokerMap[3] = "Black_Y_11"
	pokerMap[4] = "Black_Y_11"
	pokerMap[5] = "Black_N_1"
	pokerMap[6] = "Black_N_1"
	pokerMap[7] = "Black_N_1"
	pokerMap[8] = "Black_N_1"
	pokerMap[9] = "Red_Y_12"
	pokerMap[10] = "Red_Y_12"
	pokerMap[11] = "Red_Y_12"
	pokerMap[12] = "Red_Y_12"
	pokerMap[13] = "Red_N_2"
	pokerMap[14] = "Red_N_2"
	pokerMap[15] = "Red_N_2"
	pokerMap[16] = "Red_N_2"
	pokerMap[17] = "Black_Y_13"
	pokerMap[18] = "Black_Y_13"
	pokerMap[19] = "Black_Y_13"
	pokerMap[20] = "Black_Y_13"
	pokerMap[21] = "Black_N_3"
	pokerMap[22] = "Black_N_3"
	pokerMap[23] = "Black_N_3"
	pokerMap[24] = "Black_N_3"
	pokerMap[25] = "Black_Y_14"
	pokerMap[26] = "Black_Y_14"
	pokerMap[27] = "Black_Y_14"
	pokerMap[28] = "Black_Y_14"
	pokerMap[29] = "Black_N_4"
	pokerMap[30] = "Black_N_4"
	pokerMap[31] = "Black_N_4"
	pokerMap[32] = "Black_N_4"
	pokerMap[33] = "Black_Y_15"
	pokerMap[34] = "Black_Y_15"
	pokerMap[35] = "Black_Y_15"
	pokerMap[36] = "Black_Y_15"
	pokerMap[37] = "Black_N_5"
	pokerMap[38] = "Black_N_5"
	pokerMap[39] = "Black_N_5"
	pokerMap[40] = "Black_N_5"
	pokerMap[41] = "Black_Y_16"
	pokerMap[42] = "Black_Y_16"
	pokerMap[43] = "Black_Y_16"
	pokerMap[44] = "Black_Y_16"
	pokerMap[45] = "Black_N_6"
	pokerMap[46] = "Black_N_6"
	pokerMap[47] = "Black_N_6"
	pokerMap[48] = "Black_N_6"
	pokerMap[49] = "Red_Y_17"
	pokerMap[50] = "Red_Y_17"
	pokerMap[51] = "Red_Y_17"
	pokerMap[52] = "Red_Y_17"
	pokerMap[53] = "Red_N_7"
	pokerMap[54] = "Red_N_7"
	pokerMap[55] = "Red_N_7"
	pokerMap[56] = "Red_N_7"
	pokerMap[57] = "Black_Y_18"
	pokerMap[58] = "Black_Y_18"
	pokerMap[59] = "Black_Y_18"
	pokerMap[60] = "Black_Y_18"
	pokerMap[61] = "Black_N_8"
	pokerMap[62] = "Black_N_8"
	pokerMap[63] = "Black_N_8"
	pokerMap[64] = "Black_N_8"
	pokerMap[65] = "Black_Y_19"
	pokerMap[66] = "Black_Y_19"
	pokerMap[67] = "Black_Y_19"
	pokerMap[68] = "Black_Y_19"
	pokerMap[69] = "Black_N_9"
	pokerMap[70] = "Black_N_9"
	pokerMap[71] = "Black_N_9"
	pokerMap[72] = "Black_N_9"
	pokerMap[73] = "Red_Y_20"
	pokerMap[74] = "Red_Y_20"
	pokerMap[75] = "Red_Y_20"
	pokerMap[76] = "Red_Y_20"
	pokerMap[77] = "Red_N_10"
	pokerMap[78] = "Red_N_10"
	pokerMap[79] = "Red_N_10"
	pokerMap[80] = "Red_N_10"
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

func (p *PHZPoker) getPaiIndexByValue() int32 {
	//根据牌的value和花色返回index，此处的index用于统计牌的张数
	return p.GetValue()
}

func countHandPais(pais []*PHZPoker) []int {
	counts := make([]int, TotalPaiValueCount+1) //牌值1——10:小字  11——20:大字
	if pais == nil || len(pais) == 0 {
		return nil
	}
	for _, p := range pais {
		if p == nil {
			continue
		}
		if p.getPaiIndexByValue() == 0 {
			return nil
		}
		counts[p.getPaiIndexByValue()]++
	}
	return counts
}

//根据牌值找到一组牌里几张相同牌值的牌
func getSameValuePaisByValue(handPokers []*PHZPoker, value int32) []*PHZPoker {
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
func delPaiFromPokersByID(handPokers []*PHZPoker, poker *PHZPoker) []*PHZPoker {
	for i, p := range handPokers {
		if p.GetId() == poker.GetId() {
			handPokers = append(handPokers[:i], handPokers[i+1:]...)
			break
		}
	}
	return handPokers
}

//根据牌值在手牌中找相应的poker
func getPaiByValue(handPokers []*PHZPoker, value int32) *PHZPoker {
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
func delPaiFromPokers(handPokers []*PHZPoker, poker *PHZPoker) []*PHZPoker {
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