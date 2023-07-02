package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weibo/service"
	"weibo/tool"
)

func FavoriteNote(c *gin.Context) {
	var favoriteNoteService service.FavoriteService

	claims, _ := tool.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&favoriteNoteService); err == nil {
		res := favoriteNoteService.FavoriteNote(c.Request.Context(), claims.UserId)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
