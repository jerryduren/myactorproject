package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"fmt"
	"github.com/jerryduren/myactorproject/myinclude"
	"time"
)

type StatsStartStopTime struct {
	startTime time.Time
	stopTime time.Time
}

func (this * StatsStartStopTime)Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		this.stopTime = time.Now()
	case *actor.Stopping:
		//
	case *actor.Stopped:
		this.stopTime = time.Now()
		_=msg
	case *actor.Restarting:
		//
	case *actor.Restart:
		//
	}
}

type LifecycleEventActor struct {
}

var now =time.Now()

func (this *LifecycleEventActor)Receive(context actor.Context){
	switch msg:=context.Message().(type) {
	case *actor.Started:
		fmt.Println("The actor is started.")
	case *actor.Stopping:
		fmt.Println("The actor is stopping.")
		now=time.Now()
		fmt.Println(time.Now())
	case *actor.Stopped:
		fmt.Println("The actor is stopped.")
		fmt.Println(time.Now())
		/* 有点不可思议啊，停一个actor需要64ms */
		fmt.Println("Actor spend time: ",time.Since(now))
	case *actor.Restarting:
		fmt.Println("The actor is restarting ...")
	case *actor.Restart:
		fmt.Println("The actor is restarted.")
	case myinclude.HelloMessage:
		fmt.Println(msg.Who,"SAY:",msg.Say)
		}
		}

		/* 这样测试actor启停的时间是不对的 */
func DelayOfStartStop() time.Duration{
	newactor:=&StatsStartStopTime{}
	pid:=actor.Spawn(actor.FromInstance(newactor))
	pid.Stop()
	time.Sleep(time.Second)
	return newactor.stopTime.Sub(newactor.startTime)
}

func main() {
	pid:=actor.Spawn(actor.FromInstance(&LifecycleEventActor{}))
	pid.Tell(myinclude.HelloMessage{Who:"Jerry",Say:"Hello Neil." })
	time.Sleep(time.Second)

	pid.Stop()

	fmt.Println(DelayOfStartStop())
	time.Sleep(time.Second)
}
