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
	UserID    uuid_t
	FirstName string
	LastName  string
	Username  string
	Email     string
	Follows   []*Project
}

// Project Entity
type Project struct {
	ProjectID       uuid_t
	Name            string
	Description     string
	Pindrop         *Pindrop
	Timeline        *Timeline
	DiscussionBoard *DiscussionBoard
	FollowerCount   int
}

// Pindrop Entity
type Pindrop struct {
	PindropID uuid_t
	Longitude float32
	Latitude  float32
	Project   *Project
}

// Timeline Entity
type Timeline struct {
	TimelineID uuid_t
	Head       *TimelinePost
	Tail       *TimelinePost
}

// Timeline Post Entity
type TimelinePost struct {
	TimelinePostID uuid_t
	Next           *TimelinePost
	Prev           *TimelinePost
	Title          string
	Body           string
	Images         []Image
}

// Image Entity
type Image struct {
	ImageID uuid_t
	AltText string
}

// Discussion Board Entity
type DiscussionBoard struct {
	DiscussionBoardID uuid_t
	Root              *DiscussionBoardMessage
}

// Discussion Board Message Entity
type DiscussionBoardMessage struct {
	DiscussionBoardMessageID uuid_t
	Parent                   *DiscussionBoardMessage
	Children                 []*DiscussionBoardMessage
	Body                     string
}
