/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package smf

type BackupAmfInfo struct {

	BackupAmfName string `json:"backupAmfName"`

	GuamiList []Guami `json:"guamiList,omitempty"`
}
