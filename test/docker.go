package test

import (
	"fmt"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

const (
	MYSQL_IMAGE_NAME string = "mysql"
	MYSQL_IMAGE_TAG  string = "5.7"
	MYSQL_USERNAME   string = "tygbtw"
	MYSQL_PASSWORD   string = "password"
)

var (
	MysqlEnv = []string{
		fmt.Sprintf("MYSQL_USER=%s", MYSQL_USERNAME),
		fmt.Sprintf("MYSQL_PASSWORD=%s", MYSQL_PASSWORD),
		fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", MYSQL_PASSWORD),
		"MYSQL_DATABASE=test",
	}

	MysqlRunOptions = dockertest.RunOptions{
		Repository: MYSQL_IMAGE_NAME,
		Tag:        MYSQL_IMAGE_TAG,
		Env:        MysqlEnv,
	}
)

func NewPoolTest() (*dockertest.Pool, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return pool, err
	}

	pool.MaxWait = time.Minute

	err = pool.Client.Ping()

	return pool, err
}

func StartInstance(pool *dockertest.Pool, option *dockertest.RunOptions) (*dockertest.Resource, error) {
	res, err := pool.RunWithOptions(option, func(config *docker.HostConfig) {
		config.AutoRemove = true
	})

	return res, err
}
