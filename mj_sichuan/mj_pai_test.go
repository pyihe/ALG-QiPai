package mj_sichuan

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestXiPai(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			paiIndexs := xiPai(1, 36)
			//随机选出14张牌
			index := get14Pais(paiIndexs)
			pais, err := initPaisByArray(index)
			if err != nil {
				//fmt.Println(fmt.Sprintf("初始化牌失败: [%v] [%v]", err, index))
				t.Errorf("初始化牌失败: [%v] [%v]", err, index)
				return
			}
			//fmt.Println(printMjPais(pais))
			canHu, is19, jiang := canHu(pais)
			/*fmt.Println(fmt.Sprintf("canHu = [%v]  is19 = [%v]  jiang = [%v]", canHu, is19, jiang))
			fmt.Println()
			fmt.Println()*/
			if canHu {
				//t.Logf("paiIds = %v", index)
				fmt.Println(index)
				fmt.Println(printMjPais(pais))
				fmt.Println(fmt.Sprintf("canHu = [%v]  is19 = [%v]  jiang = [%v]", canHu, is19, jiang))
				fmt.Println()
				fmt.Println()
			}
		}
	}
}

func get14Pais(indexs []uint8) []uint8 {
	var result []uint8
	for i := 0; i < 14; i++ {
		rand.Seed(time.Now().UnixNano())
		v := rand.Intn(len(indexs))
		result = append(result, indexs[v])
		indexs = append(indexs[:v], indexs[v+1:]...)
	}
	return result
}

func BenchmarkXiPai(b *testing.B) {
	b.StopTimer()
	pais, _ := initPaisByArray([]uint8{21, 10, 4, 24, 12, 22, 28, 19, 2, 11, 3, 36, 27, 31})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		canHu(pais)
		/*canHu, is19, jiang := canHu(pais)
		fmt.Println(fmt.Sprintf("canHu = [%v]  is19 = [%v]  jiang = [%v]", canHu, is19, jiang))*/
	}
}
