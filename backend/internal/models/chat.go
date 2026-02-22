package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	AvatarURL string    `json:"avatar_url"`
	IsOnline  bool      `json:"is_online"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	MessageTypeText = iota
	MessageTypeEmoji
	MessageTypeImage
	MessageTypeSystem
)

type Message struct {
	ID        string     `json:"id"`
	IslandID  string     `json:"island_id"`
	SenderID  string     `json:"sender_id"`
	Content   string     `json:"content"`
	Type      int        `json:"type"`
	Timestamp time.Time  `json:"timestamp"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Island struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	OwnerID  string `json:"owner_id"`
	Capacity int    `json:"capacity"`
	Version  string `json:"version"`
}
