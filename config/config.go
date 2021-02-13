package config

//ServerConfig ...
type ServerConfig struct {
	Port           string `toml:"port"`
	MongodbAddr    string `toml:"mongodb_addr"`
	FirebaseConfig string `toml:"firebase_config"`
	KeyPem         string `toml:"key_pem"`
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
		Port:           "8080",
		MongodbAddr:    "mongodb://localhost:27003",
		FirebaseConfig: "/home/modeck/Documents/secure/serviceAccount.json",
		KeyPem:         "/home/modeck/Documents/secure/key.pem",
	}
}
