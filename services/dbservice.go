package services

import (
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

func DBInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DATABASE()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&User{}, &Session{})
	return db
}

/*  Functions specific to reading-writing DB. Ideally should have also abstracted adding to or reading from cache
	since these are tied to CacheServiceInstance, but eh, will see later. I was also thinking of adding a DBService, but I don't really
	see the use for it since rn I'm mostly tied to the caching.
	Now that I'm looking at this I'm not sure why I even separated them in the first place, but I'll keep it like this for readability
*/
func (s *CacheService) DBWrite(data interface{}) (error){
	if err := s.db.WithContext(s.ctx).Create(data).Error; err != nil {
		return fmt.Errorf("Failed to write data in database: %v", err)
	}
	return nil
}

func (s *CacheService) DBRead(entry string, data string, table interface{}) (error){
	if err := s.db.WithContext(s.ctx).Where(entry, data).First(table).Error; err != nil {
		return fmt.Errorf("Failed to read data in databse: %v", err)
	}
	return nil
}