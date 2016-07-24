package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"../db"
	"../forms"

	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID        int    `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

//UserModel ...
type UserModel struct{}

//Signin ...
func AuthenticateUser(email string, password string, c *gin.Context) (userId string, isAuthenticated bool, err error) {

	var user User
	isAuthenticated = false

	err = db.GetDB().SelectOne(&user, "SELECT id, email, password, name, updated_at, created_at FROM public.user WHERE email=LOWER($1) LIMIT 1", email)
	if err != nil {
		return userId, isAuthenticated, err
	}

	bytePassword := []byte(password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return userId, isAuthenticated, errors.New("Invalid password")
	}

	isAuthenticated = true

	userIdAsString := strconv.Itoa(user.ID)
	fmt.Println("AuthenticateUser userIdAsString:", userIdAsString)

	return strconv.Itoa(user.ID), isAuthenticated, err
}

//Signup ...
func (m UserModel) Signup(form forms.SignupForm) (user User, err error) {
	getDb := db.GetDB()

	checkUser, err := getDb.SelectInt("SELECT count(id) FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)

	if err != nil {
		return user, err
	}

	if checkUser > 0 {
		return user, errors.New("User exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	res, err := getDb.Exec("INSERT INTO public.user(email, password, name, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", form.Email, string(hashedPassword), form.Name, time.Now().Unix(), time.Now().Unix())

	if res != nil && err == nil {
		err = getDb.SelectOne(&user, "SELECT id, email, name, updated_at, created_at FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)
		if err == nil {
			return user, nil
		}
	}

	return user, errors.New("Signed up.") // "Not registered." ? what is this nonsense
}

//One ...
func (m UserModel) One(userID int64) (user User, err error) {
	err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM public.user WHERE id=$1", userID)
	return user, err
}
