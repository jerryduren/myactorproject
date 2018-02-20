package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jerryduren/myactorproject/myinclude/ggsn"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var GlobalSequenceNumber uint32=0
var gtpServiceClient ggsn.GTPServiceClient
var conn *grpc.ClientConn

func init(){
	randNumSeed:= rand.New(rand.NewSource(time.Now().UnixNano()))
	GlobalSequenceNumber = randNumSeed.Uint32()

	// Initiate Client of  GTP Service
	/* 首先与服务器建立连接  */
	hostname := "localhost:2123"
	if len(os.Args) > 1 {
		hostname = os.Args[1]+":2123"
	}
	conn, err := grpc.Dial(hostname, grpc.WithInsecure())
	//defer conn.Close()
	if err != nil {
		log.Fatal("did not establish connection with hostname %s, the cause is: %v", hostname, err)
		conn.Close()
		os.Exit(1)
	}
	gtpServiceClient = ggsn.NewGTPServiceClient(conn)
}

func NewEchoRequestMessage() *ggsn.EchoRequest{
	msgHead := ggsn.GtpMessageHeader{}
	//binary = 0x32 = 0b00110010: 001(Version，GTP v1)1(PT, GTP=1,GTP'=0)0(Reserve)0(Ext Flag)1(S flag)0(PN flag)
	msgHead.Flag =0x32
	msgHead.MessageType = ggsn.ECHOREQUEST
	msgHead.SequenceNumber,GlobalSequenceNumber =GlobalSequenceNumber,GlobalSequenceNumber+1
	msgHead.Teid = 0																	//Echo消息的TEID置始终为零
	newEchoMsg:= ggsn.EchoRequest{&msgHead,200}
	newEchoMsg.GtpHeader.Length = uint32(newEchoMsg.Size())								//消息填充完后设置, 包括GTP消息头和消息体的长度
	return &newEchoMsg
}

func TestEcho(interval time.Duration){
	if interval == 0{
		return
	}
		for {
			echoTimer :=time.NewTimer(interval)
			<-echoTimer.C
			echoRspMsg,err :=gtpServiceClient.Echo(context.Background(),NewEchoRequestMessage())
			if err != nil {
				log.Fatalf("Echo GGSN GTP Server failure: %v", err)
			}
			log.Printf(": %s", echoRspMsg.String())
		}
}

func main() {
	//defer conn.Close()
	/* 调用RPC服务 */
	/* Test Echo Service Operation */
	go TestEcho(time.Second*2)

	/* Test Crate PDP Context Service Operation */
	imsi := "460001234567890"
	crtRspMsg,err:=gtpServiceClient.CreatePdpContext(context.Background(),&ggsn.CreatePdpContextRequest{Imsi:imsi})
	if err != nil {
		log.Fatalf("Create PDP Context failure: %v", err)
	}
	log.Printf(": %s", crtRspMsg.String())

	/* Test Update PDP Context Service Operation */
	updateRspMsg,err:=gtpServiceClient.UpdatePdpContext(context.Background(),&ggsn.UpdatePdpContextRequest{Imsi:imsi})
	if err != nil {
		log.Fatalf("Update PDP Context failure: %v", err)
	}
	log.Printf(": %s", updateRspMsg.String())

	/* Test Delete PDP Context Service Operation */
	deleteRspMsg,err:=gtpServiceClient.DeletePdpContext(context.Background(),&ggsn.DeletePdpContextRequest{})
	if err != nil {
		log.Fatalf("Delete PDP Context failure: %v", err)
	}
	log.Printf(": %s", deleteRspMsg.String())

	time.Sleep(time.Second*16)
}
