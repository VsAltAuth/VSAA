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

func Read[T any](s *cache.Cache, entryname string, cacheKey string) (*T, error) {
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

func DeleteNew[T any](s *cache.Cache, entryname string, cacheKey string) error {
	var data T
	if err := Delete(DatabaseService, entryname, cacheKey, &data); err != nil {
		return err
	}
	s.Delete(cacheKey)
	return nil
}

func WriteCache[T any](s *cache.Cache, cacheKey string, data *T) (*T, error) {
	s.Set(cacheKey, data, cache.DefaultExpiration)
	return data, nil
}
