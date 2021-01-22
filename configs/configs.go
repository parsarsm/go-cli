package configs

var Configs struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Host     string
	GRPCPort string
	HTTPPort string
}

type DatabaseConfiguration struct {
	Host            string
	Port            int
	User            string
	Password        string
	DB              string
	ConnMaxLifetime int
	MaxIdleCons     int
	MaxOpenCons     int
}
