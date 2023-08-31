package models

import (
	"time"
)

type Customer struct {
	CustomerId int32 `json:"customerid" bson:"customerid"`

	FirstName string `json:"firstname" bson:"firstname"`

	LastName string `json:"lastname" bson:"lastname"`

	BankId string `json:"bankid" bson:"bankid"`

	Balance float64 `json:"balance" bson:"balance"`

	CreatedAt time.Time `json:"createdat" bson:"createdat"`

	UpdatedAt time.Time `json:"updatedat" bson:"updatedat"`

	IsActive bool `json:"isactive" bson:"isactive"`
}

type DBResponse struct {
	// ID primitive.ObjectID `json:"id" bson:"_id"`

	CustomerId int32 `json:"customerid" bson:"customerid"`

	// FirstName string `json:"firstname" bson:"firstname"`

	// LastName string `json:"lastname" bson:"lastname"`

	// BankId string `json:"bankid" bson:"bankid"`

	// Balance float64 `json:"balance" bson:"balance"`

	CreatedAt time.Time `json:"createdat" bson:"createdat"`

	// UpadtedAt string `json:"updatedat" bson:"updatedat"`

	// IsActive bool `json:"isactive" bson:"isactive"`
}
