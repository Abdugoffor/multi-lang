package language_handler

import (
	"log"
	"net/http"
	language_dto "project/module/admin/language/dto"
	language_service "project/module/admin/language/service"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type languageHandler struct {
	db              *gorm.DB
	log             *log.Logger
	languageService language_service.LanguageService
}

func NewLanguageHandler(router *echo.Group, db *gorm.DB, log *log.Logger) *languageHandler {
	handler := &languageHandler{
		db:              db,
		log:             log,
		languageService: language_service.NewLanguageService(db),
	}

	languageGroup := router.Group("/languages")
	{
		languageGroup.GET("", handler.All)
		languageGroup.GET("/:id", handler.Show)
		languageGroup.POST("", handler.Create)
		languageGroup.PUT("/:id", handler.Update)
		languageGroup.DELETE("/:id", handler.Delete)
	}

	return handler
}

func (h *languageHandler) All(ctx echo.Context) error {

	data, err := h.languageService.All(ctx)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(http.StatusOK, data)
}

func (h *languageHandler) Show(ctx echo.Context) error {

	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	data, err := h.languageService.Show(ctx, id)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(http.StatusOK, data)
}

func (h *languageHandler) Create(ctx echo.Context) error {

	var request language_dto.Create
	{
		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	data, err := h.languageService.Create(ctx, request)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(http.StatusOK, data)
}

func (h *languageHandler) Update(ctx echo.Context) error {

	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	var request language_dto.Update
	{
		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	data, err := h.languageService.Update(ctx, id, request)
	{
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
	}

	return ctx.JSON(http.StatusOK, data)
}

func (h *languageHandler) Delete(ctx echo.Context) error {

	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	{
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
	}

	if err := h.languageService.Delete(ctx, id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Language deleted successfully"})
}
