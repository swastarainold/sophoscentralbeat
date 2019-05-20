/*
 * Sophos Public API
 *
 * Swagger Specifications for public APIs
 *
 * API version: beta
 * Contact: support@sophos.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sophoscentral

// Sophos Central License Model
type CurrentLicensesResponse struct {
	Licenses []CustomerLicense `json:"licenses,omitempty"`
}
