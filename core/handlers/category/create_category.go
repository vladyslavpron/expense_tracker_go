package categoryHandler

import (
	"net/http"
	"tracker/core/logger"

	"github.com/gin-gonic/gin"
)

func (c *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var dto CreateCategoryDTO
	if err := ctx.BindJSON(&dto); err != nil {
		logger.Err("ERR: CreateCategory func" + string(err.Error()))
		ctx.JSON(http.StatusBadRequest, "Wrong DTO")
		return
	}

	category, err := c.DB.Category.Create().SetTitle(dto.Title).Save(ctx)
	if err != nil {
		logger.Err("failed creating category: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, "Can't create category")
	}
	logger.Log("category was created: " + category.String())

	ctx.JSON(http.StatusCreated, category)
}

type CreateCategoryDTO struct {
	Title string
}
