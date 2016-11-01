package routers

import (
	// 3rd party packages
	"github.com/gin-gonic/gin"

	// Local packages
	"jaha-api/controllers"
	"jaha-api/middlewares"
)

/**
 *	Returns default router instance.
 *
 *	@return *gin.Engine
 */
func GetDefaultRouter() *gin.Engine {
	var router *gin.Engine

	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Cors())

	attachDefaultRoutes(router)

	return router
}

/**
 *	Attaches route definitons to default router.
 *
 *	@param router gin.Engine - Router instance.
 *
 *	@return void
 */
func attachDefaultRoutes(router *gin.Engine) {
	router.NoRoute(controllers.DefaultController().MissingRoute)

	v1 := router.Group("v1")
	{
		category := v1.Group("categories")
		{
			category.GET("", controllers.CategoriesController().Index)
			category.POST("", controllers.CategoriesController().Create)

			category.GET(":uuid", controllers.CategoriesController().Show)
			category.PATCH(":uuid", controllers.CategoriesController().Update)
			category.DELETE(":uuid", controllers.CategoriesController().Destroy)
			category.PUT(":uuid", controllers.CategoriesController().Restore)
		}

		statement := v1.Group("statements")
		{
			statement.GET("", controllers.StatementsController().Index)
			statement.POST("", controllers.StatementsController().Create)

			statement.GET(":uuid", controllers.StatementsController().Show)
			statement.PATCH(":uuid", controllers.StatementsController().Update)
			statement.DELETE(":uuid", controllers.StatementsController().Destroy)
			statement.PUT(":uuid", controllers.StatementsController().Restore)
		}
	}

}
