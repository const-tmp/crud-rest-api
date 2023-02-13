package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nullc4t/crud-rest-api/internal/server"
	"github.com/nullc4t/crud-rest-api/pkg/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}\t${method}\t${uri}\t${status}\t${latency_human}",
	}))

	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		"localhost",
		"postgres",
		"password",
		"postgres",
		5432,
		"Europe/Kiev",
	)), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Millisecond * 500, // Slow SQL threshold
				LogLevel:                  logger.Info,            // Log level
				IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,                   // Disable color
			},
		),
	})
	if err != nil {
		e.Logger.Fatal(err)
	}

	if err = db.AutoMigrate(&server.Account{}, &server.Service{}, &server.Permission{}); err != nil {
		e.Logger.Fatal(err)
	}

	impl := server.New(db)
	auth.RegisterHandlers(e, impl)

	e.Logger.Fatal(e.Start(":8080"))
}
