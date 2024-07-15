package services

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type PgHandler struct {
	Ctx     context.Context
	Connect *pgx.Conn
	Config  *DbConfiguration
}

type DbConfiguration struct {
	URI      string
	Username string
	Password string
	Database string
}

func lookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}

func (db *DbConfiguration) NewConnection() (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	connectString := "postgres://" + db.Username + ":" + db.Password + "@" + db.URI + "/" + db.Database

	conn, err := pgx.Connect(context.Background(), connectString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn, err
}

func ParseConfiguration() *DbConfiguration {
	database := lookupEnvOrGetDefault("PG_DATABASE", "test")

	return &DbConfiguration{
		URI:      lookupEnvOrGetDefault("PG_URI", "localhost:5432"),
		Username: lookupEnvOrGetDefault("PG_USER", "pguser"),
		Password: lookupEnvOrGetDefault("PG_PASSWORD", "somepassword"),
		Database: database,
	}
}
