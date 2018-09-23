package main

import (
	"context"

	"github.com/jerryduren/myactorproject/unc/smf/smfclient"
)

func main() {
	cfg := smfclient.NewConfiguration()
	cfg.BasePath = "http://localhost:8080/nsmf-pdusession/v1"

	amf := smfclient.NewAPIClient(cfg)

	ctx := context.Context()
	reqBody:=smfclient.SmContextCreateData{}

	// 执行创建PDU Session的SM Context操作！
	rspBody, rsp, err := amf.PDUSessionsCollectionApi.PostPduSessions(ctx, reqBody)
	if err!=nil{
		responseAPI := smfclient.NewAPIResponseWithError("")
	}else {
		responseAPI := smfclient.NewAPIResponse(rsp)
	}
}
