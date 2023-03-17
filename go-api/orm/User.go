package orm

import (
	"gorm.io/gorm" // framwork ต่อกับ database ภาษา GO
)

type User struct { // สร้าง ตารางใน database ชื่อ User
	gorm.Model
	Username string
	Password string
	Fullname string
}
