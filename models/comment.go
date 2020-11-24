package models

type Comment struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserId uint   `gorm:"not null" json:"user_id"`
	PostId uint   `gorm:"not null" json:"post_id"`
	Text   string `gorm:"not null;varchar(200)" form:"text" json:"text" validate:"required"`
	Like   uint   `gorm:"not null" json:"like"`
	UnLike uint   `gorm:"not null" json:"unlike"`
}

func (*Comment) All(teamId int) (comments []Comment) {
	DB.Select("user_id, text, like, unlike").Order("id").Find(&comments)
	return comments
}

func (*Comment) Create(comment *Comment) (err error) {
	if err := DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (*Comment) Show(id string) (comment Comment) {
	if err := DB.Select("user_id, text, like, unlike").First(&comment, id).Error; err != nil {
		return
	}

	return comment
}

func (c *Comment) Delete(id string) (err error) {
	if DB.First(&c, id).Error != nil {
		return
	}
	if err := DB.Unscoped().Delete(&c).Error; err != nil {
		return err
	}

	return nil
}
