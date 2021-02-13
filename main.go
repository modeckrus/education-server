package main

import (
	"education/config"
	"education/database"
	"education/restapi"
	"flag"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

var (
	configPath   string
	serverConfig config.ServerConfig
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.toml", "path to config file")
	serverConfig = config.NewServerConfig()
	_, err := toml.DecodeFile(configPath, &serverConfig)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	r := mux.NewRouter()
	db := database.Connect(serverConfig.MongodbAddr)
	resolver := restapi.Resolver{
		DB: *db,
	}
	r.HandleFunc("/version", restapi.GetVersion)
	r.HandleFunc("/firebaseauth", resolver.FirebaseAuth)
	log.Fatal(http.ListenAndServe(":8080", r))
}
