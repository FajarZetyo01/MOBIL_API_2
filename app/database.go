package app

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	//dsn := "root:root@tcp(127.0.0.1:3306)/crud_golang?parseTime=true" //LOKAL DB
	dsn := "admin:kExAMAF22AzpH3frN7QD@tcp(database-testing.cvagmsww6fzb.ap-southeast-3.rds.amazonaws.com:3306)/mini_project?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("❌ Open DB error:", err)
		return nil
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Println("❌ Ping failed:", err)
		return nil
	}
	log.Println("✅ MySQL Connected")
	return db
}

func MonitorDB(db *sql.DB) {
	ticker := time.NewTicker(5 * time.Second)
	var wasDown bool
	go func() {
		for range ticker.C {

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			err := db.PingContext(ctx)
			cancel()
			if err != nil {
				if !wasDown {
					log.Println("⚠️ MySQL DOWN:", err)
					wasDown = true
				}
				continue
			}

			if wasDown {
				log.Println("✅ MySQL Connected")
				wasDown = false
			}
		}
	}()
}
