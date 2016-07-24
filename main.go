package main

import (
	"fmt"
	"net/http"

	"github.com/Massad/gin-boilerplate/controllers"
	"github.com/Massad/gin-boilerplate/db"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// Lame comment.
type Thingey struct {
	HasHoles bool
	HashOles string
}

func main() {
	r := gin.Default()

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("gin-boilerplate-session", store))

	r.Use(CORSMiddleware())

	// Thanks to this glob, we can grab and grok all the templatez.
	r.LoadHTMLGlob("templates/**/*")

	r.Static("/public", "./public")

	db.Init()

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/u/signin", user.Signin)
		v1.POST("/u/signup", user.Signup)
		v1.GET("/u/signout", user.Signout)

		/*** START Article ***/
		article := new(controllers.ArticleController)

		v1.POST("/a", article.Create)
		v1.GET("/a", article.All)
		v1.GET("/a/:id", article.One)
		v1.PUT("/a/:id", article.Update)
		v1.DELETE("/a/:id", article.Delete)
	}
	app := r.Group("/app")
	{
		app.GET("/posts/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // Notice that we're calling .tmpl, _not_ .html
				"title": "Posts",
				"abc":   "123", // we'll pass this to a template nested inside posts/index
				"etc":   "etcetc",
				"user":  Thingey{HasHoles: true, HashOles: "yes please"},
			})
		})
		app.GET("/users/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
				"title": "Users",
			})
		})
	}

	r.Run(":9000")
}
