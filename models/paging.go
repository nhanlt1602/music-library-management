package models

type PagingRequest struct {
	Page         int    `json:"page"`
	Size         int    `json:"size"`
	Sort         string `json:"sort"`
	PagingIgnore bool   `json:"paging_ignore"`
}
