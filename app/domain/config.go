package domain

// Config expresses config data.
type Config struct {
	grpcPort int
	rdbURL   string
	rdbName  string
	rdbUser  string
	rdbPass  string
}

// NewConfig creates *Config.
func NewConfig(
	grpcPort int,
	rdbURL, rdbName, rdbUser, rdbPass string,
) *Config {
	return &Config{
		grpcPort: grpcPort,
		rdbURL:   rdbURL,
		rdbName:  rdbName,
		rdbUser:  rdbUser,
		rdbPass:  rdbPass,
	}
}

// GetGRPCPort returns a gRPC port.
func (c *Config) GetGRPCPort() int {
	if c == nil {
		return 0
	}

	return c.grpcPort
}

// GetRDBURL returns the rdb url.
func (c *Config) GetRDBURL() string {
	if c == nil {
		return ""
	}

	return c.rdbURL
}

// GetRDBName returns the rdb name.
func (c *Config) GetRDBName() string {
	if c == nil {
		return ""
	}

	return c.rdbName
}

// GetRDBUser returns the rdb user.
func (c *Config) GetRDBUser() string {
	if c == nil {
		return ""
	}

	return c.rdbUser
}

// GetRDBPass returns the rdb password.
func (c *Config) GetRDBPass() string {
	if c == nil {
		return ""
	}

	return c.rdbPass
}
