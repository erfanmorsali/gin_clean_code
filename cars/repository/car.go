package repository

import (
	"clean_code/cars/domain"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllCars() (*[]domain.Car, error) {
	var cars []domain.Car
	r.db.Find(&cars)
	return &cars, nil
}

func (r *Repository) GetCarById(id uint) (*domain.Car, error) {
	var car domain.Car
	if err := r.db.Where("id = ?", id).First(&car).Error; err != nil {
		return nil, err
	}
	return &car, nil
}

func (r *Repository) CreateCar(car *domain.Car) (*domain.Car, error) {
	r.db.Create(&car)
	return car, nil
}

func (r *Repository) DeleteCar(car *domain.Car) {
	r.db.Delete(&car)
}
