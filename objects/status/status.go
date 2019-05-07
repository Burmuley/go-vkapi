package status

import "gitlab.com/Burmuley/go-vkapi/objects/audio"

// Status represents `status_status` API object
type Status struct {
	Audio audio.AudioFull `json:"audio"`
	Text  string          `json:"text"` // Status text
}
