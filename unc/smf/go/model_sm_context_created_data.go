/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smf

type SmContextCreatedData struct {

	PduSessionId int32 `json:"pduSessionId,omitempty"`

	SNssai *Snssai `json:"sNssai,omitempty"`

	UpCnxState *UpCnxState `json:"upCnxState,omitempty"`

	N2SmInfo *RefToBinaryData `json:"n2SmInfo,omitempty"`

	N2SmInfoType int32 `json:"n2SmInfoType,omitempty"`

	AllocatedEbiList []EbiArpMapping `json:"allocatedEbiList,omitempty"`

	HoState *HoState `json:"hoState,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
