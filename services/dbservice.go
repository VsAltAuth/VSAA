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
	ctx   context.Context
}

func DBInit() *gorm.DB {
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
		ctx: context.Background(),
	}
}

func InitCacheService(db *gorm.DB, expirationDuration time.Duration, cleanupInterval time.Duration) error {
	CacheServiceInstance := NewCacheService(db, expirationDuration, cleanupInterval)
	if CacheServiceInstance == nil {
		return fmt.Errorf("Something bad happened in InitCacheService!!!")
	}
	return nil
}

func (s *CacheService) DBWrite[T any](data *T) (error){
	if err := s.db.WithContext(s.ctx).Create(data).Error; err != nil {
		return fmt.Errorf("Failed to write data in database: %v", err)
	}
	return nil
}

func (s *CacheService) DBRead[T any](entry string, data string, table *T) (error){
	if err := s.db.WithContext(s.ctx).Where(entry, data).First(table).Error; err != nil {
		return fmt.Errorf("Failed to read data in databse: %v", err)
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
	if err := DBRead("uid = ?", uid, &user); err != nil {
		return nil, err
	}
	// Cache value we got
	s.cache.Set(cacheKey, &user, cache.DefaultExpiration)
	return &user, nil
}

func (s *CacheService) GetUserByPlayername(playername string) (*User, error) {
	cacheKey := playername
	if cached, found := s.cache.Get(cacheKey); found {
		return cached.(*User), nil
	}

	// Query DB if not found in cache
	var user User
	if err := DBRead("playername = ?", playername, &user); err != nil {
		return nil, err
	}
	// Cache value we got
	s.cache.Set(cacheKey, &user, cache.DefaultExpiration)
	return &user, nil
}

func (s *CacheService) GetUIDBySessionkey(sessionkey string) (*Session, error) {
	cacheKey := sessionkey
	if cached, found := s.cache.Get(cacheKey); found {
		return cached.(*Session), nil
	}

	// Query DB if not found in cache
	var session Session
	if err := DBRead("session = ?", sessionkey, &session); err != nil {
		return nil, err
	}
	// Cache value we got
	s.cache.Set(cacheKey, &session, cache.DefaultExpiration)
	return &session, nil
}

func (s *CacheService) WriteSession(uid string, sessionkey string, gamever string) (*Session, error) {
	cacheKey := sessionkey
	session := Session{UID: uid, Sessionkey: sessionkey, Gamever: gamever}
	if err := DBWrite(&session); err != nil {
		return nil, fmt.Errorf("Failed to create session in database: %v", err)
	}
	s.cache.Set(cacheKey, &session, cache.DefaultExpiration)
	return &session, nil
}
