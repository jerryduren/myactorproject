package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
)

type watcher struct {
}

func (this *watcher) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Terminated:
		fmt.Println(msg.Who, "is dead.")
	case *actor.Started:
		pid := actor.Spawn(actor.FromInstance(&watchee{}))
		context.Watch(pid) // 一旦监控了这个PID，这个PID死亡的时候就会给watcher发送actor.terminator消息；
	}
}

type watchee struct {
}

func (this *watchee) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *actor.Started:
		fmt.Println(context.Self().GetId(), "has started, but I choice kill self.")
		context.Self().Tell(&actor.PoisonPill{})
	}
}

func main() {
	actor.Spawn(actor.FromInstance(&watcher{}))
	time.Sleep(time.Second)
}
