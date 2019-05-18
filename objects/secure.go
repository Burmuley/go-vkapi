package objects

/////////////////////////////////////////////////////////////
// Secure related API objects	                           //
/////////////////////////////////////////////////////////////

// Level represents `secure_level` API object
type SecureLevel struct {
	Level  int `json:"level"` // Level
	UserID int `json:"uid"`   // User ID
}

// SmsNotification represents `secure_sms_notification` API object
type SecureSmsNotification struct {
	AppID   int    `json:"app_id"`  // Application ID
	Date    int    `json:"date"`    // Date when message has been sent in Unixtime
	ID      int    `json:"id"`      // Notification ID
	Message string `json:"message"` // Messsage text
	UserID  int    `json:"user_id"` // User ID
}

// TokenChecked represents `secure_token_checked` API object
type SecureTokenChecked struct {
	Date    int            `json:"date"`    // Date when access_token has been generated in Unixtime
	Expire  int            `json:"expire"`  // Date when access_token will expire in Unixtime
	Success BaseOkResponse `json:"success"` // Returns if successfully processed
	UserID  int            `json:"user_id"` // User ID
}

// Transaction represents `secure_transaction` API object
type SecureTransaction struct {
	Date   int `json:"date"`     // Transaction date in Unixtime
	ID     int `json:"id"`       // Transaction ID
	FromID int `json:"uid_from"` // From ID
	ToID   int `json:"uid_to"`   // To ID
	Votes  int `json:"votes"`    // Votes number
}
