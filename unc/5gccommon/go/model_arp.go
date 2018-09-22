/*
 * Common Data Types
 *
 * Common Data Types
 *
 * API version: 1.R15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package 5gccommon

type Arp struct {

	PriorityLevel int32 `json:"priorityLevel"`

	PreemptCap *PreemptionCapability `json:"preemptCap"`

	PreemptVuln *PreemptionVulnerability `json:"preemptVuln"`
}
