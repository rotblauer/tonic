package controllers

import (
	"strconv"

	"../forms"
	auth "../jwt"
	"../models"

	"github.com/gin-gonic/gin"
)

//ArticleController ...
type ArticleController struct{}

var articleModel = new(models.ArticleModel)

//Create ...
func (ctrl ArticleController) Create(c *gin.Context) {

	claimUserId := auth.ExtractClaims(c)["id"].(string)

	// extraPayload, exists := c.Get("JWT_PAYLOAD")
	// if !exists {
	// 	fmt.Println("error getting claim user id from payload...")
	// }

	// fmt.Println("article model extraPayload:", extraPayload)

	// castedClaims := extraPayload.(map[string]string)

	// claimUserId := castedClaims["id"]

	// // claims := auth.ExtractClaims(c)
	// if _, exists := c.Get("JWT_PAYLOAD"); !exists {
	// 	emptyClaims := make(jwt.MapClaims)

	// 	fmt.Println("There were no claims in the JWT_PAYLOAD", emptyClaims)
	// }

	// jwtClaims, _ := c.Get("JWT_PAYLOAD")

	// jc := jwtClaims.(jwt.MapClaims)

	// fmt.Println("Extreacted claims in model like in jwt: ", jc)

	// claims := jc

	// fmt.Println("extracted claims: ", claims)

	// claimUserId := claims["id"]

	// fmt.Println("creating article for userId:", claimUserId)

	// if claimUserId == nil {
	// 	c.JSON(403, gin.H{"message": "There is no userId"})
	// }

	userID, err := strconv.ParseInt(claimUserId, 10, 64)
	if err != nil {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	var articleForm forms.ArticleForm

	if c.BindJSON(&articleForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": articleForm})
		c.Abort()
		return
	}

	articleID, err := articleModel.Create(userID, articleForm)

	if articleID > 0 && err != nil {
		c.JSON(406, gin.H{"message": "Article could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Article created", "id": articleID})
}

//All ...
func (ctrl ArticleController) All(c *gin.Context) {

	claims := auth.ExtractClaims(c)
	claimUserId := claims["id"]
	if claimUserId == nil {
		c.JSON(403, gin.H{"message": "Please login first"})
	}

	userID, err := strconv.ParseInt(claimUserId.(string), 10, 64)
	if err != nil {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	data, err := articleModel.All(userID)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the articles", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl ArticleController) One(c *gin.Context) {

	claims := auth.ExtractClaims(c)
	claimUserId := claims["id"]
	if claimUserId == nil {
		c.JSON(403, gin.H{"message": "Please login first"})
	}

	userID, err := strconv.ParseInt(claimUserId.(string), 10, 64)
	if err != nil {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := articleModel.One(userID, id)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Update ...
func (ctrl ArticleController) Update(c *gin.Context) {

	claims := auth.ExtractClaims(c)
	claimUserId := claims["id"]
	if claimUserId == nil {
		c.JSON(403, gin.H{"message": "Please login first"})
	}

	userID, err := strconv.ParseInt(claimUserId.(string), 10, 64)
	if err != nil {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		var articleForm forms.ArticleForm

		if c.BindJSON(&articleForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": articleForm})
			c.Abort()
			return
		}

		err := articleModel.Update(userID, id, articleForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Article could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Article updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

//Delete ...
func (ctrl ArticleController) Delete(c *gin.Context) {

	claims := auth.ExtractClaims(c)
	claimUserId := claims["id"]
	if claimUserId == nil {
		c.JSON(403, gin.H{"message": "Please login first"})
	}

	userID, err := strconv.ParseInt(claimUserId.(string), 10, 64)
	if err != nil {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		err := articleModel.Delete(userID, id)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Article could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Article deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
