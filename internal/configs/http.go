package configs

type HTTP struct {
	Address string `env:"APP_ADDRESS"`
	Port    string `env:"APP_PORT"`
}
