package objects

/////////////////////////////////////////////////////////////
// OAuth related API objects	                           //
/////////////////////////////////////////////////////////////

// Error represents `oauth_error` API object
type OAuthError struct {
	Error            string `json:"error"`             // Error type
	ErrorDescription string `json:"error_description"` // Error description
	RedirectURI      string `json:"redirect_uri"`      // URI for validation
}