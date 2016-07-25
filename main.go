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
	authMiddleware := conf.InitJWTMiddlewareConf()

	db.Init()

	r.Static("/public", "./public")
	r.StaticFile("/", "./public/index.html") // Root page. All template rendering happens client side.

	v1 := r.Group("/v1")
	{

		/*** SET UP AUTH ***/
		auth := v1.Group("/auth")
		auth.Use(authMiddleware.MiddlewareFunc())
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		/*** START USER ***/
		user := new(controllers.UserController)
		v1.POST("/u/signup", user.Signup) // Signup -> signin transfer to be handled clientside.
		v1.POST("/signin", authMiddleware.LoginHandler)

		/*** START Article ***/
		article := new(controllers.ArticleController)

		// Public.
		v1.GET("/a", article.All)
		v1.GET("/a/:id", article.One)

		// Authy.
		auth.POST("/a", article.Create)
		auth.PUT("/a/:id", article.Update)
		auth.DELETE("/a/:id", article.Delete)

	}

	r.Run(":9000")
}
