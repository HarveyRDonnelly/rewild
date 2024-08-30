// Package entities defines the structs and interfaces for the
// application's entities.
package entities

// Import packages
import (
	"github.com/google/uuid"
)

// Alias UUID type
type uuid_t = uuid.UUID

// User Entity
type User struct {
	UserID    uuid_t     `json:"user_id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Follows   []*Project `json:"follows"`
}

// Project Entity
type Project struct {
	ProjectID       uuid_t           `json:"project_id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Pindrop         *Pindrop         `json:"pindrop"`
	Timeline        *Timeline        `json:"timeline"`
	DiscussionBoard *DiscussionBoard `json:"discussion_board"`
	FollowerCount   int              `json:"follower_count"`
}

// Pindrop Entity
type Pindrop struct {
	PindropID uuid_t  `json:"pindrop_id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Timeline Entity
type Timeline struct {
	TimelineID uuid_t         `json:"timeline_id"`
	HeadID     uuid_t         `json:"head"`
	TailID     uuid_t         `json:"tail"`
	Posts      []TimelinePost `json:"posts"`
}

// Timeline Post Entity
type TimelinePost struct {
	TimelinePostID uuid_t  `json:"timeline_post_id"`
	NextID         uuid_t  `json:"next_id"`
	PrevID         uuid_t  `json:"prev_id"`
	Title          string  `json:"title"`
	Body           string  `json:"body"`
	Images         []Image `json:"images"`
}

// Image Entity
type Image struct {
	ImageID uuid_t `json:"image_id"`
	AltText string `json:"alt_text"`
}

// Discussion Board Entity
type DiscussionBoard struct {
	DiscussionBoardID uuid_t                  `json:"discussion_board_id"`
	Root              *DiscussionBoardMessage `json:"root"`
}

// Discussion Board Message Entity
type DiscussionBoardMessage struct {
	DiscussionBoardMessageID uuid_t                    `json:"discussion_board_message_id"`
	Children                 []*DiscussionBoardMessage `json:"children"`
	Body                     string                    `json:"body"`
}
