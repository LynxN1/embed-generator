// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package pgmodel

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/sqlc-dev/pqtype"
)

type CustomBot struct {
	ID                      string
	GuildID                 string
	ApplicationID           string
	Token                   string
	PublicKey               string
	UserID                  string
	UserName                string
	UserDiscriminator       string
	UserAvatar              sql.NullString
	HandledFirstInteraction bool
	CreatedAt               time.Time
	TokenInvalid            bool
	GatewayStatus           string
	GatewayActivityType     sql.NullInt16
	GatewayActivityName     sql.NullString
	GatewayActivityState    sql.NullString
	GatewayActivityUrl      sql.NullString
}

type CustomCommand struct {
	ID                 string
	GuildID            string
	Name               string
	Description        string
	Enabled            bool
	Parameters         json.RawMessage
	Actions            json.RawMessage
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeployedAt         sql.NullTime
	DerivedPermissions pqtype.NullRawMessage
	LastUsedAt         time.Time
}

type EmbedLink struct {
	ID             string
	Url            string
	ThemeColor     sql.NullString
	OgTitle        sql.NullString
	OgSiteName     sql.NullString
	OgDescription  sql.NullString
	OgImage        sql.NullString
	OeType         sql.NullString
	OeAuthorName   sql.NullString
	OeAuthorUrl    sql.NullString
	OeProviderName sql.NullString
	OeProviderUrl  sql.NullString
	TwCard         sql.NullString
	ExpiresAt      sql.NullTime
	CreatedAt      time.Time
}

type Entitlement struct {
	ID              string
	UserID          sql.NullString
	GuildID         sql.NullString
	UpdatedAt       time.Time
	Deleted         bool
	SkuID           string
	StartsAt        sql.NullTime
	EndsAt          sql.NullTime
	Consumed        bool
	ConsumedGuildID sql.NullString
}

type Image struct {
	ID              string
	UserID          string
	GuildID         sql.NullString
	FileHash        string
	FileName        string
	FileSize        int32
	FileContentType string
	S3Key           string
}

type KvEntry struct {
	Key       string
	GuildID   string
	Value     string
	ExpiresAt sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageActionSet struct {
	ID                 string
	MessageID          string
	SetID              string
	Actions            json.RawMessage
	DerivedPermissions pqtype.NullRawMessage
	LastUsedAt         time.Time
	Ephemeral          bool
}

type SavedMessage struct {
	ID          string
	CreatorID   string
	GuildID     sql.NullString
	UpdatedAt   time.Time
	Name        string
	Description sql.NullString
	Data        json.RawMessage
}

type ScheduledMessage struct {
	ID             string
	CreatorID      string
	GuildID        string
	ChannelID      string
	MessageID      sql.NullString
	SavedMessageID string
	Name           string
	Description    sql.NullString
	CronExpression sql.NullString
	OnlyOnce       bool
	StartAt        time.Time
	EndAt          sql.NullTime
	NextAt         time.Time
	Enabled        bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CronTimezone   sql.NullString
}

type Session struct {
	TokenHash   string
	UserID      string
	GuildIds    []string
	AccessToken string
	CreatedAt   time.Time
	ExpiresAt   time.Time
}

type SharedMessage struct {
	ID        string
	CreatedAt time.Time
	ExpiresAt time.Time
	Data      json.RawMessage
}

type User struct {
	ID            string
	Name          string
	Discriminator string
	Avatar        sql.NullString
	IsTester      bool
}
