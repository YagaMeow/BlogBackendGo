package system

import (
	"blog-backend/dao/system"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct{}

func (a *ArticleApi) CreateArticle(c *gin.Context) {
	var article system.Article
	c.ShouldBindJSON(&article)

	if err := articleService.CreateArticle(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": article,
		})
	}
}
