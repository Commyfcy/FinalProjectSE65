package auth

import (
	"fmt"
	"net/http"
	"os"
	"se/jwt-api/orm"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret []byte

type RegisterBody struct { // ออกแบบข้อมูล
	Username string
	Password string
	Fullname string
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// check user ซ้า ไหม
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(200, gin.H{"status": "error", "message": "User Exists"})
		return
	}
	// encrypt password
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password),
		10) // encypt password
	// เก็บค่าลงใส่ Database
	user := orm.User{Username: json.Username, Password: string(encryptedPassword),
		Fullname: json.Fullname}
	orm.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(200, gin.H{"status": "ok", "message": "User Create Sucessful",
			"userID": user.ID})
	} else {
		c.JSON(200, gin.H{"status": "error", "message": "User Create Fail",
			"userID": user.ID})
	}
}

// สร้างการตรวจสอบ Login
type LoginBody struct { // ออกแบบข้อมูล Login
	Username string
	Password string
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// check username password
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(200, gin.H{"status": "error", "message": "User not found"})
		return
	}
	//check password ที่เข้ารหัส
	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password),
		[]byte(json.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userID": userExist.ID,
			"exp":    time.Now().Add(time.Minute * 6).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Print(tokenString, err)
		c.JSON(200, gin.H{"status": "ok", "message": "Login Success!", "token": tokenString,"user":userExist})
	} else {
		c.JSON(200, gin.H{"status": "error", "message": "password incorrect"})
	}
}
