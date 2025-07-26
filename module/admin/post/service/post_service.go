package post_service

import (
	"project/helper"
	post_dto "project/module/admin/post/dto"
	post_model "project/module/admin/post/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostService interface {
	All(ctx echo.Context) (helper.PaginatedResponse[post_dto.PostResponse], error)
	Show(ctx echo.Context, id int) (*post_dto.PostResponse, error)
	Create(ctx echo.Context, request post_dto.Create) (*post_dto.PostResponse, error)
	Update(ctx echo.Context, id int, request post_dto.Update) (*post_dto.PostResponse, error)
	Delete(ctx echo.Context, id int) error
}

type postService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) PostService {
	return &postService{db: db}
}

func (p *postService) All(ctx echo.Context) (helper.PaginatedResponse[post_dto.PostResponse], error) {
	var models []post_model.Post

	res, err := helper.Paginate(ctx, p.db, &models, 10)
	{
		if err != nil {
			return helper.PaginatedResponse[post_dto.PostResponse]{}, err
		}
	}

	var data []post_dto.PostResponse
	{
		for _, model := range models {
			data = append(data, post_dto.ToResponse(model))
		}
	}

	return helper.PaginatedResponse[post_dto.PostResponse]{Data: data, Meta: res.Meta}, nil
}

func (p *postService) Show(ctx echo.Context, id int) (*post_dto.PostResponse, error) {

	var model post_model.Post
	{
		if err := p.db.Where("id = ?", id).First(&model).Error; err != nil {
			return nil, err
		}
	}

	res := post_dto.ToResponse(model)

	return &res, nil
}

func (p *postService) Create(ctx echo.Context, request post_dto.Create) (*post_dto.PostResponse, error) {

	slug := helper.Slug(request.Title["default"])

	model := post_model.Post{
		Title:       request.Title,
		Description: request.Description,
		Content:     request.Content,
		Slug:        slug,
		IsActive:    request.IsActive,
	}

	if err := p.db.Create(&model).Error; err != nil {
		return nil, err
	}

	res := post_dto.ToResponse(model)

	return &res, nil

}

func (p *postService) Update(ctx echo.Context, id int, request post_dto.Update) (*post_dto.PostResponse, error) {
	var model post_model.Post
	{
		if err := p.db.Where("id = ?", id).First(&model).Error; err != nil {
			return nil, err
		}
	}

	model.Title = request.Title
	model.Description = request.Description
	model.Content = request.Content
	model.IsActive = request.IsActive

	if err := p.db.Save(&model).Error; err != nil {
		return nil, err
	}
	res := post_dto.ToResponse(model)

	return &res, nil
}

func (p *postService) Delete(ctx echo.Context, id int) error {

	var model post_model.Post
	{
		if err := p.db.Where("id = ?", id).First(&model).Error; err != nil {
			return err
		}
	}

	return p.db.Delete(&model).Error
}
