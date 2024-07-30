package main

type Passenger struct {
	ID    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}