package main

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/jerryduren/myactorproject/myinclude"
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
		panic("child actor generates a panic...")
	case *actor.Stopped:
		fmt.Println(context.Self(), "is stopped!")
	case *actor.Restarting:
		fmt.Println(context.Self(), "is restarting!")
	}
}

func main() {
	decider := func(child *actor.PID, reason interface{}) actor.Directive {
		fmt.Println(child.Address, "is panic.")
		fmt.Printf("%v\n", reason)
		fmt.Println("Handling failure for child...")
		return actor.RestartDirective
	}
	/* 为孩子ACTOR设置监控，一旦出错就restarting the actor */
	pid := actor.Spawn(actor.FromInstance(&ParentActor{}).WithSupervisor(actor.NewOneForOneStrategy(10, 1000, decider)))
	pid.Tell(myinclude.HelloMessage{Say: "Hello, Everyone!"})

	time.Sleep(time.Second)
}
