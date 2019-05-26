package responses

/////////////////////////////////////////////////////////////
// Authentication related responses                        //
/////////////////////////////////////////////////////////////

// AuthRestore represents `auth_restore_response` API response object
type AuthRestore struct {
	Success int    `json:"success"` // 1 if success
	Sid     string `json:"sid"`     // Parameter needed to grant access by code
}
