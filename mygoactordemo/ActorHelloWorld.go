package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/jerryduren/myactorproject/myinclude"
	"os"
	"time"
)

type HelloActor struct {
	startTime time.Time
	stopTime  time.Time
}

func (this *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case myinclude.HelloMessage:
		fmt.Println(msg.Who, "say", msg.Say)
	case *actor.Stopped:
		this.stopTime = time.Now()
	}
}
func main() {
	pid := actor.Spawn(actor.FromInstance(&HelloActor{}))
	if len(os.Args) == 2 {
		pid.Tell(myinclude.HelloMessage{Who: os.Args[1]})
	} else if len(os.Args) == 3 {
		pid.Tell(myinclude.HelloMessage{Who: os.Args[1], Say: os.Args[2]})
	} else if len(os.Args) == 1 {
		pid.Tell(myinclude.HelloMessage{Who: "", Say: "Hello, I am Jerry Du!"})
	} else {
		fmt.Println("Usge: actorhelloword <who> <say>")
	}

}
