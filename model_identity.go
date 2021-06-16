/*
 * Verification API
 *
 * # Introduction  <span class=\"subtext\"> Welcome to the Passbase Verifications API docs. This documentation will help you understand our models and the Verification API with its endpoints. Based on this you can build your own system (i.e. verification) and hook it up to Passbase.  In case of feedback or questions you can reach us under this email address: [developer@passbase.com](mailto:developer@passbase.com). </span>  A User submits a video selfie and valid identifying __Resources__ during a __Verification__ guided by the Passbase client-side integration. Once all the necessary __Resources__ are submitted, __Data points__ are extracted, digitized, and authenticated. These Data points then becomes part of the User's __Identity__. The User then consents to share __Resources__ and/or __Data points__ from their Identity with you. This information is passed to you and can be used to make decisions about a User (e.g. activate account). This table below explains our terminology further.  | Term                                    | Description | |-----------------------------------------|-------------| | [Identity](#tag/identity_model)         | A set of Data points and Resources related to and owned by one single User. This data can be accessed by you through a Verification. | | Data points                             | Any data about a User extracted from a Resource (E.g. Passport Number, or Age). | | [Resource](#tag/resource_model)         | A source document used to generate the Data points for a User (E.g. Passport). | | [User](#tag/user_model)                 | The owner of an email address associated with an Identity. | | Verification                            | A transaction through which a User consents to share Data points with you. If the Data points you request are not already available in the User's Identity, the Passbase client will ask the User to submit the necessary Resource required to extract them. | | Re-authentication (login)               | A transaction through which a User can certify the ownership of Personal data previously shared through an Authentication. |   # Authentication  <span class=\"subtext\"> There are two forms of authentication for the API: <br/>&bull; API Key <br/>&bull; Bearer JWT Token  </span> 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package passbase

type Identity struct {
	// Unique ID of the identity
	Id string `json:"id,omitempty"`
	// Current state of the identity in Passbase's systems
	Status string `json:"status,omitempty"`
	Owner *IdentityOwner `json:"owner,omitempty"`
	// Float between 0 and 1 representing our confidence that this identity is valid. 0 meaning we couldn't verify any of the information provided with accuracy, and 1 absolute confidence.
	Score float64 `json:"score,omitempty"`
	// Unix-timestamp of when the identity was created
	Created int64 `json:"created,omitempty"`
	// Unix-timestamp of when the identity was updated
	Updated int64 `json:"updated,omitempty"`
	// resources attached to a verification
	Resources []IdentityResource `json:"resources,omitempty"`
	// Customer defined arbitrary payload initially passed through the client-sdk
	Metadata *interface{} `json:"metadata,omitempty"`
	Watchlist *WatchlistResponse `json:"watchlist,omitempty"`
}
