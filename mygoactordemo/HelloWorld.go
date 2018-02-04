package main

import (
	"mypkg/myinclude"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
)

type HelloActor struct {

}

func (this *HelloActor)Receive(context actor.Context){
	switch msg:=  context.Message  ().(type){
	case myinclude.HelloMessage:
		fmt.Println(msg.Say)
	}
}
func main() {
	pid:=actor.Spawn(actor.FromInstance(&HelloActor{}))
	pid.Tell(myinclude.HelloMessage{Say:"Hello World!"})

	time.Sleep(time.Second)
}
