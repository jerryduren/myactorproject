package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/jerryduren/myactorproject/myinclude"
	"strings"
	"time"
)

type SetBecomeAndUnbecomeAct struct {
}

func (this *SetBecomeAndUnbecomeAct) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case myinclude.HelloMessage:
		fmt.Println(msg.Who, "SAY:", msg.Say)

		if strings.Contains(msg.Say, "Jerry") {
			fmt.Println("下面的消息由Jerry处理...")
			context.Become(this.JerryDo)
		} else if strings.Contains(msg.Say, "Neil") {
			fmt.Println("下面的消息由Neil处理...")
			context.Become(this.NeilDo)
		}
	}
}

func (this *SetBecomeAndUnbecomeAct) JerryDo(context actor.Context) {
	switch msg := context.Message().(type) {
	case myinclude.HelloMessage:
		fmt.Println(msg.Who, "SAY:", msg.Say)
		if strings.Contains(msg.Say, "Neil") {
			fmt.Println("下面的消息交由Neil处理...")
			context.Become(this.NeilDo)
		} else if !strings.Contains(msg.Say, "Jerry") {
			fmt.Println("下面的消息交由Default Handler处理...")
			context.Become(this.Receive)
		}
	}
}

func (this *SetBecomeAndUnbecomeAct) NeilDo(context actor.Context) {
	switch msg := context.Message().(type) {
	case myinclude.HelloMessage:
		fmt.Println(msg.Who, "SAY:", msg.Say)
		if strings.Contains(msg.Say, "Jerry") {
			fmt.Println("下面的消息交由Jerry处理...")
			context.Become(this.JerryDo)
		} else if !strings.Contains(msg.Say, "Neil") {
			fmt.Println("下面的消息交由Default Handler处理...")
			context.Become(this.Receive)
		}
	}
}

/* 演示接龙比赛，如消息的内容里面包含了Jerry，接下来由Jerry处理，如果包含了Neil就由Neil处理，否则交由Default Handler处理*/
/* actor.Context.Become(func ()(context actor.Context) 设置下一条消息的处理函数，函数定于同Receive*/
func main() {
	pid := actor.Spawn(actor.FromInstance(&SetBecomeAndUnbecomeAct{}))
	fmt.Println("下面的消息由缺省handler处理...")
	pid.Tell(myinclude.HelloMessage{Say: "Hello, Jerry"})
	pid.Tell(myinclude.HelloMessage{Who: "Jerry Du", Say: "Hello, Neil"})
	pid.Tell(myinclude.HelloMessage{Who: "Neil Ren", Say: "Hello, Jerry"})
	pid.Tell(myinclude.HelloMessage{Say: "What do I say?"})
	pid.Tell(myinclude.HelloMessage{Say: "you are Neil"})
	pid.Tell(myinclude.HelloMessage{Say: "I amn't Neil"})
	pid.Tell(myinclude.HelloMessage{Say: "you are Jerry"})

	time.Sleep(time.Second)
}
