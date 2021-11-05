package domain

import "time"

// Config expresses config data.
type Config struct {
	grpcPort             int
	rdbURL               string
	rdbName              string
	rdbUser              string
	rdbPass              string
	rdbConnectionTimeout time.Duration
}

// NewConfig creates *Config.
func NewConfig(
	grpcPort int,
	rdbURL, rdbName, rdbUser, rdbPass string,
	rdbConnectionTimeout time.Duration,
) *Config {
	return &Config{
		grpcPort:             grpcPort,
		rdbURL:               rdbURL,
		rdbName:              rdbName,
		rdbUser:              rdbUser,
		rdbPass:              rdbPass,
		rdbConnectionTimeout: rdbConnectionTimeout,
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

// GetRDBConnectionTimeout returns the timeout connecting to RDB.
func (c *Config) GetRDBConnectionTimeout() time.Duration {
	if c == nil {
		return time.Duration(0)
	}

	return c.rdbConnectionTimeout
}
