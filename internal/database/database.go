// package database

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	_ "github.com/jackc/pgx/v5/stdlib"
// 	_ "github.com/joho/godotenv/autoload"
// )

// type Service interface {
// 	Health() map[string]string
// }

// type service struct {
// 	db *sql.DB
// }

// var (
// 	database = os.Getenv("DB_DATABASE")
// 	password = os.Getenv("DB_PASSWORD")
// 	username = os.Getenv("DB_USERNAME")
// 	port     = os.Getenv("DB_PORT")
// 	host     = os.Getenv("DB_HOST")
// )

// func New() Service {
// 	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
// 	db, err := sql.Open("pgx", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	s := &service{db: db}
// 	return s
// }

// func (s *service) Health() map[string]string {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	err := s.db.PingContext(ctx)
// 	if err != nil {
// 		log.Fatalf(fmt.Sprintf("db down: %v", err))
// 	}

// 	return map[string]string{
// 		"message": "It's healthy",
// 	}
// }
package database

import (
	"context"
	"time"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"website/internal/models"
	"fmt"
	"log"
	"os"
	"gorm.io/gorm/logger"
)

type Service struct{
	db *gorm.DB
}

var DB Service

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)


func New() Service {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)

	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database.\n",err)
		os.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	db.AutoMigrate(&models.Omnibus{})

	s := Service{
		db:db,
	}
	return s
}


func (s *Service) Health() map[string]string {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    err := s.db.WithContext(ctx).Raw("SELECT 1").Error
    if err != nil {
        return map[string]string{
            "status":  "error",
            "message": fmt.Sprintf("db down: %v", err),
        }
    }

    return map[string]string{
        "status":  "ok",
        "message": "It's healthy",
    }
}
