package repositories

import (
	"mini_project/models"
	"mini_project/models/input"
)

type UserRepository interface {
	Register(userInput input.UserInput) (models.User, error)
	Login(userInput input.UserInput) (models.User, error)
}

type VehicleRepository interface {
	Create(vehicleInput input.VehicleInput) (models.Vehicle, error)
	GetByName(name string) ([]models.Vehicle, error) 
	GetAll() ([]models.Vehicle, error)
	GetById(id string) (models.Vehicle, error)
	UpdateRating(id string) (models.Vehicle, error)
	Update(vehicleInput input.VehicleInput, id string) (models.Vehicle, error)
	Delete(id string) error
}

type TransactionRepository interface {
	Update(transactionInput input.TransactionInput, id string) (models.Transaction, error)
}

type OrderRepository interface {
	Create(orderInput input.OrderInput) (models.Order, error)
	GetAll() ([]models.Order, error)
	GetHistory(id string) ([]models.Order, error)
	UpdateStatus(id string) (models.Order, error)
	UpdateRating(orderInput input.OrderInput, id string) (models.Order, error)
}