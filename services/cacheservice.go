package services

import (
	"fmt"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/patrickmn/go-cache"
)

var CacheService *cache.Cache

func InitCacheService(expirationDuration time.Duration, cleanupInterval time.Duration) {
	CacheService = cache.New(expirationDuration, cleanupInterval)
}

func Read[T any](s *cache.Cache, cacheKey string, entryname string) (*T, error) {
	if cached, found := s.Get(cacheKey); found {
		if data, ok := cached.(*T); ok {
			return data, nil
		}
		return nil, fmt.Errorf("Cache type mismatch!")
	}
	var data T
	if err := Query(DatabaseService, entryname, cacheKey, &data); err != nil {
		return nil, err
	}
	s.Set(cacheKey, &data, cache.DefaultExpiration)
	return &data, nil
}

func WriteNew[T any](s *cache.Cache, cacheKey string, data *T) (*T, error) {
	if err := Create(DatabaseService, data); err != nil {
		return nil, err
	}
	s.Set(cacheKey, data, cache.DefaultExpiration)
	return data, nil
}

func DeleteNew[T any](s *cache.Cache, cacheKey string, entryname string) error {
	var data T
	if err := Delete(DatabaseService, entryname, cacheKey, &data); err != nil {
		return err
	}
	s.Delete(cacheKey)
	return nil
}

func GetUserByUID(uid string) (*User, error) {
	user, err := Read[User](CacheService, uid, "uid")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByPlayername(playername string) (*User, error) {
	user, err := Read[User](CacheService, playername, "playername")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUIDBySessionkey(sessionkey string) (*Session, error) {
	session, err := Read[Session](CacheService, sessionkey, "sessionkey")
	if err != nil {
		return nil, err
	}
	return session, nil
}

func WriteSession(uid string, sessionkey string, gamever string) (*Session, error) {
	sessionval := Session{UID: uid, Sessionkey: sessionkey, Gamever: gamever}
	session, err := WriteNew(CacheService, sessionkey, &sessionval)
	if err != nil {
		return nil, err
	}
	return session, nil
}
