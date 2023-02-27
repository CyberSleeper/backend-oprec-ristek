package models

import (
	"github.com/CyberSleeper/backend-oprec-ristek/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type Post struct {
	gorm.Model
	Caption string `gorm:"" json:"caption"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}

func (b *Post) CreatePost() *Post {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllPosts() []Post {
	var Posts []Post
	db.Order("created_at desc").Find(&Posts)
	return Posts
}

func GetPostById(Id int64) (*Post, *gorm.DB) {
	var getPost Post
	db := db.Where("ID=?", Id).Find(&getPost)
	return &getPost, db
}

func DeletePost(Id int64) Post {
	var post Post
	db.Where("ID=?", Id).Delete(post)
	return post
}
