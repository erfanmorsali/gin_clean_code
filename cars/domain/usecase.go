package domain

type UseCase interface {
	GetCarById(id uint) (*Car, error)
	DeleteCar(car *Car)
	CreateCar(car *Car) (*Car, error)
	GetCars() (*[]Car, error)
}
