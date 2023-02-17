package gorm

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

var (
	db  *gorm.DB
	err error
)

type Account struct {
	gorm.Model
	Blocked bool
	Name    string
}

func init() {
	db, err = gorm.Open(postgres.Open(fmt.Sprintf(
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
				LogLevel:                  logger.Silent,          // Log level
				IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&Account{})
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkGenericRepoCreate(b *testing.B) {

	r := New[Account](db)

	for i := 0; i < b.N; i++ {
		_, err = r.Create(context.TODO(), Account{Name: "test"})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGORMCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err = db.Create(&Account{Name: "test"}).Error; err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGORMRawCreate(b *testing.B) {
	stmt := `INSERT INTO "accounts" ("created_at","updated_at","deleted_at","blocked","name") VALUES ('2023-02-17 06:38:18.537','2023-02-17 06:38:18.537',NULL,false,'test') RETURNING "id"`
	for i := 0; i < b.N; i++ {
		if err = db.Raw(stmt).Error; err != nil {
			b.Fatal(err)
		}
	}
}
