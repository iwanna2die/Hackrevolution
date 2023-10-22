package postgres

import (
	"HackRevolution/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var instance SqlDb

type SqlDb struct {
	SDB *sqlx.DB
}

func Connect(c *config.Config) error {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Dbname,
		c.Postgres.Password,
	)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return err
	}
	instance.SDB = db
	return nil
}

func DB() *SqlDb {
	if instance.SDB != nil {
		return &instance
	}
	panic("database not initialized")
}
