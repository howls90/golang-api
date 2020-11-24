package models

import (
	"api/libs"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostEntity interface {
	All() (posts []Post, err error)
	Create(post *Post) (err error)
	Show(id uint64) (post *Post, err error)
	Delete(id uint64) (err error)
}

type Post struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	UserId      uint64    `gorm:"" json:"user_id"`
	Title       string    `gorm:"not null;varchar(50)" form:"title" json:"title" validate:"required"`
	Description string    `gorm:"not null" form:"description" json:"description" validate:"required"`
	Comments    []Comment `gorm:"foreignKey:PostId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	// to := []string{
	// 	"marti.floriach@awl.co.jp",
	// }
	// message := "This is a really unimaginative message, I know."
	// libs.SendEmail(to, message)

	post.chachePosts(post)

	return
}

func (*Post) All() (posts []Post, err error) {
	tot := libs.RedisGet("posts")

	if err := json.Unmarshal([]byte(tot), &posts); err != nil {
		log.Error(err)
		if err := DB.Preload("Comments").Select("id, user_id, title, description").Order("id").Find(&posts).Error; err != nil {
			return nil, err
		}
	}

	return posts, nil
}

func (*Post) Create(post *Post) (err error) {
	if err := DB.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (*Post) Show(id string) (post Post) {
	if err := DB.Preload("Comments").Select("user_id, title, description").First(&post, id).Error; err != nil {
		return post
	}

	return post
}

func (p *Post) Delete(id string) (err error) {
	if DB.First(&p, id).Error != nil {
		return
	}
	if err := DB.Unscoped().Delete(&p).Error; err != nil {
		return err
	}

	return nil
}

func (*Post) chachePosts(post *Post) error {
	posts := []Post{}

	old := libs.RedisGet("posts")
	if err := json.Unmarshal([]byte(old), &posts); err != nil {
		log.Error(err)
		return err
	}
	posts = append(posts, *post)
	out, err := json.Marshal(posts)
	if err != nil {
		log.Error(err)
		return err
	}

	libs.RedisSet("posts", out)

	return nil
}
