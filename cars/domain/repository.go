package domain

type Repository interface {
	GetCarById(id uint) (*Car, error)
	DeleteCar(car *Car)
	GetAllCars() (*[]Car, error)
	CreateCar(car *Car) (*Car, error)
}
