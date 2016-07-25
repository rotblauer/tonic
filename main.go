package main

import (
	"./conf"
	"./controllers"
	"./db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(conf.CORSMiddleware())

	db.Init()

	r.StaticFile("/", "./public/index.html") // Root page. All template rendering happens client side.
	r.Static("/vendor", "./public/vendor")
	r.Static("/assets", "./public/assets")

	// the jwt middleware
	authMiddleware := conf.InitJWTMiddlewareConf()

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/u/signup", user.Signup)
		v1.POST("/signin", authMiddleware.LoginHandler)

		/*** START Article ***/
		article := new(controllers.ArticleController)

		v1.GET("/a", article.All)
		v1.GET("/a/:id", article.One)

		v1.DELETE("/a/:id", article.Delete)

		auth := v1.Group("/auth")
		auth.Use(authMiddleware.MiddlewareFunc())
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		// AUTHY Article
		auth.POST("/a", article.Create)
		auth.PUT("/a/:id", article.Update)
		auth.DELETE("/a/:id", article.Delete)

	}

	r.Run(":9000")
}
