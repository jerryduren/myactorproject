/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient

type VsmfUpdateError struct {
	Error ProblemDetails `json:"error" xml:"error"`
	Pti int32 `json:"pti,omitempty" xml:"pti"`
	N1smCause string `json:"n1smCause,omitempty" xml:"n1smCause"`
	N1SmInfoFromUe RefToBinaryData `json:"n1SmInfoFromUe,omitempty" xml:"n1SmInfoFromUe"`
	UnknownN1SmInfo RefToBinaryData `json:"unknownN1SmInfo,omitempty" xml:"unknownN1SmInfo"`
}
