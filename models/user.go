package models

type Address struct {
	State      string `json:"state" bson:"user_state"`
	City       string `json:"city" bson:"user_city"`
	PostalCode string `json:"postalcode" bson:"postal_code"`
}

type User struct {
	Id      int     `json:"id" bson:"id"`
	Name    string  `json:"name" bson:"user_name"`
	Email   string  `json:"email" bson:"user_email"`
	Age     int     `json:"age" bson:"user_age"`
	Address Address `json:"address" bson:"user_address"`
}
