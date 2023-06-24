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

type User struct {
	// The unique identifier of the user
	Id string `json:"id,omitempty"`
	// The time the user was created
	CreateTime string `json:"createTime,omitempty"`
	// The time the user was last updated
	UpdateTime string `json:"updateTime,omitempty"`
	// The first name of the user
	FirstName string `json:"firstName,omitempty"`
	// The middle name of the user
	MiddleName string `json:"middleName,omitempty"`
	// The last name of the user
	LastName string `json:"lastName,omitempty"`
	// The email of the user
	Email string `json:"email,omitempty"`
	// The country of the user
	Country string `json:"country,omitempty"`
	// The phone number of the user
	Phone string `json:"phone,omitempty"`
	// The nationality of the user
	Nationality string `json:"nationality,omitempty"`
	// The occupation of the user
	Occupation string `json:"occupation,omitempty"`
	Kyc *Kyc `json:"kyc,omitempty"`
	// The array of connected bank accounts for the user
	ConnectedBankAccounts []BankAccount `json:"connectedBankAccounts,omitempty"`
	// An array of review flags set on the user, if any
	Flags []string `json:"flags,omitempty"`
}