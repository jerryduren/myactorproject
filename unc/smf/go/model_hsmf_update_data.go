/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smf

import (
	"time"
)

type HsmfUpdateData struct {

	RequestIndication *RequestIndication `json:"requestIndication"`

	Pei string `json:"pei,omitempty"`

	VcnTunnelInfo *TunnelInfo `json:"vcnTunnelInfo,omitempty"`

	AnType *AccessType `json:"anType,omitempty"`

	UeLocation *UserLocation `json:"ueLocation,omitempty"`

	UeTimeZone string `json:"ueTimeZone,omitempty"`

	AddUeLocation *UserLocation `json:"addUeLocation,omitempty"`

	AddUeLocTime time.Time `json:"addUeLocTime,omitempty"`

	PauseCharging bool `json:"pauseCharging,omitempty"`

	Pti int32 `json:"pti,omitempty"`

	N1SmInfoFromUe *RefToBinaryData `json:"n1SmInfoFromUe,omitempty"`

	UnknownN1SmInfo *RefToBinaryData `json:"unknownN1SmInfo,omitempty"`

	QosFlowsRelNotifyList []QosFlowItem `json:"qosFlowsRelNotifyList,omitempty"`

	QosFlowsNotifyList []QosFlowNotifyItem `json:"qosFlowsNotifyList,omitempty"`

	NotifyList []PduSessionNotifyItem `json:"NotifyList,omitempty"`

	EpsBearerId []int32 `json:"epsBearerId,omitempty"`

	HoPreparationIndication bool `json:"hoPreparationIndication,omitempty"`

	RevokeEbiList []int32 `json:"revokeEbiList,omitempty"`

	Cause *Cause `json:"cause,omitempty"`
}
