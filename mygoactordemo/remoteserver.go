package main

import (
	"fmt"
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/jerryduren/myactorproject/myinclude"
	"github.com/jerryduren/myactorproject/myinclude/gtp"
)

type GtpServer struct {
}

func (this *GtpServer) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case gtp.EchoRequest:
		fmt.Println("I receive a message from ", msg.GetSender(), "which is ", msg.GetMessage())
		msg.Sender.Tell(myinclude.EchoResponse{"Result: Thanks!"})
	}
}

func main() {
	remote.Start("localhost:8091")

	pid := actor.Spawn(actor.FromInstance(&GtpServer{}))
	actor.ProcessRegistry.Add()
	actor.ProcessRegistry.Add(pid, "GTPServer")

	console.ReadLine()
}
