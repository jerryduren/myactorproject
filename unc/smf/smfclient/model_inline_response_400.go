/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smfclient
import (
	"os"
)

type InlineResponse400 struct {
	JsonData SmContextCreateError `json:"jsonData,omitempty" xml:"jsonData"`
	BinaryDataN1SmMessage *os.File `json:"binaryDataN1SmMessage,omitempty" xml:"binaryDataN1SmMessage"`
}
