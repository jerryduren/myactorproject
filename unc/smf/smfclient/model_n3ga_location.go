/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

type N3gaLocation struct {
	N3gppTai Tai `json:"n3gppTai" xml:"n3gppTai"`
	N3IwfId string `json:"n3IwfId" xml:"n3IwfId"`
	UeIpv4Addr string `json:"ueIpv4Addr,omitempty" xml:"ueIpv4Addr"`
	UeIpv6Addr string `json:"ueIpv6Addr,omitempty" xml:"ueIpv6Addr"`
	PortNumber int32 `json:"portNumber,omitempty" xml:"portNumber"`
}