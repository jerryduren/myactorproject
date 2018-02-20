package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

    "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/jerryduren/myactorproject/myinclude/ggsn"
)

var GlobalRecoveryCount  uint32 = 0

func init(){
	GlobalRecoveryCount = 10
}

// server is used to implement ggsn.GTPServiceServer interface
type GgsnGtpServer struct {
	PdpContext  map[string]ggsn.GgsnPdpContext
}

func NewEchoResponseMessage(reqMsg *ggsn.EchoRequest) *ggsn.EchoResponse{
	msgHead := ggsn.GtpMessageHeader{}
	//binary = 0x32 = 0b00110010: 001(Version，GTP v1)1(PT, GTP=1,GTP'=0)0(Reserve)0(Ext Flag)1(S flag)0(PN flag)
	msgHead.Flag = reqMsg.GetGtpHeader().Flag
	msgHead.MessageType = ggsn.ECHORESPONSE
	msgHead.SequenceNumber = reqMsg.GetGtpHeader().SequenceNumber
	msgHead.Teid = 0																	//Echo消息的TEID置始终为零
	newMsg:= ggsn.EchoResponse{&msgHead,GlobalRecoveryCount,300}

	newMsg.GtpHeader.Length = uint32(newMsg.Size())								//消息填充完后设置, 包括GTP消息头和消息体的长度

	return &newMsg
}


// Implements the interface of ggsn.GTPServiceServer
func (this *GgsnGtpServer)Echo(ctx context.Context, reqMsg *ggsn.EchoRequest) (*ggsn.EchoResponse, error){
	fmt.Println("GGSN Server received a message:",reqMsg.GetGtpHeader().MessageType.String())
	fmt.Println(reqMsg.String())

	return NewEchoResponseMessage(reqMsg), nil
}

func (this *GgsnGtpServer)CreatePdpContext(ctx context.Context, reqMsg *ggsn.CreatePdpContextRequest) (*ggsn.CreatePdpContextResponse, error){
	//fmt.Println("GGSN Server received a message:",reqMsg.GetGtpHeader().MessageType.String())
	fmt.Println(reqMsg.String())

	return &ggsn.CreatePdpContextResponse{},nil
}
func (this *GgsnGtpServer)UpdatePdpContext(ctx context.Context, reqMsg *ggsn.UpdatePdpContextRequest) (*ggsn.UpdatePdpContextResponse, error){
	//fmt.Println("GGSN Server received a message:",reqMsg.GetGtpHeader().MessageType.String())
	fmt.Println(reqMsg.String())

	return &ggsn.UpdatePdpContextResponse{},nil
}
func (this *GgsnGtpServer)DeletePdpContext(ctx context.Context, reqMsg *ggsn.DeletePdpContextRequest) (*ggsn.DeletePdpContextResponse, error){
	//fmt.Println("GGSN Server received a message:",reqMsg.GetGtpHeader().MessageType.String())
	fmt.Println(reqMsg.String())

	return &ggsn.DeletePdpContextResponse{},nil
}


func main() {
	// Establish a TCP server, which is listening on 2123 port
	lis, err := net.Listen("tcp", ":2123")
	if err != nil {
		log.Fatalf("failed to start TCP Socker: %v", err)
		os.Exit(1)
	}
	fmt.Println("GGSN GTP Server has started, which is listenning to 2123 port...")

	// 将应用层的服务与TCP连接绑定
	s := grpc.NewServer()
	ggsn.RegisterGTPServiceServer(s,&GgsnGtpServer{PdpContext:make(map[string]ggsn.GgsnPdpContext)})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("GGSN GTP Server is listening on 2123 port...")
}
