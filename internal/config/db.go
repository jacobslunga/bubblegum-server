package config

import (
	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
  dsn := "host=localhost user=postgres dbname=bubblegum sslmode=disable password=skorpan12"
  db, err := gorm.Open("postgres", dsn)

  if err != nil {
    panic(err)
  }

  return db
}
