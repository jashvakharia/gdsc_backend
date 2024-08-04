package main

import (
    "github.com/gin-gonic/gin"
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
    "net/http"
)

var db *gorm.DB
var err error

type Data struct {
    ID    uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Value string `json:"value"`
}

const apiKey = "gdsc_backend"

func main() {
    r := gin.Default()

    // Database initialization
    db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&Data{})

    r.GET("/test", getSample)
    r.POST("/post", postData)
    r.GET("/data", requireAPIKey(), getData)

    r.Run(":8080")
}

func getSample(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Welcome to the backend of GDSC."})
}

func postData(c *gin.Context) {
    var data Data
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format/content-type."})
        return
    }
    db.Create(&data)
    c.JSON(http.StatusCreated, gin.H{"message": "Your data has been saved successfully!"})
}

func getData(c *gin.Context) {
    var data []Data
    db.Find(&data)
    c.JSON(http.StatusOK, data)
}

func requireAPIKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        apiKeyHeader := c.GetHeader("x-api-key")
        if apiKeyHeader != apiKey {
            c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
            c.Abort()
            return
        }
        c.Next()
    }
}
