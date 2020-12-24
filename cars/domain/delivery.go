package domain

import "github.com/gin-gonic/gin"

type Delivery interface {
	DeleteCar(ctx *gin.Context)
	GetCarById(ctx *gin.Context)
	CreateCar(ctx *gin.Context)
	GetCars(ctx *gin.Context)
}
