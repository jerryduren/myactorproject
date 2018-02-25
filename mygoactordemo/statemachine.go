package main

import "github.com/hashicorp/consul/agent/consul/fsm"

/* 定义状态机的事件，也即状态迁移的消息 */
type M1 struct {

}

type M2 struct {

}

type M5 struct {

}

type Tea struct {

}

type Coffee struct {

}

type CalcChange struct {

}

type Cook struct {

}

/* 定义状态机的两个维度：金钱和商品 */
const (
	Pay0 fsm.FSM
)

func main() {
}
