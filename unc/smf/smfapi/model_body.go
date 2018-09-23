/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smfapi

import (
	"os"
)

type Body struct {

	JsonData *SmContextCreateData `json:"jsonData,omitempty"`

	BinaryDataN1SmMessage **os.File `json:"binaryDataN1SmMessage,omitempty"`
}