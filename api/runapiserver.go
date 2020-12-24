package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunApiServer(db *gorm.DB) error {
	r := gin.Default()
	RunCarApi(r, db)
	return r.Run(":8000")
}
