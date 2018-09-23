package  smf

import (
	"net/http"
	"time"

	"github.com/jerryduren/myactorproject/unc/smf/smfapi"
)

var UEs map[string]UE

type UE struct {
	//
	PduSessions map[int32]*PduSession

}

type PduSession struct {
	// ...
	SmContext SmContext
}

type SmContext struct {
	// ...
	Supi string `json:"supi,omitempty"`
	UnauthenticatedSupi bool `json:"unauthenticatedSupi,omitempty"`
	Pei string `json:"pei,omitempty"`
	Gpsi string `json:"gpsi,omitempty"`
	PduSessionId int32 `json:"pduSessionId,omitempty"`
	Dnn string `json:"dnn,omitempty"`
	SNssai *smfapi.Snssai `json:"sNssai,omitempty"`
	HplmnSnssai *smfapi.Snssai `json:"hplmnSnssai,omitempty"`
	ServingNfId string `json:"servingNfId"`
	Guami *smfapi.Guami `json:"guami,omitempty"`
	RequestType *smfapi.RequestType `json:"requestType,omitempty"`
	//N1SmMsg *RefToBinaryData `json:"n1SmMsg,omitempty"`
	AnType *smfapi.AccessType `json:"anType"`
	PresenceInLadn bool `json:"presenceInLadn,omitempty"`
	UeLocation *smfapi.UserLocation `json:"ueLocation,omitempty"`
	UeTimeZone string `json:"ueTimeZone,omitempty"`
	AddUeLocation *smfapi.UserLocation `json:"addUeLocation,omitempty"`
	AddUeLocTime time.Time `json:"addUeLocTime,omitempty"`
	SmContextStatusUri string `json:"smContextStatusUri"`
	HSmfUri string `json:"hSmfUri,omitempty"`
	AdditionalHsmfUri []string `json:"additionalHsmfUri,omitempty"`
	OldPduSessionId int32 `json:"oldPduSessionId,omitempty"`
	PduSessionsActivateList []int32 `json:"pduSessionsActivateList,omitempty"`
	UeEpsPdnConnection string `json:"ueEpsPdnConnection,omitempty"`
	HoState *smfapi.HoState `json:"hoState,omitempty"`
	PcfId string `json:"pcfId,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	SelMode *smfapi.DnnSelectionMode `json:"selMode,omitempty"`
	BackupAmfInfo []smfapi.BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	// ...
	UpCnxState *smfapi.UpCnxState `json:"upCnxState,omitempty"`
	//N2SmInfo *RefToBinaryData `json:"n2SmInfo,omitempty"`
	//N2SmInfoType int32 `json:"n2SmInfoType,omitempty"`
	AllocatedEbiList []smfapi.EbiArpMapping `json:"allocatedEbiList,omitempty"`
	HoState *smfapi.HoState `json:"hoState,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}

// 服务器侧的创建PDU Session的SM Context的处理函数
func CreateSmContext (w http.Response, r *http.Request)
{
	// ...
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}