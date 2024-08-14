package configs

type HTTP struct {
	Host string `env:"SERVER_HOST"`
	Port string `env:"SERVER_PORT"`
}
