package board

import (
	"fmt"
	"gitlab.com/Burmuley/go-vkapi/objects/base"
	"gitlab.com/Burmuley/go-vkapi/objects/polls"
	"gitlab.com/Burmuley/go-vkapi/objects/wall"
)

/////////////////////////////////////////////////////////////
// Board related API objects	                           //
/////////////////////////////////////////////////////////////

// DefaultOrder represents `board_default_order` API object
type DefaultOrder int

func (b *DefaultOrder) GetName() string {
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

func (b *DefaultOrder) MarshalJSON() ([]byte, error) {
	if *b < -2 || *b > 2 || *b == 0 {
		return []byte{}, fmt.Errorf("value is not in range [-2, -1, 1, 2]")
	}

	return []byte{byte(*b)}, nil
}

// Topic represents `board_topic` API object
type Topic struct {
	Comments  int          `json:"comments"`
	Created   int          `json:"created"`
	Creator   int          `json:"created_by"`
	ID        int          `json:"id"`
	Closed    base.BoolInt `json:"is_closed"`
	Fixed     base.BoolInt `json:"is_fixed"`
	Title     string       `json:"title"`
	Updated   int          `json:"updated"`
	UpdaterID int          `json:"updated_by"`
}

// TopicComment represents `board_topic_comment` API object
type TopicComment struct {
	Attachments []wall.CommentAttachment `json:"attachments"`
	Date        int                      `json:"date"`
	AuthorID    int                      `json:"from_id"`
	ID          int                      `json:"id"`
	RealOffset  int                      `json:"real_offset"`
	Text        string                   `json:"text"`
}

// TopicPoll represents `board_topic_poll` API object
type TopicPoll struct {
	AnswerID int            `json:"answer_id"`
	Answers  []polls.Answer `json:"answers"`
	Created  int            `json:"created"`
	Closed   base.BoolInt   `json:"is_closed"`
	OwnerID  int            `json:"owner_id"`
	PollID   int            `json:"poll_id"`
	Question string         `json:"question"`
	Votes    string         `json:"votes"`
}
