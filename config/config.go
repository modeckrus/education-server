package config

//ServerConfig ...
type ServerConfig struct {
	Port        string `toml:"port"`
	MongodbAddr string `toml:"mongodb_addr"`
}

// port = "8080"
// redis_addr = "localhost:6380"
// redis_pass = ""
// redis_time = 24
// mongodb_type = "local"
// mongodb_addr = "mongodb://localhost:27003"

//NewServerConfig ...
func NewServerConfig() ServerConfig {
	return ServerConfig{
		Port:        "8080",
		MongodbAddr: "mongodb://localhost:27003",
	}
}
