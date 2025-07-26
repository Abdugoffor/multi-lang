package language_dto

import (
	"project/helper"
	language_model "project/module/admin/language/model"
)

type Create struct {
	Name     string `json:"name" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type Update struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type LanguageResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToResponse(lang language_model.Language) LanguageResponse {
	return LanguageResponse{
		ID:        lang.ID,
		Name:      lang.Name,
		Slug:      lang.Slug,
		IsActive:  lang.IsActive,
		CreatedAt: helper.FormatDate(lang.CreatedAt),
		UpdatedAt: helper.FormatDate(lang.UpdatedAt),
	}
}
