package delivery

import (
	"clean_code/cars/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Delivery struct {
	useCase domain.UseCase
}

func NewDelivery(u domain.UseCase) domain.Delivery {
	return &Delivery{
		useCase: u,
	}
}

func (d *Delivery) GetCars(ctx *gin.Context) {
	cars, err := d.useCase.GetCars()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "somethings wrong",
		})
	}
	ctx.JSON(200, gin.H{
		"cars": cars,
	})
}

func (d *Delivery) CreateCar(ctx *gin.Context) {
	var car *domain.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.JSON(404, gin.H{
			"message": "invalid data",
		})
	}
	newCar, err := d.useCase.CreateCar(car)
	if err != nil {
		ctx.JSON(404, map[string]interface{}{
			"message": err,
		})
	}
	ctx.JSON(201, gin.H{
		"car": newCar,
	})
}

func (d *Delivery) GetCarById(ctx *gin.Context) {
	id := ctx.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "somethings wrong",
		})
		return
	}
	car, err := d.useCase.GetCarById(uint(newId))
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"car": car,
	})
}

func (d *Delivery) DeleteCar(ctx *gin.Context) {
	id := ctx.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "somethings wrong",
		})
		return
	}
	car, err := d.useCase.GetCarById(uint(newId))
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err,
		})
		return
	}
	d.useCase.DeleteCar(car)
	ctx.JSON(200, gin.H{
		"message": "deleted successfully",
	})
}
