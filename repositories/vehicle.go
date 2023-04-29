package repositories

import (
	"mini_project/database"
	"mini_project/models"
	"mini_project/models/input"
)

type VehicleRepositoryImp struct {
}

func InitVehicleRepository() VehicleRepository {
	return &VehicleRepositoryImp{}
}

func (ur *VehicleRepositoryImp) Create(vehicleInput input.VehicleInput) (models.Vehicle, error) {
	var vehicle models.Vehicle = models.Vehicle{
		NumberPlate: vehicleInput.NumberPlate,
		Type: vehicleInput.Type,
		Name: vehicleInput.Name,
		Price: vehicleInput.Price,
		Status: vehicleInput.Status,
		Rating: vehicleInput.Rating,
	}

	if err := database.ConnectDB().Create(&vehicle).Error; err != nil {
		return models.Vehicle{}, err
	}

	if err := database.ConnectDB().Last(&vehicle).Error; err != nil {
		return models.Vehicle{}, err
	}

    return vehicle, nil
}

func (ur *VehicleRepositoryImp) GetByName(vehicleInput input.VehicleInput) ([]models.Vehicle, error) {
	var vehicle []models.Vehicle

	if err := database.ConnectDB().Find(&vehicle, "name = ?", vehicleInput.Name).Error; err != nil {
		return []models.Vehicle{}, err
	}

	return vehicle, nil
}

func (ur *VehicleRepositoryImp) GetAll() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle

	if err := database.ConnectDB().Find(&vehicles).Error; err != nil {
		return vehicles, err
	}
	return vehicles, nil
}

func (ur *VehicleRepositoryImp) GetById(id string) (models.Vehicle, error) {
	var vehicle models.Vehicle

	if err := database.ConnectDB().First(&vehicle, id).Error; err != nil {
		return models.Vehicle{}, err
	}
	return vehicle, nil
}

func (ur *VehicleRepositoryImp) Update(vehicleInput input.VehicleInput, id string) (models.Vehicle, error) {
	vehicle, err := ur.GetById(id)

	if err != nil {
		return models.Vehicle{}, err
	}

	vehicle.NumberPlate = vehicleInput.NumberPlate
	vehicle.Type = vehicleInput.Type
	vehicle.Name = vehicleInput.Name
	vehicle.Price = vehicleInput.Price
	vehicle.Status = vehicleInput.Status
	vehicle.Rating = vehicleInput.Rating

	if err := database.ConnectDB().Save(&vehicle).Error; err != nil {
		return models.Vehicle{}, err
	}

    return vehicle, nil
}

func (ur *VehicleRepositoryImp) Delete(id string) error {
	vehicle, err := ur.GetById(id)

	if err != nil {
		return err
	}

	if err := database.ConnectDB().Delete(&vehicle).Error; err != nil {
		return err
	}

    return nil
}

