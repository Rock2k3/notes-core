package config

type HttpConfig struct {
	httpAddress string
}

func (c HttpConfig) HttpAddress() string {
	return c.httpAddress
}
