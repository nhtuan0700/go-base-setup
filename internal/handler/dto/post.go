package dto

import "mime/multipart"

type Post struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type CreatePostRequest struct {
	Title string          `form:"title" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

type CreatePostResponse struct {
}
