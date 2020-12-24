package api

import (
	delivery2 "clean_code/cars/delivery"
	"clean_code/cars/repository"
	"clean_code/cars/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunCarApi(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewRepository(db)
	useCase := usecase.NewUseCase(repo)
	handler := delivery2.NewDelivery(useCase)
	car := r.Group("/car")
	{
		car.GET("/all_cars", handler.GetCars)
		car.GET("/car_detail/:id", handler.GetCarById)
		car.POST("/create_car", handler.CreateCar)
		car.DELETE("/delete_car/:id", handler.DeleteCar)
	}
}
