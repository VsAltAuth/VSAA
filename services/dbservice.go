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
	DatabaseService := NewDatabaseService(db, context.Background())
	if DatabaseService == nil {
		return fmt.Errorf("Something bad happened in InitDatabaseService!!!")
	}
	return nil
}

/*
	  Functions specific to reading-writing DB. Now separate from CaheService =D
		Ideally I should abstract than one out too.
		How to use examples:
		var user User // type User struct
		err := DatabaseService.Read("uid = ?", "myuid", &user)
		err = DatabaseService.Write(&user)
*/
func (s *DBService) Write(data interface{}) error {
	if err := s.db.WithContext(s.ctx).Create(data).Error; err != nil {
		return fmt.Errorf("Failed to write data in database: %v", err)
	}
	return nil
}

func (s *DBService) Read(entry string, data string, table interface{}) error {
	if err := s.db.WithContext(s.ctx).Where(entry, data).First(table).Error; err != nil {
		return fmt.Errorf("Failed to read data in databse: %v", err)
	}
	return nil
}
