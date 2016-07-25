package main

import (
	"net/http"

	// "gopkg.in/appleboy/gin-jwt.v2"

	// "./conf"
	"./conf"
	"./controllers"
	"./db"
	// "./models"

	"github.com/gin-gonic/gin"
)

// Lame comment.
type Thingey struct {
	HasHoles bool
	HashOles string
}

func main() {
	r := gin.Default()

	r.Use(conf.CORSMiddleware())

	db.Init()

	// Thanks to this glob, we can grab and grok all the templatez.
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/public", "./public")

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

	app := r.Group("/")
	{
		app.GET("posts/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // Notice that we're calling .tmpl, _not_ .html
				"title": "Posts",
				"abc":   "123", // we'll pass this to a template nested inside posts/index
				"etc":   "etcetc",
				"user":  Thingey{HasHoles: true, HashOles: "yes please"},
			})
		})
		app.GET("users/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
				"title": "Users",
			})
		})
	}

	r.Run(":9000")
}
