package restapi

import (
	"education/config"
	"education/database"
)

//Resolver ...
type Resolver struct {
	DB database.DB
}

//NewResolver ...
func NewResolver(conf config.ServerConfig) *Resolver {
	db := database.Connect(conf)
	return &Resolver{
		DB: *db,
	}
}
