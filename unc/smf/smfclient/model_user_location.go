/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

type UserLocation struct {
	EutraLocation EutraLocation `json:"eutraLocation,omitempty" xml:"eutraLocation"`
	NrLocation NrLocation `json:"nrLocation,omitempty" xml:"nrLocation"`
	N3gaLocation N3gaLocation `json:"n3gaLocation,omitempty" xml:"n3gaLocation"`
}