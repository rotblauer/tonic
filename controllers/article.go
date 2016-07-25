package controllers

import (
	"errors"
	"strconv"

	"../forms"
	"../models"

	"github.com/gin-gonic/gin"
)

//ArticleController ...
type ArticleController struct{}

var articleModel = new(models.ArticleModel)

// Key "userID" is made available from the JWT middleware
// which extracts it from the claims payload 'uid'.
func getUserIdFromContext(c *gin.Context) (uid int64, err error) {
	// userID, _ := getUserIDFromJWTClaim(c)
	userid, exists := c.Get("userID")
	if !exists {
		return uid, errors.New("Unauthorized")
	}

	uid, err = strconv.ParseInt(userid.(string), 10, 64)
	return uid, err
}

//Create ...
func (ctrl ArticleController) Create(c *gin.Context) {

	userID, err := getUserIdFromContext(c)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
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

	data, err := articleModel.All()

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the articles", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl ArticleController) One(c *gin.Context) {

	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := articleModel.One(id)
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

	userID, err := getUserIdFromContext(c)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
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

	userID, err := getUserIdFromContext(c)
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
		c.Abort()
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
