package main

import (
	"github.com/jerryduren/myactorproject/myinclude"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
)

type HelloActor struct {
	startTime time.Time
	stopTime time.Time
}

func (this *HelloActor)Receive(context actor.Context){
	switch msg:=  context.Message  ().(type){
	case myinclude.HelloMessage:
		fmt.Println(msg.Say)
	case  *actor.Stopped:
		this.stopTime = time.Now()
	}
}
func main() {
	NewActor:=&HelloActor{startTime:time.Now(),stopTime:time.Now()}
	pid:=actor.Spawn(actor.FromInstance(NewActor))
	pid.Tell(myinclude.HelloMessage{Say:"Hello, I am Jerry Du!"})
	

	pid.Stop()
	time.Sleep(time.Second)
	fmt.Println(NewActor.startTime,NewActor.stopTime)
	fmt.Printf("Escape Time:%v",NewActor.stopTime.Sub(NewActor.startTime))
}
