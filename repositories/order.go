package repositories

import (
	"errors"
	"mini_project/database"
	"mini_project/models"
	"mini_project/models/input"
)

type OrderRepositoryImp struct {
}

func InitOrderRepository() OrderRepository {
	return &OrderRepositoryImp{}
}

func (ur *OrderRepositoryImp) Create(orderInput input.OrderInput) (models.Order, error) {
	var transaction models.Transaction

	if err := database.ConnectDB().Create(&transaction).Error; err != nil {
		return models.Order{}, err
	}

	var order models.Order = models.Order{
		UserID: orderInput.UserID,
		VehicleID: orderInput.VehicleID,
		TransactionID: transaction.ID,
		RentDuration: orderInput.RentDuration,
		Status: orderInput.Status,
		Transaction: transaction,
	}

	if err := database.ConnectDB().Create(&order).Error; err != nil {
		return models.Order{}, err
	}

	if err := database.ConnectDB().Last(&order).Error; err != nil {
		return models.Order{}, err
	}

    return order, nil
}

func (ur *OrderRepositoryImp) GetAll() ([]models.Order, error) {
	var orders []models.Order

	if err := database.ConnectDB().Find(&orders).Error; err != nil {
		return orders, err
	}
	return orders, nil
}

func (ur *OrderRepositoryImp) GetById(id string) (models.Order, error) {
	var order models.Order

	if err := database.ConnectDB().First(&order, id).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (ur *OrderRepositoryImp) GetHistory(id string) ([]models.Order, error) {
	var orders []models.Order

	if id == "0"{
		return orders, errors.New("user id not exists")
	}

	if err := database.ConnectDB().Find(&orders, "user_id = ?", id).Error; err != nil {
		return orders, err
	}
	return orders, nil
}

func (ur *OrderRepositoryImp) UpdateRating(orderInput input.OrderInput, id string) (models.Order, error) {
	order, err := ur.GetById(id)

	if err != nil {
		return models.Order{}, err
	}

	if order.Status != "accepted" {
		return models.Order{}, errors.New("you haven't pay your order yet")
	}

	order.OrderRate = orderInput.OrderRate

	if err := database.ConnectDB().Save(&order).Error; err != nil {
		return models.Order{}, err
	}

    return order, nil
}

func (ur *OrderRepositoryImp) UpdateStatus(id string) (models.Order, error) {
	order, err := ur.GetById(id)

	if err != nil {
		return models.Order{}, err
	}

	order.Status = "accepted"

	if err := database.ConnectDB().Save(&order).Error; err != nil {
		return models.Order{}, err
	}

    return order, nil
}
