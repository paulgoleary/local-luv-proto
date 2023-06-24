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

type UpdateWebhookRequest struct {
	// The url for the client's webhook endpoint
	Url string `json:"url,omitempty"`
	// The events to subscribe to
	Events []string `json:"events,omitempty"`
	// A name for the webhook
	Name string `json:"name,omitempty"`
}