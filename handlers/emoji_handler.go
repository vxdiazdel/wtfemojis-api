package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerContext) GetEmojis(c *gin.Context) {
	name := c.Query("name")

	emojis, err := h.Store().GetEmojis(h.Ctx(), name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": emojis})
}
