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
	"os"
)

type Body1 struct {

	JsonData *SmContextUpdateData `json:"jsonData,omitempty"`

	BinaryDataN1SmMessage **os.File `json:"binaryDataN1SmMessage,omitempty"`

	BinaryDataN2SmInformation **os.File `json:"binaryDataN2SmInformation,omitempty"`
}
