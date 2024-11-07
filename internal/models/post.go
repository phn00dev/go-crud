package models

import "time"

type Post struct {
	Id        int       `json:"id"`
	PostTitle string    `json:"post_title"`
	PostSlug  string    `json:"post_slug"`
	PostDesc  string    `json:"post_desc"`
	PostImage string    `json:"post_image"`
	ViewCount int       `json:"view_count"`
	CreatedAt time.Time `json:"created_at"`
	UserId    int       `json:"user_id"`
}
