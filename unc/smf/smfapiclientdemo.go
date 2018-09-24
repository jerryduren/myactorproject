package main

import (
	"context"
	"fmt"

	"github.com/jerryduren/myactorproject/unc/smf/smfclient"
)

func main() {
	cfg := smfclient.NewConfiguration()
	cfg.BasePath = "http://localhost:8080/nsmf-pdusession/v1"
	// add customer header
	cfg.AddDefaultHeader("myhead","w want to add a custom header.")
	ctx:=cfg.HTTPClient.request
//  cfg.Scheme = "token"

	amf := smfclient.NewAPIClient(cfg)


	pduCreateData := smfclient.SmContextCreateData{}
	pduCreateData.Supi = "460021234567890"
	pduCreateData.ServingNfId = "amdif"
	pduCreateData.AnType = smfclient.AccessType("3GPP_ACCESS")
	pduCreateData.SmContextStatusUri = ""
	pduCreateData.HSmfUri=""

	// 执行创建PDU Session的SM Context操作！
	pduCreatedData, _, err := amf.SMContextsCollectionApi.PostSmContexts(context.Background(),smfclient.Body{pduCreateData,nil})
	if err!=nil{
		fmt.Println(err)
		//responseAPI := smfclient.NewAPIResponseWithError("")
	}else {
		//responseAPI := smfclient.NewAPIResponse(rsp)
		fmt.Println("Succeed, PDU Session ID = ",pduCreatedData.PduSessionId)
	}
}
