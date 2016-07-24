package controllers

import (
	"../forms"
	"../models"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

//Signup ...
func (ctrl UserController) Signup(c *gin.Context) {
	var signupForm forms.SignupForm

	if c.BindJSON(&signupForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signupForm})
		c.Abort()
		return
	}

	user, err := userModel.Signup(signupForm)

	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if user.ID > 0 {

		// TODO: JWT

		c.JSON(200, gin.H{"message": "Success signup", "user": user})
	} else {
		c.JSON(406, gin.H{"message": "Could not signup this user", "error": err.Error()})
	}

}

//Signout ...
func (ctrl UserController) Signout(c *gin.Context) {

	// TODO: JWT

	c.JSON(200, gin.H{"message": "Signed out..."})
}
