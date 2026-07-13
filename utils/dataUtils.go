package utils

import (
	"github.com/VsAltAuth/VSAA/models"
	"github.com/VsAltAuth/VSAA/services"
)

func GetUserByUID(uid string) (*models.User, error) {
	user, err := services.Read[models.User](services.CacheService, uid, "uid")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByPlayername(playername string) (*models.User, error) {
	user, err := services.Read[models.User](services.CacheService, playername, "playername")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUIDBySessionkey(sessionkey string) (*models.Session, error) {
	session, err := services.Read[models.Session](services.CacheService, sessionkey, "sessionkey")
	if err != nil {
		return nil, err
	}
	return session, nil
}

func WriteSession(uid string, sessionkey string, gamever string) (*models.Session, error) {
	sessionval := models.Session{UID: uid, Sessionkey: sessionkey, Gamever: gamever}
	session, err := services.WriteNew(services.CacheService, sessionkey, &sessionval)
	if err != nil {
		return nil, err
	}
	return session, nil
}
