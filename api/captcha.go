package api

import (
	"github.com/gin-gonic/gin"
)

func (api *Api) GetCaptcha(c *gin.Context) {
	id, b64s, _ := api.Captcha.Generate()
	c.JSON(200, gin.H{
		"id":      id,
		"captcha": b64s,
	})
}
