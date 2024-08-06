package configs

type HTTP struct {
	Host string `env:"SERVER_HOSTs"`
	Port string `env:"SERVER_PORT"`
}
