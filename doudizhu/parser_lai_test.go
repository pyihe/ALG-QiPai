package doudizhu

import (
	"fmt"
	"testing"
)

/*
    @Create by GoLand
    @Author: hong
    @Time: 2018/8/16 22:08 
    @File: parser_lai_test.go    
*/

var (
	pokersIndex   = []int32{12, 25, 38, 	11, 24, 37, 	10, 23, 36, 	2,15, 3, 16, 13,26}
	laiZiIndex    = []int32{1}
	pokers, laiZi []*Poker
)

func init() {
	pokers = initPaiByArray(pokersIndex)
	laiZi = initPaiByArray(laiZiIndex)
	for _, p := range laiZi {
		if p != nil {
			p.Lai = true
		}
	}
}

func TestIsAirDaiDan_lai(t *testing.T) {
	pais := append(pokers, laiZi...)
	//fmt.Printf("pais.len = [%v]\n", len(pais))
	result, key := isAirDaiDui_lai(pais)
	fmt.Printf("result = [%v]	key = [%v]\n", result, key)
}
