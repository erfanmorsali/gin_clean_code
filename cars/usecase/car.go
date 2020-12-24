package usecase

import "clean_code/cars/domain"

type UseCase struct {
	repo domain.Repository
}

func NewUseCase(repo domain.Repository) domain.UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) GetCars() (*[]domain.Car, error) {
	return u.repo.GetAllCars()
}

func (u *UseCase) GetCarById(id uint) (*domain.Car, error) {
	car, err := u.repo.GetCarById(id)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (u *UseCase) CreateCar(car *domain.Car) (*domain.Car, error) {
	car, err := u.repo.CreateCar(car)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (u *UseCase) DeleteCar(car *domain.Car) {
	u.repo.DeleteCar(car)
}
