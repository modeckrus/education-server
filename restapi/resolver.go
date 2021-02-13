package restapi

import "education/database"

//Resolver ...
type Resolver struct {
	DB database.DB
}

//NewResolver ...
func NewResolver() *Resolver {
	db := database.Connect("mongodb://localhost:27003")
	return &Resolver{
		DB: *db,
	}
}
