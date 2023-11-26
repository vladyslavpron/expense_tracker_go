package categoryHandler

import (
	"net/http"
	"tracker/core/logger"

	"github.com/gin-gonic/gin"
)

func (c *CategoryHandler) GetCategories(ctx *gin.Context) {

	categories, err := c.DB.Category.Query().All(ctx)
	if err != nil {
		logger.Err("failed getting categories: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, "Can't get categories")
	}
	logger.Log("completed categories search")

	ctx.JSON(http.StatusOK, categories)
}
