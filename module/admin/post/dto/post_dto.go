package post_dto

import (
	"project/helper"
	post_model "project/module/admin/post/model"
)

type Create struct {
	Title       post_model.JSONBMap `json:"title" validate:"required,jsonb_required" gorm:"type:jsonb"`
	Description post_model.JSONBMap `json:"description" validate:"required,jsonb_required" gorm:"type:jsonb"`
	Content     post_model.JSONBMap `json:"content" validate:"required,jsonb_required" gorm:"type:jsonb"`
	IsActive    bool                `json:"is_active"`
}

type Update struct {
	Title       post_model.JSONBMap `json:"title" validate:"jsonb_required" gorm:"type:jsonb"`
	Description post_model.JSONBMap `json:"description" validate:"jsonb_required" gorm:"type:jsonb"`
	Content     post_model.JSONBMap `json:"content" validate:"jsonb_required" gorm:"type:jsonb"`
	IsActive    bool                `json:"is_active"`
}

type PostResponse struct {
	ID          int64               `json:"id"`
	Title       post_model.JSONBMap `json:"title"`
	Description post_model.JSONBMap `json:"description"`
	Content     post_model.JSONBMap `json:"content"`
	View        int64               `json:"view"`
	Slug        string              `json:"slug"`
	IsActive    bool                `json:"is_active"`
	CreatedAt   string              `json:"created_at"`
	UpdatedAt   string              `json:"updated_at"`
}

func ToResponse(post post_model.Post) PostResponse {
	return PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		Content:     post.Content,
		View:        post.View,
		Slug:        post.Slug,
		IsActive:    post.IsActive,
		CreatedAt:   helper.FormatDate(post.CreatedAt),
		UpdatedAt:   helper.FormatDate(post.UpdatedAt),
	}
}
