package interfaces

import "github.com/siddardhk24/bank_customer_service/models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
}
