package friends

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/users"
)

/////////////////////////////////////////////////////////////
// Friends related API objects	                           //
/////////////////////////////////////////////////////////////

// FriendStatusStatus represents `friends_friend_status_status` API object
type FriendStatusStatus int

func (f *FriendStatusStatus) MarshalJSON() ([]byte, error) {
	if *f < 0 || *f > 3 {
		return []byte{}, fmt.Errorf("value not in range 0..3")
	}

	return []byte{byte(*f)}, nil
}

func (f *FriendStatusStatus) GetName() string {
	switch *f {
	case 0:
		return "not a friend"
	case 1:
		return "outcoming request"
	case 2:
		return "incoming request"
	case 3:
		return "is friend"
	default:
		return "unknown"
	}
}

// FriendStatus represents `friends_friend_status` API object
type FriendStatus struct {
	FriendStatus   FriendStatusStatus `json:"friend_status"`
	ReadState      base.BoolInt       `json:"read_state"`
	RequestMessage string             `json:"request_message"`
	Sign           string             `json:"sign"`
	UserID         int                `json:"user_id"`
}

// FriendsList represents `friends_friends_list` API object
type FriendsList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// MutualFriend represents `friends_mutual_friend` API object
type MutualFriend struct {
	CommonCount   int   `json:"common_count"`
	CommonFriends []int `json:"common_friends"`
	ID            int   `json:"id"`
}

// RequestsMutual represents `friends_requests_mutual` API object
type RequestsMutual struct {
	Count int   `json:"count"`
	Users []int `json:"users"`
}

// Requests represents `friends_requests` API object
type Requests struct {
	From   string         `json:"from"`
	Mutual RequestsMutual `json:"mutual"`
	UserID int            `json:"user_id"`
}

// RequestsXtrMessage represents `friends_requests_xtr_message` API object
type RequestsXtrMessage struct {
	From    string         `json:"from"`
	Message string         `json:"message"`
	Mutual  RequestsMutual `json:"mutual"`
	UserID  int            `json:"user_id"`
}

// UserXtrLists represents `friends_user_xtr_lists` API object
type UserXtrLists struct {
	users.UserFull
	Lists []int `json:"lists"`
}

// UserXtrPhone represents `friends_user_xtr_phone` API object
type UserXtrPhone struct {
	users.UserFull
	Phone string `json:"phone"`
}
