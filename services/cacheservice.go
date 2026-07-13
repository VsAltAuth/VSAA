package services

import (
	"fmt"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/patrickmn/go-cache"
)

var CacheServiceInstance *CacheService

type CacheService struct {
	cache *cache.Cache
}

func NewCacheService(expirationDuration time.Duration, cleanupInterval time.Duration) *CacheService {
	return &CacheService{
		cache: cache.New(expirationDuration, cleanupInterval),
	}
}

func InitCacheService(expirationDuration time.Duration, cleanupInterval time.Duration) error {
	CacheService := NewCacheService(expirationDuration, cleanupInterval)
	if CacheService == nil {
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
	if err := DatabaseService.Query("uid = ?", uid, &user); err != nil {
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
	if err := DatabaseService.Query("playername = ?", playername, &user); err != nil {
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
	if err := DatabaseService.Query("session = ?", sessionkey, &session); err != nil {
		return nil, err
	}
	// Cache value we got
	s.cache.Set(cacheKey, &session, cache.DefaultExpiration)
	return &session, nil
}

func (s *CacheService) WriteSession(uid string, sessionkey string, gamever string) (*Session, error) {
	cacheKey := sessionkey
	session := Session{UID: uid, Sessionkey: sessionkey, Gamever: gamever}
	if err := DatabaseService.Write(&session); err != nil {
		return nil, fmt.Errorf("Failed to create session in database: %v", err)
	}
	s.cache.Set(cacheKey, &session, cache.DefaultExpiration)
	return &session, nil
}
