package route

import (
	"github.com/Fadhelbulloh/local-elastic/model"
	"github.com/Fadhelbulloh/local-elastic/service"
	"github.com/Fadhelbulloh/local-elastic/util"
	"github.com/gin-gonic/gin"
)

func BasicService(router *gin.Engine) {
	basic := router.Group("/basic")
	{
		basic.POST("/scroll", func(c *gin.Context) {
			var param model.ParamScroll
			if util.ErrHandler(500, c, c.BindJSON(&param)) {
				return
			}

			response := service.GetBasicScroll(param)
			if util.ErrorHandleResponse(c, response) {
				return
			}

			c.JSON(200, response)
		})

		basic.POST("/search-scroll", func(c *gin.Context) {
			var param model.ParamCatalog
			if util.ErrHandler(500, c, c.BindJSON(&param)) {
				return
			}

			response := service.GetBasicSearchScroll(param)
			if util.ErrorHandleResponse(c, response) {
				return
			}
			c.JSON(200, response)

		})

	}
}
