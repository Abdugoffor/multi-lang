package language_service

import (
	"errors"
	"project/helper"
	language_dto "project/module/admin/language/dto"
	language_model "project/module/admin/language/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LanguageService interface {
	All(ctx echo.Context) (helper.PaginatedResponse[language_dto.LanguageResponse], error)
	Show(ctx echo.Context, id int) (language_dto.LanguageResponse, error)
	Create(ctx echo.Context, request language_dto.Create) (language_dto.LanguageResponse, error)
	Update(ctx echo.Context, id int, request language_dto.Update) (language_dto.LanguageResponse, error)
	Delete(ctx echo.Context, id int) error
}

type languageService struct {
	db *gorm.DB
}

func NewLanguageService(db *gorm.DB) LanguageService {
	return &languageService{db}
}

func (s *languageService) All(ctx echo.Context) (helper.PaginatedResponse[language_dto.LanguageResponse], error) {
	var models []language_model.Language

	res, err := helper.Paginate(ctx, s.db, &models, 10)
	{
		if err != nil {
			return helper.PaginatedResponse[language_dto.LanguageResponse]{}, err
		}
	}

	var data []language_dto.LanguageResponse
	{
		for _, model := range models {
			data = append(data, language_dto.ToResponse(model))
		}
	}

	return helper.PaginatedResponse[language_dto.LanguageResponse]{Data: data, Meta: res.Meta}, nil
}

func (s *languageService) Show(ctx echo.Context, id int) (language_dto.LanguageResponse, error) {

	var model language_model.Language
	{
		if err := s.db.First(&model, id).Error; err != nil {
			return language_dto.LanguageResponse{}, err
		}
	}
	res := language_dto.ToResponse(model)
	return res, nil
}

func (s *languageService) Create(ctx echo.Context, request language_dto.Create) (language_dto.LanguageResponse, error) {

	var model language_model.Language
	{
		slug := helper.Slug(request.Name)

		// Slug bo‘yicha tilni qidiramiz
		if err := s.db.Where("slug = ?", slug).First(&model).Error; err != nil {
			// Agar yozuv topilmagan bo‘lsa — yaratamiz
			if errors.Is(err, gorm.ErrRecordNotFound) {
				model = language_model.Language{
					Name:     request.Name,
					Slug:     slug,
					IsActive: request.IsActive,
				}

				if err := s.db.Create(&model).Error; err != nil {
					return language_dto.LanguageResponse{}, err
				}
			} else {
				// Boshqa xatolik bo‘lsa — qaytaramiz
				return language_dto.LanguageResponse{}, err
			}
		}
	}

	res := language_dto.ToResponse(model)
	return res, nil
}

func (s *languageService) Update(ctx echo.Context, id int, request language_dto.Update) (language_dto.LanguageResponse, error) {
	var model language_model.Language
	{
		if err := s.db.First(&model, id).Error; err != nil {
			return language_dto.LanguageResponse{}, err
		}
	}

	model.Name = request.Name
	model.IsActive = request.IsActive
	model.Slug = helper.Slug(request.Name)

	if err := s.db.Save(&model).Error; err != nil {
		return language_dto.LanguageResponse{}, err
	}

	res := language_dto.ToResponse(model)
	return res, nil
}

func (s *languageService) Delete(ctx echo.Context, id int) error {
	var model language_model.Language
	{
		if err := s.db.First(&model, id).Error; err != nil {
			return err
		}
	}

	if err := s.db.Delete(&model).Error; err != nil {
		return err
	}

	return nil
}
