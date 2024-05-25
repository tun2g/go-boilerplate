package config

// import "os"

// type Configuration struct {
// 	GoEnv string `mapstructure:"GO_ENV"`
// 	Db    struct {
// 		Name         string `mapstructure:DATABASE_NAME"`
// 		Port         int    `mapstructure:DATABASE_PORT"`
// 		Host         string `mapstructure:DATABASE_HOST"`
// 		Driver       string `mapstructure:DATABASE_DRIVER"`
// 		User         string `mapstructure:DATABASE_USER"`
// 		Password     string `mapstructure:DATABASE_PASSWORD"`
// 		MaxOpenConns int    `mapstructure:DATABASE_MAX_OPEN_CONNS"`
// 		MaxIdleConns int    `mapstructure:DATABASE_MAX_IDLE_CONNS`
// 		ConnMaxLife  int    `mapstructure:DATABASE_CONN_MAX_LIFE`
// 	}
// }

// func LoadConfiguration() Configuration {
// 	return Configuration{
// 		GoEnv: os.Getenv("GO_ENV"),
// 		Db: {
// 			Name: os.Getenv("DATABASE_NAME"),
// 			Port: os.Getenv("DATABASE_PORT"),
// 			Host: os.Getenv("DATABASE_HOST"),
// 			Driver: os.Getenv("DATABASE_DRIVER"),
// 			User: os.Getenv("DATABASE_USER"),
// 			Password: os.Getenv("DATABASE_PASSWORD"),
// 			MaxOpenConns: os.Getenv("DATABASE_MAX_CONNS"),
// 			MaxIdleConns: os.Getenv("DATABASE_MAX_IDLE_CONNS"),
// 			ConnMaxLife: os.Getenv("DATABASE_CONN_MAX_LIFE"),			
// 		},
// 	}
// }
