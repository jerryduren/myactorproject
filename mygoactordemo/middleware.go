package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/jerryduren/myactorproject/myinclude"
	"time"
)

type HelloActor struct {
}

func (this *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case myinclude.HelloMessage:
		fmt.Println(msg)
	}
}

func messageTrace(next actor.ActorFunc) actor.ActorFunc {
	return func(c actor.Context) {
		/* Begin 把中间件需要干的事情添加着，这样发送给actor的所有用户消息都可以被解惑*/
		/* 把所有的消息都打印一遍 */
		sender := c.Sender()
		msg := c.Message()
		self := c.Self()
		if sender != nil {
			fmt.Println(sender, "sends a message to ", self, "the message is: ", msg)
		} else {
			fmt.Println(self, "receives a message:", msg)
		}

		/* End */
		/* 继续处理，就像啥都没有发生一样 */
		next(c)
	}
}

func main() {
	props := actor.FromProducer(func() actor.Actor { return &HelloActor{} })
	props = props.WithMiddleware(messageTrace)
	pid := actor.Spawn(props)
	pid.Tell(myinclude.HelloMessage{"Jerry", "Hello Neil!"})
	pid.Tell(myinclude.HelloMessage{"Du", "How are you"})

	time.Sleep(time.Second)

	//很奇怪actor都会首先收到一个空的消息！！！
}
