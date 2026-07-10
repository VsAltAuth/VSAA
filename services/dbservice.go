package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
	"github.com/patrickmn/go-cache"
)

func DATABASE() string {
	database := os.Getenv("SQLITE_PATH")
	if database == "" {
		return "database/sqlite.db"
	}
	return database
}

var CacheServiceInstance *CacheService

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

type CacheService struct {
	cache *cache.Cache
	db    *gorm.DB
}

var ctx = context.Background()

func DbInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DATABASE()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&User{}, &Session{})
	return db
}

func NewCacheService(db *gorm.DB, expirationDuration time.Duration, cleanupInterval time.Duration) *CacheService {
	return &CacheService{
		db:    db,
		cache: cache.New(expirationDuration, cleanupInterval),
	}
}

func InitCacheService(db *gorm.DB) error {
	CacheServiceInstance = NewCacheService(db, 5*time.Minute, 10*time.Minute)
	if CacheServiceInstance == nil {
		return fmt.Errorf("Something bad happened in InitCacheService!!!")
	}
	return nil
}

func (s *CacheService) GetUserByUID(uid string) (*User, error) {
	cacheKey := uid

	if cached, found := s.cache.Get(cacheKey); found {
		return cached.(*User), nil
	}

	// Query DB if not found in cache
	var user User
	if err := s.db.Where("uid = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}
	// Cache value we got
	s.cache.Set(cacheKey, &user, 5*time.Minute)
	return &user, nil
}

func (s *CacheService) GetUserByPlayername(playername string) (*User, error) {
	cacheKey := playername
	if cached, found := s.cache.Get(cacheKey); found {
		return cached.(*User), nil
	}

	// Query DB if not found in cache
	var user User
	if err := s.db.Where("playername = ?", playername).First(&user).Error; err != nil {
		return nil, err
	}
	// Cache value we got
	s.cache.Set(cacheKey, &user, 5*time.Minute)
	return &user, nil
}
