/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

type NonDynamic5qi struct {
	PriorityLevel int32 `json:"priorityLevel,omitempty" xml:"priorityLevel"`
	AverWindow string `json:"averWindow,omitempty" xml:"averWindow"`
	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty" xml:"maxDataBurstVol"`
}
