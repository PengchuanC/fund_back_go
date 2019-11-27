package models

import (
	"fund_back_go/util"
	"time"
)

type User struct {
	Uid      uint64    `gorm:"primary_key;column:uid;type:int(11)" json:"uid"`
	Username string    `gorm:"column:username;type:varchar(20)" json:"username"`
	Password string    `gorm:"column:password;type:varchar(255)" json:"password"`
	Role     uint64    `gorm:"column:role;type:int" json:"role"`
	Email    string    `gorm:"column:email;type:varchar(100)" json:"email"`
	CreateAt time.Time `gorm:"column:create_at;type:datetime" json:"create_at"`
}

func (u *User) AfterCreate() {
	u.Password = util.Encode(u.Password)
}
