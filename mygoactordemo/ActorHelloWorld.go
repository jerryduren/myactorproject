package main

import (
	"github.com/jerryduren/myactorproject/myinclude"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
	"os"
)

type HelloActor struct {
	startTime time.Time
	stopTime time.Time
}

func (this *HelloActor)Receive(context actor.Context){
	switch msg:=  context.Message  ().(type){
	case myinclude.HelloMessage:
<<<<<<< HEAD:mygoactordemo/ActorHelloWorld.go
		fmt.Println(msg.Who,"say",msg.Say)
	}
}
func main() {
	pid:=actor.Spawn(actor.FromInstance(&HelloActor{}))
	if len(os.Args)==2{
		pid.Tell(myinclude.HelloMessage{Who:os.Args[1]})
	}else if len(os.Args)==3{
		pid.Tell(myinclude.HelloMessage{Who:os.Args[1],Say:os.Args[2]})
	}else if len(os.Args)==1{
		pid.Tell(myinclude.HelloMessage{Who:"",Say:"Hello, I am Jerry Du!"})
		}else {
			fmt.Println("Usge: actorhelloword <who> <say>")
	}

		fmt.Println(msg.Say)
	case  *actor.Stopped:
		this.stopTime = time.Now()
	}
}
func main() {
	NewActor:=&HelloActor{startTime:time.Now(),stopTime:time.Now()}
	pid:=actor.Spawn(actor.FromInstance(NewActor))
	pid.Tell(myinclude.HelloMessage{Say:"Hello, I am Jerry Du!"})
	
>>>>>>> 6898f5fe8763810e790646fb055a40ca6b224491:mygoactordemo/HelloWorld.go

	pid.Stop()
	time.Sleep(time.Second)
	fmt.Println(NewActor.startTime,NewActor.stopTime)
	fmt.Printf("Escape Time:%v",NewActor.stopTime.Sub(NewActor.startTime))
}
