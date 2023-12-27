package main

import (
	"github.com/jacobslunga/bubblegum-server/internal/config"
	"github.com/jacobslunga/bubblegum-server/internal/model"
	"github.com/jacobslunga/bubblegum-server/internal/router"
)

func main() {
  db:= config.SetupDB()
  
  err := db.AutoMigrate(&model.User{}, &model.Component{})
  if err != nil {
    panic(err)
  }

  r := router.SetupRouter()
  r.Run()
}
