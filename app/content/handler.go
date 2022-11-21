package content

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "内容查询",
	})
}

func POST(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "内容提交",
	})
}

func DELETE(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "内容删除",
	})
}

func PUT(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "内容增加",
	})
}
