package models

type MySQLConfiguration struct {
	DBUser     string `mapstructure:"username"`
	DBPassword string `mapstructure:"password"`
	DBHost     string `mapstructure:"host"`
	DBPort     string `mapstructure:"port"`
	DBName     string `mapstructure:"name"`
}

type App struct {
	Port         string             `mapstructure:"service_port"`
	JWTSecretKey string             `mapstructure:"jwt_secret_key"`
	DbConfig     MySQLConfiguration `mapstructure:"database"`
}
