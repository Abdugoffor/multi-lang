package post_handler

import (
	"log"
	"net/http"
	post_dto "project/module/admin/post/dto"
	post_service "project/module/admin/post/service"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type postHandler struct {
	db          *gorm.DB
	log         *log.Logger
	postService post_service.PostService
}

func NewPostHandler(router *echo.Group, db *gorm.DB, log *log.Logger) *postHandler {
	handler := &postHandler{
		db:          db,
		log:         log,
		postService: post_service.NewPostService(db),
	}

	postGroup := router.Group("/posts")
	{
		postGroup.GET("", handler.All)
		postGroup.GET("/:id", handler.Show)
		postGroup.POST("", handler.Create)
		postGroup.PUT("/:id", handler.Update)
		postGroup.DELETE("/:id", handler.Delete)
	}

	return handler
}

func (h *postHandler) All(ctx echo.Context) error {

	data, err := h.postService.All(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, data)
}

func (h *postHandler) Show(ctx echo.Context) error {

	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	data, err := h.postService.Show(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, data)
}

func (h *postHandler) Create(ctx echo.Context) error {

	var request post_dto.Create
	{
		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	data, err := h.postService.Create(ctx, request)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(http.StatusOK, data)
}

func (h *postHandler) Update(ctx echo.Context) error {

	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	var request post_dto.Update
	{
		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	data, err := h.postService.Update(ctx, id, request)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(http.StatusOK, data)
}

func (h *postHandler) Delete(ctx echo.Context) error {

	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	if err := h.postService.Delete(ctx, id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Post deleted successfully"})
}
