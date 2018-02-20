package main

import (
	"log"
	"os"

	"github.com/jerryduren/myactorproject/myinclude/ngc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	/* 首先与服务器建立连接  */
	hostname := "localhost:50051"
	if len(os.Args) > 1 {
		hostname = os.Args[1]
	}
	conn, err := grpc.Dial(hostname, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal("did not establish connection with hostname %s, the cause is: %v", hostname, err)
		os.Exit(1)
	}
	service := ngc.NewNsmf_PDUSessionClient(conn)

	/* 调用RPC服务 */
	imsi := "460001234567890"
	if len(os.Args) > 2 {
		imsi = os.Args[2]
	}
	CrtSmCxtReq := ngc.CreateSMContextRequest{SUPI: imsi, DNN: "cmnet"}
	response, err := service.CreateSMContext(context.Background(), &CrtSmCxtReq)
	if err != nil {
		log.Fatalf("Create SM Context failure: %v", err)
	}
	log.Printf(": %s", response.String())
}
