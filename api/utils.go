package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FailNoExist(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}

func FailNoExistPostForm(c *gin.Context, formName string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf("%s post form not present", formName),
	})
}
func FailNotAuthorized(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, gin.H{
		"error": message,
	})
}

func FailOnError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
