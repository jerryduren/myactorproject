/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smfapi

type Ncgi struct {

	PlmnId *PlmnId `json:"plmnId"`

	NrCellId string `json:"nrCellId"`
}