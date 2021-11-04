package device

import (
	"context"
	"os"
	"strconv"

	"github.com/g-chicken/mah-jong/app/domain"
)

const (
	envPrefix = "MAH_JONG_"

	grpcPortEnvName = envPrefix + "GRPC_PORT"
	rdbURLEnvName   = envPrefix + "RDB_URL"
	rdbNameEnvName  = envPrefix + "RDB_NAME"
	rdbUserEnvName  = envPrefix + "RDB_USER"
	rdbPassEnvName  = envPrefix + "RDB_PASS"

	grpcPortDefault = 8080
	rdbURLDefault   = "localhost:3306"
	rdbNameDefault  = "mah_jong"
	rdbUserDefault  = "app"
	rdbPassDefault  = "hoge"
)

type configRepository struct{}

// NewConfigRepository implements domain.ConfigRepository.
func NewConfigRepository() domain.ConfigRepository {
	return &configRepository{}
}

func (r *configRepository) GetConfig(c context.Context) (*domain.Config, error) {
	grpcPort, err := r.getEnvInt(grpcPortEnvName, grpcPortDefault)
	if err != nil {
		return nil, err
	}

	return domain.NewConfig(
		grpcPort,
		r.getEnvString(rdbURLEnvName, rdbURLDefault),
		r.getEnvString(rdbNameEnvName, rdbNameDefault),
		r.getEnvString(rdbUserEnvName, rdbUserDefault),
		r.getEnvString(rdbPassEnvName, rdbPassDefault),
	), nil
}

func (r *configRepository) getEnvString(key, d string) string {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}

	return d
}

func (r *configRepository) getEnvInt(key string, d int) (int, error) {
	val, exist := os.LookupEnv(key)
	if exist {
		v, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}

		return v, nil
	}

	return d, nil
}
