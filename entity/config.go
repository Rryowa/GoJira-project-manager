package entity

type Config struct {
	Port      string `env:"DB_PORT" env-default:"5432"`
	Host      string `env:"DB_HOST" env-default:"localhost"`
	Name      string `env:"POSTGRES_DB" env-default:"demo"`
	User      string `env:"POSTGRES_USER" env-default:"postgres"`
	Password  string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
	JWTSecret string `env:"JWT_SECRET" env-default:"secret"`
}
