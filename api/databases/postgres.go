package databases

import (
	"context"
	"fmt"
	"time"

	"github.com/g-villarinho/flash-buy-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(ctx context.Context) (*gorm.DB, error) {
	dsn := getDSN(config.Env)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	slqDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	slqDB.SetMaxOpenConns(config.Env.Postgres.MaxConn)
	slqDB.SetMaxIdleConns(config.Env.Postgres.MaxIdle)
	slqDB.SetConnMaxLifetime(time.Duration(config.Env.Postgres.MaxLifeTime) * time.Second)

	if err := slqDB.PingContext(ctx); err != nil {
		_ = slqDB.Close()
		return nil, err
	}

	return db, nil
}

func getDSN(config config.Environment) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.DBName,
		config.Postgres.Password,
		config.Postgres.DBSSLMode,
	)
}
