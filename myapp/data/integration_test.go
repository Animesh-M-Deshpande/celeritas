package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "secret"
	dbName   = "celeritas_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var dummyUser = User{
	FirstName: "some",
	LastName:  "guy",
	Email:     "me@here.com",
	Active:    1,
	Password:  "password",
}

var models Models
var testDB *sql.DB

var resource *dockertest.Resource
var pool *dockertest.Pool

// 2 test mains separated using build tag

func TestMain(m *testing.M) {
	os.Setenv("DATABASE_TYPE", "postgres")

	p, err := dockertest.NewPool("")

	if err != nil {
		log.Fatalf("could not connect to docker")
	}

	pool = p

	opts := dockertest.RunOptions{
		Repository: "postgres", // image of docker // similar to docker compose YML file
		Tag:        "13.4",     //version
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRESS_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"}, // port on docker image
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err = pool.RunWithOptions(&opts)

	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not start resource:%s", err)
	}

	// check until the docker and db is ready

	if err = pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))

		if err != nil {
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatal("could not connect to docker %s", err)
	}

	// at this point, we have docker, postgress and a db
}
