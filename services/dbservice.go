package services

import (
	"context"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

func DATABASE() string {
	database := os.Getenv("SQLITE_PATH")
	if database == "" {
		return "database/sqlite.db"
	}
	return database
}

var DatabaseService *DBService

type User struct {
	gorm.Model
	UID          string
	Email        string
	HashedPass   string
	Playername   string
	Entitlements string
}

type Session struct {
	gorm.Model
	UID        string
	Sessionkey string
	Gamever    string
}

type DBService struct {
	db  *gorm.DB
	ctx context.Context
}

func DBInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DATABASE()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&User{}, &Session{})
	return db
}

func NewDatabaseService(db *gorm.DB, ctx context.Context) *DBService {
	return &DBService{
		db:  db,
		ctx: ctx,
	}
}

func InitDatabaseService(db *gorm.DB) error {
	DatabaseService = NewDatabaseService(db, context.Background())
	if DatabaseService == nil {
		return fmt.Errorf("Something bad happened in InitDatabaseService!!!")
	}
	return nil
}

/*
	    Functions specific to reading-writing DB. Now separate from CaheService =D
		How to use examples:
		var user User // type User struct
		err := Query(DatabaseService, "uid", "myuid", &user)
		err = Write(DatabaseService, &user)

		IMPORTANT!!! These should NOT be used outside of abstractions in CacheService

		TODO: rewrite gorm usage with go generics
*/
func Create[T any](s *DBService, data *T) error {
	if err := s.db.WithContext(s.ctx).Create(data).Error; err != nil {
		return fmt.Errorf("Failed to create data in database: %v", err)
	}
	return nil
}

func Query[T any](s *DBService, entryname string, cacheKey string, dest *T) error {
	var entry = entryname + " = ?"
	if err := s.db.WithContext(s.ctx).Where(entry, cacheKey).First(dest).Error; err != nil {
		return fmt.Errorf("Failed to read data in databse: %v", err)
	}
	return nil
}

func Delete[T any](s *DBService, entryname string, cacheKey string, dest *T) error {
	var entry = entryname + " = ?"
	if err := s.db.WithContext(s.ctx).Where(entry, cacheKey).Delete(dest).Error; err != nil {
		return fmt.Errorf("Failed to delete data in databse: %v", err)
	}
	return nil
}
