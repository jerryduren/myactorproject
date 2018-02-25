package main

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/jerryduren/myactorproject/myinclude"
	"strings"
)

type ParentActor struct {
}

func (this *ParentActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case myinclude.HelloMessage:
		fmt.Println("Parent receives a message: ", msg.Say)
		/* Parent actor收到一个hello message消息后创建一个子actor，并且把消息再发给child */
		context.Spawn(actor.FromInstance(&ChildActor{})).Tell(msg)
	case *actor.Stopped:
		fmt.Println(context.Self(), "is stopped!")
	case *actor.Restart:
		fmt.Println(context.Self(), "is restarted!")
	}
}

type ChildActor struct {
}

func (this *ChildActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case myinclude.HelloMessage:
		fmt.Println("Child reeives a message: ", msg.Say)
		if strings.Contains(msg.Who, "Jerry") {
			panic("Child actor generates a string panic")
		} else if strings.Contains(msg.Who, "Neil") {
			panic(10)
		}
	case *actor.Stopped:
		fmt.Println(context.Self(), "is stopped!")
	case *actor.Restarting:
		fmt.Println(context.Self(), "is restarting!")
	}
}

func main() {
	decider := func(reason interface{}) actor.Directive {
		fmt.Printf("%v\n", reason)
		fmt.Println("Handling failure for child...")
		switch reason.(type) {
		case string:
			return actor.RestartDirective
		case int:
			return actor.ResumeDirective
		}

		return actor.EscalateDirective
	}
	/* 为孩子ACTOR设置监控，一旦出错就restarting the actor */
	pid := actor.Spawn(actor.FromInstance(&ParentActor{}).WithSupervisor(actor.NewOneForOneStrategy(10, 1000, decider)))
	pid.Tell(myinclude.HelloMessage{Say: "Hello, Everyone!", Who: "Neil"})

	time.Sleep(time.Second)
}
-5