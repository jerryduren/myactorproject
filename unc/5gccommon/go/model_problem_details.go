/*
 * Common Data Types
 *
 * Common Data Types
 *
 * API version: 1.R15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package 5gccommon

type ProblemDetails struct {

	Type string `json:"type,omitempty"`

	Title string `json:"title,omitempty"`

	Status int32 `json:"status,omitempty"`

	Instance string `json:"instance,omitempty"`

	Cause string `json:"cause,omitempty"`

	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`
}
