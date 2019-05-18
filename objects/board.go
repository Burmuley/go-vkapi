package objects

import (
	"fmt"
)

/////////////////////////////////////////////////////////////
// Board related API objects	                           //
/////////////////////////////////////////////////////////////

// DefaultOrder represents `board_default_order` API object
type BoardDefaultOrder int

func (b *BoardDefaultOrder) GetName() string {
	switch *b {
	case 1:
		return "desc_updated"
	case 2:
		return "desc_created"
	case -1:
		return "asc_updated"
	case -2:
		return "asc_created"
	default:
		return "unknown"
	}
}

func (b *BoardDefaultOrder) MarshalJSON() ([]byte, error) {
	if *b < -2 || *b > 2 || *b == 0 {
		return []byte{}, fmt.Errorf("value is not in range [-2, -1, 1, 2]")
	}

	return []byte{byte(*b)}, nil
}

// Topic represents `board_topic` API object
type BoardTopic struct {
	Comments  int         `json:"comments"`
	Created   int         `json:"created"`
	Creator   int         `json:"created_by"`
	ID        int         `json:"id"`
	Closed    BaseBoolInt `json:"is_closed"`
	Fixed     BaseBoolInt `json:"is_fixed"`
	Title     string      `json:"title"`
	Updated   int         `json:"updated"`
	UpdaterID int         `json:"updated_by"`
}

// TopicComment represents `board_topic_comment` API object
type BoardTopicComment struct {
	Attachments []WallCommentAttachment `json:"attachments"`
	Date        int                     `json:"date"`
	AuthorID    int                     `json:"from_id"`
	ID          int                     `json:"id"`
	RealOffset  int                     `json:"real_offset"`
	Text        string                  `json:"text"`
}

// TopicPoll represents `board_topic_poll` API object
type BoardTopicPoll struct {
	AnswerID int           `json:"answer_id"`
	Answers  []PollsAnswer `json:"answers"`
	Created  int           `json:"created"`
	Closed   BaseBoolInt   `json:"is_closed"`
	OwnerID  int           `json:"owner_id"`
	PollID   int           `json:"poll_id"`
	Question string        `json:"question"`
	Votes    string        `json:"votes"`
}
