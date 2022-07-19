package config

type GrpcConfig struct {
	grpcUsersServiceAddress string
}

func (c GrpcConfig) GrpcUsersServiceAddress() string {
	return c.grpcUsersServiceAddress
}
