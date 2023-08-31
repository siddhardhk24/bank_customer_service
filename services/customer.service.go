package services

import (
	"context"

	"time"

	"github.com/siddhardhk24/bank_customer_service/interfaces"
	"github.com/siddhardhk24/bank_customer_service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

// CreateCustomer implements interfaces.ICustomer.
func (p *CustomerService) CreateCustomer(customer *models.Customer) (*models.DBResponse, error) {
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = customer.CreatedAt
	customer.IsActive = true

	res, err := p.CustomerCollection.InsertOne(p.ctx, &customer)

	if err != nil {
		return nil, err
	}
	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = p.CustomerCollection.FindOne(p.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
