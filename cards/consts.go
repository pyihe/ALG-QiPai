package cards

type ZWFX uint8

const (
	ZWFXDong ZWFX = iota + 1
	ZWFXNan
	ZWFXXi
	ZWFXBei
)

// GroupType 胡牌时的胡牌组合类型
type GroupType uint8

const (
	GroupTypeNone                           GroupType = iota //
	GroupTypeMJJiang                                         // 将牌
	GroupTypeMJShun                                          // 顺子
	GroupTypeMJKe                                            // 刻子
	GroupTypeDDZSingle                                       // 单牌
	GroupTypeDDZPair                                         // 对子
	GroupTypeDDZTriplet                                      // 三不带
	GroupTypeDDZTripletWithSingle                            // 三带一
	GroupTypeDDZTripletWithPair                              // 三带对
	GroupTypeDDZSequence                                     // 顺子
	GroupTypeDDZSequenceOfPair                               // 连对
	GroupTypeDDZSequenceOfTriplet                            // 飞机不带牌
	GroupTypeDDZSequenceOfTripletWithSingle                  // 飞机带单牌
	GroupTypeDDZSequenceOfTripletWithPair                    // 飞机带对子
	GroupTypeDDZQuadplexSetWithSingle                        // 四带二(单牌)
	GroupTypeDDZQuadplexSetWithPair                          // 四带二(对子)
	GroupTypeDDZBomb                                         // 普通炸弹
	GroupTypeDDZRocket                                       // 火箭(王炸)
)
