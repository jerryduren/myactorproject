package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"time"
)

type Ping struct {
	SendTime time.Time
	Payload  string
}
type PingServer struct {
}

func (this *PingServer) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Ping:
		fmt.Println(time.Now(), "I receive a Ping message from", context.Sender(), msg.Payload)
		time.Sleep(3 * time.Second)
		context.Respond(Ping{msg.SendTime, msg.Payload})
	}
}

func main() {
	pingServer := actor.Spawn(actor.FromProducer(func() actor.Actor { return &PingServer{} }))

	// RequestFuture sends a message to a given PID and returns a Future
	// func (pid *PID) RequestFuture(message interface{}, timeout time.Duration) *Future
	// Future 也是一个actor，他有自己的PID。future这ACTOR给pingServer发送一个消息，并且等待响应消息，等待时间 2s，超时重传，否则continue
	f := pingServer.RequestFuture(Ping{time.Now(), "HIJKLMN"}, 2*time.Second)
	msg, err := f.Result()
	if err == nil {
		switch msg.(type) {
		case Ping:
			//fmt.Println("Receive Ping Reply, RTT = ", time.Now().Sub(Ping(msg).SendTime), "MSG: ", msg)
		}
	} else {
		fmt.Println("Future Timeout, I will resend this message.")
		pingServer.Request(Ping{time.Now(), "abcdefg"}, f.PID())
	}

	time.Sleep(6 * time.Second)
}
