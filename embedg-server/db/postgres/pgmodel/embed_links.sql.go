// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: embed_links.sql

package pgmodel

import (
	"context"
	"database/sql"
	"time"
)

const getEmbedLink = `-- name: GetEmbedLink :one
SELECT id, url, theme_color, og_title, og_site_name, og_description, og_image, oe_type, oe_author_name, oe_author_url, oe_provider_name, oe_provider_url, tw_card, expires_at, created_at FROM embed_links WHERE id = $1
`

func (q *Queries) GetEmbedLink(ctx context.Context, id string) (EmbedLink, error) {
	row := q.db.QueryRowContext(ctx, getEmbedLink, id)
	var i EmbedLink
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.ThemeColor,
		&i.OgTitle,
		&i.OgSiteName,
		&i.OgDescription,
		&i.OgImage,
		&i.OeType,
		&i.OeAuthorName,
		&i.OeAuthorUrl,
		&i.OeProviderName,
		&i.OeProviderUrl,
		&i.TwCard,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const insertEmbedLink = `-- name: InsertEmbedLink :one
INSERT INTO embed_links (
    id,
    url,
    theme_color,
    og_title,
    og_site_name,
    og_description,
    og_image,
    oe_type,
    oe_author_name,
    oe_author_url,
    oe_provider_name,
    oe_provider_url,
    tw_card,
    expires_at,
    created_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15
) RETURNING id, url, theme_color, og_title, og_site_name, og_description, og_image, oe_type, oe_author_name, oe_author_url, oe_provider_name, oe_provider_url, tw_card, expires_at, created_at
`

type InsertEmbedLinkParams struct {
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

func (q *Queries) InsertEmbedLink(ctx context.Context, arg InsertEmbedLinkParams) (EmbedLink, error) {
	row := q.db.QueryRowContext(ctx, insertEmbedLink,
		arg.ID,
		arg.Url,
		arg.ThemeColor,
		arg.OgTitle,
		arg.OgSiteName,
		arg.OgDescription,
		arg.OgImage,
		arg.OeType,
		arg.OeAuthorName,
		arg.OeAuthorUrl,
		arg.OeProviderName,
		arg.OeProviderUrl,
		arg.TwCard,
		arg.ExpiresAt,
		arg.CreatedAt,
	)
	var i EmbedLink
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.ThemeColor,
		&i.OgTitle,
		&i.OgSiteName,
		&i.OgDescription,
		&i.OgImage,
		&i.OeType,
		&i.OeAuthorName,
		&i.OeAuthorUrl,
		&i.OeProviderName,
		&i.OeProviderUrl,
		&i.TwCard,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
