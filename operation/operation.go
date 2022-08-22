package main

import (
	"fmt"
	"time"

	"github.com/pyihe/go-selector"
)

func main() {
	event := selector.NewEvent(func() {
		fmt.Println("没有玩家选择操作, 继续摸牌...")
	})
	chi, _ := event.AddActionWithHandler(1, selector.Handler{
		selector.StateAgree: func() {
			fmt.Println("吃牌...")
		},
		selector.StateRefuse: func() {
			fmt.Println("拒绝吃牌...")
		},
		selector.StateTimeout: func() {
			fmt.Println("吃牌超时...")
		},
	})

	peng, _ := event.AddAction(2)
	peng.AddHandler(selector.StateAgree, func() {
		fmt.Println("碰牌...")
	})
	peng.AddHandler(selector.StateRefuse, func() {
		fmt.Println("拒绝碰牌...")
	})
	peng.AddHandler(selector.StateTimeout, func() {
		fmt.Println("碰牌超时...")
	})

	hu, _ := event.AddActionWithHandler(3, selector.Handler{
		selector.StateAgree: func() {
			fmt.Println("胡牌...")
		},
		selector.StateRefuse: func() {
			fmt.Println("拒绝胡牌...")
		},
		selector.StateTimeout: func() {
			fmt.Println("胡牌超时...")
		},
	})

	event.Start(5 * time.Second)

	chi.Agree()
	fmt.Println(chi.Deadline())
	peng.Refuse()
	hu.Refuse()

	time.Sleep(10 * time.Second)
}
