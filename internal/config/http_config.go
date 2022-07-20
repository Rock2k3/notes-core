package config

type httpConfig struct {
	httpAddress string
}

func (c httpConfig) HttpAddress() string {
	return c.httpAddress
}
