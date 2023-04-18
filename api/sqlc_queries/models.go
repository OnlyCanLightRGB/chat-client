// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package sqlc_queries

import (
	"encoding/json"
	"time"
)

type AuthUser struct {
	ID          int32
	Password    string
	LastLogin   time.Time
	IsSuperuser bool
	Username    string
	FirstName   string
	LastName    string
	Email       string
	IsStaff     bool
	IsActive    bool
	DateJoined  time.Time
}

type AuthUserManagement struct {
	ID        int32
	UserID    int32
	RateLimit int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ChatLog struct {
	ID        int32
	Session   json.RawMessage
	Question  json.RawMessage
	Answer    json.RawMessage
	CreatedAt time.Time
}

type ChatMessage struct {
	ID              int32
	Uuid            string
	ChatSessionUuid string
	Role            string
	Content         string
	Score           float64
	UserID          int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       int32
	UpdatedBy       int32
	IsDeleted       bool
	TokenCount      int32
	Raw             json.RawMessage
}

type ChatModel struct {
	ID            int32
	Name          string
	Label         string
	IsDefault     bool
	Url           string
	ApiAuthHeader string
	ApiAuthKey    string
}

type ChatPrompt struct {
	ID              int32
	Uuid            string
	ChatSessionUuid string
	Role            string
	Content         string
	Score           float64
	UserID          int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       int32
	UpdatedBy       int32
	IsDeleted       bool
	TokenCount      int32
}

type ChatSession struct {
	ID          int32
	UserID      int32
	Uuid        string
	Topic       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Active      bool
	Model       string
	MaxLength   int32
	Temperature float64
	TopP        float64
	MaxTokens   int32
	Debug       bool
}

type ChatSnapshot struct {
	ID           int32
	Uuid         string
	UserID       int32
	Title        string
	Summary      string
	Tags         json.RawMessage
	Conversation json.RawMessage
	CreatedAt    time.Time
}

type JwtSecret struct {
	ID       int32
	Name     string
	Secret   string
	Audience string
}

type MessageComment struct {
	ID          int32
	Uuid        string
	MessageUuid string
	UserID      int32
	Content     string
	CreatedAt   time.Time
}

type UserActiveChatSession struct {
	ID              int32
	UserID          int32
	ChatSessionUuid string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
