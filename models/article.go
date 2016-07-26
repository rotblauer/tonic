package models

import (
	"errors"
	"time"

	"../db"
	"../forms"
)

//Article ...
type Article struct {
	ID        int64    `db:"id, primarykey, autoincrement" json:"id"`
	UserID    int64    `db:"user_id" json:"userId"`
	Title     string   `db:"title" json:"title"`
	Content   string   `db:"content" json:"content"`
	UpdatedAt int64    `db:"updated_at" json:"updated_at"`
	CreatedAt int64    `db:"created_at" json:"created_at"`
	User      *JSONRaw `db:"user" json:"user"`
}

//ArticleModel ...
type ArticleModel struct{}

//Create ...
func (m ArticleModel) Create(userID int64, form forms.ArticleForm) (article Article, err error) {
	getDb := db.GetDB()

	// Ensure user exists.
	userModel := new(UserModel)
	checkUser, err := userModel.One(userID)
	if err != nil && checkUser.ID > 0 {
		return article, errors.New("User doesn't exist")
	}

	// Insert the article.
	_, err = getDb.Exec("INSERT INTO article(user_id, title, content, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", userID, form.Title, form.Content, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		return article, err
	}

	articleID, err := getDb.SelectInt("SELECT id FROM article WHERE user_id=$1 ORDER BY id DESC LIMIT 1", userID)
	// err = getDb.SelectOne(&article, "SELECT * FROM article WHERE user_id=$1 ORDER BY id DESC LIMIT 1", userID)
	article, err = m.One(articleID)
	return article, err
}

//One ...
func (m ArticleModel) One(id int64) (article Article, err error) {
	err = db.GetDB().SelectOne(&article, "SELECT a.id, a.user_id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.id=$1 GROUP BY a.id, a.title, a.content, a.updated_at, a.created_at, u.id, u.name, u.email LIMIT 1", id)
	return article, err
}

//All ...
func (m ArticleModel) All() (articles []Article, err error) {
	_, err = db.GetDB().Select(&articles, "SELECT a.id, a.user_id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM article a LEFT JOIN public.user u ON a.user_id = u.id GROUP BY a.id, a.title, a.content, a.updated_at, a.created_at, u.id, u.name, u.email ORDER BY a.id DESC")
	return articles, err
}

//Update ...
func (m ArticleModel) Update(userID int64, id int64, form forms.ArticleForm) (article Article, err error) {
	a, err := m.One(id)

	// Ensure owner of the article is the one updating it.
	if a.UserID != userID {
		return article, errors.New("Not authenticated to update this article.")
	}

	if err != nil {
		return article, errors.New("Article not found")
	}

	_, err = db.GetDB().Exec("UPDATE article SET title=$1, content=$2, updated_at=$3 WHERE id=$4", form.Title, form.Content, time.Now().Unix(), id)
	if err != nil {
		return article, err
	}

	article, err = m.One(id)
	return article, err
}

//Delete ...
func (m ArticleModel) Delete(userID, id int64) (err error) {
	a, err := m.One(id)

	// Ensure owner of the article is the one updating it.
	if a.UserID != userID {
		return errors.New("Not authenticated to delete this article.")
	}

	if err != nil {
		return errors.New("Article not found")
	}

	_, err = db.GetDB().Exec("DELETE FROM article WHERE id=$1", id)

	return err
}
