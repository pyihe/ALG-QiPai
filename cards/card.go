package cards

import (
	"sort"
)

type CardSuit int32

const (
	CardSuitNone      CardSuit = iota //
	CardSuitCharacter                 // 万
	CardSuitBamboo                    // 条
	CardSuitDot                       // 筒
	CardSuitDragon                    // 中发白
	CardSuitWind                      // 东南西北
	CardSuitSeason                    // 春夏秋冬
	CardSuitFlower                    // 梅兰竹菊
	CardSuitDiamond                   // 方块
	CardSuitClub                      // 梅花
	CardSuitHeart                     // 红桃
	CardSuitSpade                     // 黑桃
	CardSuitJoker                     // 大小王
	_CardSuitEnd
)

type Card int32

const (
	CardNone           Card = 0   //
	CardOneCharacter   Card = 101 //一万
	CardTwoCharacter   Card = 102 // 二万
	CardThreeCharacter Card = 103 // 三万
	CardFourCharacter  Card = 104 // 四万
	CardFiveCharacter  Card = 105 // 五万
	CardSixCharacter   Card = 106 // 六万
	CardSevenCharacter Card = 107 // 七万
	CardEightCharacter Card = 108 // 八万
	CardNineCharacter  Card = 109 // 九万

	CardOneBamboo   Card = 201 // 一条
	CardTwoBamboo   Card = 202 // 二条
	CardThreeBamboo Card = 203 // 三条
	CardFourBamboo  Card = 204 // 四条
	CardFiveBamboo  Card = 205 // 五条
	CardSixBamboo   Card = 206 // 六条
	CardSevenBamboo Card = 207 // 七条
	CardEightBamboo Card = 208 // 八条
	CardNineBamboo  Card = 209 // 九条

	CardOneDot   Card = 301 // 一筒
	CardTwoDot   Card = 302 // 二筒
	CardThreeDot Card = 303 // 三筒
	CardFourDot  Card = 304 // 四筒
	CardFiveDot  Card = 305 // 五筒
	CardSixDot   Card = 306 // 六筒
	CardSevenDot Card = 307 // 七筒
	CardEightDot Card = 308 // 八筒
	CardNineDot  Card = 309 // 九筒

	CardRedDragon   Card = 401 // 红中
	CardGreenDragon Card = 402 // 发财
	CardWhiteDragon Card = 403 // 白板

	CardEastWind  Card = 501 // 东风
	CardSouthWind Card = 502 // 南风
	CardWestWind  Card = 503 // 西风
	CardNorthWind Card = 504 // 北风

	CardSpring Card = 601 // 春
	CardSummer Card = 602 // 夏
	CardAutumn Card = 603 // 秋
	CardWinter Card = 604 // 冬

	CardPlum          Card = 705 // 梅
	CardOrchid        Card = 706 // 兰
	CardBamboo        Card = 707 // 竹
	CardChrysanthemum Card = 708 // 菊

	CardThreeDiamond Card = 803 // 方块3
	CardFourDiamond  Card = 804 // 方块4
	CardFiveDiamond  Card = 805 // 方块5
	CardSixDiamond   Card = 806 // 方块6
	CardSevenDiamond Card = 807 // 方块7
	CardEightDiamond Card = 808 // 方块8
	CardNineDiamond  Card = 809 // 方块9
	CardTenDiamond   Card = 810 // 方块10
	CardJackDiamond  Card = 811 // 方块J
	CardQueenDiamond Card = 812 // 方块Q
	CardKingDiamond  Card = 813 // 方块K
	CardAceDiamond   Card = 814 // 方块A
	CardTwoDiamond   Card = 815 // 方块2

	CardThreeClub Card = 903 // 梅花3
	CardFourClub  Card = 904 // 梅花4
	CardFiveClub  Card = 905 // 梅花5
	CardSixClub   Card = 906 // 梅花6
	CardSevenClub Card = 907 // 梅花7
	CardEightClub Card = 908 // 梅花8
	CardNineClub  Card = 909 // 梅花9
	CardTenClub   Card = 910 // 梅花10
	CardJackClub  Card = 911 // 梅花J
	CardQueenClub Card = 912 // 梅花Q
	CardKingClub  Card = 913 // 梅花K
	CardAceClub   Card = 914 // 梅花A
	CardTwoClub   Card = 915 // 梅花2

	CardThreeHeart Card = 1003 // 红桃3
	CardFourHeart  Card = 1004 // 红桃4
	CardFiveHeart  Card = 1005 // 红桃5
	CardSixHeart   Card = 1006 // 红桃6
	CardSevenHeart Card = 1007 // 红桃7
	CardEightHeart Card = 1008 // 红桃8
	CardNineHeart  Card = 1009 // 红桃9
	CardTenHeart   Card = 1010 // 红桃10
	CardJackHeart  Card = 1011 // 红桃J
	CardQueenHeart Card = 1012 // 红桃Q
	CardKingHeart  Card = 1013 // 红桃K
	CardAceHeart   Card = 1014 // 红桃A
	CardTwoHeart   Card = 1015 // 红桃2

	CardThreeSpade Card = 1103 // 黑桃3
	CardFourSpade  Card = 1104 // 黑桃4
	CardFiveSpade  Card = 1105 // 黑桃5
	CardSixSpade   Card = 1106 // 黑桃6
	CardSevenSpade Card = 1107 // 黑桃7
	CardEightSpade Card = 1108 // 黑桃8
	CardNineSpade  Card = 1109 // 黑桃9
	CardTenSpade   Card = 1110 // 黑桃10
	CardJackSpade  Card = 1111 // 黑桃J
	CardQueenSpade Card = 1112 // 黑桃Q
	CardKingSpade  Card = 1113 // 黑桃K
	CardAceSpade   Card = 1114 // 黑桃A
	CardTwoSpade   Card = 1115 // 黑桃2

	CardBlackJoker Card = 1216 // 小王
	CardRedJoker   Card = 1217 // 大王
)

// Suit 获取牌花色
func (c Card) Suit() CardSuit {
	return CardSuit(c / 100)
}

// Value 获取牌值
func (c Card) Value() int32 {
	return int32(c % 100)
}

// NewCard 根据花色和牌值获取一张牌
func NewCard(suit CardSuit, value int32) Card {
	return Card(int32(suit*100) + value)
}

// SortCards 给一组牌排序
func SortCards(cards []Card) {
	sort.Sort(cardList(cards))
}

// CopyCards 复制一组牌
func CopyCards(cards []Card) (copyCards []Card) {
	if len(cards) == 0 {
		return
	}
	copyCards = make([]Card, len(cards))
	copy(copyCards, cards)
	return
}

type cardList []Card

func (cl cardList) Len() int {
	return len(cl)
}

func (cl cardList) Swap(i, j int) {
	cl[i], cl[j] = cl[j], cl[i]
}

func (cl cardList) Less(i, j int) bool {
	return cl[i] < cl[j]
}
