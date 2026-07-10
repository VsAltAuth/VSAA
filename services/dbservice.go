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

type User struct {
	gorm.Model
	UID          string
	Email        string
	HashedPass   string
	Username     string
	Entitlements string
}

type Session struct {
	gorm.Model
	UID        string
	Sessionkey string
	Gamever    string
}

var ctx = context.Background()

func DbInit() {
	fmt.Println("hi from dbinit")
	db, err := gorm.Open(sqlite.Open(DATABASE()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Session{})
	//mkusr(db)
}

func mkusr(db *gorm.DB) {
	err := gorm.G[User](db).Create(ctx, &User{UID: "niggerniggernigger", Email: "megadfga111@xyecoc.com", HashedPass: "dneh", Username: "yar", Entitlements: "VIV"})
	if err != nil {
		print(err)
	}
}
