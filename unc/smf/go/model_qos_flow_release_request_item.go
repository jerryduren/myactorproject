/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smf

type QosFlowReleaseRequestItem struct {

	Qfi int32 `json:"qfi"`

	QosRules string `json:"qosRules,omitempty"`
}