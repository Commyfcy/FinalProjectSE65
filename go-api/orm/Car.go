package orm

import (
	"gorm.io/gorm" // framwork ต่อกับ database ภาษา GO
)

type Car struct { // สร้าง ตารางใน database ชื่อ User
	gorm.Model
	Carname string
	Detail  string
	Image   string
}
