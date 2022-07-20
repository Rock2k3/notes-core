package config

type grpcConfig struct {
	grpcUsersServiceAddress string
}

func (c grpcConfig) GrpcUsersServiceAddress() string {
	return c.grpcUsersServiceAddress
}
