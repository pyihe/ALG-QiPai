package paohuzi

import (
	"fmt"
)

/*
    @Create by GoLand
    @Author: hong
    @Time: 2018/10/11 11:52 
    @File: paohuzi_hu.go    
*/

//牌局过程中的数据
type GameInfo struct {
	HandPokers []*PHZPoker //手牌信息

	PengPais []*PengPai //碰牌信息
	ChiPais  []*ChiPai  //吃牌信息
	PaoPais  []*PaoPai  //跑牌信息
	WeiPais  []*WeiPai  //偎牌信息
	TiPais   []*TiPai   //提牌信息

	RoundHuXi int32 //牌局过程中吃、碰、杠、提产生的胡息总和，计算是否可以胡牌时用

	//由入口判断是否是点炮和跑胡
	IsDianPao bool
	IsPaoHu   bool
}

//胡牌时的胡牌信息
type HuInfo struct {
	CanHu       bool        //能否胡牌
	YiJuHua     []*YJH      //胡牌时手牌里的一句话
	YiJiaoPai   []*YJH      //胡牌时手里的一缴牌
	DuiZis      []*DuiZi    //胡牌时手里的对子
	KanPais     []*YiKanPai //胡牌时手里的坎牌
	TiPais      []*YiTi     //胡牌时手里的提牌
	Pengs       []*YiKanPai //胡牌时胡出来的碰牌
	HuXi        int32       //胡息数
	WinUser     uint32      //赢家
	LoseUser    []uint32    //输家
	IsZimo      bool        //是否自摸
	IsDianPao   bool        //是否点炮
	DianPaoUser uint32      //点炮的玩家
	Fan         int32       //胡牌时的番数
	WinScore    int64       //得分
	HuPai       *PHZPoker   //胡的哪张牌
	HuType      []int32     //胡牌类型
}

//缴牌value
type jiao []int

//一句话或者一缴牌
type YJH struct {
	Pais []*PHZPoker
	HuXi int32
}

//对子、将牌
type DuiZi struct {
	Pais []*PHZPoker
}

//一砍牌
type YiKanPai struct {
	Pais []*PHZPoker
	HuXi int32
}

//一提牌
type YiTi struct {
	Pais []*PHZPoker
	HuXi int32
}

//是否可以胡牌的信息
type CanHuInfo struct {
	canHu             bool
	countBigErQiShi   int    //大字二七十的数量
	countSmallErQiShi int    //小字二七十的数量
	countBigYiErSan   int    //大字一二三的数量
	countSmallYiErSan int    //小字一二三的数量
	jiangs            []int  //将牌 如果有
	jiaos             []jiao //绞牌
	yijuhuas          []int  //一句话
	kans              []int  //一坎牌
	pengs             []int  //碰牌
	totalHuXi         int32  //总胡息
}

const (
	TotalPaiValueCount = 20 //牌张
	PaiValueSmall      = 10 //小字
	CanHuHuXiLimit     = 15 //满足胡牌条件的胡息数
)

//递归判断是否可以胡牌
func CanHu(huxi int32, count []int, length int, zimoPaiValue int32, jiangs []int) *CanHuInfo {
	info := &CanHuInfo{}

	info.totalHuXi = huxi
	info.jiangs = jiangs

	if len(jiangs) > 1 {
		info.canHu = false
		return info
	}

	//递归完所有的牌
	if length == 0 {
		if info.totalHuXi >= CanHuHuXiLimit && len(jiangs) <= 1 {
			//胡息满足条件，并且将牌数最多为1表示胡了
			info.canHu = true
			return info
		}
		info.canHu = false
		return info
	}

	/*这里注意胡息的优先级:
		先找大胡息数的组合
	*/

	//大字坎 6胡息
	for i := PaiValueSmall + 1; i <= TotalPaiValueCount; i++ {
		if int(zimoPaiValue) == i && count[i] >= 3 {

			count[i] -= 3
			info.totalHuXi += 6

			info = CanHu(info.totalHuXi, count, length-3, zimoPaiValue, info.jiangs)
			if info.canHu {
				info.kans = append(info.kans, i)
				return info
			}

			info.totalHuXi -= 6
			count[i] += 3
		}
	}

	//大字二七十 6胡息
	if count[12] > 0 && count[17] > 0 && count[20] > 0 {
		count[12] -= 1
		count[17] -= 1
		count[20] -= 1

		info.totalHuXi += 6 //大字二七十

		info = CanHu(info.totalHuXi, count, length-3, zimoPaiValue, info.jiangs)
		if info.canHu {
			info.countBigErQiShi++
			return info
		}

		info.totalHuXi -= 6 //大字二七十

		count[12] += 1
		count[17] += 1
		count[20] += 1
	}
	//大字一二三 6胡息
	if count[11] > 0 && count[12] > 0 && count[13] > 0 {
		count[11] -= 1
		count[12] -= 1
		count[13] -= 1

		info.totalHuXi += 6 //大字一二三

		info = CanHu(info.totalHuXi, count, length-3, zimoPaiValue, info.jiangs)
		if info.canHu {
			info.countBigYiErSan++
			return info
		}

		info.totalHuXi -= 6 //大字一二三

		count[11] += 1
		count[12] += 1
		count[13] += 1
	}

	//小字坎 3胡息
	for i := 1; i <= PaiValueSmall; i++ {
		if int(zimoPaiValue) == i && count[i] >= 3 {

			count[i] -= 3
			info.totalHuXi += 3

			info = CanHu(info.totalHuXi, count, length-3, zimoPaiValue, info.jiangs)
			if info.canHu {
				info.kans = append(info.kans, i)
				return info
			}

			info.totalHuXi -= 3
			count[i] += 3
		}
	}

	//大字碰 3胡息
	for i := PaiValueSmall + 1; i <= TotalPaiValueCount; i++ {
		if int(zimoPaiValue) != i && count[i] >= 3 {

			count[i] -= 3
			info.totalHuXi += 3

			info = CanHu(info.totalHuXi, count, length-3, zimoPaiValue, info.jiangs)
			if info.canHu {
				info.pengs = append(info.pengs, i)
				return info
			}

			info.totalHuXi -= 3
			count[i] += 3
		}
	}

	//小字二七十 3胡息
	if count[2] > 0 && count[7] > 0 && count[10] > 0 {
		count[2] -= 1
		count[7] -= 1
		count[10] -= 1

		info.totalHuXi += 3 //小字二七十

		info = CanHu(info.totalHuXi, count, length-3, zimoPaiValue, info.jiangs)
		if info.canHu {
			info.countSmallErQiShi++
			return info
		}

		info.totalHuXi -= 3 //小字二七十

		count[2] += 1
		count[7] += 1
		count[10] += 1
	}
	//小字一二三 3胡息
	if count[1] > 0 && count[2] > 0 && count[3] > 0 {
		count[1] -= 1
		count[2] -= 1
		count[3] -= 1

		info.totalHuXi += 3 //小字一二三

		info = CanHu(info.totalHuXi, count, length-3, zimoPaiValue, info.jiangs)
		if info.canHu {
			info.countSmallYiErSan++
			return info
		}

		info.totalHuXi -= 3 //小字一二三

		count[1] += 1
		count[2] += 1
		count[3] += 1
	}

	//小字碰 1胡息
	for i := 1; i <= PaiValueSmall; i++ {
		if int(zimoPaiValue) != i && count[i] >= 3 {

			count[i] -= 3
			info.totalHuXi += 1

			info = CanHu(info.totalHuXi, count, length-3, zimoPaiValue, info.jiangs)
			if info.canHu {
				info.pengs = append(info.pengs, i)
				return info
			}

			info.totalHuXi -= 1
			count[i] += 3
		}
	}

	//是否是一句话（顺），这里应该分开判断
	//小字 一句话
	for i := 1; i < PaiValueSmall-1; i++ {
		if count[i] > 0 && count[i+1] > 0 && count[i+2] > 0 {
			count[i] -= 1
			count[i+1] -= 1
			count[i+2] -= 1
			info = CanHu(huxi, count, length-3, zimoPaiValue, info.jiangs)
			if info.canHu {
				info.yijuhuas = append(info.yijuhuas, i)
				return info
			}
			count[i] += 1
			count[i+1] += 1
			count[i+2] += 1
		}
	}
	//大字 一句话
	for i := PaiValueSmall + 1; i < TotalPaiValueCount-1; i++ {
		if count[i] > 0 && count[i+1] > 0 && count[i+2] > 0 {
			count[i] -= 1
			count[i+1] -= 1
			count[i+2] -= 1
			info = CanHu(huxi, count, length-3, zimoPaiValue, info.jiangs)
			if info.canHu {
				info.yijuhuas = append(info.yijuhuas, i)
				return info
			}
			count[i] += 1
			count[i+1] += 1
			count[i+2] += 1
		}
	}

	//是否是一绞牌 根据小字去组合大字
	for i := 1; i <= PaiValueSmall; i++ {
		if count[i] > 0 && count[PaiValueSmall+i] >= 2 {
			count[i] -= 1
			count[PaiValueSmall+i] -= 2
			info = CanHu(huxi, count, length-3, zimoPaiValue, info.jiangs)
			if info.canHu {
				jiao := jiao{i, PaiValueSmall + i, PaiValueSmall + i}
				info.jiaos = append(info.jiaos, jiao)
				return info
			}
			count[i] += 1
			count[PaiValueSmall+i] += 2

		} else if count[i] >= 2 && count[PaiValueSmall+i] > 0 {
			count[i] -= 2
			count[PaiValueSmall+i] -= 1
			info = CanHu(huxi, count, length-3, zimoPaiValue, info.jiangs)
			if info.canHu {
				jiao := jiao{i, i, PaiValueSmall + i}
				info.jiaos = append(info.jiaos, jiao)
				return info
			}
			count[i] += 2
			count[PaiValueSmall+i] += 1
		}
	}

	//找对牌
	for i := 1; i <= TotalPaiValueCount; i++ {
		if count[i] >= 2 {
			count[i] -= 2

			info.jiangs = append(info.jiangs, i)

			info = CanHu(huxi, count, length-2, zimoPaiValue, info.jiangs)
			if info.canHu {
				return info
			}

			//不能胡 移除该对子
			newJiangs := []int{}
			continues := 0
			for _, j := range info.jiangs {
				if j == i && continues <= 0 {
					//只过滤一个对子
					continues++
					continue
				}
				newJiangs = append(newJiangs, j)
			}
			info.jiangs = newJiangs

			count[i] += 2
		}
	}
	info.canHu = false
	return info
}

//尝试胡牌
func TryHu(userGameData *GameInfo, checkPai *PHZPoker) (*HuInfo, error) {
	pais := userGameData.HandPokers

	huInfo := &HuInfo{}
	isPaoHu := userGameData.IsPaoHu     //是否是跑胡
	isDianPao := userGameData.IsDianPao //是否是点炮

	//找出牌组里原始坎牌
	var srcKan []*YiKanPai
	totalHuXi := int32(userGameData.RoundHuXi)
	if len(pais) >= 3 {
		count := CountHandPais(pais)
		for paiValue, paiCount := range count {
			//找手牌里既有的坎牌
			if paiCount == 3 {
				kan := &YiKanPai{}
				pokers := GetSameValuePaisByValue(pais, int32(paiValue))
				kan.Pais = append(kan.Pais, pokers...)
				if isPaoHu && checkPai != nil && int32(paiValue) == checkPai.GetValue() {
					//TODO 如果是跑胡，且checkPai和坎牌一样，则不能加胡息
				} else {
					if paiValue <= 10 {
						kan.HuXi = 3
					} else {
						kan.HuXi = 6
					}
				}

				//累加总胡息
				totalHuXi += kan.HuXi
				for _, p := range pokers {
					pais = DelPaiFromPokersByID(pais, p)
				}
				srcKan = append(srcKan, kan)
			}
		}
	}
	//组合checkPai
	checkPokers := pais

	zimoPaiValue := int32(-1)

	if checkPai != nil {
		checkPokers = append(pais, checkPai)
		if !isDianPao {
			//自摸的牌值
			zimoPaiValue = checkPai.GetValue()
		}
	}
	counts := CountHandPais(checkPokers)
	for paiValue, paiCount := range counts {
		//找原始的提牌
		if paiCount == 4 {
			pokers := GetSameValuePaisByValue(pais, int32(paiValue))
			for _, p := range pokers {
				pais = DelPaiFromPokersByID(pais, p)
			}
		}
	}


	huInfo.KanPais = append(huInfo.KanPais, srcKan...)
	//如果找完原始的坎牌，手里没有牌了，则牌型上是可以胡了
	if (checkPokers == nil || len(checkPokers) == 0) && totalHuXi >= CanHuHuXiLimit {
		huInfo.CanHu = true
		//fmt.Println(fmt.Sprintf("check玩家是否可以胡时，玩家手里只有坎牌，可以胡牌..."))
		return huInfo, nil
	}

	canHuInfo := CanHu(totalHuXi, CountHandPais(checkPokers), len(checkPokers), zimoPaiValue, nil)

	fmt.Println(fmt.Sprintf("CanHu的CheckPokers结果是:[%+v]", canHuInfo))
	huInfo.CanHu = canHuInfo.canHu
	huInfo.HuXi = canHuInfo.totalHuXi

	//可以胡牌时，将手牌组合成对应的胡牌组合
	if huInfo.CanHu {
		//将牌
		if len(canHuInfo.jiangs) > 0 {
			jiangPai1 := GetPaiByValue(checkPokers, int32(canHuInfo.jiangs[0]))
			if jiangPai1 != nil {
				checkPokers = DelPaiFromPokers(checkPokers, jiangPai1)
			}

			jiangPai2 := GetPaiByValue(checkPokers, int32(canHuInfo.jiangs[0]))
			if jiangPai2 != nil {
				checkPokers = DelPaiFromPokers(checkPokers, jiangPai2)
			}

			jiangPais := []*PHZPoker{jiangPai1, jiangPai2}
			huInfo.DuiZis = append(huInfo.DuiZis, &DuiZi{Pais: jiangPais})
		}

		//一坎牌
		for _, kan := range canHuInfo.kans {
			kanPais := GetSameValuePaisByValue(checkPokers, int32(kan))
			huInfo.KanPais = append(huInfo.KanPais, &YiKanPai{Pais: kanPais})
			for _, delPai := range kanPais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
		}

		//一碰牌
		for _, kan := range canHuInfo.pengs {
			kanPais := GetSameValuePaisByValue(checkPokers, int32(kan))
			huInfo.Pengs = append(huInfo.Pengs, &YiKanPai{Pais: kanPais})
			for _, delPai := range kanPais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
		}

		//一绞牌
		for _, jiao := range canHuInfo.jiaos {
			yjh := &YJH{}
			for _, jiaoPaiValue := range jiao {
				if pai := GetPaiByValue(checkPokers, int32(jiaoPaiValue)); pai != nil {
					yjh.Pais = append(yjh.Pais, pai)
				}
			}
			huInfo.YiJiaoPai = append(huInfo.YiJiaoPai, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
		}

		//一句话
		for _, yijuhua := range canHuInfo.yijuhuas {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(yijuhua)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(yijuhua+1)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(yijuhua+2)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
		}
		//大字一二三
		for i := canHuInfo.countBigYiErSan; i > 0; i-- {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(11)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(12)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(13)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
		}

		//大字二七十
		for i := canHuInfo.countBigErQiShi; i > 0; i-- {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(12)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(17)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(20)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
		}

		//小字一二三
		for i := canHuInfo.countSmallYiErSan; i > 0; i-- {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(1)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的小字一:[%v]", Cards2String(yjh.Pais)))
			if pai := GetPaiByValue(checkPokers, int32(2)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的小字二:[%v]", Cards2String(yjh.Pais)))
			if pai := GetPaiByValue(checkPokers, int32(3)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的小字三:[%v]", Cards2String(yjh.Pais)))

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
		}

		//小字二七十
		for i := canHuInfo.countSmallErQiShi; i > 0; i-- {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(2)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(7)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(10)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
		}

	}

	return huInfo, nil
}
