/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

type TunnelInfo struct {
	Ipv4Addr string `json:"ipv4Addr,omitempty" xml:"ipv4Addr"`
	Ipv6Addr string `json:"ipv6Addr,omitempty" xml:"ipv6Addr"`
	GtpTeid string `json:"gtpTeid" xml:"gtpTeid"`
}
