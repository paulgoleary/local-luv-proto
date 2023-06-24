/*
 * Ratio API
 *
 * API endpoints and models for using the Ratio service
 *
 * API version: 1.0.0
 * Contact: support@ratio.me
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AchLimit struct {
	// The currency for the limit
	Currency string `json:"currency,omitempty"`
	// The maximum allowable sum of ACH transaction amounts
	Limit string `json:"limit,omitempty"`
	// The current sum of ACH transaction amounts by the ACH limit type (e.g. daily)
	Used string `json:"used,omitempty"`
	// The remaining allowable sum, calculated as the delta between the limit and used values
	Remaining string `json:"remaining,omitempty"`
}