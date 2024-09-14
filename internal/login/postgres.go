package login

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/proyectogolang/internal/login/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres(cfg config.Database) (Postgres, error) {
	gdb, err := connect(cfg)
	if err != nil {
		return Postgres{}, err
	}
	if err = applyMigrations(cfg); err != nil {
		return Postgres{}, fmt.Errorf("an error ocurred applying migrations: %w", err)
	}
	return Postgres{
		DB: gdb,
	}, nil
}

func connect(cfg config.Database) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: cfg.LogLevel,
		},
	)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	instance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	db, err := instance.DB()
	if err != nil {
		return nil, err
	}

	maxIdleSize, err := strconv.Atoi(env.GetEnv("MAX_IDLE_SIZE"))
	if err != nil {
		return nil, err
	}

	maxOpenSize, err := strconv.Atoi(env.GetEnv("MAX_OPEN_SIZE"))
	if err != nil {
		return nil, err
	}

	maxLifeTime, err := time.ParseDuration(env.GetEnv("MAX_LIFE_TIME") + "ms")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleSize)
	db.SetMaxOpenConns(maxOpenSize)
	db.SetConnMaxLifetime(maxLifeTime)

	return instance, nil
}

func applyMigrations(cfg config.Database) error {
	currentPath := fmt.Sprintf("file:///%s/migrations", path.GetMainPath())
	url := generatePgUrl(cfg)
	m, err := migrate.New(currentPath, url)
	if err != nil {
		return err
	}

	return m.Up()
}

func generatePgUrl(cfg config.Database) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}
