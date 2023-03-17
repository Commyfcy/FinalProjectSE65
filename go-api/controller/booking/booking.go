package booking


import (
	"net/http"
	"se/jwt-api/orm"
	"time"
	"github.com/gin-gonic/gin"
)



type BookingBody struct{
	UserID string
	CarID  string
	Start  time.Time
	End    time.Time
}

func BookingCar(c *gin.Context) {
	var json BookingBody
	if err := c.ShouldBindJSON(&json); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
	}
	layout := "2006-01-02"
start, err := time.Parse(layout, json.Start.Format(layout))
if err != nil {
// handle error
}
end, err := time.Parse(layout, json.End.Format(layout))
if err != nil {
// handle error
}
// Query the database using Gorm
var results []orm.Booking
orm.Db.Where("car_id = ? AND start BETWEEN ? AND ?", json.CarID, start,
end).Find(&results)
// Check if the booking already exists
if len(results) > 0 {
c.JSON(400, gin.H{"status": "error", "message": "Booking Exists"})
return
}
// Create the booking
booking := orm.Booking{UserID: json.UserID, CarID: json.CarID, Start: start,
End: end}
orm.Db.Create(&booking)
c.JSON(200, gin.H{"status": "success", "data": booking})

}