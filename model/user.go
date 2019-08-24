package model

type User struct {
	UserID    int    `gorm:"column:user_id;type:binint(20);primary_key"`
	AccountID string `gorm:"column:account_id;type:varchar(30);unique_index"`
	Password  string `gorm:"column:pwd;type:varchar(20)"`
	UserName  string `gorm:"column:user_name;type:varchar(30)"`
	Mail      string `gorm:"column:mail;type:varchar(80)"`
	Phone     string `gorm:"column:phone;type:varchar(20)"`
	FBID      string `gorm:"column:fb_id;type:varchar(20)" `
	IsDeleted bool   `gorm:"column:is_deleted;type:tinyint(1)"`
}
