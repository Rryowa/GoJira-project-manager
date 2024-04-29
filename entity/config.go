package entity

type Config struct {
	Port      string `env:"DB_PORT" env-default:"5432"`
	Host      string `env:"DB_HOST" env-default:"localhost"`
	Name      string `env:"DB_NAME" env-default:"demo"`
	User      string `env:"DB_USER" env-default:"postgres"`
	Password  string `env:"DB_PASSWORD" env-default:"postgres"`
	JWTSecret string `env:"DB_JWT_SECRET" env-default:"random"`
}
