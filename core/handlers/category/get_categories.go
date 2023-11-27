package categoryHandler

import (
	"context"
	"net/http"
	"tracker/core/logger"
	"tracker/ent"

	"github.com/gin-gonic/gin"
)

func (c *CategoryHandler) GetCategoriesHandler(ctx *gin.Context) {

	categories, err := c.GetCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Can't get categories")
	}

	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryHandler) GetCategories(ctx context.Context) ([]*ent.Category, error) {

	categories, err := c.DB.Category.Query().All(ctx)
	if err != nil {
		logger.Err("failed getting categories: " + err.Error())
		return nil, err
	}
	logger.Log("completed categories search")

	return categories, nil
}
