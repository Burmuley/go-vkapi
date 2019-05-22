package objects

/////////////////////////////////////////////////////////////
// Friends related API objects	                           //
/////////////////////////////////////////////////////////////

// FriendStatusStatus represents `friends_friend_status_status` API object
type FriendsFriendStatusStatus int

func (f *FriendsFriendStatusStatus) MarshalJSON() ([]byte, error) {
	return GetIntFromRange(int(*f), 0, 3, true)
}

func (f *FriendsFriendStatusStatus) GetName() string {
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
type FriendsFriendStatus struct {
	FriendStatus   FriendsFriendStatusStatus `json:"friend_status"`
	ReadState      BaseBoolInt               `json:"read_state"`
	RequestMessage string                    `json:"request_message"`
	Sign           string                    `json:"sign"`
	UserID         int                       `json:"user_id"`
}

// FriendsList represents `friends_friends_list` API object
type FriendsFriendsList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// MutualFriend represents `friends_mutual_friend` API object
type FriendsMutualFriend struct {
	CommonCount   int   `json:"common_count"`
	CommonFriends []int `json:"common_friends"`
	ID            int   `json:"id"`
}

// RequestsMutual represents `friends_requests_mutual` API object
type FriendsRequestsMutual struct {
	Count int   `json:"count"`
	Users []int `json:"users"`
}

// Requests represents `friends_requests` API object
type FriendsRequests struct {
	From   string                `json:"from"`
	Mutual FriendsRequestsMutual `json:"mutual"`
	UserID int                   `json:"user_id"`
}

// RequestsXtrMessage represents `friends_requests_xtr_message` API object
type FriendsRequestsXtrMessage struct {
	From    string                `json:"from"`
	Message string                `json:"message"`
	Mutual  FriendsRequestsMutual `json:"mutual"`
	UserID  int                   `json:"user_id"`
}

// UserXtrLists represents `friends_user_xtr_lists` API object
type FriendsUserXtrLists struct {
	*UsersUserFull
	Lists []int `json:"lists"`
}

// UserXtrPhone represents `friends_user_xtr_phone` API object
type FriendsUserXtrPhone struct {
	*UsersUserFull
	Phone string `json:"phone"`
}
