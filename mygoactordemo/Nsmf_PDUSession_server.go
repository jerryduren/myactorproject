package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/jerryduren/myactorproject/myinclude/ngc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement ngc.Nsmf_PDUSessionServer interface
type Nsmf_PDUSession_Server struct {
	CtxList map[string]SMContextActor
}

type SMContextActor struct {
	SUPI              string
	Result_Indication string
	Cause             uint32
	PDU_Session_ID    uint64
	N2_SM_Info        []byte
	N1_SM_Container   []byte
}

// SayHello implements helloworld.GreeterServer
func (s *Nsmf_PDUSession_Server) CreateSMContext(ctx context.Context, in *ngc.CreateSMContextRequest) (*ngc.CreateSMContextResponse, error) {
	fmt.Println("... Print all SM Context ...")
	for key, value := range s.CtxList {
		fmt.Println(key, value)
	}
	fmt.Println("I receive a Create SM Context Ruqeust message.")
	fmt.Println(in.String())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s.CtxList[in.SUPI] = SMContextActor{SUPI: in.GetSUPI(), Result_Indication: "successed", Cause: 200, PDU_Session_ID: r.Uint64()}
	return &ngc.CreateSMContextResponse{Result_Indication: "Successful", Cause: 200}, nil
}

func main() {
	fmt.Println("I am listenning to 50051 port...")
	port := ":50051"
	if len(os.Args) > 2 {
		port = os.Args[1]
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ngc.RegisterNsmf_PDUSessionServer(s, &Nsmf_PDUSession_Server{CtxList: make(map[string]SMContextActor)})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
